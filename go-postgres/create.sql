drop schema if exists gopg cascade;
create schema gopg;
set search_path='gopg';

create table kv(
    id serial,
    key text,
    value text,
    constraint pk primary key(id, key)
);
