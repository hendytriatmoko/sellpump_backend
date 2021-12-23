package daos

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"user_microservices/databases"
	"user_microservices/helper"
	"user_microservices/middleware"
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
	user.Verifikasi = "N"
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
	verifikasi.Status = "N"
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
	if params.Email != "" {
		err = err.Where("email = ?", params.Email)
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
		path := "/user/"
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

func (m *User) UserCheckAkun(params models.CheckAkunUser) error {

	checkakun := models.CheckAkunRead{}
	var check bool
	today := m.helper.GetTimeNow()

	check = databases.DatabaseSellPump.DB.Table("verifikasi v").
		Where("v.email = ?", params.Email).Find(&checkakun).RecordNotFound()

	if check == true {
		err := errors.New("Email Tidak Ditemukan")
		return err
	}else if check == false {
		if today > checkakun.ExpiredAt {
			err := errors.New("Verification Expired")
			return err
		}else if today < checkakun.ExpiredAt {
			if checkakun.Status == "N" {
				err := errors.New("Silahkan Verifikasi Email Anda")
				return err
			}
		}
	}

	return error(nil)

}

func (m *User) UserResendVerification(params models.CheckAkunUser) (models.CheckAkunRead, error) {

	updateverifikasi := models.CheckAkunRead{}

	ti := time.Now()
	ti_n := ti.AddDate(0, 0, 7)
	next := string(ti_n.Format("2006-01-02 15:04:05.999999"))
	updateverifikasi.ExpiredAt = next

	err := databases.DatabaseSellPump.DB.Table("verifikasi").Where("email = ?", params.Email).Update(&updateverifikasi).Error

	if err != nil {
		return models.CheckAkunRead{}, err
	}

	errx := databases.DatabaseSellPump.DB.Table("verifikasi").Where("email = ?", params.Email).Find(&updateverifikasi).Error

	if errx != nil {
		return models.CheckAkunRead{}, errx
	}

	erry := m.helper.SendEmailVerifikasi(updateverifikasi.Email, updateverifikasi.IdUser, updateverifikasi.IdVerifikasi)
	if erry != nil {
		return updateverifikasi, erry
	}

	return updateverifikasi, nil

}

func (m *User) UserForgotPassword(params models.CheckAkunUser) (models.UserGet, error) {

	updateverifikasi := models.UserGet{}

	errx := databases.DatabaseSellPump.DB.Table("user").Where("email = ?", params.Email).Find(&updateverifikasi).Error

	if errx != nil {
		return models.UserGet{}, errx
	}

	erry := m.helper.SendForgotPassword(updateverifikasi.Email)
	if erry != nil {
		return updateverifikasi, erry
	}

	return updateverifikasi, nil

}

func (m *User) LoginCheck(params models.UserToken) error {

	checkakun := models.UserGet{}
	var check bool

	check = databases.DatabaseSellPump.DB.Table("user").
		Where("email = ?", params.Email).Find(&checkakun).RecordNotFound()

	if check == true {
		err := errors.New("Email Tidak Ditemukan")
		return err
	}

	return error(nil)

}

func (m *User) UserGetLogin(params models.UserLogin) ([]models.UserGet, error) {

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

func (m *User) Signin(params models.UserToken) ([]models.UserGet, string, error) {

	userGet := models.GetUser{}
	userRead := []models.UserGet{}
	updateToken := models.UserUpdate{}
	var token string
	var er error

	err := m.LoginCheck(params)

	if err != nil {
		return userRead, "", err
	}

	if params.Email != "" {
		userGet.Email = params.Email
	}
	if params.Password != "" {
		userGet.Password = params.Password
	}

	userRead, err = m.UserGet(userGet)

	if err != nil {
		return userRead, "", err
	}

	if userRead[0].Verifikasi == "N" {
		err = errors.New("Akun Anda di Nonaktifkan, Tidak Dapat di Akses")
		return userRead, "", err
	}

	//token, er := m.helper.GetToken(userRead[0].IdUser)
	password,_ := DecryptPassword(userRead[0].Password)

	if userRead[0].Email == params.Email && params.Password == password && userRead[0].Verifikasi == "Y"  {
		fmt.Println("cocok")
		token, er = middleware.CreateAuth(userRead[0].IdUser, "user", "none", "none")

		if er != nil {
			return userRead, "", er
		}

		updateToken.Token = token

		err = databases.DatabaseSellPump.DB.Table("user").Where("id_user = ?", userRead[0].IdUser).Update(&updateToken).Error

		if err != nil {
			return userRead, "", err
		}
	}else {
		fmt.Println("tidak cocok")
	}



	return userRead, token, nil

}