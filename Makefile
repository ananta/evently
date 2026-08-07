# ==========================
# Configuration
# ==========================
DB_URL=postgres://evently_user:evently_password@db:5432/evently_db?sslmode=disable
MIGRATIONS=./migrations

# ==========================
# Docker
# ==========================
up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

# ==========================
# Migration
# ==========================
migrate-up:
	docker compose run --rm migration \
		-path=/migrations \
		-database="$(DB_URL)" \
		up

migrate-down:
	docker compose run --rm migration \
		-path=/migrations \
		-database="$(DB_URL)" \
		down 1

migrate-force:
	docker compose run --rm migration \
		-path=/migrations \
		-database="$(DB_URL)" \
		force $(VERSION)

migrate-version:
	docker compose run --rm migration \
		-path=/migrations \
		-database="$(DB_URL)" \
		version

create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make create-migration name=create_users"; \
		exit 1; \
	fi
	docker run --rm \
		-v $(PWD)/migrations:/migrations \
		migrate/migrate \
		create -ext sql -dir /migrations -seq $(name)
