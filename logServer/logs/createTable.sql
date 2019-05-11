create table t_orelog (
logid SERIAL primary key,
logtime varchar(100) default now(),
admin varchar(100) not null,
module varchar(100) not null,
logsql varchar(255) not null
);