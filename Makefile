config-path=./config/local.yaml # enter your config file path
storage-path=./storage/auth.db 
migrations-path=./migrations

run:
	go run ./cmd/auth-service/main.go --config-path=${config-path}

migrate:
	go run ./cmd/migrator --storage-path=${storage-path} --migrations-path=${migrations-path}
	
	


