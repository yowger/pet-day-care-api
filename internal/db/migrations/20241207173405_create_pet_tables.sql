-- migrate:up
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL CHECK (name != ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
);
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL CHECK (char_length(first_name) > 2),
    last_name VARCHAR(32) NOT NULL CHECK (char_length(last_name) > 2),
    email VARCHAR(64) UNIQUE NOT NULL CHECK (email != ''),
    phone_number VARCHAR(20) NOT NULL,
    password VARCHAR(250) NOT NULL CHECK (password != ''),
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
);
CREATE TABLE species (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL CHECK (name != ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
);
CREATE TABLE breeds (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    species_id int NOT NULL REFERENCES species(id) ON DELETE CASCADE,
    UNIQUE (name, species_id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
);
CREATE TABLE pets (
    id SERIAL PRIMARY KEY,
    birth_date DATE NOT NULL,
    name VARCHAR(32) NOT NULL CHECK (name != ''),
    species_id int NOT NULL REFERENCES species(id) ON DELETE CASCADE,
    breed_id int NOT NULL REFERENCES breeds(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
);
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    pet_id int NOT NULL REFERENCES pets(id) ON DELETE CASCADE,
    start_time DATE NOT NULL,
    end_time DATE NOT NULL,
    comments TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
);
-- migrate:down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS pets;
DROP TABLE IF EXISTS breeds;
DROP TABLE IF EXISTS species;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS bookings;