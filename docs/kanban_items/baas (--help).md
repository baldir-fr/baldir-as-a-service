- create a CLI
- basic feature : display help

## Acceptance criterias

An executable shell command that provides the following result.

```shell
go build
```

```sh
$ ./baas

BaaS is a CLI tool to interact with up-to-date professional informations about Marc Bouvier.
Such informations are : resume, contact infos, news, curated lists, tip of the day...

Usage:  
  baas [command]

Available Commands:

Flags:
  -h, --help             help for baas

Use "baas [command] --help" for more information about a command.


```

Alternate commands for the same thing
```shell
./baas --help
./baas -h
```