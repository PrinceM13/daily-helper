# 🚪 Daily Helper – Useful Docker & Dev Commands

This guide helps you run and manage the project with Docker.

## 🚀 Start & Stop Containers

| Action               | Command                           |
| -------------------- | --------------------------------- |
| Start (foreground)   | `docker compose up`               |
| Start (background)   | `docker compose up -d`            |
| Rebuild & start      | `docker compose up --build`       |
| Stop all             | `docker compose down`             |
| View logs            | `docker compose logs -f`          |
| Logs (frontend only) | `docker compose logs -f frontend` |
| Restart frontend     | `docker compose restart frontend` |

## 🧰 Frontend Commands (in /frontend)

```bash
pnpm install
pnpm dev      # Run in dev mode
pnpm build    # Build for production
```

## 🔧 Backend Commands (in /backend)

```bash
go mod tidy
go run main.go
```

## 🔑 Environment Variables

Place environment variables in `.env.local` for local development.

## 📁 Structure

```
/backend       # Go API
/frontend      # Next.js frontend
/docker-compose.yml
```

---

Feel free to expand this if your workflows grow.
