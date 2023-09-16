# Variables
IMAGE_NAME = region
CONTAINER_NAME = region
DOCKERFILE = Dockerfile

# Comandos
build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

docker-regional:
	docker run -p 8080:80 $(CONTAINER_NAME)
stop:
	docker stop $(CONTAINER_NAME)
rm:
	docker rm $(CONTAINER_NAME)
clean:	stop rm
	docker rmi $(IMAGE_NAME)
logs:
	docker logs -f $(CONTAINER_NAME)