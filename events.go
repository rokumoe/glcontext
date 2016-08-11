package glcontext

import "C"

type (
	RedrawFunc func()
	FocusFunc  func()
	ResizeFunc func(w int32, h int32)
	TouchFunc  func(x int32, y int32)
)

var (
	OnRedraw    RedrawFunc
	OnResize    ResizeFunc
	OnGotFocus  FocusFunc
	OnLostFocus FocusFunc
	OnTouchDown TouchFunc
	OnTouchMove TouchFunc
	OnTouchUp   TouchFunc
)

//export onResize
func onResize(w, h int32) {
	f := OnResize
	if f != nil {
		f(w, h)
	}
}

//export onGotFocus
func onGotFocus() {
	f := OnGotFocus
	if f != nil {
		f()
	}
}

//export onLostFocus
func onLostFocus() {
	f := OnLostFocus
	if f != nil {
		f()
	}
}

//export onRedraw
func onRedraw() {
	f := OnRedraw
	if f != nil {
		f()
	}
}

//export onTouchDown
func onTouchDown(x, y int32) {
	f := OnTouchDown
	if f != nil {
		f(x, y)
	}
}

//export onTouchMove
func onTouchMove(x, y int32) {
	f := OnTouchMove
	if f != nil {
		f(x, y)
	}
}

//export onTouchUp
func onTouchUp(x, y int32) {
	f := OnTouchUp
	if f != nil {
		f(x, y)
	}
}
