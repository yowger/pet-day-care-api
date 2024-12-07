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
-- migrate:down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS pets;
DROP TABLE IF EXISTS breeds;
DROP TABLE IF EXISTS species;
DROP TABLE IF EXISTS roles;