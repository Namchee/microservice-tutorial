CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) UNIQUE NOT NULL,
    "name" VARCHAR(128) NOT NULL,
    bio VARCHAR(256) NOT NUll
);

INSERT INTO "user" (username, "name", bio) VALUES
('namchee', 'Cristopher Namchee', 'I am Batman');