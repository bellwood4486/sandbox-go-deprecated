create table tweets
(
    id           serial primary key,
    content      text    not null,
    user_name    text    not null,
    comment_num  integer not null default 0,
    star_num     integer not null default 0,
    re_tweet_num integer not null default 0
);
