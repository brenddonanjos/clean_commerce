createmigration: 
	migrate create -ext=sql -dir=db/migrations -seq init

migrate: 
	migrate -path=db/migrations -database "mysql://root:root@tcp(localhost:3306)/commerce_ai" -verbose up

migratedown:
	migrate -path=db/migrations -database "mysql://root:root@tcp(localhost:3306)/commerce_ai" -verbose down

.PHONY: createmigration migrate migratedown