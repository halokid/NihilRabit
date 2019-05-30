package nodes

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

}