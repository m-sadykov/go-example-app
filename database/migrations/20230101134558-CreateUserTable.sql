
-- +migrate Up
create table public.user (
    id int primary key,
    name varchar not null,
    email varchar not null,
    password varchar not null,
    cratedAt timestamp without time zone,
    updatedAt timestamp without time zone,
    deletedAt timestamp without time zone
);
-- +migrate Down
drop table user;
