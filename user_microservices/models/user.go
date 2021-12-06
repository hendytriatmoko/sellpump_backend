package models

import "mime/multipart"

type CreateUser struct {
	Nama        string `json:"nama" form:"nama"`
	NoTelp		string `json:"no_telp" form:"no_telp"`
	Email       string `json:"email" form:"email"`
	Foto      	*multipart.FileHeader `json:"foto" form:"foto"`
	Password    string `json:"password" form:"password"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}
type UserCreate struct {
	IdUser		string `json:"id_user" form:"id_user"`
	Nama        string `json:"nama" form:"nama"`
	NoTelp		string `json:"no_telp" form:"no_telp"`
	Email       string `json:"email" form:"email"`
	Foto      	string `json:"foto" form:"foto"`
	Password    string `json:"password" form:"password"`
	Status    	string `json:"status" form:"status"`
	Verifikasi 	bool   `json:"verifikasi" form:"verifikasi"`
	CreatedAt   string `json:"created_at" form:"created_at"`
}

type VerifikasiUser struct {
	IdVerifikasi	string `json:"id_verifikasi" form:"id_verifikasi"`
	IdUser			string `json:"id_user" form:"id_user"`
	Email       	string `json:"email" form:"email"`
	Status    		bool `json:"status" form:"status"`
	CreatedAt   	string `json:"created_at" form:"created_at"`
	ExpiredAt   	string `json:"expired_at" form:"expired_at"`
}

type GetUser struct {
	IdUser       		string `json:"id_user" form:"id_user"`
}

type UserGet struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Nama         	string `json:"nama" form:"nama"`
	Email      		string `json:"email" form:"email"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Foto      		string `json:"foto" form:"foto"`
	Status         	string `json:"status" form:"status"`
	Verifikasi  	string `json:"verifikasi" form:"verifikasi"`
	CreatedAt       string `json:"created_at" form:"created_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
	DeletedAt       string `json:"deleted_at" form:"deleted_at"`
}


type UpdateUser struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Nama         	string `json:"nama" form:"nama"`
	Email      		string `json:"email" form:"email"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Password    	string `json:"password" form:"password"`
	Foto      		*multipart.FileHeader `json:"foto" form:"foto"`
}

type UserUpdate struct {
	IdUser       	string `json:"id_user" form:"id_user"`
	Nama         	string `json:"nama" form:"nama"`
	Email      		string `json:"email" form:"email"`
	NoTelp         	string `json:"no_telp" form:"no_telp"`
	Password    	string `json:"password" form:"password"`
	Foto      		*string `json:"foto" form:"foto"`
	Status         	string `json:"status" form:"status"`
	Verifikasi  	string `json:"verifikasi" form:"verifikasi"`
	CreatedAt       string `json:"created_at" form:"created_at"`
	UpdatedAt       string `json:"updated_at" form:"updated_at"`
}

type DeleteUser struct {
	IdUser 		string `json:"id_user" form:"id_user"`
	DeletedAt  	string `json:"deleted_at" form:"deleted_at"`
}