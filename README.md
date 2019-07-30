[![Build Status](https://travis-ci.com/logingood/gooz.svg?branch=master)](https://travis-ci.com/logingood/gooz)

# Description

Go version of Zendesk Coding Challenge, code name GOOZ.

To build cli we have used [spf13 cobra](https://github.com/spf13/cobra)

* Search is organized using basic hash table on maps
* Results are printed as tables, for short version use `--related=fasle` flag.
  Tables for all related objects will be printed otherwise.
* Hash table takes interface{} as input.
* Search module attempts to detect types and makes everything `string`.
* Table module draws tables and also detects type.
* We validate input against provided schema for safety reasons.
* We use [sync.RWMutex](https://golang.org/pkg/sync/#RWMutex) to ensure thread
  safety, however this tool is single threaded. No guarantee if it works
  correctly for parallel execution because we use [golang
  maps](https://blog.golang.org/go-maps-in-action). See concurrency section.

- [Challenge check list](https://github.com/logingood/gooz#challenge-check-list)
- [Install and build](https://github.com/logingood/gooz#install-and-build)
- [CLI and Configuration](https://github.com/logingood/gooz#cli-and-configuration)
- [Docker build](https://github.com/logingood/gooz#dockerfile)
- [Test coverage](https://github.com/logingood/gooz#test-coverage)
- [Performance](https://github.com/logingood/gooz#performance)

# Challenge check list

- [x] Separation of concerns: every module is implemented as an independent
  package (e.g. `index`, `search`, `table`, `cli`, `backend` and etc)
- [x] Simplicity: we used the most basic type of data structure: a hash table
  and a list to organize a search index
- [x] Test coverage: every module is tested separately, a test coverage report
  is provided. We haven't tested cobra cli because it comes from the library
  and all business logic is tested separately in the corresponding modules.
- [x] Performance analysis is provided. For large subsets 13.79 lookups/ms
  performance should be expected.
- [x] Errors are handled and logged, corresponding tests are provided.

# Install and build

You can use `go get` command:

```
go get -u github.com/logingood/gooz
```

Alternatively use [Docker](https://github.com/logingood/gooz#dockerfile) or
build it yourself:
```
git clone https://github.com/logingood/gooz
cd gooz
go run . --help
```

# CLI and Configuration

By default `gooz` will look for the tables in `./data/tickets.json`,
`./data/users.json` and `./data/organizations.json`.  To change table's path
you can use `--organizations_path`, `--tickets_path` and `--users_path` flags.

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

Backend implements an interface `Store`, you need to implement three methods in
order to use interface.

* `Open() error`
* `Read() (map[string]interface{}, error)`
* `Close() error`

A module that is called `zfile` implements this interface to read supplied
`josn` file from the file system and marshal it into `[]map[string]interface{}`.
However we can extend this tool to use `sqlite`, `mysql`, `redis` and etc.

# Test coverage

```
ok      github.com/logingood/gooz       0.020s  coverage: 0.0% of statements
ok      github.com/logingood/gooz/cmd   0.068s  coverage: 20.3% of statements
ok      github.com/logingood/gooz/internal/backend      (cached)        coverage: 0.0% of statements
ok      github.com/logingood/gooz/internal/backend/zfile        (cached)        coverage: 94.7% of statements
ok      github.com/logingood/gooz/internal/helpers      (cached)        coverage: 97.3% of statements
ok      github.com/logingood/gooz/internal/index        0.023s  coverage: 100.0% of statements
ok      github.com/logingood/gooz/internal/search       0.020s  coverage: 87.0% of statements
ok      github.com/logingood/gooz/internal/table        (cached)        coverage: 0.0% of statements
```

# Performance

You can use the generator flag to generate perfdata, e.g.
```
gooz generate --size 100
```

By default it will create files 100K records each in `perfdata/` directory. 100K of users, organizations and tickets. The size of the each file is around ~30MB.
Let's assume that we search a bool flag e.g. true/false and say it takes true value for everything. Hence we need to iterate through the list of 100K records. Then for the each ID we search of organizations and tickets. If it is a user search we are going to run 2 lookups per user (lookup assignee_id and submitter_id) + building indexes. It should be around 100K + 100K(OrgID) + 100K(assignee_id) + 100K(submitter_id) ~ 400K lookups, it takes around 30s.

Lookup rate  400K/29s = 13.79 lookups/ms.

```
time go run . users role agent --related=false --organizations_path perfdata/organizations.json --tickets_path perfdata/tickets.json --users_path perfdata/users.json >/dev/null
go run . users role agent --related=false --organizations_path  --tickets_pat  30.60s user 17.50s system 143% cpu 33.515 total

```

For 10K records, lookup rate would be 13 lookups/ms

```
time go run . users role agent --related=false --organizations_path perfdata/organizations.json --tickets_path perfdata/tickets.json --users_path perfdata/users.json >/dev/null
go run . users role agent --related=false --organizations_path  --tickets_pat  3.66s user 1.96s system 149% cpu 3.754 total
<Paste>
```

For thousand records (each file 1000 records) we have time 1.13s. Files are 290K each. Lookup rate 3.5 records/ms.

```
time go run . users role agent --related=false --organizations_path perfdata/organizations.json --tickets_path perfdata/tickets.json --users_path perfdata/users.json >/dev/null
go run . users role agent --related=false --organizations_path  --tickets_pat  1.13s user 0.49s system 153% cpu 1.055 total
```

That is interesting that we can lookup faster if saerch 100K records files because we are building Indexes for invokation of the CLI. Hence for bigger amount of records we see better performance as it takes some time to build
the hash table.

Adding concurrency and statically typing everything can give extra perormance, however here we have used `map[string]interface{}`.
