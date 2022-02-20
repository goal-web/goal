FROM alpine

RUN mkdir "/opt/app"

COPY ./bin_linux /opt/app/app

ENV http.port 8008

EXPOSE 8008

WORKDIR /opt/app

ENTRYPOINT /opt/app/app run