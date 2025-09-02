package routes

import (
	"danantara/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine,
	karyawanHandler *handler.KaryawanHandler,
	gajiPokokHandler *handler.GajiPokokHandler,
	kehadiranHandler *handler.KehadiranHandler,
	statusPembayaranGajiHandler *handler.StatusPembayaranGajiHandler) {

	api := router.Group("/api")
	{
		api.POST("/karyawan", karyawanHandler.CreateKaryawan)
		api.GET("/karyawan/:id", karyawanHandler.GetKaryawanByID)
		api.GET("/karyawan", karyawanHandler.GetAllKaryawan)
		api.PUT("/karyawan/:id", karyawanHandler.UpdateKaryawan)
		api.DELETE("/karyawan/:id", karyawanHandler.DeleteKaryawan)

		api.POST("/gajipokok", gajiPokokHandler.CreateGajiPokok)
		api.GET("/gajipokok/:id", gajiPokokHandler.GetGajiPokokByID)
		api.GET("/gajipokok/jabatan/:jabatan", gajiPokokHandler.GetGajiPokokByJabatan)
		api.GET("/gajipokok", gajiPokokHandler.GetAllGajiPokok)
		api.PUT("/gajipokok/:id", gajiPokokHandler.UpdateGajiPokok)
		api.DELETE("/gajipokok/:id", gajiPokokHandler.DeleteGajiPokok)

		api.POST("/kehadiran", kehadiranHandler.CreateKehadiran)
		api.GET("/kehadiran/:id", kehadiranHandler.GetKehadiranByID)
		api.GET("/kehadiran/karyawan/:idKaryawan", kehadiranHandler.GetKehadiranByKaryawanID)
		api.GET("/kehadiran", kehadiranHandler.GetAllKehadiran)
		api.PUT("/kehadiran/:id", kehadiranHandler.UpdateKehadiran)
		api.DELETE("/kehadiran/:id", kehadiranHandler.DeleteKehadiran)

		api.POST("/statuspembayaranganji", statusPembayaranGajiHandler.CreateStatusPembayaranGaji)
		api.GET("/statuspembayaranganji/:id", statusPembayaranGajiHandler.GetStatusPembayaranGajiByID)
		api.GET("/statuspembayaranganji/karyawan/:idKaryawan", statusPembayaranGajiHandler.GetStatusPembayaranGajiByKaryawanID)
		api.GET("/statuspembayaranganji", statusPembayaranGajiHandler.GetAllStatusPembayaranGaji)
		api.PUT("/statuspembayaranganji/:id", statusPembayaranGajiHandler.UpdateStatusPembayaranGaji)
		api.DELETE("/statuspembayaranganji/:id", statusPembayaranGajiHandler.DeleteStatusPembayaranGaji)
	}
}
