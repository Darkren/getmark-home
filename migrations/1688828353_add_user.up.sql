CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    login character varying (60),
    password character varying (60),
    name character varying (100),
    email character varying (100)
);