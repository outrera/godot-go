package class

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewHBoxContainerFromPointer(ptr gdnative.Pointer) HBoxContainer {
func NewHBoxContainerFromPointer(ptr gdnative.Pointer) HBoxContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HBoxContainer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Horizontal box container. See [BoxContainer].
*/
type HBoxContainer struct {
	BoxContainer
	owner gdnative.Object
}

func (o *HBoxContainer) BaseClass() string {
	return "HBoxContainer"
}