# The DEFAULT_GOAL variable specifies the default target that will be built
# when you run the make command without any arguments.
.DEFAULT_GOAL := all

.PHONY: check_pkgx
check_pkgx:
	@echo "$(YELLOW)Checking if pkgx is installed...$(NC)"
	@if ! command -v pkgx > /dev/null; then \
		echo "$(RED)Error: pkgx is not installed. Please install it and try again.$(NC)"; \
		echo "$(RED)Consult the docs here: https://docs.pkgx.sh/$(NC)"; \
		echo "$(BLUE)The easiest way to install is with:$(NC)"; \
		echo "$(BLUE)curl -fsS https://pkgx.sh | sh$(NC)"; \
		exit 1; \
	fi
	@echo "pkgx is installed."

.PHONY: all
all: check_pkgx
	pkgx task lint gen
