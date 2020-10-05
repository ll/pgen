# Golang struct generator for PostgreSQL

CLI tool for generation Golang structures by table definition from PostgreSQL.

## Install
```
go mod tidy
go build ./... 
cp pgen $GOPATH/bin
```


## Help
```
Usage:
  pgen [OPTIONS]

Application Options:
  -s, --server=         Server name or IP address (default: 127.0.0.1)
  -p, --port=           Port (default: 5432)
  -u, --user=           Database user.
  -w, --password=       Database password.
  -d, --database=       Database name.
  -t, --tables=         Tables to export.
      --ssl=            SSL mode (require|verify-full|verify-ca|disable) (default: disable)
      --pkg=            Package name for generated files.

Help Options:
  -h, --help            Show this help message
```
