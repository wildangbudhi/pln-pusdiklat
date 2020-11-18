FROM golang

WORKDIR /src
COPY . .
RUN go build -o apigateway
CMD ["./apigateway"]