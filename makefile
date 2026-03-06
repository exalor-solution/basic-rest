build:
	docker build -t rest-basic .
run:
	docker run --rm -p 8080:8080 rest-basic