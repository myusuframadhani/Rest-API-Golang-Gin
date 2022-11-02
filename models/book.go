package models

type Buku struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Judul    string `json:"judul"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
	Harga    uint64 `json:"harga"`
}
