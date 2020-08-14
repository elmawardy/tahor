FROM golang:1.13.4
RUN go get github.com/githubnemo/CompileDaemon

ADD ./ app/
COPY dev.sh /docker-entrypoint-initdb.d/dev.sh
RUN sed -i 's/\r$//g' /docker-entrypoint-initdb.d/dev.sh
RUN chmod 777 /docker-entrypoint-initdb.d/dev.sh

CMD ["/docker-entrypoint-initdb.d/dev.sh"]

EXPOSE 8000