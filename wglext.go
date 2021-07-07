// Copyright 2021 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

import (
	"syscall"
	"unsafe"
)

// wglChoosePixelFormatARB
const (
	WGL_NUMBER_PIXEL_FORMATS_ARB    = 0x2000
	WGL_DRAW_TO_WINDOW_ARB          = 0x2001
	WGL_DRAW_TO_BITMAP_ARB          = 0x2002
	WGL_ACCELERATION_ARB            = 0x2003
	WGL_NEED_PALETTE_ARB            = 0x2004
	WGL_NEED_SYSTEM_PALETTE_ARB     = 0x2005
	WGL_SWAP_LAYER_BUFFERS_ARB      = 0x2006
	WGL_SWAP_METHOD_ARB             = 0x2007
	WGL_NUMBER_OVERLAYS_ARB         = 0x2008
	WGL_NUMBER_UNDERLAYS_ARB        = 0x2009
	WGL_TRANSPARENT_ARB             = 0x200A
	WGL_TRANSPARENT_RED_VALUE_ARB   = 0x2037
	WGL_TRANSPARENT_GREEN_VALUE_ARB = 0x2038
	WGL_TRANSPARENT_BLUE_VALUE_ARB  = 0x2039
	WGL_TRANSPARENT_ALPHA_VALUE_ARB = 0x203A
	WGL_TRANSPARENT_INDEX_VALUE_ARB = 0x203B
	WGL_SHARE_DEPTH_ARB             = 0x200C
	WGL_SHARE_STENCIL_ARB           = 0x200D
	WGL_SHARE_ACCUM_ARB             = 0x200E
	WGL_SUPPORT_GDI_ARB             = 0x200F
	WGL_SUPPORT_OPENGL_ARB          = 0x2010
	WGL_DOUBLE_BUFFER_ARB           = 0x2011
	WGL_STEREO_ARB                  = 0x2012
	WGL_PIXEL_TYPE_ARB              = 0x2013
	WGL_COLOR_BITS_ARB              = 0x2014
	WGL_RED_BITS_ARB                = 0x2015
	WGL_RED_SHIFT_ARB               = 0x2016
	WGL_GREEN_BITS_ARB              = 0x2017
	WGL_GREEN_SHIFT_ARB             = 0x2018
	WGL_BLUE_BITS_ARB               = 0x2019
	WGL_BLUE_SHIFT_ARB              = 0x201A
	WGL_ALPHA_BITS_ARB              = 0x201B
	WGL_ALPHA_SHIFT_ARB             = 0x201C
	WGL_ACCUM_BITS_ARB              = 0x201D
	WGL_ACCUM_RED_BITS_ARB          = 0x201E
	WGL_ACCUM_GREEN_BITS_ARB        = 0x201F
	WGL_ACCUM_BLUE_BITS_ARB         = 0x2020
	WGL_ACCUM_ALPHA_BITS_ARB        = 0x2021
	WGL_DEPTH_BITS_ARB              = 0x2022
	WGL_STENCIL_BITS_ARB            = 0x2023
	WGL_AUX_BUFFERS_ARB             = 0x2024
	WGL_NO_ACCELERATION_ARB         = 0x2025
	WGL_GENERIC_ACCELERATION_ARB    = 0x2026
	WGL_FULL_ACCELERATION_ARB       = 0x2027
	WGL_SWAP_EXCHANGE_ARB           = 0x2028
	WGL_SWAP_COPY_ARB               = 0x2029
	WGL_SWAP_UNDEFINED_ARB          = 0x202A
	WGL_TYPE_RGBA_ARB               = 0x202B
	WGL_TYPE_COLORINDEX_ARB         = 0x202C
	WGL_SAMPLE_BUFFERS_ARB          = 0x2041
	WGL_SAMPLES_ARB                 = 0x2042
)

// wglCreateContextAttribsARB
const (
	WGL_CONTEXT_DEBUG_BIT_ARB                 = 0x0001
	WGL_CONTEXT_FORWARD_COMPATIBLE_BIT_ARB    = 0x0002
	WGL_CONTEXT_MAJOR_VERSION_ARB             = 0x2091
	WGL_CONTEXT_MINOR_VERSION_ARB             = 0x2092
	WGL_CONTEXT_LAYER_PLANE_ARB               = 0x2093
	WGL_CONTEXT_FLAGS_ARB                     = 0x2094
	WGL_CONTEXT_PROFILE_MASK_ARB              = 0x9126
	WGL_CONTEXT_CORE_PROFILE_BIT_ARB          = 0x0001
	WGL_CONTEXT_COMPATIBILITY_PROFILE_BIT_ARB = 0x0002
)

var wglChoosePixelFormatARB uintptr
var wglCreateContextAttribsARB uintptr

// InitWglExt initializes function pointers for WGL extensions. If it is called
// without a current WGL context, all function pointers will be nil.
func InitWglExt() {
	wglChoosePixelFormatARB = WglGetProcAddress(syscall.StringBytePtr("wglChoosePixelFormatARB"))
	wglCreateContextAttribsARB = WglGetProcAddress(syscall.StringBytePtr("wglCreateContextAttribsARB"))
}

func HasWglChoosePixelFormatARB() bool    { return wglChoosePixelFormatARB != 0 }
func HasWglCreateContextAttribsARB() bool { return wglCreateContextAttribsARB != 0 }

func WglChoosePixelFormatARB(hdc HDC, piAttribIList *int32, pfAttribFList *float32, nMaxFormats uint32, piFormats *int32, nNumFormats *uint32) bool {
	ret, _, _ := syscall.Syscall6(wglChoosePixelFormatARB, 6,
		uintptr(hdc),
		uintptr(unsafe.Pointer(piAttribIList)),
		uintptr(unsafe.Pointer(pfAttribFList)),
		uintptr(nMaxFormats),
		uintptr(unsafe.Pointer(piFormats)),
		uintptr(unsafe.Pointer(nNumFormats)))

	return ret != 0
}

func WglCreateContextAttribsARB(hdc HDC, hShareContext HGLRC, attribList *int32) HGLRC {
	ret, _, _ := syscall.Syscall(wglCreateContextAttribsARB, 3,
		uintptr(hdc),
		uintptr(hShareContext),
		uintptr(unsafe.Pointer(attribList)))

	return HGLRC(ret)
}
