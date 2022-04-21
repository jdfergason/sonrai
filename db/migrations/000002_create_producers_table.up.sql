CREATE TABLE producer (
    id UUID PRIMARY KEY,
    user TEXT,
    slug TEXT,
    name TEXT,
    description TEXT,
    kind TEXT,
    command TEXT,
    arguments JSONB,
    environment JSONB,
    secrets JSONB,
    log_configuration JSONB,
    schedule TEXT,
    lastRun TIMESTAMP,
    recordTypes JSONB,
    created TIMESTAMP,
    updated TIMESTAMP
);
