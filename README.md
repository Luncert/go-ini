# go-ini

ini file parser for golang.


## Example

```
package main

import "ioutils"
import parser "github.com/Luncert/go-ini"

confFile, _ := ioutils.ReadFile("./test.ini")
cfg := parser.ParseIni(string(confFile))
addr := cfg.Section("server").Variable("address").String()
port := cfg.Section("server").Variable("port").Int()
fmt.Printf("server listen on %s:%d", addr, port)
```