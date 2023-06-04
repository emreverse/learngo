CREATE TABLE emre (
    id int,
    name varchar(255),
    lastlogin timestamp,
    PRIMARY KEY (id)
);

-- pgbackrest --stanza=pg-test-01-stanza --type=time --target="2023-04-03 11:09:32.888117+03" restore

-- pg_ctl -D /pgsql-data/ start