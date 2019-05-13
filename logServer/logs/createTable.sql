create table t_orelog (
logid SERIAL primary key,
create_time timestamp default current_timestamp,
admin varchar(100) not null,
module varchar(100) not null,
logsql varchar(1000) not null
);