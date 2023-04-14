FROM alpine

COPY ./bin_linux /opt/app/app

ENV http.port 8008

EXPOSE 8008

WORKDIR /app

ENTRYPOINT /app/goal run