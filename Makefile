# include .env file get out 
ifneq (,&(wildcard .env))
	include .env
	export
endif

.PHONY: build run
APP_NAME := server
BUILD_DIR := app/src/
MAIN_FILE := app/src/main.go
ALERT_FILE := app/src/alert.go

build:
	@echo "Building $(APP_NAME)..."
	cd $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE) $(ALERT_FILE)
run:
	@echo "Running application..."
	cd $(BUILD_DIR) && ./$(APP_NAME)

debug:
	@echo "PASSWORD : $(PASSWORD)"
