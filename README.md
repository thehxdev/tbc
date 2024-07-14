# TBC
TBC is a command-line client for [Txtban](https://github.com/thehxdev/txtban) server application.


## Build
You have to install go compiler. Then run:
```bash
go build -ldflags='-s' .
```
This command will produce `tbc` executable that you can use.


## Usage
Use `help` command-line flag to get a help message:
```bash
./tbc help
```

1. Create a `config.json` file and set `TBC_CONF` environment variable equal to full path to the `config.json` file.
(There is an example `config.json` file in [config.example.json](config.example.json))
2. If you already signed up, fill the empty strings in then example config above.
3. Then you can use `tbc` to interact with txtban server.
4. If you don't have an account use `tbc useradd -password <PASSWORD>` command to signup.
