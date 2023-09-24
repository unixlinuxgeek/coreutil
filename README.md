![Go](https://github.com/unixlinuxgeek/logos/blob/main/100x100/go.png?raw=true) 

### GNU Core Utilities helpers for go 

Installing dependencies:

```shell
go get github.com/unixlinuxgeek/coreutil
```

### Usage

Create app.go fie and insert:
```shell
import github.com/unixlinuxgeek/coreutil

func main() {
  i := coreutil.Installed("ffmpeg")  // i == true
 
  // Do something...
}
```

```shell
$ go run app.go
ffmpeg is installed in "/usr/bin/ffmpeg" 
```
