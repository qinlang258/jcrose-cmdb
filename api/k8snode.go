package api

import (
	"jcrose-cmdb/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type K8sNodeListReq struct {
	g.Meta `path:"/api/k8s/node/list" tags:"K8S Node相关" method:"post" summary:"查看所有Node的信息"`
}

type K8sNodeListRes struct {
	Items []*model.K8sNodeListOutput
}

type K8sNodeDescribeReq struct {
	g.Meta   `path:"/api/k8s/node/describenode" tags:"K8S Node相关" method:"post" summary:"查看Node的describe"`
	NodeName string `v:"required" description:"必须输入node的名字"`
}

type K8sNodeDescribeRes struct {
	Items []*model.K8sNodeDescribeOutput
}
