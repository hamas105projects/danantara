package handler

import (
	"danantara/entity"
	"danantara/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatusPembayaranGajiHandler struct {
	statusPembayaranGajiUseCase *usecase.StatusPembayaranGajiUseCase
}

func NewStatusPembayaranGajiHandler(statusPembayaranGajiUseCase *usecase.StatusPembayaranGajiUseCase) *StatusPembayaranGajiHandler {
	return &StatusPembayaranGajiHandler{statusPembayaranGajiUseCase}
}

func (h *StatusPembayaranGajiHandler) CreateStatusPembayaranGaji(c *gin.Context) {
	var spg entity.StatusPembayaranGaji
	if err := c.ShouldBindJSON(&spg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.statusPembayaranGajiUseCase.CreateStatusPembayaranGaji(&spg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Status pembayaran gaji berhasil ditambahkan", "status_pembayaran_gaji": spg})
}

func (h *StatusPembayaranGajiHandler) GetStatusPembayaranGajiByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	spg, err := h.statusPembayaranGajiUseCase.GetStatusPembayaranGajiByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data status pembayaran gaji tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_pembayaran_gaji": spg})
}

func (h *StatusPembayaranGajiHandler) GetStatusPembayaranGajiByKaryawanID(c *gin.Context) {
	idKaryawanStr := c.Param("idKaryawan")
	idKaryawan, err := strconv.ParseUint(idKaryawanStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Karyawan tidak valid"})
		return
	}

	spg, err := h.statusPembayaranGajiUseCase.GetStatusPembayaranGajiByKaryawanID(uint(idKaryawan))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data status pembayaran gaji tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_pembayaran_gaji": spg})
}

func (h *StatusPembayaranGajiHandler) GetAllStatusPembayaranGaji(c *gin.Context) {
	spg, err := h.statusPembayaranGajiUseCase.GetAllStatusPembayaranGaji()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_pembayaran_gaji": spg})
}

func (h *StatusPembayaranGajiHandler) UpdateStatusPembayaranGaji(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var spg entity.StatusPembayaranGaji
	if err := c.ShouldBindJSON(&spg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	spg.ID = uint(id)

	if err := h.statusPembayaranGajiUseCase.UpdateStatusPembayaranGaji(&spg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data status pembayaran gaji berhasil diperbarui", "status_pembayaran_gaji": spg})
}

func (h *StatusPembayaranGajiHandler) DeleteStatusPembayaranGaji(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := h.statusPembayaranGajiUseCase.DeleteStatusPembayaranGaji(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data status pembayaran gaji berhasil dihapus"})
}
