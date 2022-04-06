package main

import (
	"context"
	"flag"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/utils/pointer"
	"path/filepath"
)

const (
	NAMESPACE       = "test-clientset"
	DEPLOYMENT_NAME = "client-test-deployment"
	SERVICE_NAME    = "client-test-service"
)

func main() {
	var kubeconfig *string
	// home是家目录，如果能取得家目录的值，就可以用来做默认值
	if home := homedir.HomeDir(); home != "" {
		// 如果输入了kubeconfig参数，该参数的值就是kubeconfig文件的绝对路径，
		// 如果没有输入kubeconfig参数，就用默认路径~/.kube/config
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		// 如果取不到当前用户的家目录，就没办法设置kubeconfig的默认目录了，只能从入参中取
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	operate := flag.String("operate", "create", "operate type: create or clean")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("operate is %v\n", *operate)
	if *operate == "clean" {
		clean(clientset)
	} else {
		createNameSpace(clientset)
		createDeployment(clientset)
		createService(clientset)
	}
}

func clean(clientset *kubernetes.Clientset) {
	emptyDeleteOptions := metav1.DeleteOptions{}
	// delete deployment
	if err := clientset.AppsV1().Deployments(NAMESPACE).Delete(context.TODO(), DEPLOYMENT_NAME, emptyDeleteOptions); err != nil {
		panic(err)
	}
	fmt.Printf("deleted a deployment: %s\n", DEPLOYMENT_NAME)

	// delete service
	if err := clientset.CoreV1().Services(NAMESPACE).Delete(context.TODO(), SERVICE_NAME, emptyDeleteOptions); err != nil {
		panic(err)
	}
	fmt.Printf("deleted a service: %s\n", SERVICE_NAME)

	// delete namespace
	if err := clientset.CoreV1().Namespaces().Delete(context.TODO(), NAMESPACE, emptyDeleteOptions); err != nil {
		panic(err)
	}
	fmt.Printf("deleted a namaspace: %s\n", NAMESPACE)
}

func createNameSpace(clientset *kubernetes.Clientset) {
	namespaceClient := clientset.CoreV1().Namespaces()
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: NAMESPACE,
		},
	}

	result, err := namespaceClient.Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("created a namespace: %s", result.GetName())
}

func createDeployment(clientset *kubernetes.Clientset) {
	deploymentClient := clientset.AppsV1().Deployments(NAMESPACE)
	deploment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: DEPLOYMENT_NAME,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "tomcat",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "tomcat",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            "tomcat",
							Image:           "tomcat:8.0.18-jre8",
							ImagePullPolicy: "IfNotPresent",
							Ports: []v1.ContainerPort{
								{
									Name:          "http",
									Protocol:      v1.ProtocolSCTP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deploymentClient.Create(context.TODO(), deploment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("created a deployment: %s", result.GetName())
}

func createService(clientset *kubernetes.Clientset) {
	serviceClient := clientset.CoreV1().Services(NAMESPACE)
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: SERVICE_NAME,
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name:     "http",
					Port:     8080,
					NodePort: 30080,
				},
			},
			Selector: map[string]string{
				"app": "tomcat",
			},
			Type: v1.ServiceTypeNodePort,
		},
	}

	result, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("created a service: %s\n", result.GetName())
}
