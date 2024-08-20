-- Migration up script for email-events
SET statement_timeout = 0;

--bun:split

CREATE TABLE email.events (
    id SERIAL PRIMARY KEY,
    message_id INT NOT NULL REFERENCES email.messages(id) ON DELETE CASCADE,
    template_id INT NOT NULL REFERENCES email.templates(id) ON DELETE CASCADE,
    event_type VARCHAR(50) NOT NULL,
    payload  JSON DEFAULT NULL,
    event_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);