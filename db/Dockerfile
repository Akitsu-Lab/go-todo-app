FROM mysql:8.0

ENV MYSQL_DATABASE todo

COPY ./config/my.conf /etc/mysql/conf.d/my.cnf
COPY init/* /docker-entrypoint-initdb.d/
CMD ["mysqld"]
