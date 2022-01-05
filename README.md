# Golang file probe package

Probe util with file for golang

## Install

```shell
go get github.com/dollarsignteam/go-fileprobe
```

## Usage

Get datetime before or after specific datetime

```go
import (
  "log"

  "github.com/dollarsignteam/go-fileprobe"
)

func main() {
  fp := fileprobe.NewHandler()
  if err := fp.Create(); err != nil {
    log.Fatal(err)
  }
  defer fp.Remove()
}
```

## Author

Dollarsign

## License

Licensed under the MIT License - see the [LICENSE][1] file for details.

[1]: https://github.com/dollarsignteam/go-fileprobe/blob/main/LICENSE
