FROM mysql:8.0

ENV MYSQL_DATABASE todo

COPY init/* /docker-entrypoint-initdb.d/
CMD ["mysqld"]
