create table users
(
    id            SERIAL PRIMARY KEY,
    name          TEXT                                     NOT NULL,
    mail          TEXT UNIQUE,
    password_hash TEXT,
    salt          TEXT,
    role          smallint check (role >= 0 and role <= 1) not null default 0,
    created_at    timestamptz                                       default timezone('europe/moscow'::text, now()),
    updated_at    timestamptz                                       default timezone('europe/moscow'::text, now())
);

create index inx_users_mail on users (mail);

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
