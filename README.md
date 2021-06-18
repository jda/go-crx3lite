# go-crx3lite
A stripped down respin of [mmadfox/go-crx3](https://github.com/mmadfox/go-crx3) 
with support for reading existing crx3 files and processing them in memory instead 
of being based around reading/writing files on disk.

[![Documentation](https://godoc.org/github.com/jda/go-crx3lite?status.svg)](https://pkg.go.dev/github.com/jda/go-crx3lite)

Provides a sets of tools packing, unpacking, zip, unzip, download, gen id, etc...

## Table of contents
+ [Installation](#installation)
+ [Commands](#commands)
+ [Examples](#examples)
  - [Unpack chrome extension into current directory](#unpack)
  - [Download a chrome extension from the web store](#download)
  - [Unzip an extension to the directory](#unzip)
  - [IsDir, IsZip, IsCRX3 helpers](#isdir-iszip-iscrx3)
  - [Load or save private key](#newprivatekey-loadprivatekey-saveprivatekey)
+ [License](#license)

### Commands
```shell script
make proto 
make covertest
``` 

### Examples

#### Unpack
##### Unpack chrome extension into current directory
```go
import crx3 "github.com/jda/go-crx3"

if err := crx3.Extension("/path/to/ext.crx").Unpack(); err != nil {
   panic(err)
}
```
```shell script
$ crx3 unpack /path/to/ext.crx 
```

#### Download 
##### Download a chrome extension from the web store
```go
import crx3 "github.com/jda/go-crx3"

extensionID := "blipmdconlkpinefehnmjammfjpmpbjk"
filepath := "/path/to/ext.crx"
if err := crx3.DownloadFromWebStore(extensionID,filepath); err != nil {
    panic(err)
}
```
```shell script
$ crx3 download blipmdconlkpinefehnmjammfjpmpbjk [-o /custom/path]
$ crx3 download https://chrome.google.com/webstore/detail/lighthouse/blipmdconlkpinefehnmjammfjpmpbjk
```

#### Unzip
##### Unzip an extension to the current directory
```go
import crx3 "github.com/jda/go-crx3"

if err := crx3.Extension("/path/to/ext.zip").Unzip(); err != nil {
    panic(err)
}
```
```shell script
$ crx3 unzip /path/to/ext.zip [-o /custom/path] 
``` 

## License
go-crx3lite is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/jda/go-crx3/blob/master/LICENSE)

### Credits
go-crx3lite is derived from go-crx3 by MediaBuyerBot which was released under the Apache 2.0 license.
