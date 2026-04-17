cd ../

go build -o app .

docker build -t go-app:v1 .