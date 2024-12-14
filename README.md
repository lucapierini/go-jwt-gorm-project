# go-jwt-gorm-project

~/go/bin/compiledaemon --command="./api" 
docker run --name "postgress-jwt-container" -e POSTGRES_USER=userTest -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB="gorm-jwt-auth" -p 5432:5432 -d postgres