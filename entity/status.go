package entity

import "time"

type StatusPembayaranGaji struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	IDKaryawan        uint      `gorm:"not null"`
	IDGaji            uint      `gorm:"not null"`
	TanggalPembayaran time.Time `gorm:"not null"`
	Status            string    `gorm:"type:enum('Pending', 'Done');not null"`
	StatusPrestasi    string    `gorm:"type:varchar(100)"`
	TunjanganPrestasi uint      `gorm:"default:0"`
	JenisPelanggaran  string    `gorm:"type:text"`
	PotonganGaji      uint      `gorm:"default:0"`
	TotalGaji         uint      `gorm:"not null"`
}
