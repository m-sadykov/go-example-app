
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

alter table only public.users
 add constraint "users_email_uniq_constraint" unique (email);

-- +migrate Down
alter table public.users drop constraint "users_email_uniq_constraint";
drop table public.users;
