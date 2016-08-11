// +build windows

package glcontext

// #cgo LDFLAGS: -lopengl32 -lgdi32
import "C"

/*
#include "context_win32.h"

int glc_init()
{
	return 0;
}

int glc_create_context(struct wwin *ww)
{
	PIXELFORMATDESCRIPTOR pfd;
	int pf;

	ww->dc = GetDC(ww->h);
	if (ww->dc == NULL) {
		return 1;
	}
	ZeroMemory(&pfd, sizeof(pfd));
	pfd.nSize = sizeof(pfd);
	pfd.nVersion = 1;
	pfd.dwFlags = PFD_DRAW_TO_WINDOW | PFD_SUPPORT_OPENGL | PFD_DOUBLEBUFFER;
	pfd.iPixelType = PFD_TYPE_RGBA;
	pfd.cColorBits = 32;
	pfd.iLayerType = PFD_MAIN_PLANE;
	pf = ChoosePixelFormat(ww->dc, &pfd);
	if (pf == 0) {
		return 2;
	}
	if (!SetPixelFormat(ww->dc, pf, &pfd)) {
		return 3;
	}
	ww->glrc = wglCreateContext(ww->dc);
	if (ww->glrc == NULL) {
		return 4;
	}
	return 0;
}

void glc_destroy_context(struct wwin *ww)
{
	wglMakeCurrent(NULL, NULL);
	wglDeleteContext(ww->glrc);
	ReleaseDC(ww->h, ww->dc);
}

void glc_make_current(struct wwin *ww)
{
	wglMakeCurrent(ww->dc, ww->glrc);
}

void glc_swap_context(struct wwin *ww)
{
	SwapBuffers(ww->dc);
}
*/
import "C"
