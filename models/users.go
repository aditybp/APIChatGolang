package models

type User struct {
	Id          uint   `json:"id_user"`
	NoHandphone string `json:"no_handphone"`
	Nama        string `json:"nama_user" `
	Password    []byte `json:"password"`
}
