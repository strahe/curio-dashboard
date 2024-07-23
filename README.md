# curio-dashboard

Status: **under heavy development**

## Description
`curio-dashboard` is a dashboard for [curio](https://github.com/filecoin-project/curio), it provides a web interface to monitor the status of curio.

## Development

### Prerequisites
```shell
# Install system dependencies, refer to: https://lotus.filecoin.io/lotus/install/linux/#building-from-source

git clone https://github.com/strahe/curio-dashboard.git

cd curio-dashboard

make all
```

### Run 
```shell
# Run Server

# Set the environment variable `FULLNODE_API_INFO` to the Lotus API endpoint.
export FULLNODE_API_INFO=""

# For the current phase, 
# the dashboard only allows connections to read-only databases (slave nodes). 
# If you insist on connecting to a read-write database (master node), 
# add the `target_session_attrs=any` parameter to the URL.
export CURIO_HARMONYDB_URL="postgres://yugabyte:yugabyte@localhost:5433/curio?search_path=curio"

# A database dedicated for dashboard use.
export CURIO_APPDB_URL="postgres://yugabyte:yugabyte@localhost:5433/dashboard"

./curio-dashboard --debug run

# Run Frontend
cd web && yarn dev

# Open http://localhost:3000
# http://localhost:9091 # for production
# http://127.0.0.1:9091/playground # for graphql playground
```
