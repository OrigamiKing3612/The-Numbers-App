create table cars (
    id varchar(36) primary key not null,
    last_name text not null,
    number integer not null,
    checked_in integer(1) not null default 0
);
