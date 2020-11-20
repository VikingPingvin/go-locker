# Welcome to Locker!

**Locker** is a GO application meant to be used in CI/CD pipelines as an **artifact blob storage**.

Uses TCP/IP connection for sending file metadata and byte stream.

Application can be used both as a server with DB connection and as an agent to send / consume artifacts

Application built with the help of Cobra.

## Status
[![vikingpingvin](https://img.shields.io/circleci/build/gh/VikingPingvin/go-locker/master)](https://app.circleci.com/pipelines/github/VikingPingvin/go-locker?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/VikingPingvin/go-locker)](https://goreportcard.com/report/github.com/VikingPingvin/go-locker)


## Usage example

> go build .
### Start server
> ./locker server
### Starg agent
> ./locker agent --file=<path-to-file>




