<div class="button-locker">LOCKER</div> 

# Changelog


---
# 1.0.0 (Planned - 2021 Q1)
## FEATURE

- Basic Auth (HTPASSWD)
- Basic web frontend with a tree-like view
- Fully logging capability implemented with rotated logfiles
- Parts of code is covered by tests
- **(Post 1.0)** *Project based profiles in config -> No need to pass multiple* artifact paths.
---
# 0.4 (Unreleased)
## Feature Backlog
- SQLITE database with ORM setup
- Ability for Agent to request artifact from the Server
- Rework logging to use io.multiwriter (STDOUT and configured log file)

# 0.3 (2020.12.25)
## Feature
- Basic Docker support via **Dockerfile** and **docker-compose.yml**
- Artifact Root Directory Configuration is implemented

## Fixes
- Fixed agent panic when send_concurrent is false

# 0.2 (2020.12.08)
## Feature
- Ability for Agent to send multiple files with a single configuration
- Agent and Server can be configured with .yml files and ENV variables (Not complete functionality)

## BUGS
- If the Agent cannot connect to the Server (bad configuration) Locker still crashes, without properly handling the error.
- Agent Panics when send_concurrent is 'false', after the first file is sent
---
# 0.1 (2020.12.03)
## FEATURE 

- Send file over TCP /w Protobuf message to statically configured server
- Server compares recieved file and MetaData Packet
- Recieved file is re-created in `./out/namespace/project/job-id` folder structure

## BUGS

- Locker throws OS specific error when server is not running

