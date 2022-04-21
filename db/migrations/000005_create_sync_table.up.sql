CREATE TABLE sync (
    id UUID PRIMARY KEY,
    user TEXT,
    slug TEXT,
    name TEXT,
    description TEXT,
    kind TEXT,
    repository TEXT,
    command TEXT,
    arguments JSONB,
    environment JSONB,
    secrets JSONB,
    schedule TEXT,
    lastRun TIMESTAMP,
    recordTypes JSONB,
    created TIMESTAMP,
    updated TIMESTAMP
);
