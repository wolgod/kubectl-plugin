package kubeclient

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// macos
//const kubeconfig = "/Users/test/.kube/config"

// linux
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

func (client *Client) ListPodPage(ctx context.Context) (*v1.PodList, error) {

	allPodList := &v1.PodList{}

	listOptions := metav1.ListOptions{
		Limit: 500, // 设置每批返回的Pod数量
	}
	podList, err := client.Clientset.CoreV1().Pods("").List(ctx, listOptions)
	if err != nil {
		return nil, err
	}
	if len(podList.Items) == 0 {
		return allPodList, nil
	}
	allPodList.Items = append(allPodList.Items, podList.Items...)
	i := 0
	if podList.GetContinue() != "" {
		continueToken := podList.GetContinue()
		// 使用continue字段进行分批查询
		for continueToken != "" {
			i++
			sprintf := fmt.Sprintf("第%d次Fetching more pods...", i)
			fmt.Println(sprintf)
			// 更新ListOptions中的Continue字段
			listOptions.Continue = continueToken

			// 使用更新后的ListOptions获取下一批Pods
			pods, err := client.Clientset.CoreV1().Pods("").List(ctx, listOptions)
			if err != nil {
				panic(err)
			}
			allPodList.Items = append(allPodList.Items, pods.Items...)
			// 更新continueToken
			continueToken = pods.GetContinue()
		}
	}
	return allPodList, nil
}
