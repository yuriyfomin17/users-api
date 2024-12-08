CREATE TABLE IF NOT EXISTS USERS
(
    id         varchar(255) primary key,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    email      varchar(255) not null unique,
    age        int,
    created_at timestamp(6) not null DEFAULT CURRENT_TIMESTAMP
)