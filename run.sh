#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/...
/server
