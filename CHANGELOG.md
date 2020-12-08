<div class="button-locker">LOCKER</div> 

# Changelog


---
# 1.0.0 (Planned - 2021 Q1)
## FEATURE

- Basic Auth (HTPASSWD)
- Basic web frontend with a tree-like view
- Parts of code is covered by tests
- **(Post 1.0)** *Project based profiles in config -> No need to pass multiple* artifact paths.
---
# 0.3 (Unreleased)
## Feature
- SQLITE database with ORM setup
- Ability for Agent to request artifact from the Server

# 0.2 (2020.12.08)
## Feature
- Ability for Agent to send multiple files with a single configuration
- Agent and Server can be configured with .yml files and ENV variables (Not complete functionality)

## BUGS
- If the Agent cannot connect to the Server (bad configuration) Locker still crashes, without properly handling the error.
---
# 0.1 (2020.12.03)
## FEATURE 

- Send file over TCP /w Protobuf message to statically configured server
- Server compares recieved file and MetaData Packet
- Recieved file is re-created in `./out/namespace/project/job-id` folder structure

## BUGS

- Locker throws OS specific error when server is not running

