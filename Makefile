.PHONY: docker-up docker-down docker-build docker-logs docker-ps docker-restart migrate-up

.DEFAULT_GOAL := help

help: 
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  docker-up      Start all services with Docker Compose"
	@echo "  docker-down    Stop and remove all containers, networks, and volumes"
	@echo "  docker-build   Rebuild all microservice Docker images"
	@echo "  docker-logs    Follow logs from all Docker containers"
	@echo "  docker-ps      List running Docker containers"
	@echo "  docker-restart Restart all Docker services"

docker-up: 
	docker compose up -d --build

docker-down: 
	docker compose down -v

docker-build: 
	docker compose build

docker-logs: 
	docker compose logs -f

docker-ps: 
	docker compose ps

docker-restart: 
	docker compose restart
