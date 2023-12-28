package tools

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func InitClientSet(kubeConfigPath string) (*kubernetes.Clientset, error) {
	if kubeConfigPath == "" {
		config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
		if err != nil {
			klog.Error(context.Background(), err)
		}
		return kubernetes.NewForConfig(config)
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}
	return kubernetes.NewForConfig(config)
}
