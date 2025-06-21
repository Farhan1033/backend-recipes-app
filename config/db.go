package config

import (
	// "backend-recipes/models"
	"fmt"
	"log"
	"os"

	// "strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load file .env jika bukan production
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Gagal memuat file .env, lanjut pakai environment bawaan")
		}
	}

	// === Konfigurasi Database ===
	// host := "localhost"
	// user := "postgres"
	// pass := "postgres"
	// dbname := "recipe_app"
	// port := 5432
	// timezone := "Asia/Jakarta"

	// Konversi port ke string
	// portStr := strconv.Itoa(port)

	// Format DSN PostgreSQL
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
	//	host, portStr, user, pass, dbname, timezone)

	// Gunakan string DSN di bawah ini jika ingin mengubah konfigurasi via environment:
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("TIME_ZONE"),
	)

	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
	// 	host, portStr, user, pass, dbname, timezone,
	// )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	// Auto migrate semua model
	// err = db.AutoMigrate(
	// 	&models.Category{},
	// 	&models.Ingredient{},
	// 	&models.Recipe{},
	// 	&models.RecipeIngredient{},
	// )
	if err != nil {
		log.Fatal("Gagal melakukan auto migrate:", err)
	}

	DB = db
	fmt.Println("✅ Database terhubung dengan PostgreSQL!")
}
