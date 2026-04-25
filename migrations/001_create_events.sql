CREATE TABLE events {
    id SERIAL PRIMARY KEY,
    user_id TEXT,
    event_type TEXT,
    timestamp TIMESTAMP,
    properties JSONB
};
