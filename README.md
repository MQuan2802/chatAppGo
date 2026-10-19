go get go.mongodb.org/mongo-driver/v2/mongo

go get ://github.com

docker run --name chat-app \
  -e POSTGRES_PASSWORD=mysecurepassword \
  -p 5432:5432 \
  -d postgres


docker exec -it chat-app psql -U postgres

docker build -t ChatAppGo:1.27.2 .

