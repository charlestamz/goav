// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"

	"github.com/charlestamz/goav/avutil"
)

//Utility function to access log2_chroma_w log2_chroma_h from the pixel format AvPixFmtDescriptor.
func AvcodecGetChromaSubSample(pix_fmt avutil.PixelFormat, h, v *int) {
	panic("deprecated")
	//C.avcodec_get_chroma_sub_sample((C.enum_AVPixelFormat)(pix_fmt), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(v)))
}

//Return a value representing the fourCC code associated to the pixel format pix_fmt, or 0 if no associated fourCC code can be found.
func AvcodecPixFmtToCodecTag(pix_fmt avutil.PixelFormat) uint {
	return uint(C.avcodec_pix_fmt_to_codec_tag((C.enum_AVPixelFormat)(pix_fmt)))
}

//Find the best pixel format to convert to given a certain source pixel format.
func AvcodecFindBestPixFmtOfList(pix_fmt_list *avutil.PixelFormat, src_pix_fmt avutil.PixelFormat, a int, l *int) avutil.PixelFormat {
	return (avutil.PixelFormat)(C.avcodec_find_best_pix_fmt_of_list((*C.enum_AVPixelFormat)(pix_fmt_list), (C.enum_AVPixelFormat)(src_pix_fmt), C.int(a), (*C.int)(unsafe.Pointer(l))))
}

func AvcodecFindBestPixFmtOf2(dst1, dst2, src avutil.PixelFormat, a int, l *int) avutil.PixelFormat {
	return (avutil.PixelFormat)(C.avcodec_find_best_pix_fmt_of_2((C.enum_AVPixelFormat)(dst1), (C.enum_AVPixelFormat)(dst2), (C.enum_AVPixelFormat)(src), C.int(a), (*C.int)(unsafe.Pointer(l))))
}
