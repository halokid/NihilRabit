package main

import (
  "github.com/fsnotify/fsnotify"
  . "github.com/r00tjimmy/ColorfulRabbit"
)

func (n *Nodes) Notify() {
  // 监控log文件的变化
  watch, err := fsnotify.NewWatcher()
  CheckError(err)

  defer watch.Close()

  logFile := n.getNodeLog()
  err = watch.Add(logFile)
  CheckError(err)

  // 开启一个 goroutine 来处理监控对象事件
  go func() {
    for {
      select {
      case ev := <- watch.Events:
        {
          if ev.Op & fsnotify.Write == fsnotify.Write {
            logx.DebugPrint("write log: ", ev.Name)
          }
        }
      case err := <- watch.Errors:
         {
           logx.DebugPrint("error: ", err)
           return
         }
      }
    }
  }()

  // 循环这个函数的主gorontine, 那么由此函数派生的gorontine都存在
  select {}
}











