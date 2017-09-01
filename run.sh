docker-compose up -d & \
go run gateway/main.go & \ 
go run client/main.go & \
go run server/main.go &