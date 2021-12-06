package models

type CreateKategori struct {
	Kategori         string `json:"kategori" form:"kategori"`
	SubKategori      string `json:"sub_kategori" form:"sub_kategori"`
	CreatedAt        string `json:"created_at" form:"created_at"`
}

type GetKategori struct {
	IdKategori       string `json:"id_kategori" form:"id_kategori"`
	Kategori         string `json:"kategori" form:"kategori"`
	SubKategori      string `json:"sub_kategori" form:"sub_kategori"`
	CreatedAt        string `json:"created_at" form:"created_at"`
	UpdatedAt        string `json:"updated_at" form:"updated_at"`
	DeletedAt        string `json:"deleted_at" form:"deleted_at"`
}

type UpdateKategori struct {
	IdKategori       string `json:"id_kategori" form:"id_kategori"`
	UpdatedAt        string `json:"updated_at" form:"updated_at"`
	Kategori         string `json:"kategori" form:"kategori"`
	SubKategori      string `json:"sub_kategori" form:"sub_kategori"`
}

type DeleteKategori struct {
	IdKategori string `json:"id_kategori" form:"id_kategori"`
	DeletedAt  string `json:"deleted_at" form:"deleted_at"`
}