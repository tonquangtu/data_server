DROP TABLE IF EXISTS public.user;
CREATE TABLE public.user
(
    id                serial primary key not null ,
    username          varchar(64)         not null,
    password_hash     varchar(255)        not null,
    created_at        timestamp not null,
    updated_at        timestamp not null,
    deleted_at        timestamp null
);


DROP TABLE IF EXISTS public.device;
CREATE TABLE public.device
(
    id                serial primary key NOT NULL,
    user_id           int NOT NULL,
    device_name       varchar(255) not null,
    device_detail     varchar(512) not null,
    created_at        timestamp not null,
    updated_at        timestamp not null,
    deleted_at        timestamp null
);


DROP TABLE IF EXISTS public.note;
CREATE TABLE public.note
(
    id                serial primary key not null,
    user_id           int not null,
    device_id         int not null,
    data              text null,
    created_at        timestamp not null,
    updated_at        timestamp not null,
    deleted_at        timestamp null
);