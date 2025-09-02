package handler

import (
	"net/http"
	"strconv"

	"danantara/entity"
	"danantara/usecase"

	"github.com/gin-gonic/gin"
)

type KaryawanHandler struct {
	karyawanUseCase *usecase.KaryawanUseCase
}

func NewKaryawanHandler(useCase *usecase.KaryawanUseCase) *KaryawanHandler {
	return &KaryawanHandler{
		karyawanUseCase: useCase,
	}
}

func (h *KaryawanHandler) CreateKaryawan(c *gin.Context) {
	var karyawan entity.Karyawan
	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.karyawanUseCase.CreateKaryawan(&karyawan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Karyawan berhasil dibuat", "karyawan": karyawan})
}

func (h *KaryawanHandler) GetKaryawanByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	karyawan, err := h.karyawanUseCase.GetKaryawanByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, karyawan)
}

func (h *KaryawanHandler) GetAllKaryawan(c *gin.Context) {
	karyawan, err := h.karyawanUseCase.GetAllKaryawan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, karyawan)
}

func (h *KaryawanHandler) UpdateKaryawan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var karyawan entity.Karyawan
	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	karyawan.IDKaryawan = uint(id)

	if err := h.karyawanUseCase.UpdateKaryawan(&karyawan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Karyawan berhasil diperbarui", "karyawan": karyawan})
}

func (h *KaryawanHandler) DeleteKaryawan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := h.karyawanUseCase.DeleteKaryawan(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Karyawan berhasil dihapus"})
}
