package main

import (
	"RestfulAPI/config"
	"RestfulAPI/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Inisialisasi koneksi database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to MySQL database successfully!")

	// Definisikan port yang ingin digunakan
	// port := os.Getenv("PORT")
	// if port == "" {
	// port = "8080" // Gunakan port 8080 jika PORT tidak didefinisikan
	// }

	http.HandleFunc("/get/user", routes.UserHandler)

	fmt.Printf("Server is running on port %s...\n", 8080)
	// Jalankan server HTTP
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// logMiddleware adalah fungsi penengah yang menambahkan log sebelum menangani permintaan
