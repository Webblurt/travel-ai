dev:
	npm run dev

build:
	docker build -t travel-ai .

run:
	docker run -p 8080:8080 travel-ai

compose:
	docker-compose up --build

clean:
	docker system prune -f
