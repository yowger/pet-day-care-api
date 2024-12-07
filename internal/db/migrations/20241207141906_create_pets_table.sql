-- migrate:up
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL CHECK (name != ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
);
-- initial roles
INSERT INTO roles (name)
VALUES ('owner'),
    ('staff');
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
-- initial species
INSERT INTO species (name)
VALUES ('dog'),
    ('cat');
CREATE TABLE breeds (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    species_id int NOT NULL REFERENCES species(id) ON DELETE CASCADE,
    UNIQUE (name, species_id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
);
-- initial breeds, 
-- todo, seed instead
INSERT INTO breeds (name, species_id)
VALUES (
        'Labrador Retriever',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'Golden Retriever',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'Bulldog',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'Beagle',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'Poodle',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'German Shepherd',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'Chihuahua',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'Pug',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'Boxer',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'Rottweiler',
        (
            SELECT id
            FROM species
            WHERE name = 'dog'
        )
    ),
    (
        'Persian',
        (
            SELECT id
            FROM species
            WHERE name = 'cat'
        )
    ),
    (
        'Siamese',
        (
            SELECT id
            FROM species
            WHERE name = 'cat'
        )
    ),
    (
        'Maine Coon',
        (
            SELECT id
            FROM species
            WHERE name = 'cat'
        )
    ),
    (
        'Bengal',
        (
            SELECT id
            FROM species
            WHERE name = 'cat'
        )
    ),
    (
        'Sphynx',
        (
            SELECT id
            FROM species
            WHERE name = 'cat'
        )
    ),
    (
        'Abyssinian',
        (
            SELECT id
            FROM species
            WHERE name = 'cat'
        )
    );
CREATE TABLE pets (
    id SERIAL PRIMARY KEY,
    age TIMESTAMP NOT NULL,
    name VARCHAR(32) NOT NULL CHECK (name != ''),
    species_id int NOT NULL REFERENCES species(id) ON DELETE CASCADE,
    breed_id int NOT NULL REFERENCES breeds(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
);
-- create index
CREATE INDEX idx_users_created_at ON users (created_at DESC);
CREATE INDEX idx_pets_owner_id ON pets (owner_id);
CREATE INDEX idx_pets_species_id ON pets (species_id);
CREATE INDEX idx_pets_breed_id ON pets (breed_id);
CREATE INDEX idx_breeds_species_id ON breeds (species_id);
CREATE INDEX idx_roles_name ON roles (name);
CREATE INDEX idx_pets_species_breed_id ON pets (species_id, breed_id);
CREATE INDEX idx_species_name ON species (name);
-- migrate:down
-- drop index
DROP INDEX IF EXISTS idx_species_name;
DROP INDEX IF EXISTS idx_pets_species_breed_id;
DROP INDEX IF EXISTS idx_roles_name;
DROP INDEX IF EXISTS idx_breeds_species_id;
DROP INDEX IF EXISTS idx_pets_breed_id;
DROP INDEX IF EXISTS idx_pets_species_id;
DROP INDEX IF EXISTS idx_pets_owner_id;
DROP INDEX IF EXISTS idx_users_created_at;
-- drop tables
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS pets;
DROP TABLE IF EXISTS breeds;
DROP TABLE IF EXISTS species;
DROP TABLE IF EXISTS roles;