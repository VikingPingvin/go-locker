# Welcome to Locker!

**Locker** is a GO application meant to be used in CI/CD pipelines as a generic **artifact blob storage**.

Uses TCP/IP connection for sending file metadata and payload using protobuf between the Agent and the Server, which stores it in a database according to `namespace/project/artifact`.

Application built with the help of Cobra.

## Status
[![vikingpingvin](https://img.shields.io/circleci/build/gh/VikingPingvin/go-locker/master)](https://app.circleci.com/pipelines/github/VikingPingvin/go-locker?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/VikingPingvin/go-locker)](https://goreportcard.com/report/github.com/VikingPingvin/go-locker)


## Usage example


### Start server
> ./locker server

The server app handles initial configuration and database connection and waits for incoming connections.
### Start agent
> ./locker agent --file="path-to-file"

The agent parses the file given under `--file` and collects metadata, such as sha256 hash.
It sends the metadata to the server, followed by payload messages containing the raw binary file contents.

## Building
To build from source:
> go build .

