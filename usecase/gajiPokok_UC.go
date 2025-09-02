package usecase

import (
	"danantara/entity"
	"errors"

	"gorm.io/gorm"
)

type GajiPokokUseCase struct {
	db *gorm.DB
}

func NewGajiPokokUseCase(db *gorm.DB) *GajiPokokUseCase {
	return &GajiPokokUseCase{db: db}
}

func (uc *GajiPokokUseCase) GetGajiPokokByJabatan(jabatan string) (*entity.GajiPokok, error) {
	var gajiPokok entity.GajiPokok
	result := uc.db.Where("jabatan = ?", jabatan).First(&gajiPokok)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("gaji pokok untuk jabatan tidak ditemukan")
		}
		return nil, result.Error
	}
	return &gajiPokok, nil
}

func (uc *GajiPokokUseCase) GetGajiPokokByID(id uint) (*entity.GajiPokok, error) {
	var gajiPokok entity.GajiPokok
	result := uc.db.First(&gajiPokok, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("gaji pokok tidak ditemukan")
		}
		return nil, result.Error
	}
	return &gajiPokok, nil
}

func (uc *GajiPokokUseCase) GetAllGajiPokok() ([]entity.GajiPokok, error) {
	var gajiPokok []entity.GajiPokok
	result := uc.db.Find(&gajiPokok)
	if result.Error != nil {
		return nil, result.Error
	}
	return gajiPokok, nil
}

func (uc *GajiPokokUseCase) CreateGajiPokok(gajiPokok *entity.GajiPokok) error {
	var existingGajiPokok entity.GajiPokok
	if err := uc.db.Where("jabatan = ?", gajiPokok.Jabatan).First(&existingGajiPokok).Error; err == nil {
		return errors.New("gaji pokok untuk jabatan ini sudah ada")
	}
	result := uc.db.Create(gajiPokok)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *GajiPokokUseCase) UpdateGajiPokok(gajiPokok *entity.GajiPokok) error {
	result := uc.db.Save(gajiPokok)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *GajiPokokUseCase) DeleteGajiPokok(id uint) error {
	result := uc.db.Delete(&entity.GajiPokok{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
