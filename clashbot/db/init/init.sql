CREATE SCHEMA IF NOT EXISTS clashbot;

CREATE TABLE IF NOT EXISTS clashbot.telegram_users (
    id BIGINT PRIMARY KEY,
    first_name VARCHAR(255),
    username VARCHAR(255),
    date BIGINT,
    clash_tag VARCHAR(20),
    is_admin BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);