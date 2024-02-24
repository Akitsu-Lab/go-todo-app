FROM mysql:8.0

# docker-compose.ymlで指定したらいらないかも？
ENV MYSQL_ROOT_PASSWORD rootpass
ENV MYSQL_DATABASE todo
ENV MYSQL_USER user
ENV MYSQL_PASSWORD pass

COPY init/* /docker-entrypoint-initdb.d/
CMD ["mysqld"]
