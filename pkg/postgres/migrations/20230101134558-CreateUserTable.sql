
-- +migrate Up
create table public.users (
    id serial primary key,
    name varchar not null,
    email varchar not null,
    password varchar not null,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
-- +migrate Down
drop table public.users;
