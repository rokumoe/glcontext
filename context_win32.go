// +build windows

package glcontext

/*
#include "context_win32.h"
#include "goexport.h"

#define GET_LO(d) ((int)(short)LOWORD(d))
#define GET_HI(d) ((int)(short)HIWORD(d))

static int _touched = 0;
static struct wwin _main_win;

static LRESULT WINAPI on_wnd_proc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam)
{
	PAINTSTRUCT ps;

	switch (uMsg) {
	case WM_MOUSEMOVE:
		if (_touched) {
			onTouchMove(GET_LO(lParam), GET_HI(lParam));
		}
		break;
	case WM_SIZE:
		onResize(GET_LO(lParam), GET_HI(lParam));
		break;
	case WM_LBUTTONDOWN:
		if (!_touched) {
			_touched = 1;
			SetCapture(hWnd);
			onTouchDown(GET_LO(lParam), GET_HI(lParam));
		}
		break;
	case WM_LBUTTONUP:
		if (_touched && GetCapture() == hWnd) {
			onTouchUp(GET_LO(lParam), GET_HI(lParam));
			ReleaseCapture();
			_touched = 0;
		}
		break;
	case WM_PAINT:
		BeginPaint(hWnd, &ps);
		onRedraw();
		EndPaint(hWnd, &ps);
		break;
	case WM_SETFOCUS:
		onGotFocus();
		break;
	case WM_KILLFOCUS:
		onLostFocus();
		break;
	case WM_DESTROY:
		PostQuitMessage(0);
		break;
	default:
		return DefWindowProcA(hWnd, uMsg, wParam, lParam);
	}
	return 0;
}

static int create_window(const char *name, int x, int y, int w, int h, struct wwin *ww)
{
	HINSTANCE instance;
	WNDCLASSA wc;

	instance = (HINSTANCE)GetModuleHandle(NULL);
	wc.style = CS_OWNDC;
	wc.lpfnWndProc = on_wnd_proc;
	wc.cbClsExtra = 0;
	wc.cbWndExtra = 0;
	wc.hInstance = instance;
	wc.hIcon = LoadIcon(NULL, IDI_APPLICATION);
	wc.hCursor = LoadCursor(NULL, IDC_ARROW);
	wc.hbrBackground = NULL;
	wc.lpszMenuName = NULL;
	wc.lpszClassName = name;
	if (!RegisterClassA(&wc)) {
		return 1;
	}
	ww->h = CreateWindowA(name, name, WS_OVERLAPPEDWINDOW | WS_CLIPCHILDREN | WS_CLIPSIBLINGS,
						  x, y, w, h, NULL, NULL, instance, NULL);
	if (ww->h == NULL) {
		return 2;
	}
	return 0;
}

static void destroy_window(struct wwin *ww)
{
}

int context_init()
{
	return 0;
}

int context_main_init(const char *name, int x, int y, int w, int h)
{
	int r;
	r = create_window(name, x, y, w, h, &_main_win);
	if (r != 0) {
		return 1;
	}
	ShowWindow(_main_win.h, SW_NORMAL);
	return 0;
}

void context_main_exit()
{
	glc_destroy_context(&_main_win);
	destroy_window(&_main_win);
}

void context_wait_events()
{
	WaitMessage();
}

int context_process_events()
{
	int reason;
	MSG msg;

	reason = 0;
	while (PeekMessageA(&msg, NULL, 0, 0, PM_REMOVE)) {
		if (msg.message == WM_QUIT) {
			reason = 1;
		} else {
			TranslateMessage(&msg);
			DispatchMessageA(&msg);
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
