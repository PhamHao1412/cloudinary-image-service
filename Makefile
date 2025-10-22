# ---------------------------
# ⚙️ CONFIG
# ---------------------------
APP_NAME = image-service
DOCKER_COMPOSE = docker-compose
PORT = 8085

# ---------------------------
# 🏗️ BUILD
# ---------------------------
build:
	@echo "🔨 Building Docker image..."
	docker build -t $(APP_NAME):latest .

compose-build:
	@echo "🔨 Building Docker Compose services..."
	$(DOCKER_COMPOSE) build

# ---------------------------
# ▶️ RUN
# ---------------------------
run:
	@echo "🚀 Running Docker container..."
	docker run --name $(APP_NAME) -p $(PORT):$(PORT) --env-file .env $(APP_NAME):latest

compose-up:
	@echo "🚀 Starting with Docker Compose..."
	$(DOCKER_COMPOSE) up -d

# ---------------------------
# ⏹️ STOP / CLEANUP
# ---------------------------
stop:
	@echo "🛑 Stopping container..."
	-docker stop $(APP_NAME)
	-docker rm $(APP_NAME)

compose-down:
	@echo "🧹 Stopping Docker Compose..."
	$(DOCKER_COMPOSE) down

# ---------------------------
# 🧼 CLEAN IMAGES
# ---------------------------
clean:
	@echo "🧽 Removing image..."
	-docker rmi $(APP_NAME):latest

# ---------------------------
# 🧩 DEV HELPER
# ---------------------------
logs:
	@echo "📜 Showing logs..."
	docker logs -f $(APP_NAME)

ps:
	@echo "📦 Containers running:"
	docker ps --filter "name=$(APP_NAME)"
