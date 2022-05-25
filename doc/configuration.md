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
docker_args = ["-v", "name:/sdf"]
env.env1 = "val"
tags = ["tag"]
on.error.action = "retry"
on.error.max_retries = 3
```

### Transformer

### Monitor

### Sync

## Server Configuration
