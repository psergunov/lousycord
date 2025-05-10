FROM golang:1.24

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/app

EXPOSE 8080

CMD ["./server"]