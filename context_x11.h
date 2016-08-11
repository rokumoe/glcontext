#ifndef _CONTEXT_X11_H_
#define _CONTEXT_X11_H_

#include <X11/Xlib.h>

#ifndef USE_EGL
#include <GL/glx.h>
#else
#include <EGL/egl.h>
#endif

struct xwin {
    XVisualInfo *vi;
    Window win;
    Colormap cm;
#ifdef USE_EGL
    EGLContext ctx;
    EGLConfig cfg;
    EGLSurface surface;
#else
	GLXContext ctx;
	GLXFBConfig cfg;
#endif
};

int glc_init();
int glc_choose_visual(struct xwin *xw);
int glc_create_context(struct xwin *xw);
void glc_destroy_context(struct xwin *xw);
void glc_make_current(struct xwin *xw);
void glc_swap_context(struct xwin *xw);

extern Display *_display;

#endif /* !_CONTEXT_X11_H_ */
