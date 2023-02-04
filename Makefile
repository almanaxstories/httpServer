migrateup:
	migrate -path /home/mike11/Documents/VSCodeProjects/httpServer/db/migrations -database "postgresql://admin:qwert@localhost:2023/usersInfoDB?sslmode=disable"  up

migratedown:
	migrate -path /home/mike11/Documents/VSCodeProjects/httpServer/db/migrations -database "postgresql://admin:qwert@localhost:2023/usersInfoDB?sslmode=disable" down

.PHONY: migrateup migratedown