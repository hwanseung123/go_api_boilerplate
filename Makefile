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
	@# cat => it can be read pid file
	@if [ -f $(PID_FILE) ]; then \
		PID=$$(cat $(PID_FILE)); \
		if ps -p $$PID > /dev/null 2>&1; then \
			kill -9 $$PID; \
			sleep 1; \
		fi; \
	fi
	@rm -f $(PID_FILE)
	@PID_PORT=$$(lsof -ti :8080); \
	echo "$$PID_PORT"

	@if lsof -i :8080 >/dev/null 2>&1; then \
		echo "Port 8080 kill..."; \
		fuser -k 8080/tcp 2>/dev/null 2>&1; \
		sleep 1; \
	fi

	@# go build & run
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE) $(ALERT_FILE)

	@# Separte => standard output , session signal , turmainl (tuermain parent process , go server child process (standart output, session sginal)
	@#  (nohub ~) => Bundle them and execute (./server , app.log , pid_file)
	(nohup sh -c "cd $(BUILD_DIR) && stdbuf -oL ./$(APP_NAME)" > app.log 2>&1 </dev/null & echo $$! > $(PID_FILE))

