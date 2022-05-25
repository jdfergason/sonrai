CREATE TABLE activity (
    id UUID PRIMARY KEY,
    user TEXT,
    event_time TIMESTAMP,
    event_type TEXT,
    message TEXT
);

CREATE TABLE jobs (
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
    schedule TEXT,
    lastRun TIMESTAMP,
    recordTypes JSONB,
    tags JSONB,
    created TIMESTAMP,
    updated TIMESTAMP
);

CREATE TABLE job_run (
    run_id UUID PRIMARY KEY,
    job_id UUID NOT NULL,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    error_count INT,
    expires TIMESTAMP,
    FOREIGN KEY(job_id) REFERENCES job(id) ON DELETE CASCADE
);

CREATE TABLE job_log (
    job_id UUID,
    run_id UUID NOT NULL,
    log_time TIMESTAMP,
    log_level TEXT,
    source TEXT,
    message TEXT,
    expires TIMESTAMP,
    FOREIGN KEY(job_id) REFERENCES job(id) ON DELETE CASCADE,
    FOREIGN KEY(run_id) REFERENCES job_run(run_id) ON DELETE CASCADE
);

CREATE TABLE secrets (
    user TEXT,
    name TEXT,
    secret TEXT
);
