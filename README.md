curl -X POST -d '{"email":"jane.doe@mail.com", "password": "123456"}' http://localhost:8080/v1/auth/login
psql -h 127.0.0.1 -U alextanhongpin -d grpc-openid -W