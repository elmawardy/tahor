FROM postgres:11
ENV POSTGRES_PASSWORD 123456
ENV DBNAME tahor

COPY init-db.sh /docker-entrypoint-initdb.d/init-db.sh
RUN sed -i 's/\r$//g' /docker-entrypoint-initdb.d/init-db.sh
RUN chmod 777 /docker-entrypoint-initdb.d/init-db.sh
RUN /docker-entrypoint-initdb.d/init-db.sh

EXPOSE 5432