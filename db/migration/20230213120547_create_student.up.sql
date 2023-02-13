CREATE TABLE student (
                         id bigserial not null primary key,
                         login varchar not null unique,
                         encrypted_password varchar not null
);