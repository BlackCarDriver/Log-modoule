create table t_opelog (
   logid                SERIAL    not null,
   types                VARCHAR(50)   not null,
   operator             VARCHAR(50)   not null,
   logtime               TIMESTAMP    not null  default now(),
   operation            VARCHAR(1000)  not null,
   constraint PK_T_OPELOG primary key (logid)
);

--grant select, insert on t_opelog to logmoduleuser;

--insert into t_opelog(types,operator,operation)values('成功','blackcardriver','登录');


--select count(*) from t_opelog where types='成功';

--select types,operator,logtime,operation from t_opelog where types='成功' offset 0 limit 55