package k8snode

import (
	"context"
	"jcrose-cmdb/api"
	"jcrose-cmdb/internal/model"
	"jcrose-cmdb/internal/service"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func initClientSet() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}
	return kubernetes.NewForConfig(config)
}

type sK8sNode struct {
	k8sClient *kubernetes.Clientset
}

func New() *sK8sNode {
	return &sK8sNode{}
}

func init() {
	node := New()
	clientset, err := initClientSet()
	if err != nil {
		klog.Error(context.TODO(), err)
	}
	node.k8sClient = clientset
	service.RegisterK8SNode(node)
}

func (s *sK8sNode) K8sNodeList(ctx context.Context, req *api.K8sNodeListReq) ([]*model.K8sNodeListOutput, error) {
	nodeList, err := s.k8sClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Error(ctx, err)
	}

	var node model.K8sNodeListOutput
	var k8snodeList []*model.K8sNodeListOutput
	for _, valus := range nodeList.Items {
		node.Name = valus.Name
		for _, values := range valus.Status.Conditions {
			node.Status = string(values.Type)
		}
		node.Version = valus.Status.NodeInfo.KubeletVersion
		node.OsImage = valus.Status.NodeInfo.OSImage
		node.KernelVersion = valus.Status.NodeInfo.KernelVersion
		node.ContainerRuntime = valus.Status.NodeInfo.ContainerRuntimeVersion
	}

	k8snodeList = append(k8snodeList, &node)
	return k8snodeList, err
}
