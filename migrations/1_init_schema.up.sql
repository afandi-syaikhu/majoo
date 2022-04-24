CREATE TABLE Users (
    id bigserial PRIMARY KEY,
    name varchar(45) DEFAULT NULL,
    user_name varchar(45) DEFAULT NULL,
    password varchar(225) DEFAULT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by bigint NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by bigint NOT NULL
);

CREATE TABLE Merchants (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    merchant_name varchar(40) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by bigint NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by bigint NOT NULL
);

CREATE TABLE Outlets (
    id bigserial PRIMARY KEY,
    merchant_id bigint NOT NULL,
    outlet_name varchar(40) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by bigint NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by bigint NOT NULL
);

CREATE TABLE Transactions (
    id bigserial PRIMARY KEY,
    merchant_id bigint NOT NULL,
    outlet_id bigint NOT NULL,
    bill_total double precision NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by bigint NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by bigint NOT NULL
);