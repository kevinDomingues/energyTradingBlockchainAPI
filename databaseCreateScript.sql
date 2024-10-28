CREATE TABLE Admin (
    adminID UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    blockchain_user UUID NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    city VARCHAR(100),
    postal_code VARCHAR(20),
    user_type SMALLINT NOT NULL DEFAULT 0
);

CREATE TABLE energyProducers (
    id UUID PRIMARY KEY,
    energy_source VARCHAR(255),
    production_capacity DECIMAL(10, 2),
    FOREIGN KEY (id) REFERENCES users(id)
);