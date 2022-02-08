from mysql:latest

USER mysql
COPY migrate/* /docker-entrypoint-initdb.d/

ENTRYPOINT ["mysqld"]
