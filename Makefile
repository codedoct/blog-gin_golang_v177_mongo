run_local:
	go run main.go

seed_data:
	mongoimport --host localhost:27017 --db codedoct_gin_golang177 --collection users --file db/seed_user.json --jsonArray

run_docker:
	docker stop basecodeapiserver || true && docker rm basecodeapiserver || true
	docker build --tag basecode-api:dev .
	docker run --name basecodeapiserver -d -p 4001:4001 basecode-api:dev
