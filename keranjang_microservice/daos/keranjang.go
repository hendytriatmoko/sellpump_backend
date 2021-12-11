package daos

import (
	"keranjang_microservice/databases"
	"keranjang_microservice/helper"
	"keranjang_microservice/models"
	_ "os"
	_ "path/filepath"
	_ "strings"
)

type Keranjang struct {
	helper helper.Helper
}

func (m *Keranjang) KeranjangGet(params models.KeranjangGet) ([]models.KeranjangGet, error) {

	keranjang := []models.KeranjangGet{}

	err := databases.DatabaseSellPump.DB.Table("keranjang").Find(&keranjang).Error

	if err != nil {
		return []models.KeranjangGet{}, err
	}

	return keranjang, nil
}

func (m *Keranjang) KeranjangCreate(params models.CreateKeranjang) (models.KeranjangCreate, error) {

	keranjang := models.KeranjangCreate{}

	keranjang.IdKeranjang = m.helper.StringWithCharset()
	keranjang.IdProduk = params.IdProduk
	keranjang.IdUser = params.IdUser
	keranjang.CreatedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("keranjang").Create(&keranjang).Error

	if err != nil {
		return models.KeranjangCreate{}, err
	}

	return keranjang, nil
}

func (m *Keranjang) KeranjangDelete(params models.DeleteKeranjang) (models.DeleteKeranjang, error) {

	keranjang := models.DeleteKeranjang{}

	keranjang.DeletedAt = m.helper.GetTimeNow()

	err := databases.DatabaseSellPump.DB.Table("keranjang").Where("id_keranjang = ?", params.IdKeranjang).Update(&keranjang).Error

	if err != nil {
		return models.DeleteKeranjang{}, err
	}

	return keranjang, nil

}