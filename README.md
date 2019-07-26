[![Build Status](https://travis-ci.com/logingood/gooz.svg?branch=master)](https://travis-ci.com/logingood/gooz)

# Description

Go version of Zendesk Coding Challenge, code name GOOZ.

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
# Pluggable backend

Backend implements interface `Store`, you need to implement three methods in order to use interface.
At the moment only Filesystem with json files is implemented as per the task, however you can replace it
with other backends that implements similar interface without requirement to rewrite whole code.
E.g. using sqlx with sqlite/mysql and etc.

# Test coverage

```
ok      github.com/logingood/gooz       (cached)        coverage: 0.0% of statements
ok      github.com/logingood/gooz/cmd   0.013s  coverage: 26.1% of statements
ok      github.com/logingood/gooz/internal/backend      (cached)        coverage: 0.0% of statements
ok      github.com/logingood/gooz/internal/backend/zfile        (cached)        coverage: 100.0% of statements
ok      github.com/logingood/gooz/internal/config       0.018s  coverage: 0.0% of statements
ok      github.com/logingood/gooz/internal/schema       0.016s  coverage: 0.0% of statements
ok      github.com/logingood/gooz/internal/search       (cached)        coverage: 62.3% of statements
```

# Performance
