CREATE TABLE IF NOT EXISTS plan (
    name TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS period (
    name TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS form (
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT NOT NULL UNIQUE,
    plan TEXT NOT NULL,
    period TEXT NOT NULL,
    online_service BOOLEAN DEFAULT FALSE,
    larger_storage BOOLEAN DEFAULT FALSE,
    customizable_profile BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (plan) REFERENCES plan (name),
    FOREIGN KEY (period) REFERENCES period (name)
);

INSERT INTO plan (name)
VALUES
    ('arcade'),
    ('advanced'),
    ('pro');

INSERT INTO period (name)
VALUES
    ('monthly'),
    ('yearly');
