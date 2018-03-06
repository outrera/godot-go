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

//func NewConcavePolygonShapeFromPointer(ptr gdnative.Pointer) ConcavePolygonShape {
func NewConcavePolygonShapeFromPointer(ptr gdnative.Pointer) ConcavePolygonShape {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ConcavePolygonShape{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Concave polygon shape resource, which can be set into a [PhysicsBody] or area. This shape is created by feeding a list of triangles.
*/
type ConcavePolygonShape struct {
	Shape
	owner gdnative.Object
}

func (o *ConcavePolygonShape) BaseClass() string {
	return "ConcavePolygonShape"
}

/*
        Return the faces (an array of triangles).
	Args: [], Returns: PoolVector3Array
*/
func (o *ConcavePolygonShape) GetFaces() gdnative.PoolVector3Array {
	//log.Println("Calling ConcavePolygonShape.GetFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConcavePolygonShape", "get_faces")

	// Call the parent method.
	// PoolVector3Array
	retPtr := gdnative.NewEmptyPoolVector3Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector3ArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Set the faces (an array of triangles).
	Args: [{ false faces PoolVector3Array}], Returns: void
*/
func (o *ConcavePolygonShape) SetFaces(faces gdnative.PoolVector3Array) {
	//log.Println("Calling ConcavePolygonShape.SetFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector3Array(faces)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConcavePolygonShape", "set_faces")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}