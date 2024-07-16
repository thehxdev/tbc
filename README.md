# TBC
TBC is a CLI client for [Txtban](https://github.com/thehxdev/txtban) services.


## Build
You have to install go compiler GNU make. Then run:
```bash
make
```
This command will build `tbc`.

### Install
```bash
make install PREFIX=/usr/local/bin
```
Command above will install `tbc` executable to the directory specified by `PREFIX` variable.


## Usage
Use `help` command-line flag to get a help message:
```bash
./tbc help
```

1. Create a `config.json` file and set `TBC_CONF` environment variable equal to full path to the `config.json` file.
You can use [the example config file](config.example.json). (You can leave `user` fields empty)
2. If you don't have an account use `tbc useradd -password <PASSWORD>` command to signup.
3. Then you can use `tbc` to interact with txtban server.
