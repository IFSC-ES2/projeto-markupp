COMPOSE_FILE=docker-compose.yaml
BACKEND_WORKDIR=/src
DB_CONFIG_PKG=./internal/storage
DOCKER_COMPOSE=docker compose

.PHONY: all test-db-config compose-config compose-env docker-up docker-test docker-down run

all: compose-env compose-config docker-up docker-test docker-down

# Testa o módulo de banco de dados do backend dentro de um contêiner Docker
test-db-config:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) run --rm backend sh -c "cd $(BACKEND_WORKDIR) && go test $(DB_CONFIG_PKG)"

# Valida o arquivo docker-compose e a interpolação de variáveis de ambiente
compose-config:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) config

# Configura variáveis de ambiente no ambiente do Compose
compose-env:
	@echo "Criando arquivo .env com valores padrão..."
	@printf 'BACKEND_PORT=8080\nDATA_VOLUME=markupp_data\nGO_MOD_CACHE=go_mod_cache\n' > .env
	@echo ".env criado/atualizado com sucesso."

# Sobe o container Docker do backend em modo destacado
docker-up:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up -d --build backend

# Executa todos os testes Go do backend dentro do container Docker em execução
docker-test:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) run --rm backend sh -c "cd $(BACKEND_WORKDIR) && go test ./..."

# Desce e remove o container Docker usado nos testes
docker-down:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down

# Roda a aplicação completa com air via Docker
run: compose-env compose-config
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up --build backend
