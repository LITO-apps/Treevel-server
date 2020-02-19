build-image:
	docker build . -t treevel-server

run:
	docker run -p 8080:8080 -d treevel-server:latest
