#!/bin/bash

# docker network create progra_avanzada
# --network progra_avanzada
docker run -d --name local-mongo \
	-e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
	-e MONGO_INITDB_ROOT_PASSWORD=secret \
    -p 27017:27017 \
	mongo