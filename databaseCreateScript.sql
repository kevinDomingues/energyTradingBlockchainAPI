CREATE TABLE Admin (
    adminID UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE location (
    locationId SERIAL PRIMARY KEY,
    postal_code VARCHAR(10) NOT NULL,
    city VARCHAR(50) NOT NULL,
    UNIQUE (postal_code, city)
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

CREATE TABLE energyTypes (
    id SERIAL PRIMARY KEY,
    energy_type VARCHAR(255) NOT NULL
);

CREATE TABLE consumptions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    consumption_year INT NOT NULL,
    consumption_month INT NOT NULL,
    energy_type_id INT NOT NULL REFERENCES energyTypes(id),
    energy_consumed DOUBLE PRECISION NOT NULL,
    UNIQUE (user_id, consumption_year, consumption_month, energy_type_id)
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

INSERT INTO energyTypes (energy_type) VALUES
('Solar'),
('Wind'),
('Hydropower'),
('Geothermal'),
('Biomass');