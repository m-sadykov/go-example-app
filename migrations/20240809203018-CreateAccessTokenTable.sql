
-- +migrate Up
create table public.access_tokens (
    id serial primary key,
    token varchar not null,
    user_id serial,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

alter table only public.access_tokens
    add constraint "user_id_fk" foreign key (user_id) references public.users(id) on delete cascade;

-- +migrate Down
alter table public.access_tokens drop constraint "user_id_fk";

drop table public.access_tokens;
