package kubeclient

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//macos
//const kubeconfig = "/Users/yiche/.kube/config"

//linux
const kubeconfig = "/root/.kube/config"

type Client struct {
	Clientset *kubernetes.Clientset
}

func NewClient() (*Client, error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}

	client := &Client{
		Clientset: clientset,
	}
	return client, nil
}

func (client *Client) ListNode(ctx context.Context, labels string) (*v1.NodeList, error) {
	nodeList, err := client.Clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		return nil, err
	}
	return nodeList, nil
}

func (client *Client) ListPod(ctx context.Context) (*v1.PodList, error) {
	podList, err := client.Clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return podList, nil
}
