create schema if not exists core;

create table if not exists core.transactions
(
    id integer not null,
    date varchar not null,
    transaction double precision not null
);