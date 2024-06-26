# Menggunakan image golang sebagai base image
FROM golang:1.22-alpine

# Mengatur working directory
WORKDIR /app

ENV HOST 0.0.0.0

# Mengcopy go.mod dan go.sum lalu install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Mengcopy kode sumber ke dalam container
COPY . .

# Build aplikasi
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Menjalankan aplikasi
CMD ["./main"]