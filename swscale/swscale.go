// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package swscale performs highly optimized image scaling and colorspace and pixel format conversion operations.
//Rescaling: is the process of changing the video size. Several rescaling options and algorithms are available.
//Pixel format conversion: is the process of converting the image format and colorspace of the image.
package swscale

//#cgo pkg-config: libswscale libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <string.h>
//#include <stdint.h>
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"

	"github.com/charlestamz/goav/avutil"
)

type (
	Context C.struct_SwsContext
	Filter  C.struct_SwsFilter
	Vector  C.struct_SwsVector
	Class   C.struct_AVClass
)

const (
	SWS_FAST_BILINEAR = 1
	SWS_BILINEAR      = 2
	SWS_BICUBIC       = 4
	SWS_X             = 8
	SWS_POINT         = 0x10
	SWS_AREA          = 0x20
	SWS_BICUBLIN      = 0x40
	SWS_GAUSS         = 0x80
	SWS_SINC          = 0x100
	SWS_LANCZOS       = 0x200
	SWS_SPLINE        = 0x400
)

//Return the LIBSWSCALE_VERSION_INT constant.
func SwscaleVersion() uint {
	return uint(C.swscale_version())
}

//Return the libswscale build-time configuration.
func SwscaleConfiguration() string {
	return C.GoString(C.swscale_configuration())
}

//Return the libswscale license.
func SwscaleLicense() string {
	return C.GoString(C.swscale_license())
}

//Return a pointer to yuv<->rgb coefficients for the given colorspace suitable for sws_setColorspaceDetails().
func SwsGetcoefficients(c int) *int {
	return (*int)(unsafe.Pointer(C.sws_getCoefficients(C.int(c))))
}

//Return a positive value if pix_fmt is a supported input format, 0 otherwise.
func SwsIssupportedinput(p avutil.PixelFormat) int {
	return int(C.sws_isSupportedInput((C.enum_AVPixelFormat)(p)))
}

//Return a positive value if pix_fmt is a supported output format, 0 otherwise.
func SwsIssupportedoutput(p avutil.PixelFormat) int {
	return int(C.sws_isSupportedOutput((C.enum_AVPixelFormat)(p)))
}

func SwsIssupportedendiannessconversion(p avutil.PixelFormat) int {
	return int(C.sws_isSupportedEndiannessConversion((C.enum_AVPixelFormat)(p)))
}

////Scale the image slice in srcSlice and put the resulting scaled slice in the image in dst.
func SwsScale(ctxt *Context, src [8]*uint8, str [8]int32, y, h int, d [8]*uint8, ds [8]int32) int {
	cctxt := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	csrc := (**C.uint8_t)(unsafe.Pointer(&src[0]))
	cstr := (*C.int)(unsafe.Pointer(&str))
	cd := (**C.uint8_t)(unsafe.Pointer(&d[0]))
	cds := (*C.int)(unsafe.Pointer(&ds))
	return int(C.sws_scale(cctxt, csrc, cstr, C.int(y), C.int(h), cd, cds))
}

func SwsSetcolorspacedetails(ctxt *Context, it *int, sr int, t *int, dr, b, c, s int) int {
	cit := (*C.int)(unsafe.Pointer(it))
	ct := (*C.int)(unsafe.Pointer(t))
	return int(C.sws_setColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, C.int(sr), ct, C.int(dr), C.int(b), C.int(c), C.int(s)))
}

func SwsGetcolorspacedetails(ctxt *Context, it, sr, t, dr, b, c, s *int) int {
	cit := (**C.int)(unsafe.Pointer(it))
	csr := (*C.int)(unsafe.Pointer(sr))
	ct := (**C.int)(unsafe.Pointer(t))
	cdr := (*C.int)(unsafe.Pointer(dr))
	cb := (*C.int)(unsafe.Pointer(b))
	cc := (*C.int)(unsafe.Pointer(c))
	cs := (*C.int)(unsafe.Pointer(s))
	return int(C.sws_getColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, csr, ct, cdr, cb, cc, cs))
}

func SwsGetdefaultfilter(lb, cb, ls, cs, chs, cvs float32, v int) *Filter {
	return (*Filter)(unsafe.Pointer(C.sws_getDefaultFilter(C.float(lb), C.float(cb), C.float(ls), C.float(cs), C.float(chs), C.float(cvs), C.int(v))))
}

func SwsFreefilter(f *Filter) {
	C.sws_freeFilter((*C.struct_SwsFilter)(f))
}

//Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
func SwsConvertpalette8topacked32(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
func SwsConvertpalette8topacked24(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//Get the Class for swsContext.
func SwsGetClass() *Class {
	return (*Class)(C.sws_get_class())
}
