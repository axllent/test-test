# Golp

**This app is very much in beta at the moment!**

Golp automates build workflows, compiling SASS and JavaScript into configurable "dist" directories. It also handles dynamic copying of static assets.

Golp is not a Gulp drop-in replacement, but aims to solve many of the same problems that Gulp does. It is fast, simple, and runs from a single binary.

Internally it uses [esbuild](https://github.com/evanw/esbuild) & [golibsass](https://github.com/bep/golibsass) to compile JavaScript & SASS.


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
  -c  --config string   config file (default "./golp.yaml")
  -h, --help            help for golp
  -v, --verbose         verbose logging
```

## Installation

There are some pre-build binaries available for Linux and MacOS available in the releases.

Golp relies on CGO for the [golibsass](https://github.com/bep/golibsass) compilation, making cross-platform / multi-arch pre-built binaries very challenging.

If your system has go, gcc & g++ installed, you can install it easily from source with: `go install github.com/axllent/golp@latest`


## Config file

Typically your config file will be found in your project root folder, and named `golp.yaml`. An alternative config can be specified using the `-c` flag.

Please note that all `styles`, `scripts` and `copy` src files are relative to your config file.

Run `golp config` to view an example config file.


### Example config file

```yaml
# If directories are specified, they will be deleted each time when golp is run.
# Note that individual `dist` direectories will be deleted automatically too.
clean: 
  - themes/site/dist # optional

# SASS & CSS files
styles:
  - src:
      # process all *.scss files in this folder
      - themes/site/src/sass/*.scss
      # process all *.css files in this folder and child folders
      - themes/site/src/sass/**.css 
      # add a specific file
      - node_modules/@dashboardcode/bsmultiselect/dist/css/BsMultiSelect.css
    # output directory for all src files
    dist: themes/site/dist/css/

# JavaScript files
scripts:
  - src:
      - node_modules/@popperjs/core/dist/umd/popper.min.js
      - node_modules/bootstrap/dist/js/bootstrap.min.js
      - node_modules/axios/dist/axios.min.js
      - node_modules/@dashboardcode/bsmultiselect/dist/js/BsMultiSelect.min.js
      - node_modules/vuedraggable/dist/vuedraggable.umd.min.js
      - node_modules/sortablejs/Sortable.min.js
      - node_modules/vue/dist/vue.global.prod.js
    # merge all files into a single file (only scripts & styles supported)
    dist: themes/site/dist/js/libs.js 
    # optional name for the console output
    name: libs

  - src:
      - themes/site/src/js/**.js
    dist: themes/site/dist/js
    name: site scripts

# All other files
copy:
  - src:
      - themes/site/src/images/**
    dist: themes/site/dist/images
    name: images

  - src: 
      - themes/site/src/fonts/**
    dist: themes/site/dist/fonts/
    name: fonts
```
