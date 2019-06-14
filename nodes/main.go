package main

import (
  "flag"
  . "github.com/r00tjimmy/ColorfulRabbit"
)

var logx = &Logx{
  DebugFlag:  true,
}

var (
  debug   bool
)


type NodeIter interface {
  getNodeLog() string
  Notify() bool
  TailLog() string
  HitLog() bool
  getNodeRec() string
}


type Nodes struct {
  mIp           string
  mPort         int
  appName       string
}

func (n *Nodes) getNodeLog() string {
  // get node log file path
  return "./jvmIdx.log"
}

func init() {
  flag.BoolVar(&debug,"debug",false,"debug mode")
}

func main() {
  flag.Parse()
  if debug {
    logx.DebugFlag = true
  }

  n := &Nodes{
    mIp:        "127.0.0.1",
    mPort:      9527,
    appName:    "jvmIdx",
  }

  n.Notify()
}












