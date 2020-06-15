package engine

import (
	"container/list"
)

var resources *list.List = list.New()

type Resource interface {
	release()
}

func addResource(_resource Resource) {
	resources.PushBack(_resource)
}

func freeResources() {
	for i := resources.Front(); i != nil; i = i.Next() {
		resource, valid := i.Value.(Resource)

		if valid {
			resource.release()
		}
	}
}
