package daos

import (
	"os"
	"path/filepath"
	"strings"
	"time"
	"user_microservices/databases"
	"user_microservices/helper"
	"user_microservices/models"
)

type User struct {
	helper helper.Helper
}

func (m *User) UserCreate(params models.CreateUser) (models.UserCreate, error) {

	user := models.UserCreate{}

	if params.Foto != nil {
		path := "/user/"

		pathImage := "./files/"+path
		ext := filepath.Ext(params.Foto.Filename)
		filename := strings.Replace(params.Nama," ","_", -1)+params.NoTelp+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Foto, pathImage+filename)
		if errx != nil{
			return models.UserCreate{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		//user.Foto = new(string)
		user.Foto = url
	}

	user.IdUser = m.helper.StringWithCharset()
	user.Nama = params.Nama
	user.NoTelp = params.NoTelp
	user.Email = params.Email
	user.Password,_ = EncryptPassword(params.Password)
	user.Status = "pembeli"
	user.Verifikasi = false
	user.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("user").Create(&user).Error

	if err != nil {
		return models.UserCreate{}, err
	}

	dataverifikasi := models.VerifikasiUser{}
	dataverifikasi.IdUser = user.IdUser
	dataverifikasi.Email = user.Email
	errx := m.VerifikasiUser(dataverifikasi)
	if errx != nil {
		return models.UserCreate{}, errx
	}

	return user, nil
}

func (m *User) VerifikasiUser(params models.VerifikasiUser) error {
	verifikasi := models.VerifikasiUser{}

	verifikasi.IdVerifikasi = m.helper.StringWithCharset()
	verifikasi.IdUser = params.IdUser
	verifikasi.Email = params.Email
	verifikasi.Status = false
	verifikasi.CreatedAt = m.helper.GetTimeNow()

	ti := time.Now()
	ti_n := ti.AddDate(0, 0, 7)
	next := string(ti_n.Format("2006-01-02 15:04:05.999999"))
	verifikasi.ExpiredAt = next

	err := databases.DatabaseSellPump.DB.Table("verifikasi").Create(&verifikasi).Error

	if err != nil {
		return err
	}

	err = m.helper.SendEmailVerifikasi(verifikasi.Email, verifikasi.IdUser, verifikasi.IdVerifikasi)

	if err != nil {
		return err
	}

	return nil
}

func (m *User) UserGet(params models.GetUser) ([]models.UserGet, error) {

	user := []models.UserGet{}

	err := databases.DatabaseSellPump.DB.Table("user")
	if params.IdUser != "" {
		err = err.Where("id_user = ?", params.IdUser)
	}

	err = err.Find(&user)

	errx := err.Error


	if errx != nil {
		return []models.UserGet{}, errx
	}

	return user, nil
}

func (m *User) UserUpdate(params models.UpdateUser) ([]models.UserGet, error) {

	user := models.UserUpdate{}
	getuser := []models.UserGet{}

	if params.Foto != nil {
		path := "/produk/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.Foto.Filename)
		filename := strings.Replace(params.Nama," ","_", -1)+params.NoTelp+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(params.Foto, pathImage+filename)
		if errx != nil{
			return []models.UserGet{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		user.Foto = new(string)
		*user.Foto = url
	}
	user.UpdatedAt = m.helper.GetTimeNow()
	user.Nama = params.Nama
	user.Email = params.Email
	user.NoTelp = params.NoTelp
	if params.Password != "" {
		user.Password,_ = EncryptPassword(params.Password)
	}

	err := databases.DatabaseSellPump.DB.Table("user").Where("id_user = ?", params.IdUser).Update(&user).Error

	if err != nil {
		return []models.UserGet{}, err
	}

	paramuser := models.GetUser{}
	paramuser.IdUser = params.IdUser
	getuser,errx := m.UserGet(paramuser)
	if errx != nil {
		return []models.UserGet{}, errx
	}
	return getuser, nil

}

func (m *User) UserDelete(params models.DeleteUser) (models.DeleteUser, error) {

	user := models.DeleteUser{}

	user.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("user").Where("id_user = ?", params.IdUser).Update(&user).Error

	if err != nil {
		return models.DeleteUser{}, err
	}

	return user, nil

}