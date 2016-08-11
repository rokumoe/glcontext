// +build linux,!android

package glcontext

// #cgo LDFLAGS: -lX11
import "C"

/*
#include "context_x11.h"
#include "goexport.h"

Display *_display = NULL;

static Atom _wm_delete_window;

static int _inited = 0;
static struct xwin _main_win;
static int _touched = 0;

static int create_window(const char *name, int x, int y, int w, int h, struct xwin *xw)
{
    Window root;
    Window win;
    XSetWindowAttributes swa;

    root = RootWindow(_display, xw->vi->screen);
    xw->cm = XCreateColormap(_display, root, xw->vi->visual, AllocNone);
    swa.colormap = xw->cm;
    swa.border_pixel = 0;
    swa.event_mask = StructureNotifyMask | ExposureMask | ButtonPressMask | ButtonReleaseMask | Button1MotionMask | FocusChangeMask;
    win = XCreateWindow(_display, root, x, y, w, h, 0, xw->vi->depth, InputOutput, xw->vi->visual, CWBackPixel | CWColormap | CWEventMask, &swa);
    if (win == 0) {
        return 1;
    }
    XStoreName(_display, win, name);
    XSetWMProtocols(_display, win, &_wm_delete_window, 1);
    xw->win = win;
    return 0;
}

static void destroy_window(struct xwin *xw)
{
    XUnmapWindow(_display, xw->win);
    XFree(xw->vi);
    XDestroyWindow(_display, xw->win);
    XFreeColormap(_display, xw->cm);
    XFlush(_display);
}

int context_init()
{
    if (_inited) {
        return 0;
    }
    _display = XOpenDisplay(NULL);
    if (_display == NULL) {
        return 1;
    }
    _wm_delete_window = XInternAtom(_display, "WM_DELETE_WINDOW", False);
    if (glc_init() != 0) {
        return 2;
    }
    _inited = 1;
    return 0;
}

int context_main_init(const char *name, int x, int y, int w, int h)
{
    if (glc_choose_visual(&_main_win) != 0) {
        return 1;
    }
    if (create_window(name, x, y, w, h, &_main_win) != 0) {
        return 2;
    }
    XMapWindow(_display, _main_win.win);
    return 0;
}

void context_main_exit()
{
    glc_destroy_context(&_main_win);
    destroy_window(&_main_win);
}

void context_wait_events()
{
    XEvent e;
    XPeekEvent(_display, &e);
}

int context_process_events()
{
    int reason = 0;
    XEvent e;
    int n;

	n = XPending(_display);
    while (n--) {
        XNextEvent(_display, &e);
        switch (e.type) {
		case MotionNotify:
			if (_touched) {
				onTouchMove(e.xmotion.x, e.xmotion.y);
			}
			break;
		case ButtonPress:
			if (!_touched && e.xbutton.button == Button1) {
				_touched = 1;
				onTouchDown(e.xbutton.x, e.xbutton.y);
			}
			break;
		case ButtonRelease:
			if (_touched && e.xbutton.button == Button1) {
				onTouchUp(e.xbutton.x, e.xbutton.y);
				_touched = 0;
			}
			break;
        case ConfigureNotify:
			onResize(e.xconfigure.width, e.xconfigure.height);
            break;
		case Expose:
			onRedraw();
			break;
		case FocusIn:
			onGotFocus();
			break;
		case FocusOut:
			onLostFocus();
			break;
        case DestroyNotify:
            reason = 1;
            break;
        case ClientMessage:
            if ((Atom)e.xclient.data.l[0] == _wm_delete_window) {
                reason = 2;
            }
            break;
        }
    }
    return reason;
}

int context_set_current()
{
	int r;

	r = glc_create_context(&_main_win);
	if (r != 0) {
		return r;
	}
	glc_make_current(&_main_win);
	return 0;
}

void context_swap_buffer()
{
	glc_swap_context(&_main_win);
}
*/
import "C"
