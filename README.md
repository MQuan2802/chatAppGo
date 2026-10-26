go get go.mongodb.org/mongo-driver/v2/mongo

go get ://github.com

docker run --name chat-app \
  -e POSTGRES_PASSWORD=mysecurepassword \
  -p 5432:5432 \
  -d postgres


docker exec -it chat-app psql -U postgres

docker build -t chat-app-go:1.27.2 .
//wont run because the file is injected as an environment variable not a file 
docker run --env-file .env chat-app-go:1.27.2
//mount a volume when running container and inject file 
docker run -v $(pwd)/.env:/.env chat-app-go:1.27.2

# 1. Stop containers and completely wipe the old database volume
docker compose down -v

# 2. Start everything up fresh with your new .env values
docker compose up --build -d
  

