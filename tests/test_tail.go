package main

import (
  "io"
  "fmt"
  "os"

  "github.com/papertrail/go-tail/follower"
)

func main() {
  t, err := follower.New("yourlogfile.log", follower.Config{
    Whence: io.SeekEnd,
    Offset: 0,
    Reopen: true,
  })

  if err != nil {
    fmt.Println(err)
  }

  for line := range t.Lines() {
    //fmt.Println(line)
    fmt.Println(line.String())
  }

  if t.Err() != nil {
    fmt.Fprintln(os.Stderr, t.Err())
  }
}

