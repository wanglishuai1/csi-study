package service

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type IdentityService struct {
}

func NewIdentityService() *IdentityService {
	return &IdentityService{}
}

var _ csi.IdentityServer = &IdentityService{}

func (i *IdentityService) GetPluginCapabilities(ctx context.Context, request *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	capList := []csi.PluginCapability_Service_Type{
		csi.PluginCapability_Service_CONTROLLER_SERVICE,               //控制器
		csi.PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS, //拓扑约束
	}
	var caps []*csi.PluginCapability
	for _, capObj := range capList {
		c := &csi.PluginCapability{
			Type: &csi.PluginCapability_Service_{
				Service: &csi.PluginCapability_Service{
					Type: capObj,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.GetPluginCapabilitiesResponse{Capabilities: caps}, nil
}

// 健康检查相关
func (i *IdentityService) Probe(ctx context.Context, request *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	status := &wrappers.BoolValue{
		Value: true,
	}
	return &csi.ProbeResponse{
		Ready: status,
	}, nil
}

func (i *IdentityService) GetPluginInfo(ctx context.Context, request *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return &csi.GetPluginInfoResponse{
		Name:          "mycsi.jtthink.com",
		VendorVersion: "v1.0",
	}, nil
}
