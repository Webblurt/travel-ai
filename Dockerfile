# FRONTEND BUILD
FROM node:20 AS frontend-build
WORKDIR /app
COPY apps/frontend ./apps/frontend
WORKDIR /app/apps/frontend
RUN npm install && npm run build

# BACKEND BUILD 
FROM golang:1.25 AS backend-build
WORKDIR /app
COPY apps/backend ./apps/backend
WORKDIR /app/apps/backend
RUN go mod tidy && go build -o /app/bin/travel-ai ./cmd/travel-ai

# FRONTEND RUNTIME 
FROM node:20-slim AS frontend
WORKDIR /app
COPY --from=frontend-build /app/apps/frontend/build ./build
RUN npm install -g serve
EXPOSE 3000
CMD ["serve", "-s", "build", "-l", "3000"]

# BACKEND RUNTIME
FROM debian:bookworm-slim AS backend
WORKDIR /app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=backend-build /app/bin/travel-ai ./travel-ai
COPY apps/backend/config ./config
COPY apps/backend/migrations ./migrations
COPY .env .env
EXPOSE 8080
CMD ["./travel-ai"]
