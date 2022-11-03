package controllers

type InputTambahBuku struct {
	Judul    string `json:"judul" binding:"required"`
	Penulis  string `json:"penulis" binding:"required"`
	Penerbit string `json:"penerbit" binding:"required"`
}

type InputUpdateBuku struct {
	Judul    string `json:"judul"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}
