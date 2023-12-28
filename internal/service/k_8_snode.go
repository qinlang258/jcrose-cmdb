// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"jcrose-cmdb/api"
	"jcrose-cmdb/internal/model"
)

type (
	IK8SNode interface {
		K8sNodeList(ctx context.Context, req *api.K8sNodeListReq) ([]*model.K8sNodeListOutput, error)
	}
)

var (
	localK8SNode IK8SNode
)

func K8SNode() IK8SNode {
	if localK8SNode == nil {
		panic("implement not found for interface IK8SNode, forgot register?")
	}
	return localK8SNode
}

func RegisterK8SNode(i IK8SNode) {
	localK8SNode = i
}