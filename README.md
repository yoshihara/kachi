# kachi

A simple app for time tracking.


## How to use

```sh
$ kachi init # create ~/.kachi directory
$ kachi start taskName
$ kachi current
$ kachi stop
$ kachi stats -s 1.5 # -s option specifies the scale for estimated time of each task
(at the end of day)
$ kachi refresh
```

## How to develop

```sh
$ git clone
$ cd repository
$ make bundle_install
$ go run main.go
$ make build_* # choose your target OS (darwin, windows, linux)
```
