CREATE TABLE IF NOT EXISTS craftingmaterials (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    year integer NOT NULL,
    price integer NOT NULL
);