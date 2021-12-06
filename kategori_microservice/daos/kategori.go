package daos

import (
	"kategori_microservice/databases"
	"kategori_microservice/helper"
	"kategori_microservice/models"
)

type Kategori struct {
	helper helper.Helper
}

func (m *Kategori) KategoriCreate(params models.CreateKategori) (models.CreateKategori, error) {

	kategori := models.CreateKategori{}


	kategori.Kategori = params.Kategori
	kategori.SubKategori = params.SubKategori
	kategori.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("kategori").Create(&kategori).Error

	if err != nil {
		return models.CreateKategori{}, err
	}

	return kategori, nil
}

func (m *Kategori) KategoriGet() ([]models.GetKategori, error) {

	kategori := []models.GetKategori{}

	err := databases.DatabaseSellPump.DB.Table("kategori").Find(&kategori).Error

	if err != nil {
		return []models.GetKategori{}, err
	}

	return kategori, nil
}

func (m *Kategori) KategoriUpdate(params models.UpdateKategori) (models.UpdateKategori, error) {

	kategori := models.UpdateKategori{}

	kategori.UpdatedAt = m.helper.GetTimeNow()
	kategori.Kategori = params.Kategori
	kategori.SubKategori = params.SubKategori

	err := databases.DatabaseSellPump.DB.Table("kategori").Where("id_kategori = ?", params.IdKategori).Update(&kategori).Error

	if err != nil {
		return models.UpdateKategori{}, err
	}

	return kategori, nil

}

func (m *Kategori) KategoriDelete(params models.DeleteKategori) (models.DeleteKategori, error) {

	kategori := models.DeleteKategori{}

	kategori.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("kategori").Where("id_kategori = ?", params.IdKategori).Update(&kategori).Error

	if err != nil {
		return models.DeleteKategori{}, err
	}

	return kategori, nil

}