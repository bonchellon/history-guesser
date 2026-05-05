# History Guesser (Tauri 2 + Go) MVP

Production-oriented multiplayer desktop MVP for Steam-like gameplay inspired by GeoGuessr/WenWare.

## Stack
- Client: Tauri 2, Vite, React, TypeScript, Three.js, MapLibre GL, Zustand, Tailwind CSS
- Server: Go, Chi, Gorilla WebSocket, PostgreSQL, Redis-ready
- Data: Goose migrations, sql seed placeholders, sqlc-compatible layout

## Repository structure
- `client/` desktop app
- `server/` backend API + websocket room/match authority
- `shared/` strict protocol typings

## Features implemented
- 360° panorama viewer on inside-facing sphere (`BackSide`) with drag rotate + wheel zoom.
- Clickable MapLibre guess map storing lat/lon.
- BC/AD timeline slider storing selected year.
- WebSocket protocol contracts and reconnect logic.
- Authoritative server room/round/timer flow, anti-cheat range checks, duplicate submit guard.
- Scoring with smooth exponential decay by distance/year delta.
- DB migrations for all requested entities.
- Seed rounds with placeholder panorama CDN URLs.
- Tauri desktop shell ready for Steam wrapping and Windows bundle target.

## Local run
1. Copy env:
   ```bash
   cp .env.example .env
   ```
2. Start infra + server:
   ```bash
   cd server
   docker compose up --build
   ```
3. Migrate + seed:
   ```bash
   make migrate
   make seed
   ```
4. Run client:
   ```bash
   cd client
   npm install
   npm run dev
   ```
5. Run Tauri desktop dev:
   ```bash
   cd client
   npm run tauri dev
   ```

## Steam integration approach
- Keep Steam behind adapter interface (`client/src/features/steam`, `server/internal/steam`) with `dev` fallback identity.
- In dev mode: use synthetic user IDs and names.
- In production: wire Steamworks SDK (Rust side for client + server Web API verification).
- TODO only credentials/App ID/achievements IDs/leaderboard IDs.

## Build
```bash
make build
```

## Notes
- Replace placeholder panorama URLs in `server/sql/seed.sql` with CDN assets.
- Add sqlc generated queries from `server/sql/queries.sql` as next step.
- Add Redis-backed room state replication for multi-instance horizontal scaling.
