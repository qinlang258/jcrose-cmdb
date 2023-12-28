package controller

import (
	"context"
	"jcrose-cmdb/api"
	"jcrose-cmdb/internal/service"
	"k8s.io/klog/v2"
)

type cK8snode struct {
}

var K8snode = cK8snode{}

func (c *cK8snode) K8sNodeList(ctx context.Context, req *api.K8sNodeListReq) (*api.K8sNodeListRes, error) {
	nodeList, err := service.K8SNode().K8sNodeList(ctx, req)
	if err != nil {
		klog.Error(ctx, err)
	}

	res := &api.K8sNodeListRes{
		Items: nodeList,
	}

	return res, err
}
