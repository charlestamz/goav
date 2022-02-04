// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/channel_layout.h>
//#include <libavutil/pixdesc.h>
//#include <stdlib.h>
//#include <errno.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type (
	Options       C.struct_AVOptions
	AvTree        C.struct_AVTree
	Rational      C.struct_AVRational
	MediaType     C.enum_AVMediaType
	AvPictureType C.enum_AVPictureType
	PixelFormat   C.enum_AVPixelFormat
	File          C.FILE
)

const (
	AV_TIME_BASE   = C.AV_TIME_BASE
	AV_NOPTS_VALUE = C.AV_NOPTS_VALUE
)

var AV_TIME_BASE_Q Rational = NewRational(1, AV_TIME_BASE)

const (
	AVMEDIA_TYPE_UNKNOWN    = C.AVMEDIA_TYPE_UNKNOWN
	AVMEDIA_TYPE_VIDEO      = C.AVMEDIA_TYPE_VIDEO
	AVMEDIA_TYPE_AUDIO      = C.AVMEDIA_TYPE_AUDIO
	AVMEDIA_TYPE_DATA       = C.AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_SUBTITLE   = C.AVMEDIA_TYPE_SUBTITLE
	AVMEDIA_TYPE_ATTACHMENT = C.AVMEDIA_TYPE_ATTACHMENT
	AVMEDIA_TYPE_NB         = C.AVMEDIA_TYPE_NB
)

// MediaTypeFromString returns a media type from a string
func MediaTypeFromString(i string) MediaType {
	switch i {
	case "audio":
		return AVMEDIA_TYPE_AUDIO
	case "subtitle":
		return AVMEDIA_TYPE_SUBTITLE
	case "video":
		return AVMEDIA_TYPE_VIDEO
	default:
		return -1
	}
}

const (
	AV_CH_FRONT_LEFT           = C.AV_CH_FRONT_LEFT
	AV_CH_FRONT_RIGHT          = C.AV_CH_FRONT_RIGHT
	AV_CH_LAYOUT_STEREO        = C.AV_CH_LAYOUT_STEREO
	AV_CH_LAYOUT_MONO          = C.AV_CH_LAYOUT_MONO
	AV_CH_LAYOUT_SURROUND      = C.AV_CH_LAYOUT_SURROUND
	AV_CH_LAYOUT_QUAD          = C.AV_CH_LAYOUT_QUAD
	AV_CH_LAYOUT_5POINT0_BACK  = C.AV_CH_LAYOUT_5POINT0_BACK
	AV_CH_LAYOUT_5POINT1_BACK  = C.AV_CH_LAYOUT_5POINT1_BACK
	AV_CH_LAYOUT_6POINT1       = C.AV_CH_LAYOUT_6POINT1
	AV_CH_LAYOUT_7POINT1       = C.AV_CH_LAYOUT_7POINT1
	AV_CH_LAYOUT_HEXADECAGONAL = C.AV_CH_LAYOUT_HEXADECAGONAL
)

const (
	AVERROR_EAGAIN    = -(C.EAGAIN)
	AVERROR_EIO       = -(C.EIO)
	AVERROR_EOF       = C.AVERROR_EOF
	AVERROR_EPERM     = -(C.EPERM)
	AVERROR_EPIPE     = -(C.EPIPE)
	AVERROR_ETIMEDOUT = -(C.ETIMEDOUT)
)

const (
	MAX_AVERROR_STR_LEN        = 255
	MAX_CHANNEL_LAYOUT_STR_LEN = 64
)

const (
	AV_PICTURE_TYPE_NONE = C.AV_PICTURE_TYPE_NONE
	AV_PICTURE_TYPE_I    = C.AV_PICTURE_TYPE_I
	AV_PICTURE_TYPE_B    = C.AV_PICTURE_TYPE_B
	AV_PICTURE_TYPE_P    = C.AV_PICTURE_TYPE_P
)

// AvutilVersion Return the LIBAvUTIL_VERSION_INT constant.
func AvutilVersion() uint {
	return uint(C.avutil_version())
}

// AvutilConfiguration Return the libavutil build-time configuration.
func AvutilConfiguration() string {
	return C.GoString(C.avutil_configuration())
}

// AvutilLicense Return the libavutil license.
func AvutilLicense() string {
	return C.GoString(C.avutil_license())
}

// AvGetMediaTypeString Return a string describing the media_type enum, NULL if media_type is unknown.
func AvGetMediaTypeString(mt MediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

// AvGetPictureTypeChar Return a single letter to describe the given picture type pict_type.
func AvGetPictureTypeChar(pt AvPictureType) string {
	return string(C.av_get_picture_type_char((C.enum_AVPictureType)(pt)))
}

// AvXIfNull Return x default pointer in case p is NULL.
func AvXIfNull(p, x int) {
	C.av_x_if_null(unsafe.Pointer(&p), unsafe.Pointer(&x))
}

// AvIntListLengthForSize Compute the length of an integer list.
func AvIntListLengthForSize(e uint, l int, t uint64) uint {
	return uint(C.av_int_list_length_for_size(C.uint(e), unsafe.Pointer(&l), (C.uint64_t)(t)))
}

// AvFopenUtf8 Open a file using a UTF-8 filename.
func AvFopenUtf8(p, m string) *File {
	cp := C.CString(p)
	defer C.free(unsafe.Pointer(cp))
	cm := C.CString(m)
	defer C.free(unsafe.Pointer(cm))
	f := C.av_fopen_utf8(cp, cm)
	return (*File)(f)
}

// AvGetTimeBaseQ Return the fractional representation of the internal time base.
func AvGetTimeBaseQ() Rational {
	return (Rational)(C.av_get_time_base_q())
}

func AvGetChannelLayoutNbChannels(channelLayout uint64) int {
	return int(C.av_get_channel_layout_nb_channels(C.uint64_t(channelLayout)))
}

func AvGetPixFmtName(pixFmt PixelFormat) string {
	s := C.av_get_pix_fmt_name((C.enum_AVPixelFormat)(pixFmt))
	if s == nil {
		return fmt.Sprintf("unknown pixel format with value %d", pixFmt)
	}
	return C.GoString(s)
}

func AvGetChannelLayoutString(nbChannels int, channelLayout uint64) string {
	bufSize := C.size_t(MAX_CHANNEL_LAYOUT_STR_LEN)
	buf := (*C.char)(C.malloc(bufSize))
	if buf == nil {
		return fmt.Sprintf("unknown channel layout with code %d", channelLayout)
	}
	defer C.free(unsafe.Pointer(buf))
	C.av_get_channel_layout_string(buf, C.int(bufSize), C.int(nbChannels), C.uint64_t(channelLayout))
	return C.GoString(buf)
}

func AvStrerr(errcode int) string {
	errbufSize := C.size_t(MAX_AVERROR_STR_LEN)
	errbuf := (*C.char)(C.malloc(errbufSize))
	if errbuf == nil {
		return fmt.Sprintf("unknown error with code %d", errcode)
	}
	defer C.free(unsafe.Pointer(errbuf))
	ret := C.av_strerror(C.int(errcode), errbuf, errbufSize)
	if ret < 0 {
		return fmt.Sprintf("unknown error with code %d", errcode)
	}
	return C.GoString(errbuf)
}
