create table telegram_bot_valute.message_reply
(
    id           int auto_increment
        primary key,
    user_id      int                             null,
    message_text text collate utf8mb4_unicode_ci null,
    date_action  datetime                        null
)
    charset = utf8;

create index message_reply_user_id_index
    on telegram_bot_valute.message_reply (user_id);

create table telegram_bot_valute.message_user
(
    id           int auto_increment
        primary key,
    user_id      int                  null,
    message_text text charset utf8mb4 not null,
    date_action  datetime             null
)
    charset = utf8;

create index message_user_user_id_index
    on telegram_bot_valute.message_user (user_id);

create table telegram_bot_valute.summary
(
    id          int auto_increment
        primary key,
    user_id     int         not null,
    chat_id     int         null,
    time_action varchar(32) null,
    constraint summary_user_id_uindex
        unique (user_id)
);

create table telegram_bot_valute.users
(
    id          int auto_increment
        primary key,
    user_id     bigint       null,
    nickname    varchar(128) null,
    first_name  varchar(128) not null comment 'Имя',
    last_name   varchar(128) not null comment 'Фамилия',
    date_action datetime     not null
)
    charset = utf8;

create table telegram_bot_valute.valutes
(
    id       int auto_increment
        primary key,
    date_val date  null,
    usd      float null,
    eur      float null,
    gbp      float null,
    constraint valutes_date_uindex
        unique (date_val)
);

