CREATE TABLE producer_run (
    producer_id UUID,
    run_number INT,
    run_reference TEXT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    records_produced INT,
    error_count INT,
    expires TIMESTAMP
);

CREATE TABLE producer_log (
    producer_id UUID,
    run_number INT,
    log_time TIMESTAMP,
    log_level TEXT,
    source TEXT,
    message TEXT,
    expires TIMESTAMP
);

CREATE TABLE transformer_run (
    transformer_id UUID,
    run_number INT,
    run_reference TEXT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    records_produced INT,
    error_count INT,
    expires TIMESTAMP
);

CREATE TABLE transformer_log (
    transformer_id UUID,
    run_number INT,
    log_time TIMESTAMP,
    log_level TEXT,
    source TEXT,
    message TEXT,
    expires TIMESTAMP
);

CREATE TABLE sync_run (
    sync_id UUID,
    run_number INT,
    run_reference TEXT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    records_produced INT,
    error_count INT,
    expires TIMESTAMP
);

CREATE TABLE sync_log (
    sync_id UUID,
    run_number INT,
    log_time TIMESTAMP,
    log_level TEXT,
    source TEXT,
    message TEXT,
    expires TIMESTAMP
);

CREATE TABLE record_timeseries (
    record_type TEXT,
    event_time TIMESTAMP,
    produced INT,
    errors INT,
    source_kind TEXT,
    source_id UUID
);
