package ordentperpustakaan

import (
	"log"
	"ordentperpustakaan/config"
	"ordentperpustakaan/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getenvWithDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	if _, exists := os.LookupEnv("POSTGRES_HOST"); !exists {
		if err := godotenv.Load(".env"); err != nil {
			log.Printf("no .env found (%v) — expecting env vars to be set", err)
		} else {
			log.Println("Loaded .env")
		}
	} else {
		log.Println("Running with environment variables (docker)")
	}

	if err := config.InitPostgres(); err != nil {
		log.Fatalf("Postgres init failed: %v", err)
	}
	log.Println("Postgres connected")

	if err := config.ConnectRedis(); err != nil {
		log.Fatalf("Redis init failed: %v", err)
	}
	log.Println("Redis connected")

	if err := config.AutoMigrate(); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	log.Println("AutoMigrate finished")

	r := gin.Default()

	if _, err := os.Stat("./docs/swagger.yaml"); err == nil {
		r.StaticFile("/swagger.yaml", "./docs/swagger.yaml")
	}

	routes.InitRoutes(r)

	port := getenvWithDefault("PORT", "8080")
	log.Printf("Server running on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server EXIT: %v", err)
	}
}
