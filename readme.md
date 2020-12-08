# Welcome to Locker!

**Locker** is blazing fast GO application meant to be used in CI/CD pipelines as a generic **artifact blob storage**.

Uses TCP/IP connection for sending file metadata and payload using protobuf between the Agent and the Server, which stores it in a database according to `namespace/project/artifact`.

## Status
[![vikingpingvin](https://img.shields.io/circleci/build/gh/VikingPingvin/go-locker/master)](https://app.circleci.com/pipelines/github/VikingPingvin/go-locker?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/VikingPingvin/go-locker)](https://goreportcard.com/report/github.com/VikingPingvin/go-locker)
[![Go Version](https://img.shields.io/github/go-mod/go-version/VikingPingvin/go-locker)](https://img.shields.io/github/go-mod/go-version/VikingPingvin/go-locker)


## Usage example


### Start server
> ./locker server

To start Locker in server mode use the `server` command.  
The server app handles initial configuration and database connection and waits for incoming connections.
### Start agent
> ./locker agent --file="path-to-file" --namespace="namespace/project/job-id

To start Locker in Agent mode, use the `agent` command.  
The agent parses the file given under the `--file` flag and collects metadata, such as sha256 hash.
It sends the metadata to the server, followed by payload messages containing the raw binary file contents.

### Configuring
#### With YML files
Reference the config files with the `-c` global flag.  
Contents of the files:
```yml
#agent-config
agentconfig:
  server_ip: "127.0.0.1"
  server_port: "27001"
  send_concurrent: true
  log_path: "./locker_agent_log.txt" #Not yet implemented
```
```yml
#server-config
serverconfig:
  server_ip: "127.0.0.1"
  server_port: "27001"
  artifacts_root_dir: "./out" #not yet implemented
  log_path: "./locker_server_log.txt" #not yet implemented
```
### With ENV variables
If no ENV vars are supplied, the following are used as defaults, if not set otherwise:
```yml
# AGENT
LOCKER_SERVER_IP = "127.0.0.1"
LOCKER_SERVER_PORT = "27001"
LOCKER_AGENT_CONCURRENT = true
LOCKER_AGENT_LOG = "./locker-agent.log"
```
```yml
# AGENT
LOCKER-SERVER-IP = "127.0.0.1"
LOCKER-SERVER-PORT = "27001"
LOCKER-SERVER-LOG = "./locker-server.log"
LOCKER-SERVER-ARTIFACTS-ROOT = "."
```
## Building
To build from source:
> go build .

