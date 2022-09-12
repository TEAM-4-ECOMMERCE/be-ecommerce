FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o e-commerce-app

CMD ["./e-commerce-app"]