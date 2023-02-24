To run the app you need to:
1. set a project's folder as a current working directory in your CLI
2. run ```docker-compose up --build``` to build containerized app
3. open a new terminal window, set a project's folder as a current working directory and run ```  docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations -database "postgres://user:qwert@127.0.0.1:5433/usersList?sslmode=disable" up 1 ``` to apply the first DB migration which creates a table from scratch with necessary parameters
4. proceed to ``` http://127.0.0.1:8090/ ``` via your webbrowser and start using the app

Other possible commands to use:

5. running ``` docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations -database "postgres://user:qwert@127.0.0.1:5433/usersList?sslmode=disable" down 1 ``` from the project's directory to migrate DB down. It will clear the table but not delete it.
6. running ``` docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations -database "postgres://user:qwert@127.0.0.1:5433/usersList?sslmode=disable" drop -f ``` to drop the whole DB. 