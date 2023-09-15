# Variables
IMAGE_NAME = go-containerized:latest
CONTAINER_NAME = go-containerized:latest
DOCKERFILE = Dockerfile

# Comandos
build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

run:
	docker run -t $(CONTAINER_NAME)
stop:
	docker stop $(CONTAINER_NAME)
rm:
	docker rm $(CONTAINER_NAME)
clean:	stop rm
	docker rmi $(IMAGE_NAME)
logs:
	docker logs -f $(CONTAINER_NAME)