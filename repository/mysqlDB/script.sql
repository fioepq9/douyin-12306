DROP DATABASE IF EXISTS douyin_12306;

CREATE DATABASE douyin_12306;

USE douyin_12306;

create table user
(
    user_id  BIGINT,
    token    VARCHAR(255) not null,
    username CHAR(32)     not null,
    password CHAR(32)     not null,
    constraint user_pk
        primary key (username)
);

alter table user
    modify user_id BIGINT auto_increment;
