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
