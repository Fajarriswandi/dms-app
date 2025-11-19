# DMS App - Document Management System

Aplikasi Document Management System dengan stack:
- **Frontend**: Vue 3 + TypeScript + Vite
- **Backend**: Go (Golang)
- **CI/CD**: GitHub Actions dengan Docker

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- Node.js 20+ (untuk development frontend)
- Go 1.23+ (untuk development backend)

### Development Setup

#### Option 1: Docker Compose (Recommended untuk pertama kali)

```bash
# Development mode dengan hot reload
docker-compose -f docker-compose.dev.yml up --build

# Atau production mode
docker-compose up --build
```

**Akses:**
- Frontend: http://localhost:5173 (dev) atau http://localhost:3000 (prod)
- Backend API: http://localhost:8080
- Health Check: http://localhost:8080/health

#### Option 2: Local Development (Lebih cepat untuk development)

**Backend:**
```bash
cd backend
go mod download
go run main.go
```

**Frontend:**
```bash
cd frontend
npm install
npm run dev
```

## 📁 Project Structure

```
dms-app/
├── backend/          # Go backend API
│   ├── main.go      # Entry point
│   ├── go.mod       # Go dependencies
│   └── Dockerfile   # Production Docker image
├── frontend/         # Vue 3 frontend
│   ├── src/         # Source code
│   ├── package.json # Node dependencies
│   └── Dockerfile   # Production Docker image
├── .github/
│   └── workflows/   # CI/CD pipelines
└── docker-compose.yml # Local development setup
```

## 🔧 Development Commands

### Backend
```bash
cd backend
go run main.go          # Run server
go test ./...           # Run tests
golangci-lint run       # Lint code
```

### Frontend
```bash
cd frontend
npm run dev             # Development server
npm run build           # Build for production
npm run lint            # Lint code
npm run test:unit       # Run tests
```

## 🐳 Docker Commands

```bash
# Build images
docker-compose build

# Start services
docker-compose up

# Start in background
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f

# Rebuild and restart
docker-compose up --build
```

## 🚢 CI/CD

Pipeline otomatis berjalan saat:
- Push ke branch `main`
- Push tag versi (v1.0.0, v2.1.3, dll)

**Fitur CI/CD:**
- ✅ Lint & Test (Frontend & Backend)
- ✅ Security Scan (Trivy)
- ✅ Build Docker Images
- ✅ Push ke GitHub Container Registry
- ✅ Automatic Version Tagging
- ✅ Generate Changelog
- ✅ Create GitHub Release (saat push tag)

## 📝 Release Process

```bash
# 1. Buat tag versi
git tag v1.0.0
git push origin v1.0.0

# 2. CI/CD akan otomatis:
#    - Build images dengan tag v1.0.0
#    - Generate changelog
#    - Create GitHub release
#    - Push images ke registry
```

## 🔍 Health Check

```bash
# Backend health check
curl http://localhost:8080/health

# Expected response: OK
```

## 📦 Port Configuration

- **Frontend (Dev)**: 5173
- **Frontend (Prod)**: 3000
- **Backend API**: 8080

**Note**: Pastikan port-port ini tidak digunakan oleh aplikasi lain.

## 🛠️ Troubleshooting

### Port sudah digunakan
```bash
# Cek port yang digunakan
lsof -i :8080
lsof -i :5173

# Atau ubah port di docker-compose.yml
```

### Docker build error
```bash
# Clean build
docker-compose down
docker system prune -f
docker-compose up --build
```

### Frontend tidak connect ke backend
- Pastikan `VITE_API_URL` di frontend sesuai dengan backend URL
- Cek CORS settings di backend jika diperlukan

## 📚 Tech Stack

- **Frontend**: Vue 3, TypeScript, Vite, Pinia, Vue Router
- **Backend**: Go 1.23, Standard Library
- **Container**: Docker, Docker Compose
- **CI/CD**: GitHub Actions
- **Security**: Trivy Scanner

## 🤝 Contributing

1. Buat branch dari `main`
2. Develop fitur
3. Test & lint
4. Push dan buat PR
5. Setelah merge, CI/CD akan otomatis build

## 📄 License

[Your License Here]

