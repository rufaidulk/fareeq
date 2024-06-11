#!/bin/bash

COMPOSE_FILE="./deploy/docker-compose.yml"

init() {
	echo "Builiding the containers...."
    docker-compose up --build
}

case "$1" in
    init)
		init
		;;
    up)
		docker-compose --progress=plain -f $COMPOSE_FILE up
		;;
    clean)
		docker-compose down -v
		;;
	*)
		echo "Invalid option\n"
		;;
esac
