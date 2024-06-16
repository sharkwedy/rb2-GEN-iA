-- Creating the PostgreSQL schema for storing person entities
CREATE TABLE IF NOT EXISTS people (
    id UUID PRIMARY KEY,
    apelido VARCHAR(32) NOT NULL UNIQUE,
    nome VARCHAR(100) NOT NULL,
    nascimento DATE NOT NULL,
    stack VARCHAR[]
);

CREATE INDEX idx_people_apelido ON people (apelido);
CREATE INDEX idx_people_nome ON people (nome);