package models

type GetKeranjang struct {
	IdUser      string `json:"id_user" form:"id_user"`
	Email      	string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
}

type KeranjangGet struct {
	IdKeranjang    	string `json:"id_keranjang" form:"id_keranjang"`
	IdUser         	string `json:"id_user" form:"id_user"`
	IdProduk   		string `json:"id_produk" form:"id_produk"`
	NamaProduk      string `json:"nama_produk" form:"nama_produk"`
	HargaProduk    	float64 `json:"harga_produk" form:"harga_produk"`
	CreatedAt       string `json:"created_at" form:"created_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
	DeletedAt       string `json:"deleted_at" form:"deleted_at"`
}

type CreateKeranjang struct {
	IdKeranjang string `json:"id_keranjang" form:"id_keranjang"`
	IdUser      string `json:"id_user" form:"id_user"`
	IdProduk   	string `json:"id_produk" form:"id_produk"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type KeranjangCreate struct {
	IdKeranjang string `json:"id_keranjang" form:"id_keranjang"`
	IdUser      string `json:"id_user" form:"id_user"`
	IdProduk   	string `json:"id_produk" form:"id_produk"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type DeleteKeranjang struct {
	IdKeranjang   	string `json:"id_keranjang" form:"id_keranjang"`
	DeletedAt  		string `json:"deleted_at" form:"deleted_at"`
}