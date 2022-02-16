FROM postgres
ENV POSTGRES_DB todo_database
ENV POSTGRES_PASSWORD admin1234
ENV POSTGRES_USER todo

COPY schema.sql /docker-entrypoint-initdb.d/init.sql
