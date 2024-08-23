-- Migration up script for email-templates
SET statement_timeout = 0;

--bun:split

CREATE TABLE email.templates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    text_content TEXT,
    html_content TEXT,
    placeholders JSON DEFAULT NULL
);