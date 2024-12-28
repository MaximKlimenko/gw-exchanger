build:
	go build -o main ./cmd

run:
	./main

docker-build:
	docker build -t gw-exchanger .

docker-run:
	docker run --rm -p 50051:50051 --env-file=config.env gw-exchanger