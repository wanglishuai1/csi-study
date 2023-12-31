package bootstrap

import (
	"csi/pkg/service"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
	"k8s.io/klog"
	"net"
	"os"
)

type MyDriver struct {
	NodeId string
}

func NewMyDriver(nodeId string) *MyDriver {
	return &MyDriver{NodeId: nodeId}
}

func (d *MyDriver) Start() {
	ctlSvc := service.NewControllerService()
	idenSvc := service.NewIdentityService()
	nodeSvc := service.NewNodeService(d.NodeId)

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(DumpLog),
	}

	grpcServer := grpc.NewServer(opts...)
	csi.RegisterControllerServer(grpcServer, ctlSvc)
	csi.RegisterIdentityServer(grpcServer, idenSvc)
	csi.RegisterNodeServer(grpcServer, nodeSvc)

	proto := "unix"
	addr := "/csi/csi.sock"

	if err := os.Remove(addr); err != nil && !os.IsNotExist(err) {
		klog.Fatalf("Failed to remove %s, error: %s", addr, err.Error())
	}
	// TODO 本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
	//把协议 定死为  unix:///csi/csi.sock
	listener, err := net.Listen(proto, addr)
	if err != nil {
		klog.Fatalf("Failed to listen: %v", err)
	}

	grpcServer.Serve(listener)
}
