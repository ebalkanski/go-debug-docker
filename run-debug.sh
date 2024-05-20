#!/bin/bash

go build -gcflags "all=-N -l" -o /server ./cmd/...
dlv --listen=:40000 --headless=true --api-version=2 --accept-multiclient exec /server
