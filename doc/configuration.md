# Configuration

## Project Configuration

A project consists of a set of jobs defined in a TOML file.

### Job

```toml
[[job]]
name = "Job Name (required)"
slug = "Short URL safe name (generated from name if not provided)"
description = "Description of job"
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

## Server Configuration
