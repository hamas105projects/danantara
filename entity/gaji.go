package entity

type GajiPokok struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Jabatan   string `gorm:"type:varchar(100);not null"`
	GajiPokok uint   `gorm:"not null"`
}
