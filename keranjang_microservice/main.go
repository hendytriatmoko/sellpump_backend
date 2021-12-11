package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"io"
	"keranjang_microservice/common"
	"keranjang_microservice/controllers"
	"keranjang_microservice/databases"
	"keranjang_microservice/middleware"
	"keranjang_microservice/models"
	"log"
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
	keranjang := controllers.Keranjang{}

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
	api := m.router.Group("/sellpump/api/keranjang")
	store := cookie.NewStore([]byte(middleware.JwtKey()))
	api.Use(sessions.Sessions("backend", store))
	//api.Use(middleware.Auth)
	{
		api.Static("photo/", "./files")
		v1 := api.Group("/v1")
		//v1.Use(middleware.Auth)
		{
			userEP := v1.Group("keranjang")
			{
				authUserEP := userEP.Group("")
				authUserEP.Use(middleware.Auth)

				authUserEP.GET("/getkeranjang", keranjang.GetDataKeranjang)
				authUserEP.POST("/create", keranjang.KeranjangCreate)
				authUserEP.DELETE("/delete", keranjang.KeranjangDelete)
			}
		}

	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	m.router.Run(common.Config.Port)
}
