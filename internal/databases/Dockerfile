FROM postgres:latest

COPY initialDDL.sql /docker-entrypoint-initdb.d/1.sql

EXPOSE 5432

CMD ["postgres"]