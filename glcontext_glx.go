// +build linux,!android,!egl

package glcontext

// #cgo LDFLAGS: -lGL
import "C"

/*
#include "context_x11.h"

static int _config_attrs[] = {
	GLX_USE_GL,
	GLX_RGBA,
	GLX_DEPTH_SIZE,
	16,
	GLX_DOUBLEBUFFER,
	None
};

int glc_init()
{
    return 0;
}

int glc_choose_visual(struct xwin *xw)
{
	int stub;
	XVisualInfo *vi;

	if (!glXQueryExtension(_display, &stub, &stub)) {
		return 1;
	}
	vi = glXChooseVisual(_display, DefaultScreen(_display), _config_attrs);
	if (vi == NULL) {
		return 2;
	}
    xw->vi = vi;
    return 0;
}

int glc_create_context(struct xwin *xw)
{
    xw->ctx = glXCreateContext(_display, xw->vi, None, GL_TRUE);
    if (xw->ctx == NULL) {
        return 1;
    }
    return 0;
}

void glc_destroy_context(struct xwin *xw)
{
	glXDestroyContext(_display, xw->ctx);
}

void glc_make_current(struct xwin *xw)
{
	glXMakeCurrent(_display, xw->win, xw->ctx);
}

void glc_swap_context(struct xwin *xw)
{
	glXSwapBuffers(_display, xw->win);
}
*/
import "C"
