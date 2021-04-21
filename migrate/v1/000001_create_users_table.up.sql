CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    user_name VARCHAR (50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);
