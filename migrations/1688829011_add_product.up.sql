CREATE TABLE IF NOT EXISTS products (
    barcode character varying (36) PRIMARY KEY,
    name character varying(200),
    "desc" text,
    cost bigint,
    user_id bigint,

    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);