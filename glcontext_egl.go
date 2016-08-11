// +build linux,!android,egl

package glcontext

// #cgo CFLAGS: -DUSE_EGL
// #cgo LDFLAGS: -lEGL
import "C"

/*
#include "context_x11.h"

static const EGLint _config_attrs[] = {
    EGL_RENDERABLE_TYPE, EGL_OPENGL_ES2_BIT,
    EGL_SURFACE_TYPE, EGL_WINDOW_BIT,
    EGL_RED_SIZE, 8,
    EGL_GREEN_SIZE, 8,
    EGL_BLUE_SIZE, 8,
    EGL_ALPHA_SIZE, 8,
    EGL_DEPTH_SIZE, 16,
    EGL_NONE
};

static const EGLint _context_attrs[] = {
    EGL_CONTEXT_CLIENT_VERSION, 2,
    EGL_NONE
};

static EGLDisplay _egldisplay = NULL;

int glc_init()
{
    EGLDisplay egldisplay;

    egldisplay = eglGetDisplay(_display);
    if (egldisplay == NULL) {
        return 1;
    }
    if (eglInitialize(egldisplay, &stub, &stub) == EGL_FALSE) {
        return 2;
    }
    _egldisplay = egldisplay;
    return 0;
}

int glc_choose_visual(struct xwin *xw)
{
	EGLConfig config;
    int config_nums;
	EGLint visualid;
    XVisualInfo vi;
    int vi_nums;
    XVisualInfo *pvi;

    if (eglChooseConfig(_egldisplay, _config_attrs, &config, 1, &config_nums) == EGL_FALSE) {
        return 1;
    }
    if (eglGetConfigAttrib(_egldisplay, config, EGL_NATIVE_VISUAL_ID, &visualid) == EGL_FALSE) {
        return 2;
    }
    vi.visualid = vid;
    pvi = XGetVisualInfo(_display, VisualIDMask, &vi, &vi_nums);
    if (pvi == NULL) {
        return 3;
    }
    xw->cfg = config;
    xw->vi = pvi;
    return 0;
}

int glc_create_context(struct xwin *xw)
{
    xw->surface = eglCreateWindowSurface(_egldisplay, xw->cfg, xw->win, NULL);
    if (xw->surface == NULL) {
        return 1;
    }
    eglBindAPI(EGL_OPENGL_ES_API);
    xw->ctx = eglCreateContext(_egldisplay, xw->cfg, EGL_NO_CONTEXT, _context_attrs);
    if (xw->ctx == NULL) {
        return 2;
    }
    return 0;
}

int glc_make_current(struct xwin *xw)
{
    if ( == EGL_FALSE) {
        return 1;
    }
    return 0;
}

void glc_destroy_context(struct xwin *xw)
{
    if (eglGetCurrentContext() == xw->ctx) {
        eglMakeCurrent(_egldisplay, NULL, NULL, NULL);
    }
    eglDestroyContext(_egldisplay, xw->ctx);
    eglDestroySurface(_egldisplay, xw->surface);
}

void glc_make_current(struct xwin *xw)
{
    eglMakeCurrent(_egldisplay, xw->surface, xw->surface, xw->ctx);
}

void glc_swap_context(struct xwin *xw)
{
    eglSwapBuffers(_egldisplay, xw->surface);
}
*/
import "C"
