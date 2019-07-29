[![Build Status](https://travis-ci.com/logingood/gooz.svg?branch=master)](https://travis-ci.com/logingood/gooz)

# Description

Go version of Zendesk Coding Challenge, code name GOOZ.

To build cli we have used [spf13 cobra](https://github.com/spf13/cobra)

* Search is organized using basic hash table on maps
* Hash table takes interface{} as input
* Search module attempts to detect types and makes everything `string`
* Table module draw tables and also detects type
* We validate input against provided schema for safety reasons
* We use [sync.RWMutex](https://golang.org/pkg/sync/#RWMutex) to ensure thread
  safety, however this tool is single threaded. No guarantee if it will work
  correctly for parallel execution because we use [golang
  maps](https://blog.golang.org/go-maps-in-action). See concurrency section.

# Challenge check list

- [x] Separation of concerns: every module is implemented as independent package (e.g. `index`, `search`, `table`, `cli` and etc)
- [x] Simplicity: we used the most basic type of data structure - hash table and list to organize search index
- [x] Test coverage: every module is tested separate, test coverage provided. CLI tests are in progress
- [ ] Performance analysis TBD
- [x] Errors are handled and logged, corresponding tests provided

# Install and build

You can use `go get` command:

```
go get -u github.com/logingood/gooz
```

Alternatively use [Docker](https://github.com/logingood/gooz#dockerfile) or build yourself:
```
git clone https://github.com/logingood/gooz
cd gooz
go run . --help
```

# CLI and Configuration

By default `gooz` will look for tables in `./data/tickets.json`, `./data/users.json` and `./data/organizations.json`.
To alter this you can use `--organizations_path`, `--tickets_path` and `--users_path` flags and supply to cli.

Search tickets by any field, e.g. subject

```
gooz tickets subject "A Problem in South Africa"
```

For short output you can use flag `--related=false`, it is true by default and prints all related informationt to found items
```
gooz users suspended true --related=false
```

Search organization by any field, e.g.

```
gooz organizations created_at  "2016-05-21T11:10:28 -10:00"
```

Search users by any field
```
gooz users alias "Miss Rosanna"`
```

Search by empty field, just use ""

```
gooz tickets assignee_id ""
```

Check all availabel options and configuration keys with

```
gooz help
```

```
Usage:
  gooz [command]

Available Commands:
  help          Help about any command
  organizations Search organizations table
  tickets       Tickets search
  users         search users table

Flags:
      --config string               config file (default is $HOME/.gooz.yaml)
  -h, --help                        help for gooz
      --organizations_path string   path to your organizations.json, default is data/organizations.json (default "data/organizations.json")
      --tickets_path string         path to your tickets.json, default is data/tickets.json (default "data/tickets.json")
      --users_path string           path to your users.json, default is data/users.json (default "data/users.json")
```

# Dockerfile

One of the easiest ways to build and run the code is a docker command

```
git clone https://github.com/logingood/gooz
cd gooz

docker build . -t gooz
```

Run the container
```
docker rm -f gooz >/dev/null 2>&1;  docker run -ti --name gooz gooz:latest /bin/gooz --help
```


# Pluggable backend

Backend implements interface `Store`, you need to implement three methods in order to use interface.
At the moment only Filesystem with json files is implemented as per the task, however you can replace it
with other backends that implements similar interface without requirement to rewrite whole code.
E.g. using sqlx with sqlite/mysql and etc.

# Test coverage

```
ok      github.com/logingood/gooz       0.043s  coverage: 0.0% of statements
ok      github.com/logingood/gooz/cmd   0.016s  coverage: 18.1% of statements
ok      github.com/logingood/gooz/internal/backend      (cached)        coverage: 0.0% of statements
ok      github.com/logingood/gooz/internal/backend/zfile        (cached)        coverage: 94.7% of statements
ok      github.com/logingood/gooz/internal/helpers      0.028s  coverage: 97.3% of statements
ok      github.com/logingood/gooz/internal/index        (cached)        coverage: 100.0% of statements
ok      github.com/logingood/gooz/internal/search       0.023s  coverage: 87.0% of statements
ok      github.com/logingood/gooz/internal/table        0.014s  coverage: 0.0% of statements
```

# Performance
