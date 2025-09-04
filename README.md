# Tana

Tana is a simple CLI that can store keys and their values.

## Installation

If you have Go installed you can use `go install`:

```bash
go install github.com/kostya-zero/tana
```

Otherwise, download binary from the releases section.

## Usage

```bash
# Add key to store
tana set hello "Hello, world!"

# Get key
tana get hello

# List all keys
tana list

# Update key
tana update hello "Hello, John!"

# Delete key 
tana delete hello
```
