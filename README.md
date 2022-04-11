# Golp

**This app is very much in beta at the moment!**

Golp automates build workflows, compiling SASS and JavaScript into configurable "dist" directories. It also handles dynamic copying of static assets.

Golp is not a Gulp drop-in replacement, but aims to solve many of the same problems that Gulp does. It is fast, simple, and runs from a single binary.


## Usage
```
Usage:
  golp [command]

Available Commands:
  build       Compile & copy your assets (single)
  clean       Clean (delete) your dist directories
  config      View an example config file
  version     Display the current version & update information
  watch       Build & watch src directories for changes

Flags:
      --config string   config file (default "./golp.yaml")
  -h, --help            help for golp
  -v, --verbose         verbose logging
```
