# Elichika with Docker
The Docker container offers a lightweight and simplistic approach to deploying Elichika across different architectures and operating systems using the `debian:latest` base image.

The latest release docker image can be found [here](https://hub.docker.com/r/arina999999997/elichika).

Docker must be installed, along with Docker Compose to create and deploy the container. More information can be found [here](https://docs.docker.com/engine/install/).

## How to deploy
Navigate to the `docker` directory and run the following:
```
docker compose up -d
```

A container will be generated and expose ports required to accessing the server via `server_address:8080/webui/admin`.

## Updating container
Before proceeding with this, please ensure that `userdata.db` is properly backed up or exported with the WebUI. The docker container can be spun down and rebuilt with a new image:
```
# Copy user data
docker container cp elichika:/elichika/userdata.db .

# Delete existing image
docker compose down
docker rmi elichika:latest
docker compose up -d

# Place user data inside container
docker container cp userdata.db elichika:/elichika

# Restart container with new changes
docker container restart elichika
```

Optionally, the update can be ran in place:
```
docker container exec -it elichika bash /root/update_elichika

# Restart container with new changes
docker container restart elichikas
```

## GitHub Workflow
Upon commits to `main`, a GitHub Workflow is generated to build and push new images to DockerHub. As of this time, `arm64` and `amd64` are the supported architectures.

The image is tested by successfully deploying the container and accessing the WebUI endpoint.
