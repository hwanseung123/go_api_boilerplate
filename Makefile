# if the .env file currently exist  => export .env file => use .go file  
# if the .env file does not currently exist => return ""

ifneq (,&(wildcard .env))
	include .env
	export
endif



#########################################################################3
.PHONY: build run

APP_NAME := server
BUILD_DIR := app/src/
MAIN_FILE := app/src/main.go
ALERT_FILE := app/src/alert.go
PID_FILE := server.pid

build:
	@echo "Building $(APP_NAME)..."
	cd $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE) $(ALERT_FILE)
run:
	@echo "Running application..."
	cd $(BUILD_DIR) && ./$(APP_NAME)

debug:
	@echo "PASSWORD : $(PASSWORD)"

reload:
	@echo "Automatic...build&run"
	if [ -f $(PID_FILE) ]; then \
		kill -9 $$(cat $(PID_FILE)) 2>	/dev/null || true; \
		rm -f $(PID_FILE); \
	fi

	(nohup sh -c "cd $(BUILD_DIR) && ./$(APP_NAME)" > app.log 2>&1 </dev/null & echo $$! > $(PID_FILE))
	pwd

