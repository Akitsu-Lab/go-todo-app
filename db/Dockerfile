FROM mysql:8.0

# docker-compose.ymlで指定したらいらないかも？
ENV MYSQL_ROOT_PASSWORD ${DBPASS}
ENV MYSQL_DATABASE todo
ENV MYSQL_USER ${DBUSER}
ENV MYSQL_PASSWORD ${DBPASS}

COPY init/* /docker-entrypoint-initdb.d/
CMD ["mysqld"]
