# docker build -t dply-server .
# docker tag dply-server herryg91/dply-server && docker push herryg91/dply-server
# docker run -d  -p 30080:80 -p 39090:9090 --name dply-server dply-server
FROM golang:1.16-buster as builder

ARG GITHUB_TOKEN

RUN mkdir /app
ADD . /app

WORKDIR /app
RUN GOOS=linux go build -o /bin/goapp

WORKDIR /app/handler/cli/create-user
RUN GOOS=linux go build -o /bin/create-user

WORKDIR /app/handler/cli/edit-user
RUN GOOS=linux go build -o /bin/edit-user

WORKDIR /app/handler/cli/image-delete
RUN GOOS=linux go build -o /bin/image-delete

WORKDIR /app/handler/cli/set-default-port
RUN GOOS=linux go build -o /bin/set-default-port

WORKDIR /app/handler/cli/set-default-affinity
RUN GOOS=linux go build -o /bin/set-default-affinity

WORKDIR /app/handler/cli/set-user-to-active
RUN GOOS=linux go build -o /bin/set-user-to-active

WORKDIR /app/handler/cli/set-user-to-inactive
RUN GOOS=linux go build -o /bin/set-user-to-inactive


FROM debian:buster-slim
RUN apt-get update -y \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        openssl \
        bash \
        curl \
        wget \
        tar \
        gzip \
        vim \
        nano \
    && update-ca-certificates \
    && apt-get clean \
    && rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /bin/goapp goapp-output
COPY --from=builder /bin/create-user create-user
COPY --from=builder /bin/edit-user edit-user
COPY --from=builder /bin/image-delete image-delete
COPY --from=builder /bin/set-default-port set-default-port
COPY --from=builder /bin/set-default-affinity set-default-affinity
COPY --from=builder /bin/set-user-to-active set-user-to-active
COPY --from=builder /bin/set-user-to-inactive set-user-to-inactive

EXPOSE 80
EXPOSE 9090
EXPOSE 9800

ENTRYPOINT ["./goapp-output"]