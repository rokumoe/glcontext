#ifndef _CONTEXT_WIN32_H_
#define _CONTEXT_WIN32_H_

#define WIN32_LEAN_AND_MEAN
#include <Windows.h>

struct wwin {
	HWND h;
	HDC dc;
	HGLRC glrc;
};

int glc_init();
int glc_create_context(struct wwin *ww);
void glc_destroy_context(struct wwin *ww);
void glc_make_current(struct wwin *ww);
void glc_swap_context(struct wwin *ww);

#endif /* !_CONTEXT_WIN32_H_ */
