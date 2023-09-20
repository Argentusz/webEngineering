FROM golang

RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd
RUN go build -o main .
COPY localhost.crt .
COPY localhost.csr .
COPY localhost.key .
WORKDIR /app
CMD ["./cmd/main"]