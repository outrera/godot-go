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

//func NewEditorScriptFromPointer(ptr gdnative.Pointer) EditorScript {
func NewEditorScriptFromPointer(ptr gdnative.Pointer) EditorScript {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorScript{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Scripts extending this class and implementing its [code]_run()[/code] method can be executed from the Script Editor's [code]File -> Run[/code] menu option (or by pressing [code]CTRL+Shift+X[/code]) while the editor is running. This is useful for adding custom in-editor functionality to Godot. For more complex additions, consider using [EditorPlugin]s instead. Note that extending scripts need to have [code]tool mode[/code] enabled. Example script: [codeblock] tool extends EditorScript func _run(): print("Hello from the Godot Editor!") [/codeblock] Note that the script is run in the Editor context, which means the output is visible in the console window started with the Editor (STDOUT) instead of the usual Godot [i]Output[/i] dock.
*/
type EditorScript struct {
	Reference
	owner gdnative.Object
}

func (o *EditorScript) BaseClass() string {
	return "EditorScript"
}

/*
        This method is executed by the Editor when [code]File -> Run[/code] is used.
	Args: [], Returns: void
*/
func (o *EditorScript) X_Run() {
	//log.Println("Calling EditorScript.X_Run()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorScript", "_run")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds [code]node[/code] as a child of the root node in the editor context. WARNING: The implementation of this method is currently disabled.
	Args: [{ false node Object}], Returns: void
*/
func (o *EditorScript) AddRootNode(node Object) {
	//log.Println("Calling EditorScript.AddRootNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorScript", "add_root_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the [EditorInterface] singleton instance.
	Args: [], Returns: EditorInterface
*/
func (o *EditorScript) GetEditorInterface() EditorInterface {
	//log.Println("Calling EditorScript.GetEditorInterface()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorScript", "get_editor_interface")

	// Call the parent method.
	// EditorInterface
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewEditorInterfaceFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the Editor's currently active scene.
	Args: [], Returns: Node
*/
func (o *EditorScript) GetScene() Node {
	//log.Println("Calling EditorScript.GetScene()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorScript", "get_scene")

	// Call the parent method.
	// Node
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewNodeFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}