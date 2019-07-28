[![Build Status](https://travis-ci.com/logingood/gooz.svg?branch=master)](https://travis-ci.com/logingood/gooz)

# Description

Go version of Zendesk Coding Challenge, code name GOOZ.

To build cli we have used [spf13 cobra](https://github.com/spf13/cobra)

* Search is organized using basic hash table on maps
* Hash table takes interface{} as input
* Search module attempts to detect types and makes everything `string`
* Table module draw tables and also detects type
* We validate input against provided schema for safety reasons

# Challenge check list

- [x] Separation of concerns: every module is implemented as independent package (e.g. `index`, `search`, `table`, `cli` and etc)
- [x] Simplicity: we used the most basic type of data structure - hash table and list to organize search index
- [x] Test coverage: every module is tested separate, test coverage provided. CLI tests are in progress
- [ ] Performance analysis TBD
- [x] Errors are handled and logged, corresponding tests provided


# Dockerfile

The easiest way to build and run the code is a dockek

```
git clone https://github.com/logingood/gooz
cd gooz

docker build . -t gooz
```

Run the container
```
docker rm -f gooz >/dev/null 2>&1;  docker run -ti --name gooz gooz:latest /bin/gooz --help
```

# CLI

Search tickets by any field, e.g. subject

```
gooz tickets subject "A Problem in South Africa"
```

Search organization by any field, e.g.

```
gooz organizations created_at  "2016-05-21T11:10:28 -10:00"
```

Search users by any field
```
gooz users alias "Miss Rosanna"`
```


# Pluggable backend

Backend implements interface `Store`, you need to implement three methods in order to use interface.
At the moment only Filesystem with json files is implemented as per the task, however you can replace it
with other backends that implements similar interface without requirement to rewrite whole code.
E.g. using sqlx with sqlite/mysql and etc.

# Test coverage

```
ok      github.com/logingood/gooz       0.023s  coverage: 0.0% of statements
ok      github.com/logingood/gooz/cmd   0.024s  coverage: 19.8% of statements
ok      github.com/logingood/gooz/internal/backend      0.023s  coverage: 0.0% of statements
ok      github.com/logingood/gooz/internal/backend/zfile        0.028s  coverage: 94.7% of statements
ok      github.com/logingood/gooz/internal/index        0.023s  coverage: 100.0% of statements
ok      github.com/logingood/gooz/internal/search       0.015s  coverage: 87.0% of statements
ok      github.com/logingood/gooz/internal/table        0.014s  coverage: 0.0% of statements
```

# Performance
