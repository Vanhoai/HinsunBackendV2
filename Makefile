.PHONY: help build build-amd64 run stop logs clean restart shell check-arch

# Variables
IMAGE_NAME := vanhoaiadv/hinsun-backend
CONTAINER_NAME := hinsun-backend
VERSION := 1.0.0
PLATFORM := linux/amd64

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build Docker image for amd64
	@echo "🔨 Building Docker image for amd64..."
	docker build --platform $(PLATFORM) -t $(IMAGE_NAME):$(VERSION) .
	@echo "✅ Build completed!"

build-multi: ## Build for multiple platforms using buildx
	@echo "🔨 Building for multiple platforms..."
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		-t $(IMAGE_NAME):$(VERSION) \
		--load \
		.

check-arch: ## Check image architecture
	@echo "📋 Image architecture info:"
	@docker image inspect $(IMAGE_NAME):$(VERSION) --format='Architecture: {{.Architecture}}'
	@docker image inspect $(IMAGE_NAME):$(VERSION) --format='OS: {{.Os}}'
	@docker image inspect $(IMAGE_NAME):$(VERSION) --format='Size: {{.Size}} bytes'

run: ## Run container with .env file
	@echo "🚀 Starting container..."
	docker run -d \
		-p 8080:8080 \
		--env-file .env \
		--name $(CONTAINER_NAME) \
		--restart unless-stopped \
		$(IMAGE_NAME):$(VERSION)
	@echo "✅ Container started!"
	@echo "📚 Swagger: http://localhost:8080/swagger/index.html"

stop: ## Stop container
	@echo "🛑 Stopping container..."
	docker stop $(CONTAINER_NAME) || true

logs: ## Show container logs
	docker logs -f $(CONTAINER_NAME)

clean: ## Remove container and image
	@echo "🧹 Cleaning up..."
	docker stop $(CONTAINER_NAME) 2>/dev/null || true
	docker rm $(CONTAINER_NAME) 2>/dev/null || true
	docker rmi $(IMAGE_NAME):$(VERSION) 2>/dev/null || true
	@echo "✅ Cleanup completed!"

restart: stop run ## Restart container

shell: ## Open shell in container
	docker exec -it $(CONTAINER_NAME) sh

rebuild: clean build run ## Rebuild and run container
