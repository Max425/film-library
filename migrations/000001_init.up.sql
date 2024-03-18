create table users
(
    id            serial primary key,
    name          text                                     not null,
    mail          text unique,
    password_hash text,
    salt          text,
    role          smallint check (role >= 0 and role <= 1) not null default 0,
    created_at    timestamptz                                       default timezone('europe/moscow'::text, now()),
    updated_at    timestamptz                                       default timezone('europe/moscow'::text, now())
);

create index inx_users_mail on users (mail);

insert into users values (1,'admin','admin','32393837663431632d646231662d346264332d613139302d336466653830326239383232d033e22ae348aeb5660fc2140aec35850c4da997','2987f41c-db1f-4bd3-a190-3dfe802b9822',1);
insert into users values (2,'user','user','34643462643365352d303133392d343838642d623830312d65643163323562626333636112dea96fec20593566ab75692c9949596833adc9','4d4bd3e5-0139-488d-b801-ed1c25bbc3ca',0);


create table actor
(
    id         serial primary key,
    name       varchar(255) not null,
    gender     varchar(10)  not null,
    birth_date date         not null,
    created_at timestamptz default timezone('europe/moscow'::text, now()),
    updated_at timestamptz default timezone('europe/moscow'::text, now())
);

create table film
(
    id           serial primary key,
    title        varchar(150)                                       not null,
    description  text,
    release_date date                                               not null,
    rating       decimal(3, 1) check (rating >= 0 and rating <= 10) not null,
    created_at   timestamptz default timezone('europe/moscow'::text, now()),
    updated_at   timestamptz default timezone('europe/moscow'::text, now())
);

create index idx_film_title on film (title);
create index idx_film_release_date on film (release_date);
create index idx_film_rating on film (rating);


create table film_actor
(
    film_id  int references film (id) on delete cascade,
    actor_id int references actor (id) on delete cascade,
    primary key (film_id, actor_id)
);
create index idx_film_actor_film_id_actor_id on film_actor (film_id, actor_id);
