createmigration: 
	migrate create -ext=sql -dir=db/migrations -seq init

migrate: 
	migrate -path=services/payment/internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/commerce_payment" -verbose up
	migrate -path=services/user/internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/commerce_user" -verbose up

migratedown:
	migrate -path=services/payment/internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/commerce_payment" -verbose down
	migrate -path=services/user/internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/commerce_user" -verbose down

.PHONY: createmigration migrate migratedown