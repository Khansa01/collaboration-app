# Collaboration App

A real-time document collaboration tool built with modern technologies.

## Tech Stack

**Frontend**

- React + Vite
- TailwindCSS
- Connect-Web (gRPC)
- Yjs (CRDT for real-time sync)

**Backend**

- Golang
- gRPC (Connect-Go)
- PostgreSQL (Supabase)

**Infrastructure**

- Railway (Backend)
- Vercel (Frontend)
- Supabase (Database)

## Features

- 🔐 Authentication (Register & Login)
- 📄 Create & manage documents
- ✏️ Real-time collaborative editing
- 👥 Presence indicator (see who's online)
- 🖱️ Live cursor tracking

## Architecture

FE (React Vite) → Connect-Web → BE (Golang + gRPC) → PostgreSQL

## Getting Started

### Backend

```bash
cd be
cp .env.example .env
go run cmd/server/main.go
```

### Frontend

```bash
cd fe
cp .env.example .env
npm install
npm run dev
```

## Live Demo

- Frontend: Coming soon
- Backend: https://collaboration-app.up.railway.app
