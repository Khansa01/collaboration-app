# Collabify

A real-time collaborative document editor built with Go, gRPC, WebSocket, and React.

🔗 **Live Demo**: [collaboration-app-by-khansa.vercel.app](https://collaboration-app-by-khansa.vercel.app)

---

## Features

- 📝 Real-time collaborative editing (Yjs + WebSocket)
- 👥 Presence indicator — see who's online in a document
- 💾 Auto-save — changes saved automatically after 2 seconds
- 🔐 Authentication with JWT
- 📄 Document CRUD

---

## Tech Stack

| Layer    | Technology                              |
| -------- | --------------------------------------- |
| Frontend | React + Vite, TipTap, Yjs, Tailwind CSS |
| Backend  | Go, gRPC (ConnectRPC), WebSocket        |
| Database | PostgreSQL (Supabase)                   |
| Deploy   | Vercel (FE), Railway (BE)               |

---

## Project Structure

collaboration-app/
├── fe/ # React Vite frontend
└── be/ # Go backend

---

## Getting Started

### Prerequisites

- Node.js 18+
- Go 1.25+
- PostgreSQL (or Supabase account)

### Frontend

```bash
cd fe
npm install
cp .env.example .env
npm run dev
```

`.env`:

VITE_API_URL=http://localhost:8080
VITE_WS_URL=ws://localhost:8080

### Backend

```bash
cd be
cp .env.example .env
go mod download
go run cmd/server/main.go
```

`.env`:

DB_HOST=
DB_PORT=5432
DB_NAME=postgres
DB_USER=
DB_PASSWORD=
JWT_SECRET=

### Docker (Backend)

```bash
cd be
docker build -t collabify-be .
docker run -p 8080:8080 --env-file .env collabify-be
```

---

## Deployment

| Service  | Platform | URL                                                                                      |
| -------- | -------- | ---------------------------------------------------------------------------------------- |
| Frontend | Vercel   | [collaboration-app-by-khansa.vercel.app](https://collaboration-app-by-khansa.vercel.app) |
| Backend  | Railway  | collaboration-app.up.railway.app                                                         |
| Database | Supabase | PostgreSQL                                                                               |

---

## API

- **gRPC (ConnectRPC)** — Auth, Document, Presence services
- **WebSocket** — `/ws/{docId}?token=<jwt>` for real-time collaboration

---

## License

MIT
