package controller

import (
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/cuityhj/gok8s/handler"
	"github.com/cuityhj/gok8s/predicate"
)

type Controller interface {
	Watch(obj runtime.Object) error
	Start(stop <-chan struct{}, handler handler.EventHandler, predicates ...predicate.Predicate)
}
