# CoreOS backend administrative tools.

[![Build Status](https://travis-ci.org/coreos/core-admin.png)](https://travis-ci.org/coreos/core-admin)

## Building

```
go build
```

## Running and Getting Help

```
./core-admin
```

## Hacking

```
$ mkdir core-admin
$ cd core-admin
$ export GOPATH=`pwd`
$ go get github.com/coreos/core-admin/...
```

# ...hack...hack..hack...

```
$ vim src/github.com/coreos/core-admin/update/types/types.go
```

# rebuild ./bin/core-admin

```
$ go install github.com/coreos/core-admin/...
```
