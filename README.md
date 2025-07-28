# 🖼️ Image Optimization API

Un servidor de optimización de imágenes en tiempo real construido con Go que permite redimensionar, comprimir y servir imágenes desde cualquier URL con un sistema de caché inteligente.

## ✨ Características

- 🚀 **Optimización en tiempo real** - Redimensiona y comprime imágenes al vuelo
- 📦 **Sistema de caché inteligente** - Almacena imágenes procesadas para respuestas ultra-rápidas
- 🌐 **Soporte multi-formato** - JPEG, PNG, GIF, WebP
- ⚡ **Alto rendimiento** - Respuestas en microsegundos para imágenes cacheadas
- 🔧 **Configuración flexible** - Parámetros de calidad y tamaño personalizables
- 🛡️ **Headers inteligentes** - Evita bloqueos de CORS y referer

## 🚀 Inicio Rápido

### Instalación

```bash
git clone <repository-url>
cd api-rest-#1
go mod tidy
go run main.go
```

### Configuración

Crea un archivo `.env` en la raíz del proyecto:

```env
TIME_CACHE=1h
CACHE_DIR=./cache/images
MAX_CACHE_SIZE=1000
API_TOKEN=your-secret-api-token-here
PORT=4440
```

## 📖 Uso

### Formato de URL

```
http://localhost:4000/api/image/[parámetros]/[url-de-imagen]
```

### Parámetros Disponibles

- `w_[número]` - Ancho en píxeles (máximo 2000)
- `q_[número]` - Calidad JPEG (1-100)
- `origin=[dominio]` - Dominio de origen para headers

### Ejemplos de Uso

```bash
# Redimensionar a 400px de ancho con calidad 90% (con autenticación)
curl -H "Authorization: Bearer your-api-token" \
  "http://localhost:4000/api/image/w_400,q_90/example.com/image.jpg"

# Solo redimensionar (calidad por defecto 90%)
curl -H "Authorization: Bearer your-api-token" \
  "http://localhost:4000/api/image/w_300/example.com/photo.png"

# Solo cambiar calidad (ancho por defecto 400px)
curl -H "Authorization: Bearer your-api-token" \
  "http://localhost:4000/api/image/q_75/example.com/picture.webp"

# Ruta sin autenticación (compatibilidad)
http://localhost:4000/w_400,q_90/example.com/image.jpg
```

## 🛠️ API Endpoints

| Endpoint | Método | Autenticación | Descripción |
|----------|--------|---------------|-------------|
| `/api/image/*` | GET | ✅ Requerida | Optimiza y sirve imágenes |
| `/api/info?url=[url]` | GET | ✅ Requerida | Obtiene información de una imagen |
| `/api/cache/stats` | GET | ✅ Requerida | Estadísticas del sistema de caché |
| `/api/health` | GET | ✅ Requerida | Estado del servidor |
| `/*` | GET | ❌ No requerida | Optimización sin autenticación (compatibilidad) | Testing Only

### 🔐 Autenticación

Las rutas `/api/*` requieren autenticación mediante token Bearer:

```bash
# Formato del header
Authorization: Bearer your-api-token

# Ejemplo de uso
curl -H "Authorization: Bearer your-api-token" \
  "http://localhost:4000/api/health"
```

### Ejemplo de Respuesta - Info

```json
{
  "width": 1920,
  "height": 1080,
  "format": "jpeg"
}
```

### Ejemplo de Respuesta - Cache Stats

```json
{
  "cache_size": "45.67 MB",
  "cache_dir": "./cache/images"
}
```

## 🎯 Formatos Soportados

### Entrada
- **JPEG** (.jpg, .jpeg)
- **PNG** (.png)
- **GIF** (.gif)
- **WebP** (.webp)

### Salida
- **JPEG optimizado** (mejor rendimiento y compatibilidad)

## ⚙️ Variables de Entorno

| Variable | Descripción | Valor por Defecto |
|----------|-------------|-------------------|
| `TIME_CACHE` | Duración del caché | `1h` |
| `CACHE_DIR` | Directorio de caché | `./cache/images` |
| `MAX_CACHE_SIZE` | Tamaño máximo en MB | `1000` |
| `API_TOKEN` | Token de seguridad para API | *(opcional)* |
| `PORT` | Puerto del servidor | `4441` |

## 🔧 Configuración Avanzada

### Duración del Caché

Formatos soportados para `TIME_CACHE`:
- `30min` - 30 minutos
- `2h` - 2 horas
- `1d` - 1 día
- `1w` - 1 semana
- `1m` - 1 mes
- `1y` - 1 año

### Limpieza Automática

El sistema automáticamente:
- Elimina archivos expirados según `TIME_CACHE`
- Libera espacio cuando se alcanza `MAX_CACHE_SIZE`
- Mantiene los archivos más recientes

## 📊 Rendimiento

- **Primera petición**: ~1-2 segundos (descarga + procesamiento)
- **Peticiones cacheadas**: ~1ms (desde disco)
- **Throughput**: Miles de imágenes por segundo
- **Memoria**: Uso eficiente con streaming

## 🛡️ Características de Seguridad

### 🔐 Autenticación por Token API
- **Protección de endpoints**: Todas las rutas `/api/*` requieren autenticación
- **Token Bearer**: Formato estándar `Authorization: Bearer <token>`
- **Configuración flexible**: Token opcional para desarrollo, obligatorio para producción
- **Respuestas de error claras**: Mensajes específicos para diferentes tipos de errores de autenticación

### 🛡️ Otras Medidas de Seguridad
- Validación de URLs de entrada
- Límites de tamaño de imagen
- Headers CORS configurados
- Timeouts de descarga
- Sanitización de parámetros
- Rutas de compatibilidad sin autenticación para migración gradual

## 📦 Dependencias

- `github.com/go-chi/chi/v5` - Router HTTP
- `github.com/disintegration/imaging` - Procesamiento de imágenes
- `github.com/joho/godotenv` - Variables de entorno
- `golang.org/x/image/webp` - Soporte WebP

## 🤝 Contribuir

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.

## 🚀 Deploy

### Docker

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 4441
CMD ["./main"]
```
---

⭐ **¡Dale una estrella si te gusta el proyecto!**