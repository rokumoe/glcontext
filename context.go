package glcontext

/*
int context_init();
int context_main_init(const char *name, int x, int y, int w, int h);
void context_main_exit();
void context_wait_events();
int context_process_events();
int context_set_current();
void context_swap_buffer();
*/
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"
)

func MainInit(name string, x, y, w, h int) int {
	p := (*reflect.StringHeader)(unsafe.Pointer(&name)).Data
	return int(C.context_main_init((*C.char)(unsafe.Pointer(p)), C.int(x), C.int(y), C.int(w), C.int(h)))
}

func SetCurrent() int {
	return int(C.context_set_current())
}

func MainExit() {
	C.context_main_exit()
}

func MainLoop() int {
	r := 0
	for r == 0 {
		runtime.Gosched()
		C.context_wait_events()
		r = ProcessEvents()
	}
	return r
}

func WaitEvents() {
	C.context_wait_events()
}

func ProcessEvents() int {
	return int(C.context_process_events())
}

func SwapBuffer() {
	C.context_swap_buffer()
}

func init() {
	r := int(C.context_init())
	if r != 0 {
		panic("init failed")
	}
}
