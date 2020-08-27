FROM golang

ENV CLUSTERIP "3.90.70.181"
ENV PORT ":8081"

RUN mkdir /build

COPY . /build
COPY . /build

WORKDIR /build

RUN go mod download

RUN go build -o cassandraclient

EXPOSE 8081

ENTRYPOINT ["./cassandraclient"]

