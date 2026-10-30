.PHONY: build run

APP_NAME := server
BUILD_DIR := app/src/
MAIN_FILE := app/src/main.go

build:
	@echo "Building $(APP_NAME)..."
	cd $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
run:
	@echo "Running application..."
	cd $(BUILD_DIR) && ./$(APP_NAME)
