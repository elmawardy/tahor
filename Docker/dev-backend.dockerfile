FROM golang:latest
RUN go get github.com/githubnemo/CompileDaemon

COPY dev.sh /docker-entrypoint-initdb.d/dev.sh
RUN sed -i 's/\r$//g' /docker-entrypoint-initdb.d/dev.sh
RUN chmod 777 /docker-entrypoint-initdb.d/dev.sh

CMD ["/docker-entrypoint-initdb.d/dev.sh"]

EXPOSE 8000