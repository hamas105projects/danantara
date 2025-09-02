package entity

type Karyawan struct {
	IDKaryawan uint   `gorm:"primaryKey;autoIncrement;column:id_karyawan"`
	Nama       string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:varchar(255);not null"`
	NomorHp    string `gorm:"type:varchar(20);not null"`
}
