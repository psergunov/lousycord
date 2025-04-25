# Stage 1: Build React app
FROM node:23-slim AS frontend-builder

WORKDIR /app/frontend

COPY frontend/package*.json ./
RUN npm install

COPY frontend/ ./
RUN npm run build


# Stage 2: Build Go app
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

COPY backend/go.mod ./
RUN go mod download

COPY backend/ ./
COPY --from=frontend-builder /app/frontend/build ./frontend/build

RUN go build -o server ./cmd/app


# Stage 3: Final image
FROM alpine:3.21 

WORKDIR /app

COPY --from=backend-builder /app/certs/ ./certs
COPY --from=backend-builder /app/server .
COPY --from=backend-builder /app/frontend/build ./frontend/build


EXPOSE 8080

CMD ["./server"]