-- migrate:up
CREATE INDEX idx_users_created_at ON users (created_at DESC);
CREATE INDEX idx_pets_species_id ON pets (species_id);
CREATE INDEX idx_pets_breed_id ON pets (breed_id);
CREATE INDEX idx_breeds_species_id ON breeds (species_id);
CREATE INDEX idx_roles_name ON roles (name);
CREATE INDEX idx_pets_species_breed_id ON pets (species_id, breed_id);
CREATE INDEX idx_species_name ON species (name);
-- migrate:down
DROP INDEX IF EXISTS idx_species_name;
DROP INDEX IF EXISTS idx_pets_species_breed_id;
DROP INDEX IF EXISTS idx_roles_name;
DROP INDEX IF EXISTS idx_breeds_species_id;
DROP INDEX IF EXISTS idx_pets_breed_id;
DROP INDEX IF EXISTS idx_pets_species_id;
DROP INDEX IF EXISTS idx_users_created_at;