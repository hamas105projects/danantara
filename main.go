package main

import (
	"fmt"
	"log"
	"os"

	"danantara/handler"
	"danantara/routes"
	"danantara/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Koneksi ke database gagal: %v", err)
	}

	karyawanUseCase := usecase.NewKaryawanUseCase(db)
	gajiPokokUseCase := usecase.NewGajiPokokUseCase(db)
	kehadiranUseCase := usecase.NewKehadiranUseCase(db)
	statusPembayaranGajiUseCase := usecase.NewStatusPembayaranGajiUseCase(db)

	karyawanHandler := handler.NewKaryawanHandler(karyawanUseCase)
	gajiPokokHandler := handler.NewGajiPokokHandler(gajiPokokUseCase)
	kehadiranHandler := handler.NewKehadiranHandler(kehadiranUseCase)
	statusPembayaranGajiHandler := handler.NewStatusPembayaranGajiHandler(statusPembayaranGajiUseCase)

	router := gin.Default()
	routes.SetupRoutes(router, karyawanHandler, gajiPokokHandler, kehadiranHandler, statusPembayaranGajiHandler)

	log.Printf("Server berjalan di port %s", serverPort)
	if err := router.Run(":" + serverPort); err != nil {
		log.Fatalf("Server gagal dijalankan: %v", err)
	}
}
