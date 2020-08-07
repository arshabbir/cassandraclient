FROM golang

ENV CLUSTERIP "54.90.57.182"
ENV PORT "8080"

RUN mkdir /build

COPY go.mod .
COPY go.sum .

WORKDIR /build

RUN go mod download

RUN go build -o cassandraclient

EXPOSE 8080

ENTRYPOINT [".\cassandraclient"]

