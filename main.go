package main

import (
	"context"
	"os"

	"fmt"
	"net/http"

	"github.com/rancher/multi-cluster-app/server"
	"github.com/rancher/norman/store/proxy"
	"github.com/rancher/types/config"
	"k8s.io/client-go/tools/clientcmd"
)

var VERSION = "v0.0.0-dev"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return err
	}

	management, err := config.NewManagementContext(*kubeConfig)
	if err != nil {
		return err
	}
	client, err := proxy.NewClientGetterFromConfig(*kubeConfig)
	if err != nil {
		panic(err)
	}
	management.ClientGetter = client

	handler, err := server.NewMultiClusterAppServer(context.Background(), management)
	if err != nil {
		return err
	}

	fmt.Println("Listening on 0.0.0.0:4567")
	return http.ListenAndServe("0.0.0.0:4567", handler)
}
