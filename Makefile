migrateup:
	migrate -path $(CURRENT_DIR)/db/migrations -database "postgresql://admin:qwert@localhost:2023/usersInfoDB?sslmode=disable"  up

migratedown:
	migrate -path $(CURRENT_DIR)/db/migrations -database "postgresql://admin:qwert@localhost:2023/usersInfoDB?sslmode=disable" down

.PHONY: migrateup migratedown