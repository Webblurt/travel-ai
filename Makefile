dev:
	npm run dev

build:
	docker build -t travel-ai .

run:
	docker run -p 8080:8080 travel-ai

db:
	docker-compose -f docker-compose.db.yml up -d

swagger:
	docker-compose -f docker-compose.swagger.yml up -d

clean:
	docker system prune -f
