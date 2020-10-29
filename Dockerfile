FROM golang:latest

WORKDIR /app

COPY . /app

EXPOSE 4747

CMD ["go", "run", "main.go"]