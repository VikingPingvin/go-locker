# Welcome to Locker!

**Locker** is blazing fast GO application meant to be used in CI/CD pipelines as a generic **artifact blob storage**.

Uses TCP/IP connection for sending file metadata and payload using protobuf between the Agent and the Server, which stores it in a database according to `namespace/project/artifact`.

## Status
[![vikingpingvin](https://img.shields.io/circleci/build/gh/VikingPingvin/go-locker/master)](https://app.circleci.com/pipelines/github/VikingPingvin/go-locker?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/VikingPingvin/go-locker)](https://goreportcard.com/report/github.com/VikingPingvin/go-locker)
[![Go Version](https://img.shields.io/github/go-mod/go-version/VikingPingvin/go-locker)](https://img.shields.io/github/go-mod/go-version/VikingPingvin/go-locker)

## Benchmarks
**Avg.Time** is the duration from the start of the TCP connection, until the end of it, and as such, includes the complete file chunking, sending, reassemblyng server-side,
hashing and sending an ACK message to the agent, followed by the closing of the TCP connection.

<details>
<summary>For reference the HDD used for localhost benchmarking </summary>
<br>
Western Digital Caviar Blue 3.5 1TB 7200rpm 64MB SATA3
<hr>
</details>

#### Localhost Agent -> Localhost Server

|File size|Avg. Time| Write Speed |
|---------|---------|-------------|
|3.59 Gb  | 162s    | ~22 MB/s    |
|615 MB   | 23.6s   | ~26 MB/s    |
|393 MB   | 14.5s   | ~27 MB/s    |
|196 MB   | 8s      | ~24.5 MB/s  |
|47.5 MB  | 1.82s   | ~26 MB/s    |

#### Localhost Agent -> Localhost Docker Server
With **Bind** mounted output directory.
|File size|Avg. Time| Write Speed |
|---------|---------|-------------|
|196 MB   | 15s     | ~13 MB/s    |
|47.5 MB  | 3.7s    | ~12.8 MB/s  |

#### Localhost Agent -> Cloud VM (Linode Nanode 1GB) native Server App

|File size|Avg. Time| Write Speed |
|---------|---------|-------------|
|196 MB   | 77s     | ~2.5 MB/s    |
|47.5 MB  | 22s     | ~2.1 MB/s   |

### Localhost Agent -> Cloud VM using Docker Server
Not tested

## Usage example


### Start server
> .\locker.exe server -c="./cfg-server.yml"

To start Locker in server mode use the `server` command.  
The server app handles initial configuration and database connection and waits for incoming connections.
### Start agent
> ./locker agent --file="path-to-file" --namespace="namespace/project/job-id

To start Locker in Agent mode, use the `agent` command.  
You can supply multiple files for the `--file` flag separated with **comma**.  
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
  artifacts_root_dir: "./out"
  log_path: "./locker_server_log.txt" #not yet implemented
```

If no config file is supplies via `-c` flag, the above values are used as default.
### With ENV variables
If no ENV vars are supplied, the following are used as defaults, if not set otherwise:
```yml
# AGENT
LOCKER_SERVER_IP = "127.0.0.1" # The server to connect to
LOCKER_SERVER_PORT = "27001"
LOCKER_AGENT_CONCURRENT = true
LOCKER_AGENT_LOG = "./locker-agent.log"
```
```yml
# SERVER
LOCKER_SERVER_IP = "127.0.0.1"
LOCKER_SERVER_PORT = "27001"
LOCKER_SERVER_LOG = "./locker-server.log"
LOCKER_SERVER_ARTIFACTS_ROOT = "."
```

## Using with Docker
An example `docker-compose.yml` is supplied here:
[Docker Compose Readme](Docker/readme.md)  
Disclaimer: The example docker related files are not final and WILL change over the time, only there to test functionality.

## Building
To build from source:
> go build .

To cross-compile, follow the Go compiler instruction for `GOOS` and `GOARCH` enviroment variables.
