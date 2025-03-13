package watcher

import (
	"github.com/moweilong/go-example/nightwatch/pkg/store"
	"k8s.io/client-go/kubernetes"
)

type Config struct {
	Store store.IStore

	Clientset kubernetes.Interface
}
