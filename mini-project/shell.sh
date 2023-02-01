# Menjalankan migrations
go run ./cmd/migrations

# Menjalankan API
go run ./cmd/api

# Melakukan compile project menjadi binary
go build -o main ./cmd/api

# Build image docker
docker build -t api-project:1.0 .

# Membuat container dikarenakan file configurasi dan repository disana maka tidak perlu menggunakan .env ekternal
docker create --name svc-rest-api -p 4002:4002 api-project:1.0
docker run -d --name svc-rest-api -p 4002:4002 api-project:1.0 # Langsung jalan
# Menjalankan container
docker start svc-rest-api

# Melakukan pengecekan container
docker ps || docker logs -f svc-rest-api

# Menghentikan container
docker stop svc-rest-api || docker rm svc-rest-api
