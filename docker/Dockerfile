FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY backend/ ./backend/
COPY config.yaml .
COPY data.json .
RUN go build -o main ./backend/graph
CMD ["./main"]