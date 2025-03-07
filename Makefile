run:
	go run main.go

run_docker:
	docker stop movie-festival-api || true
	docker rm movie-festival-api || true
	docker build --progress=plain -t movie-festival-api:$(shell git rev-parse --short HEAD) .
	docker run -d --name movie-festival-api -p 8000:8000 movie-festival-api:$(shell git rev-parse --short HEAD)

run_db_migrate_up:
	go run database/migrate/migrate.go -t up

run_db_migrate_down:
	go run database/migrate/migrate.go -t down

run_db_seed_admin:
	go run database/seed/admin/admin.go
