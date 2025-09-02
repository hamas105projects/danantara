package usecase

import (
	"danantara/entity"
	"errors"

	"gorm.io/gorm"
)

type KaryawanUseCase struct {
	db *gorm.DB
}

func NewKaryawanUseCase(db *gorm.DB) *KaryawanUseCase {
	return &KaryawanUseCase{db: db}
}

func (uc *KaryawanUseCase) GetKaryawanByID(id uint) (*entity.Karyawan, error) {
	var karyawan entity.Karyawan
	result := uc.db.First(&karyawan, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("karyawan tidak ditemukan")
		}
		return nil, result.Error
	}
	return &karyawan, nil
}

func (uc *KaryawanUseCase) GetKaryawanByEmail(email string) (*entity.Karyawan, error) {
	var karyawan entity.Karyawan
	result := uc.db.Where("email = ?", email).First(&karyawan)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("karyawan dengan email tersebut tidak ditemukan")
		}
		return nil, result.Error
	}
	return &karyawan, nil
}

func (uc *KaryawanUseCase) GetAllKaryawan() ([]entity.Karyawan, error) {
	var karyawan []entity.Karyawan
	result := uc.db.Find(&karyawan)
	if result.Error != nil {
		return nil, result.Error
	}
	return karyawan, nil
}

func (uc *KaryawanUseCase) CreateKaryawan(karyawan *entity.Karyawan) error {
	var existingKaryawan entity.Karyawan
	if err := uc.db.Where("email = ?", karyawan.Email).First(&existingKaryawan).Error; err == nil {
		return errors.New("karyawan dengan email ini sudah ada")
	}
	result := uc.db.Create(karyawan)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *KaryawanUseCase) UpdateKaryawan(karyawan *entity.Karyawan) error {
	result := uc.db.Save(karyawan)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *KaryawanUseCase) DeleteKaryawan(id uint) error {
	result := uc.db.Delete(&entity.Karyawan{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
