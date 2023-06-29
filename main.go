package main

import (
	"csi/pkg/bootstrap"
	"flag"
	"k8s.io/klog/v2"
)

var (
	nodeId = ""
)

func main() {
	flag.StringVar(&nodeId, "nodeid", "", "--nodeid=xxx")
	klog.InitFlags(nil)
	flag.Parse()
	driver := bootstrap.NewMyDriver(nodeId)
	driver.Start()
}
