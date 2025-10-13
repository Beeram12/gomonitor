# Makefile for go monitor project

#---- Variables ----
binary_name=gomonitor
package_path=./cmd/gomonitor
build_dir=./bin

# ---- Help ----
.DEFAULT_GOAL := help

.PHONY:	help
help:
	@echo "Available Commands:"
	@echo "  make build      - Compiles the application into the $(BUILD_DIR) directory."
	@echo "  make run        - Builds and runs the application"
	@echo "  make clean      - Removes build artifacts"

# ---- Build Targets ----
.PHONY:	build
build:
	@mkdir -p $(build_dir)
	@echo "Building $(binary_name)..."
	@go build -o $(build_dir)/$(binary_name) $(package_path)
	@echo "$(binary_name) built sucessfully"

.PHONY:	run
run: build
	@echo "Starting $(binary_name)"
	@./$(build_dir)/${binary_name} start

.PHONY:	clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(build_dir)
	@rm -f config.yaml


