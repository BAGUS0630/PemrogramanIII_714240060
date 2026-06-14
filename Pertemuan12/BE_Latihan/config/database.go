package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	// Coba load dari folder saat ini
	err := godotenv.Load(".env")
	if err != nil {
		// Jika gagal (saat testing), coba load dari satu folder di atasnya
		err = godotenv.Load("../.env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("SUPABASE_DSN tidak ditemukan di dalam file .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database : %v", err)
	}

	DB = db
	log.Println("koneksi ke postgreSQL BERHASIL")

}
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("DB belum diinisialisasi")
	}
	return DB
}
