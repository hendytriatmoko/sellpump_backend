package daos

import (
	"merk_microservices/databases"
	"merk_microservices/helper"
	"merk_microservices/models"
	"os"
	"path/filepath"
	"strings"
)

type Merk struct {
	helper helper.Helper
}

func (m *Merk) MerkCreate(params models.CreateMerk) (models.MerkCreate, error) {

	insertmerk := models.MerkCreate{}

	path := "/merk/"
	pathImage := "./files/"+path
	ext := filepath.Ext(params.GambarMerk.Filename)
	filename := strings.Replace(params.NamaMerk," ","_", -1)+ext

	os.MkdirAll(pathImage, 0777)
	errx := m.helper.SaveUploadedFile(params.GambarMerk, pathImage+filename)
	if errx != nil{
		return models.MerkCreate{},errx
	}

	url := string(filepath.FromSlash(path+filename))

	insertmerk.NamaMerk = params.NamaMerk
	insertmerk.GambarMerk = url
	insertmerk.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("merk").Create(&insertmerk).Error

	if err != nil {
		return models.MerkCreate{}, err
	}

	return insertmerk, nil
}


func (m *Merk) MerkUpdate(params models.UpdateMerk) (models.MerkUpdate, error) {

	merk := models.MerkUpdate{}

	if params.GambarMerk != nil {
		path := "/merk/"
		pathImage := "./files/"+path
		ext := filepath.Ext(params.GambarMerk.Filename)
		filename := strings.Replace(params.NamaMerk," ","_", -1)+ext

		os.MkdirAll(pathImage, 0777)
		errx := m.helper.SaveUploadedFile(*params.GambarMerk, pathImage+filename)
		if errx != nil{
			return models.MerkUpdate{},errx
		}

		url := string(filepath.FromSlash(path+filename))

		merk.GambarMerk = new(string)
		*merk.GambarMerk = url
	}

	merk.UpdatedAt = m.helper.GetTimeNow()
	merk.NamaMerk = new(string)
	*merk.NamaMerk = params.NamaMerk

	err := databases.DatabaseSellPump.DB.Table("merk").Where("id_merk = ?", params.IdMerk).Update(&merk).Error

	if err != nil {
		return models.MerkUpdate{}, err
	}

	return merk, nil

}

func (m *Merk) MerkGet() ([]models.GetMerk, error) {

	merk := []models.GetMerk{}

	err := databases.DatabaseSellPump.DB.Table("merk").Find(&merk).Error

	if err != nil {
		return []models.GetMerk{}, err
	}

	return merk, nil
}

func (m *Merk) MerkDelete(params models.DeleteMerk) (models.DeleteMerk, error) {

	merk := models.DeleteMerk{}

	merk.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("merk").Where("id_merk = ?", params.IdMerk).Update(&merk).Error

	if err != nil {
		return models.DeleteMerk{}, err
	}

	return merk, nil

}