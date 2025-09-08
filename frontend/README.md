# 🛠 Daily Helper

A full-stack project using Next.js (frontend) and Go (backend), containerized with Docker and orchestrated using Docker Compose.

## 🚀 Getting Started

### 1. Prerequisites

- Docker & Docker Compose installed
- (Optional) Node.js + pnpm installed locally for development outside of Docker
- (Optional) Go installed for backend local dev

### 2. Clone the repository

```bash
git clone https://github.com/PrinceM13/daily-helper.git
cd daily-helper
```

### 3. Start all services (frontend & backend)

> This spins up both the Go backend and the Next.js frontend using Docker Compose.

```bash
docker compose up --build
```

- Frontend: [http://localhost:3000](http://localhost:3000)
- Backend: [http://localhost:8080](http://localhost:8080)
- Health check: [http://localhost:8080/health](http://localhost:8080/health)

### 4. Environment Variables

- Create a .env file for Docker Compose:

  ```
  FRONTEND_PORT=3000
  BACKEND_PORT=8080
  ```

- Create a frontend/.env.local (used by Next.js):

  ```env
  NEXT_PUBLIC_API_URL=http://localhost:8080
  ```

> You can edit these based on your environment (e.g. dev/staging).

---

## 💂 Folder Structure

```bash
daily-helper/
│
├── backend/          # Go backend
│   └── main.go
│
├── frontend/         # Next.js frontend
│   ├── app/
│   ├── public/
│   └── ...
│
├── docker-compose.yml
└── README.md
```

---

## 📦 Useful Commands

### Frontend (Next.js)

```bash
# Run locally (without Docker)
pnpm dev

# Build for production
pnpm build
```

### Backend (Go)

```bash
# Run locally (without Docker)
go run main.go
```

---

## 🧪 Health Check

```http
GET http://localhost:8080/health
Response: OK
```

---

## 🧰 Troubleshooting

- ❗ If you see Turbopack errors:
- Try deleting `.next` and restarting:

  ```bash
  rm -rf .next
  docker compose down
  docker compose up --build
  ```

- Make sure you’re using the correct Node version (e.g. via nvm):

  ```bash
  nvm use 20
  ```

---

## 🙋‍♀️ For New Contributors

Welcome! 🎉 To get started:

1. Install Docker and clone this repo
2. Run docker compose up
3. Start contributing from either frontend or backend
4. Feel free to ask questions or suggest improvements!

---
