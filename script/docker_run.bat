cd ../

CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build -o app .

docker build -t go-app:v1 .