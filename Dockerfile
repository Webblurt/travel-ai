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
COPY --from=frontend-build /app/apps/frontend/build ./frontend
COPY --from=backend-build /app/bin/api ./api
EXPOSE 8080
CMD ["./api"]
