package models

import "mime/multipart"

type CreateMerk struct {
	NamaMerk         string 				`json:"nama_merk" form:"nama_merk"`
	GambarMerk       multipart.FileHeader 	`json:"gambar_merk" form:"gambar_merk"`
	CreatedAt        string 				`json:"created_at" form:"created_at"`
}
type MerkCreate struct {
	NamaMerk         string 	`json:"nama_merk" form:"nama_merk"`
	GambarMerk       string 	`json:"gambar_merk" form:"gambar_merk"`
	CreatedAt        string 	`json:"created_at" form:"created_at"`
}
type UpdateMerk struct {
	IdMerk       	 string 				`json:"id_merk" form:"id_merk"`
	UpdatedAt        string 				`json:"updated_at" form:"updated_at"`
	NamaMerk         string 				`json:"nama_merk" form:"nama_merk"`
	GambarMerk       *multipart.FileHeader  `json:"gambar_merk" form:"gambar_merk"`
}
type MerkUpdate struct {
	UpdatedAt        string 	`json:"updated_at" form:"updated_at"`
	NamaMerk         *string 	`json:"nama_merk" form:"nama_merk"`
	GambarMerk       *string		`json:"gambar_merk" form:"gambar_merk"`
}


type GetMerk struct {
	IdMerk       string `json:"id_merk" form:"id_merk"`
	NamaMerk     string `json:"nama_merk" form:"nama_merk"`
	GambarMerk   string `json:"gambar_merk" form:"gambar_merk"`
	CreatedAt    string `json:"created_at" form:"created_at"`
	UpdatedAt    string `json:"updated_at" form:"updated_at"`
	DeletedAt    string `json:"deleted_at" form:"deleted_at"`
}



type DeleteMerk struct {
	IdMerk 		string `json:"id_merk" form:"id_merk"`
	DeletedAt  	string `json:"deleted_at" form:"deleted_at"`
}