curl -X POST -d '{"email":"john.doe@mail.com", "password": "123456"}' http://localhost:8080/v1/auth/login


psql -h 0.0.0.0 -U alextanhongpin -d grpc-openid -W