package models

import (
	"mime/multipart"
)

type CreateProduk struct {
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		float64 `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	int `json:"stok_produk" form:"stok_produk"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	GambarProduk      	*multipart.FileHeader `json:"gambar_produk" form:"gambar_produk"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}
type ProdukCreate struct {
	IdProduk			string `json:"id_produk" form:"id_produk"`
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		float64 `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	int `json:"stok_produk" form:"stok_produk"`
	Tayangan      		int `json:"tayangan" form:"tayangan"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	GambarProduk      	string `json:"gambar_produk" form:"gambar_produk"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
}
type UpdateProduk struct {
	IdProduk       	 	string `json:"id_produk" form:"id_produk"`
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		float64 `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	int `json:"stok_produk" form:"stok_produk"`
	Tayangan      		int `json:"tayangan" form:"tayangan"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	GambarProduk      	*multipart.FileHeader `json:"gambar_produk" form:"gambar_produk"`
}
type ProdukUpdate struct {
	IdProduk			string `json:"id_produk" form:"id_produk"`
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		float64 `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	int `json:"stok_produk" form:"stok_produk"`
	Tayangan      		int `json:"tayangan" form:"tayangan"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	GambarProduk      	*string `json:"gambar_produk" form:"gambar_produk"`
	UpdatedAt        	string `json:"updated_at" form:"updated_at"`
}

type GetProduk struct {
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	Search       		string `json:"search" form:"search"`
}


type ProdukGet struct {
	IdProduk       		string `json:"id_produk" form:"id_produk"`
	NamaProduk      	string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    		string `json:"harga_produk" form:"harga_produk"`
	KodeProduk      	string `json:"kode_produk" form:"kode_produk"`
	TipeProduk      	string `json:"tipe_produk" form:"tipe_produk"`
	BeratProduk     	string `json:"berat_produk" form:"berat_produk"`
	IdMerkProduk   	 	string `json:"id_merk_produk" form:"id_merk_produk"`
	NamaMerk   	 		string `json:"nama_merk" form:"nama_merk"`
	GambarMerk   	 	string `json:"gambar_merk" form:"gambar_merk"`
	DeskripsiProduk 	string `json:"deskripsi_produk" form:"deskripsi_produk"`
	StokProduk      	string `json:"stok_produk" form:"stok_produk"`
	Tayangan      		string `json:"tayangan" form:"tayangan"`
	IdKategoriProduk    string `json:"id_kategori_produk" form:"id_kategori_produk"`
	Kategori    		string `json:"kategori" form:"kategori"`
	SubKategori	    	string `json:"sub_kategori" form:"sub_kategori"`
	GambarProduk      	string `json:"gambar_produk" form:"gambar_produk"`
	CreatedAt        	string `json:"created_at" form:"created_at"`
	UpdatedAt        	string `json:"updated_at" form:"updated_at"`
	DeletedAt        	string `json:"deleted_at" form:"deleted_at"`
}





type DeleteProduk struct {
	IdProduk   string `json:"id_produk" form:"id_produk"`
	DeletedAt  string `json:"deleted_at" form:"deleted_at"`
}