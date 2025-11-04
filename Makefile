# --- Configuration ---
BACKEND_IMAGE=travel-ai-backend
FRONTEND_IMAGE=travel-ai-frontend
BACKEND_PORT=8080
FRONTEND_PORT=3000

# --- Backend ---
backend:
	@if [ -z "$$(docker images -q $(BACKEND_IMAGE))" ]; then \
		echo "Image $(BACKEND_IMAGE) not found, building..."; \
		docker build -t $(BACKEND_IMAGE) --target backend .; \
	else \
		echo "Image $(BACKEND_IMAGE) already exists, skipping build."; \
	fi
	@echo "Starting backend container..."
	@docker run -d --restart=unless-stopped -p $(BACKEND_PORT):8080 --name $(BACKEND_IMAGE) $(BACKEND_IMAGE) || true

# --- Frontend ---
frontend:
	@if [ -z "$$(docker images -q $(FRONTEND_IMAGE))" ]; then \
		echo "Image $(FRONTEND_IMAGE) not found, building..."; \
		docker build -t $(FRONTEND_IMAGE) --target frontend .; \
	else \
		echo "Image $(FRONTEND_IMAGE) already exists, skipping build."; \
	fi
	@echo "Starting frontend container..."
	@docker run -d --restart=unless-stopped -p $(FRONTEND_PORT):3000 --name $(FRONTEND_IMAGE) $(FRONTEND_IMAGE) || true

# --- Compose ---
start:
	docker-compose start

up:
	docker-compose up --build

logs:
	docker-compose logs -f

db:
	docker-compose -f docker-compose.db.yml up -d

swagger:
	docker-compose -f docker-compose.swagger.yml up -d

# --- Utility ---
ps:
	@docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

logs-backend:
	@docker logs -f $(BACKEND_IMAGE)

logs-frontend:
	@docker logs -f $(FRONTEND_IMAGE)

restart:
	@docker restart $(BACKEND_IMAGE) $(FRONTEND_IMAGE)

stop:
	@docker stop $(BACKEND_IMAGE) $(FRONTEND_IMAGE) || true
	@docker rm $(BACKEND_IMAGE) $(FRONTEND_IMAGE) || true

build:
	docker build -t $(BACKEND_IMAGE) --target backend .
	docker build -t $(FRONTEND_IMAGE) --target frontend .

rebuild: stop build

test-backend:
	docker run --rm $(BACKEND_IMAGE) go test ./...

clean:
	docker system prune -f

nuke:
	@echo "Removing all containers, images, and cache..."
	-docker stop $(BACKEND_IMAGE) $(FRONTEND_IMAGE)
	-docker rm $(BACKEND_IMAGE) $(FRONTEND_IMAGE)
	-docker rmi $(BACKEND_IMAGE) $(FRONTEND_IMAGE)
	docker system prune -af

help:
	@echo ""
	@echo "Available targets:"
	@echo "  make backend        - build and run backend container"
	@echo "  make frontend       - build and run frontend container"
	@echo "  make start/up       - start all services with docker-compose"
	@echo "  make ps             - list running containers"
	@echo "  make logs           - docker-compose logs"
	@echo "  make logs-backend   - show backend container logs"
	@echo "  make logs-frontend  - show frontend container logs"
	@echo "  make restart        - restart backend and frontend containers"
	@echo "  make stop           - stop and remove backend and frontend containers"
	@echo "  make rebuild        - rebuild both images"
	@echo "  make test-backend   - run Go tests inside backend image"
	@echo "  make clean          - prune unused Docker resources"
	@echo "  make nuke           - remove all containers, images, and cache"
	@echo "  make help           - show this help message"
	@echo ""
