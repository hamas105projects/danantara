package entity

type Kehadiran struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	IDKaryawan uint   `gorm:"not null"`
	Bulan      string `gorm:"type:varchar(20);not null"`
	Tahun      uint   `gorm:"not null"`
	Hadir      uint   `gorm:"not null"`
	Alpa       uint   `gorm:"not null"`
	Sakit      uint   `gorm:"not null"`
	Izin       uint   `gorm:"not null"`
}
