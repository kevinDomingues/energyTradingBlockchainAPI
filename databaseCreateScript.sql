CREATE TABLE Admin (
    adminID UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE location (
    locationId SERIAL PRIMARY KEY,
    postalCode VARCHAR(10) NOT NULL,
    city VARCHAR(50) NOT NULL,
    UNIQUE (postalCode, city)
)
;
CREATE TABLE users (
    id UUID PRIMARY KEY,
    blockchain_user UUID NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    city VARCHAR(100),
    postal_code VARCHAR(20),
    user_type SMALLINT NOT NULL DEFAULT 0,
    UNIQUE (email)
);

CREATE TABLE consumption (
    userId UUID REFERENCES users(id) ON DELETE CASCADE,
    consumptionYear INT NOT NULL,
    consumptionMonth INT NOT NULL,
    energyConsumed DOUBLE PRECISION NOT NULL,
    PRIMARY KEY (userId, consumptionYear, consumptionMonth)
);

CREATE TABLE energyProducers (
    id UUID PRIMARY KEY,
    energy_source VARCHAR(255),
    production_capacity DECIMAL(10, 2),
    FOREIGN KEY (id) REFERENCES users(id)
);

CREATE TABLE RegulatoryAuthority (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255) NOT NULL,
    apiURL VARCHAR(255)
);