# RICDMS


## Building

### Local build and Run

Chekout code for [RICDMS](https://gerrit.o-ran-sc.org/r/admin/repos/ric-plt/ricdms) repository from gerrit.
```sh
$ git clone ssh://subhash_singh@gerrit.o-ran-sc.org:29418/ric-plt/ricdms
```

build locally
```sh
$ make build
```

Run the executable
```sh
$./ricdms
{"ts":1684321663015,"crit":"INFO","id":"ricdms","mdc":{},"msg":"Logger is initialized without config file()."}
{"ts":1684321663023,"crit":"INFO","id":"ricdms","mdc":{},"msg":"Starting server at : 0.0.0.0:8000"}
2023/05/17 11:07:43 Serving r i c d m s at http://[::]:8000
```

It will start the RICDMS on port `:8000`

### 