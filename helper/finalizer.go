package helper

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cuityhj/cement/slice"
)

func AddFinalizer(obj metav1.Object, finalizer string) {
	finalizers := obj.GetFinalizers()
	if slice.SliceIndex(finalizers, finalizer) == -1 {
		finalizers = append(finalizers, finalizer)
		obj.SetFinalizers(finalizers)
	}
}

func RemoveFinalizer(obj metav1.Object, finalizer string) {
	finalizers := obj.GetFinalizers()
	slice.SliceRemove(&finalizers, finalizer)
	obj.SetFinalizers(finalizers)
}

func HasFinalizer(obj metav1.Object, finalizer string) bool {
	return slice.SliceIndex(obj.GetFinalizers(), finalizer) != -1
}
