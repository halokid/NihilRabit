package main

import (
  "fmt"
  "github.com/hpcloud/tail"
)

func main() {
  //t, err := tail.TailFile("yourlogfile.log", tail.Config{Follow: true})
  //seek := &tail.SeekInfo{
  //  Offset:  0,
  //  Whence:  5,
  //}
  t, err := tail.TailFile(
                         "yourlogfile.log",
                          tail.Config{
                            Follow: false,
                            //MaxLineSize: 10,
                            //Location:  seek,
                         },
                         )
  if err != nil {
    fmt.Println(err)
  }
  for line := range t.Lines {
    //fmt.Println(i)
    fmt.Println(line.Text)
  }
}
