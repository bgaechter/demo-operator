package controller

import (
	"github.com/bgaechter/demo-operator/pkg/controller/demo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, demo.Add)
}
