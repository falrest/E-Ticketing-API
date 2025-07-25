CREATE DATABASE eticket;

-- Admin user
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) UNIQUE,
  password TEXT
);

-- Terminal
CREATE TABLE terminals (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  location VARCHAR(255)
);
