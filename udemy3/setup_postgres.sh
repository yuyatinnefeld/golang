# create a server
export IMG_NAME=postgres:14-alpine
export CONTAINER_NAME=postgres-container
export POSTGRES_USER=root
export POSTGRES_PWD=secret

docker pull $IMAGE_NAME
docker images
docker run --name $CONTAINER_NAME -p 5432:5432 -e POSTGRES_USER=$POSTGRES_USER -e POSTGRES_PASSWORD=$POSTGRES_PWD -d postgres:14-alpine
docker ps

# create connection
docker exec -it $CONTAINER_NAME psql -U $POSTGRES_USER
docker logs $CONTAINER_NAME