
docker stop server
docker stop client

docker rm server
docker rm client

docker build -t cda-test .

docker run -dit --name server -e host=server --network cda -e type=server cda-test
docker run -dit --name client -e host=server --network cda cda-test