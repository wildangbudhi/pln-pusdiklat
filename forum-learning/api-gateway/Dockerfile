FROM golang

WORKDIR /src
COPY . .
RUN mkdir config
RUN go build -o apigateway
CMD ["./apigateway"]