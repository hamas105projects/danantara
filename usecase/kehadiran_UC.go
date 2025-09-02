package usecase

import (
	"danantara/entity"
	"errors"

	"gorm.io/gorm"
)

type KehadiranUseCase struct {
	db *gorm.DB
}

func NewKehadiranUseCase(db *gorm.DB) *KehadiranUseCase {
	return &KehadiranUseCase{db: db}
}

func (uc *KehadiranUseCase) GetKehadiranByID(id uint) (*entity.Kehadiran, error) {
	var kehadiran entity.Kehadiran
	result := uc.db.First(&kehadiran, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data kehadiran tidak ditemukan")
		}
		return nil, result.Error
	}
	return &kehadiran, nil
}

func (uc *KehadiranUseCase) GetKehadiranByKaryawanIDAndBulanTahun(idKaryawan uint, bulan string, tahun int) (*entity.Kehadiran, error) {
	var kehadiran entity.Kehadiran
	result := uc.db.Where("id_karyawan = ? AND bulan = ? AND tahun = ?", idKaryawan, bulan, tahun).First(&kehadiran)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("data kehadiran untuk karyawan, bulan, dan tahun tersebut tidak ditemukan")
		}
		return nil, result.Error
	}
	return &kehadiran, nil
}

func (uc *KehadiranUseCase) GetAllKehadiran() ([]entity.Kehadiran, error) {
	var kehadiran []entity.Kehadiran
	result := uc.db.Find(&kehadiran)
	if result.Error != nil {
		return nil, result.Error
	}
	return kehadiran, nil
}

func (uc *KehadiranUseCase) CreateKehadiran(kehadiran *entity.Kehadiran) error {
	var existingKehadiran entity.Kehadiran
	if err := uc.db.Where("id_karyawan = ? AND bulan = ? AND tahun = ?", kehadiran.IDKaryawan, kehadiran.Bulan, kehadiran.Tahun).First(&existingKehadiran).Error; err == nil {
		return errors.New("data kehadiran untuk karyawan, bulan, dan tahun ini sudah ada")
	}
	result := uc.db.Create(kehadiran)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *KehadiranUseCase) UpdateKehadiran(kehadiran *entity.Kehadiran) error {
	result := uc.db.Save(kehadiran)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (uc *KehadiranUseCase) DeleteKehadiran(id uint) error {
	result := uc.db.Delete(&entity.Kehadiran{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
