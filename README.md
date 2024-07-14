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

### Start
```shell
# Run Server

export FULLNODE_API_INFO=""
export CURIO_HARMONYDB_URL="postgres://yugabyte:yugabyte@localhost:5433/curio?search_path=curio"
# For data safety, please use a read-only slave node of yugabyte, do not use the master node.
# If you insist on using the master node (ensure you know what you're doing), 
# add '&target_session_attrs=any' to the end of CURIO_HARMONYDB_URL

./curio-dashboard --debug

# Run Frontend
cd web && yarn dev

# Open http://localhost:3000
# Open http://localhost:9091 # for production
```
