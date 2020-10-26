CREATE TABLE public.friends
(
    id integer NOT NULL DEFAULT nextval('friends_id_seq'::regclass),
    user_one_email text COLLATE pg_catalog."default" NOT NULL,
    user_two_email text COLLATE pg_catalog."default" NOT NULL,
    update_status boolean NOT NULL,
    CONSTRAINT friends_pkey PRIMARY KEY (id, user_one_email, user_two_email)
)