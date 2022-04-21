CREATE TABLE activity (
    id UUID PRIMARY KEY,
    user TEXT,
    event_time TIMESTAMP,
    event_type TEXT,
    message TEXT
);