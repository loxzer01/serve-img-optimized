package routes

import (
	"fmt"
	"net/http"

	"github.com/loxzer01/serve-img-optimized/images"
)

// OptimizeImageHandler maneja las peticiones de optimización de imágenes
func OptimizeImageHandler(optimizer *images.ImageOptimizer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Procesar imagen usando el optimizador
		processedData, contentType, err := optimizer.OptimizeImage(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error processing image: %v", err), http.StatusInternalServerError)
			return
		}

		// Configurar headers de respuesta
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Escribir la imagen procesada
		w.Write(processedData)
	}
}

// ImageInfoHandler obtiene información de una imagen
func ImageInfoHandler(optimizer *images.ImageOptimizer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageURL := r.URL.Query().Get("url")
		if imageURL == "" {
			http.Error(w, "URL parameter is required", http.StatusBadRequest)
			return
		}

		info, err := optimizer.GetImageInfo(imageURL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting image info: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"width":%v,"height":%v,"format":"%v"}`,
			info["width"], info["height"], info["format"])
	}
}

// CacheStatsHandler obtiene estadísticas del caché
func CacheStatsHandler(optimizer *images.ImageOptimizer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats := optimizer.GetCacheStats()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"cache_size":"%v","cache_dir":"%v"}`,
			stats["cache_size"], stats["cache_dir"])
	}
}
