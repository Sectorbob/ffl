
compile:
	@GOOS=linux GOARCH=amd64 go build -o target/service

containerize: compile
	@docker build -t ffl-data-service:latest .

run: containerize
	@docker run -it --rm ffl-data-service:latest

cleanup-images:
	#docker rmi $$(docker images $| grep "^<none>" $| awk "{print $$3}")