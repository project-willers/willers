from mysql:latest

COPY migrate/* /docker-entrypoint-initdb.d/

ENTRYPOINT ["mysqld"]
CMD ["--initialize","--user=root"]
