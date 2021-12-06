package main

import (
	"fmt"
	_ "github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"io"
	"log"
	"merk_microservices/common"
	"merk_microservices/controllers"
	"merk_microservices/databases"
	"merk_microservices/models"
	"os"
	//_ "./docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	// Load config file
	err = common.LoadConfig()
	if err != nil {
		fmt.Println("error", err.Error())
		return err
	}

	// Initialize User database
	err = databases.DatabaseSellPump.Init()
	if err != nil {
		fmt.Println("error db", err.Error())
		return err
	}


	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		f, err := os.OpenFile("logs/gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	//gin.SetMode(gin.ReleaseMode)
	m.router = gin.Default()
	m.router.Use(cors.AllowAll())

	return nil
}

func main() {
	merk := controllers.Merk{}

	m := Main{}
	// Initialize server
	if m.initServer() != nil {
		return
	}

	defer databases.DatabaseSellPump.DB.Close()

	m.router.NoRoute(func(c *gin.Context) {
		response := models.Response{}
		response.ApiMessage = "Page Not Found"
		c.JSON(404, response)
	})

	f, err := os.OpenFile("user.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	//public := m.router.Group("/warna/api/")
	//{
	//	public.Static("files/", "./files/doc/")
	//	public.Static("files1/", "/files/doc/")
	//	public.Static("files2/", "files/doc/")
	//	public.Static("files3/", "./doc/")
	//	public.Static("files4/", "/doc/")
	//	public.Static("files5/", "doc/")
	//	public.Static("files6/", "/")
	//}
	// Simple group: v1
	api := m.router.Group("/sellpump/api/merk")
	//api.Use(middleware.Auth)
	{
		api.Static("photo/", "./files")
		v1 := api.Group("/v1")
		//v1.Use(middleware.Auth)
		{
			v1.GET("/getmerk", merk.GetDataMerk)
			v1.POST("/create", merk.MerkCreate)
			v1.PUT("/update", merk.MerkUpdate)
			v1.DELETE("/delete", merk.MerkDelete)
		}

	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	m.router.Run(common.Config.Port)
}
