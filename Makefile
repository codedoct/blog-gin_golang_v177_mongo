run_local:
	go run main.go

seed_data:
	mongoimport --host localhost:27017 --db codedoct_gin_golang177 --collection users --file db/seed_user.json --jsonArray
