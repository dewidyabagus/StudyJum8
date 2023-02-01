# Set the baseImage to use for subsequent instructions.
# FROM must be the first instruction in a Dockerfile
FROM golang:1.18.3-alpine3.16

# Set the working directory for any subsequent ADD, COPY, CMD, ENTRYPOINT, or RUN 
# instructions that follow it in the Dockerfile.
WORKDIR /app

# Copy files or folders from source to the dest path in the image's filesystem.
# Hanya menyalin modules untuk kebutuhan membuat layer download modules
# ketika tidak ada perubahan dan dilakukan build ulang akan menggunakan 
# cache pada build image sebelumnya
COPY go.mod go.sum ./

# Execute any commands on top of the current image as a new layer and commit the results.
# Mendownload modules yang sudah disalin sebelumnya
RUN go mod download -x

# Copy all source code
COPY . .

# Compile source code menjadi binary file
RUN go build -o main ./cmd/api

RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

# Menjalankan file binary
CMD [ "./main" ]