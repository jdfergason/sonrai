# Configuration

## Project Configuration

A project consists of a set of producers, transformers, metrics, monitors and syncs. It is configured via JSON or TOML. Each type is an array of 
tables in TOML.

### Producer

```toml
[[producer]]
name = "Producer Name (required)"
slug = "Short URL safe name (generated from name if not provided)"
description = "Description of producer"
schedule = "@every 5m"
type = "docker"
image = "my/image"
args = ["1", "${MY_SECRET}", "3"]
tags = ["tag"]
on.error.action = "retry"
on.error.max_retries = 3

[producer.env]
env1 = "val"
```

NOTE: environment may also be defined using the `.` syntax or as an inline table, like so:

```toml
[[producer]]
env.env1 = "val"
```

and,

```toml
[[producer]]
env = { env1 = "val" }  
```

### Transformer

### Monitor

### Sync

## Server Configuration
