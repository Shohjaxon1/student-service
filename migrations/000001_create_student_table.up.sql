CREATE TYPE gender_type AS ENUM('male', 'female');

CREATE TABLE IF NOT EXISTS students (
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    date_of_birth DATE,
    gender gender_type NOT NULL,
    phone_number VARCHAR(20),
    address TEXT,
    points INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
