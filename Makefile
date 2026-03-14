.PHONY: run test build clean migrate docker-up docker-down

# Default target
run:
	go run main.go

test:
	go test -v ./...

build:
	go build -o coffee-shop main.go

clean:
	rm -f coffee-shop

migrate:
	@echo "Running migrations..."
	# Add your migration tool command here, e.g., golang-migrate
	# migrate -path migrations -database "$(DB_URL)" up

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f
