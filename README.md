# apicl

`apicl` is a CLI tool that helps you to list, fetch and update APIS for APIARY.

[![Build Status](https://travis-ci.org/nephyer/apicl.svg?branch=master)](https://travis-ci.org/nephyer/apicl)

## Description

With `apicl` you can manage Apiary.o

## Installation

You can use Go 1.1X as requirement due to Go Modules. In order to  build this  you will ned:

*Debian/Ubuntu*
```bash
$ sudo apt install  build-essential git make
```

Now check out the git repository: 

```
$ git clone  https://github.com/kdjaldkdj/apicl/
$ cd apicl && make
```

After the tool is buil, you will need to export you APIARY_API_KEY, otherwise the application
will  not start.

```
_build/linux/apicl 
Please, configure your APIARY_API_KEY before using this tool.
```

## Usage

```NAME:
   apicl - Manage apiary

USAGE:
   apicl [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     fetch, f    Fetch API information
     list, l     List APIs.
     publish, p  Publish API
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

Currrent implementation supports 
 * list
 * fetch

## Fetch 

An example fetching a given API. You must use the subdomain name as name.

```
_build/linux/apicl  fetch  --api-name sre3
```

## TODO

Curent implemntation for publishing gets a 400 error, but designed usage would be like this.

```
USAGE:
   apicl publish [command options] [arguments...]

OPTIONS:
   --api-name value, -a value  
   --filename value, -f value
```
Example

`$ apicl publish --api-name publishapi2 --filename publishapi2`

Filename holds your API Blueprint to update.



## License


MIT License

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

