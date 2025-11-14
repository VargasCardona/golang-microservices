package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guitar-service/crud"
	"github.com/guitar-service/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=db user=postgres password=secret dbname=guitar_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// Auto-migrate (dev only)
	if err := db.AutoMigrate(&crud.GuitarModel{}); err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Table 'guitar_models' ready")

	repo := crud.NewGuitarRepository(db)
	svc := services.NewGuitarService(repo)
	handler := services.NewHandler(svc)

	r := chi.NewRouter()

	// === CRUD ROUTES ===
	r.Post("/guitars", handler.CreateGuitar)
	r.Get("/guitars/{id}", handler.GetGuitar)
	r.Get("/guitars", handler.ListGuitars)        // ← LIST
	r.Put("/guitars/{id}", handler.UpdateGuitar) // ← UPDATE
	r.Delete("/guitars/{id}", handler.DeleteGuitar) // ← DELETE

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
