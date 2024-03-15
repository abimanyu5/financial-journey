-- +migrate Up

CREATE TABLE masters (
    id bigserial NOT NULL,
    name VARCHAR(256),
    description VARCHAR(256),
    user_id INT,
    income_type INT,
    created_at DATE,
    updated_at DATE
);

CREATE TABLE transactions (
    id bigserial NOT NULL,
    master_id INT,
    goal_id INT,
    user_id INT,
    amount INT,
    title VARCHAR(256),
    description VARCHAR(256),
    created_at DATE,
    updated_at DATE
);

CREATE TABLE goals (
    id bigserial NOT NULL,
    id_user INT,
    amount INT,
    amount_goal INT,
    name VARCHAR(256),
    description VARCHAR(256),
    created_at DATE,
    updated_at DATE
);

-- CREATE TABLE reports (
--     id bigserial NOT NULL,
--     id_user INT,
--     total_amount INT,
--     amount_income INT,
--     amount_outcome INT,
--     amount_different INT,
--     created_at DATE,
--     updated_at DATE
-- );

CREATE TABLE users (
	id bigserial NOT NULL,
	name varchar(256) NULL,
	password varchar(256) NULL,
    email VARCHAR(256) NULL,
	role varchar(256) NULL,
	created_at date,
	updated_at date
);