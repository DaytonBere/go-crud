start:
	# Uses CompileDaemon to start the server
	CompileDaemon -command="./go-crud"

migrate-up:
	# Runs the migrate.go file to auto migrate all the tables
	go run migrate/migrate.go