package handler

import (
	"danantara/entity"
	"danantara/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KehadiranHandler struct {
	kehadiranUseCase *usecase.KehadiranUseCase
}

func NewKehadiranHandler(kehadiranUseCase *usecase.KehadiranUseCase) *KehadiranHandler {
	return &KehadiranHandler{kehadiranUseCase}
}

func (h *KehadiranHandler) CreateKehadiran(c *gin.Context) {
	var kehadiran entity.Kehadiran
	if err := c.ShouldBindJSON(&kehadiran); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.kehadiranUseCase.CreateKehadiran(&kehadiran); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Data kehadiran berhasil ditambahkan", "kehadiran": kehadiran})
}

func (h *KehadiranHandler) GetKehadiranByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	kehadiran, err := h.kehadiranUseCase.GetKehadiranByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data kehadiran tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kehadiran": kehadiran})
}

func (h *KehadiranHandler) GetKehadiranByKaryawanID(c *gin.Context) {
	idKaryawanStr := c.Param("idKaryawan")
	idKaryawan, err := strconv.ParseUint(idKaryawanStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Karyawan tidak valid"})
		return
	}

	kehadiran, err := h.kehadiranUseCase.GetKehadiranByID(uint(idKaryawan))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data kehadiran tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kehadiran": kehadiran})
}

func (h *KehadiranHandler) GetAllKehadiran(c *gin.Context) {
	kehadiran, err := h.kehadiranUseCase.GetAllKehadiran()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kehadiran": kehadiran})
}

func (h *KehadiranHandler) UpdateKehadiran(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var kehadiran entity.Kehadiran
	if err := c.ShouldBindJSON(&kehadiran); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	kehadiran.ID = uint(id)

	if err := h.kehadiranUseCase.UpdateKehadiran(&kehadiran); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data kehadiran berhasil diperbarui", "kehadiran": kehadiran})
}

func (h *KehadiranHandler) DeleteKehadiran(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := h.kehadiranUseCase.DeleteKehadiran(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data kehadiran berhasil dihapus"})
}
