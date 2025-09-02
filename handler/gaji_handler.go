package handler

import (
	"danantara/entity"
	"danantara/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GajiPokokHandler struct {
	gajiPokokUseCase *usecase.GajiPokokUseCase
}

func NewGajiPokokHandler(gajiPokokUseCase *usecase.GajiPokokUseCase) *GajiPokokHandler {
	return &GajiPokokHandler{gajiPokokUseCase}
}

func (h *GajiPokokHandler) CreateGajiPokok(c *gin.Context) {
	var gaji entity.GajiPokok
	if err := c.ShouldBindJSON(&gaji); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.gajiPokokUseCase.CreateGajiPokok(&gaji); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Gaji pokok berhasil ditambahkan", "gaji": gaji})
}

func (h *GajiPokokHandler) GetGajiPokokByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	gaji, err := h.gajiPokokUseCase.GetGajiPokokByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data gaji pokok tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"gaji": gaji})
}

func (h *GajiPokokHandler) GetGajiPokokByJabatan(c *gin.Context) {
	jabatan := c.Param("jabatan")
	gaji, err := h.gajiPokokUseCase.GetGajiPokokByJabatan(jabatan)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"gaji": gaji})
}

func (h *GajiPokokHandler) GetAllGajiPokok(c *gin.Context) {
	gaji, err := h.gajiPokokUseCase.GetAllGajiPokok()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"gaji": gaji})
}

func (h *GajiPokokHandler) UpdateGajiPokok(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var gaji entity.GajiPokok
	if err := c.ShouldBindJSON(&gaji); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gaji.ID = uint(id)

	if err := h.gajiPokokUseCase.UpdateGajiPokok(&gaji); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data gaji pokok berhasil diperbarui", "gaji": gaji})
}

func (h *GajiPokokHandler) DeleteGajiPokok(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := h.gajiPokokUseCase.DeleteGajiPokok(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data gaji pokok berhasil dihapus"})
}
