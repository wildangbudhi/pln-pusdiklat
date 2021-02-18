FROM golang

WORKDIR /src
COPY . .

RUN go mod tidy
RUN bash ./unit_test.sh

RUN go build -o identity
CMD ["./identity"]