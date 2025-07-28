package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/loxzer01/serve-img-optimized/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Advertencia: No se pudo cargar el archivo .env. Se usarán las variables de entorno del sistema si existen.")
	}

	r := routes.NewRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "4441"
	}
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, r)

}
