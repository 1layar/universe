-- Migration up script for emails
SET statement_timeout = 0;

--bun:split
CREATE SCHEMA IF NOT EXISTS email;

--bun:split
CREATE TABLE email.messages (
    id SERIAL PRIMARY KEY,
    to_email VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    text_body TEXT,
    html_body TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    send_at TIMESTAMP DEFAULT NULL,
    open_at TIMESTAMP DEFAULT NULL,
    read_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);