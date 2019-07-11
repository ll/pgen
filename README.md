# Golang struct generator for PostgreSQL

CLI tool for generation Golang structures by table definition from PostgreSQL.

## Build
```
go build ./... 
```


## Help
```
Usage:
  go-pg-generator [OPTIONS]

Application Options:
  -s, --server=         Server name or IP address (default: 127.0.0.1)
  -p, --port=           Port (default: 5432)
  -u, --user=           Database user.
  -w, --password=       Database password.
  -d, --database=       Database name.
  -t, --tables=         Tables to export.
      --ssl=            SSL mode (require|verify-full|verify-ca|disable) (default: disable)
  -f, --file-per-table  Save each structure to its own .go file.
      --package=        Package name for generated files.

Help Options:
  -h, --help            Show this help message
```
