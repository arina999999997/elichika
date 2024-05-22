# Elichika with Docker
The Docker container offers a lightweight and simplistic approach to deploying Elichika across different architectures and operating systems using the `golang:alpine` base image.

Docker must be installed, along with Docker Compose to create and deploy the container. More information can be found [here](https://docs.docker.com/engine/install/).

## How to deploy
Navigate to the `docker` directory and run the following:
```
docker compose build
docker compose up -d
```

Additionally, the server can be deployed on a different GitHub branch:
```
docker compose up -d -e BRANCH=<branch_name>
```

A container will be generated and expose ports required to accessing the server via `server_address:8080/webui/admin`.

## Updating container
Before proceeding with this, please ensure that `userdata.db` is properly backed up or exported with the WebUI. The docker container can be spun down and rebuilt with a new image:
```
# Copy user data
docker container cp llas:/elichika/userdata.db .

# Delete existing image
docker compose down
docker rmi llas:latest
docker compose build
docker compose up -d

# Place user data inside container
docker container cp userdata.db llas:/elichika
```

Optionally, the update can be ran in place:
```
docker container exec -d llas bash /elichika/update_elichika
```

