package models

// Pasien merepresentasikan struktur data pasien dalam database
type Pasien struct {
	Id       int    `json:"id"`       // ID unik pasien
	Nama     string `json:"nama"`     // Nama pasien
	Alamat   string `json:"alamat"`   // Alamat pasien
	NoHP     string `json:"no_hp"`    // Nomor HP pasien
	Penyakit string `json:"penyakit"` // Penyakit yang diderita pasien
}
