# FRONTEND 
FROM node:20 AS frontend-build
WORKDIR /app
COPY apps/frontend ./apps/frontend
WORKDIR /app/apps/frontend
RUN npm install && npm run build

# BACKEND 
FROM golang:1.25 AS backend-build
WORKDIR /app
COPY apps/backend ./apps/backend
WORKDIR /app/apps/backend
RUN go mod tidy && go build -o /app/bin/api ./cmd/api

# FINAL IMAGE 
FROM debian:bookworm-slim
WORKDIR /app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=frontend-build /app/apps/frontend/build ./frontend
COPY --from=backend-build /app/bin/api ./api
COPY .env .env  
EXPOSE 8080
CMD ["./api"]
