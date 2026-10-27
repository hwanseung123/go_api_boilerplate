.PHONY: build run

APP_NAME := app
BUILD_DIR := bin
MAIN_FILE := app/src/main.go

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
run:
	@echo "Running application..."
	cd $(BUILD_DIR) && ./$(APP_NAME)


