package model

// Package model berisi representasi entitas yang dilibatkan dalam service

// Makul mewakilkan data MataKuliah
type Makul struct {
	Nama      string `json:"nama_matakuliah"`
	JumlahSKS int    `json:"jumlah_sks"`
	UUID      string `json:"uuid"`
}

// MakulResponse akan digunakan sebagai format response json
type MakulResponse struct {
	Count int     `json:"total"`
	Data  []Makul `json:"data"`
}
