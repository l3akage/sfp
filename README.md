# SFP

Tool to read/write SFP informations/firmware

## Install
```bash
go get -u github.com/l3akage/sfp
```

## Usage
```bash
Small Form-factor Pluggable Inter-Integrated Circuit Bus Reader/Writer

Usage:
  sfp [flags]
  sfp [command]

Available Commands:
  crack       Crack SFP password
  dump        Dump firmware
  fixcrc      Fix Checksum
  help        Help about any command
  info        Read SFP
  version     Print the version number of sfp

Flags:
      --debug           Enable debugging messages
      --device string   path to i2c device (default is /dev/i2c-1) (default "/dev/i2c-1")
  -h, --help            help for sfp

Use "sfp [command] --help" for more information about a command.
```

## License
(c) Martin Poppen, 2019. Licensed under [MIT](LICENSE) license.
