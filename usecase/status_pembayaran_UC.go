package usecase

import (
	"danantara/entity"
	"errors"

	"gorm.io/gorm"
)

type StatusPembayaranGajiUseCase struct {
	db *gorm.DB
}

func NewStatusPembayaranGajiUseCase(db *gorm.DB) *StatusPembayaranGajiUseCase {
	return &StatusPembayaranGajiUseCase{db: db}
}

func (uc *StatusPembayaranGajiUseCase) CreateStatusPembayaranGaji(pembayaran *entity.StatusPembayaranGaji) error {
	result := uc.db.Create(pembayaran)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *StatusPembayaranGajiUseCase) GetStatusPembayaranGajiByID(id uint) (*entity.StatusPembayaranGaji, error) {
	var pembayaran entity.StatusPembayaranGaji
	result := uc.db.First(&pembayaran, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data pembayaran tidak ditemukan")
		}
		return nil, result.Error
	}
	return &pembayaran, nil
}

func (uc *StatusPembayaranGajiUseCase) GetStatusPembayaranGajiByKaryawanID(idKaryawan uint) ([]entity.StatusPembayaranGaji, error) {
	var pembayaran []entity.StatusPembayaranGaji
	result := uc.db.Where("id_karyawan = ?", idKaryawan).Find(&pembayaran)
	if result.Error != nil {
		return nil, result.Error
	}
	return pembayaran, nil
}

func (uc *StatusPembayaranGajiUseCase) GetAllStatusPembayaranGaji() ([]entity.StatusPembayaranGaji, error) {
	var pembayaran []entity.StatusPembayaranGaji
	result := uc.db.Find(&pembayaran)
	if result.Error != nil {
		return nil, result.Error
	}
	return pembayaran, nil
}

func (uc *StatusPembayaranGajiUseCase) UpdateStatusPembayaranGaji(pembayaran *entity.StatusPembayaranGaji) error {
	result := uc.db.Save(pembayaran)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *StatusPembayaranGajiUseCase) DeleteStatusPembayaranGaji(id uint) error {
	result := uc.db.Delete(&entity.StatusPembayaranGaji{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
