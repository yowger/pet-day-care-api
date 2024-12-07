-- migrate:up
INSERT INTO roles (name)
VALUES ('owner'),
    ('staff');
INSERT INTO species (name)
VALUES ('dog'),
    ('cat');
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
-- migrate:down