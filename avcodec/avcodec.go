// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libswresample
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/frame.h>
//void register_codecs(){
//#if LIBAVCODEC_VERSION_MAJOR < 58
//	  av_register_all();
//	  avcodec_register_all();
//#endif
//}
//
//
//
import "C"
import (
	"unsafe"

	"github.com/charlestamz/goav/avutil"
)

type (
	Codec                         C.struct_AVCodec
	Context                       C.struct_AVCodecContext
	CodecParameters               C.struct_AVCodecParameters
	Descriptor                    C.struct_AVCodecDescriptor
	Parser                        C.struct_AVCodecParser
	ParserContext                 C.struct_AVCodecParserContext
	Frame                         C.struct_AVFrame
	MediaType                     C.enum_AVMediaType
	Packet                        C.struct_AVPacket
	BitStreamFilter               C.struct_AVBitStreamFilter
	BitStreamFilterContext        C.struct_AVBitStreamFilterContext
	Rational                      C.struct_AVRational
	Class                         C.struct_AVClass
	AvHWAccel                     C.struct_AVHWAccel
	AvPacketSideData              C.struct_AVPacketSideData
	AvPanScan                     C.struct_AVPanScan
	Picture                       C.struct_AVPicture
	AvProfile                     C.struct_AVProfile
	AvSubtitle                    C.struct_AVSubtitle
	AvSubtitleRect                C.struct_AVSubtitleRect
	RcOverride                    C.struct_RcOverride
	AvBufferRef                   C.struct_AVBufferRef
	AvAudioServiceType            C.enum_AVAudioServiceType
	AvChromaLocation              C.enum_AVChromaLocation
	CodecId                       C.enum_AVCodecID
	AvColorPrimaries              C.enum_AVColorPrimaries
	AvColorRange                  C.enum_AVColorRange
	AvColorSpace                  C.enum_AVColorSpace
	AvColorTransferCharacteristic C.enum_AVColorTransferCharacteristic
	AvDiscard                     C.enum_AVDiscard
	AvFieldOrder                  C.enum_AVFieldOrder
	AvPacketSideDataType          C.enum_AVPacketSideDataType
	AvSampleFormat                C.enum_AVSampleFormat
)

const (
	AV_PKT_FLAG_KEY     = C.AV_PKT_FLAG_KEY
	AV_PKT_FLAG_CORRUPT = C.AV_PKT_FLAG_CORRUPT
	AV_PKT_FLAG_DISCARD = C.AV_PKT_FLAG_DISCARD
)

const (
	AVDISCARD_NONE     = C.AVDISCARD_NONE     // discard nothing
	AVDISCARD_DEFAULT  = C.AVDISCARD_DEFAULT  // discard useless packets like 0 size packets in avi
	AVDISCARD_NONREF   = C.AVDISCARD_NONREF   // discard all non reference
	AVDISCARD_BIDIR    = C.AVDISCARD_BIDIR    // discard all bidirectional frames
	AVDISCARD_NONINTRA = C.AVDISCARD_NONINTRA // discard all non intra frames
	AVDISCARD_NONKEY   = C.AVDISCARD_NONKEY   // discard all frames except keyframes
	AVDISCARD_ALL      = C.AVDISCARD_ALL      // discard all
)

const (
	AVMEDIA_TYPE_ATTACHMENT = C.AVMEDIA_TYPE_ATTACHMENT
	AVMEDIA_TYPE_AUDIO      = C.AVMEDIA_TYPE_AUDIO
	AVMEDIA_TYPE_DATA       = C.AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_NB         = C.AVMEDIA_TYPE_NB
	AVMEDIA_TYPE_SUBTITLE   = C.AVMEDIA_TYPE_SUBTITLE
	AVMEDIA_TYPE_UNKNOWN    = C.AVMEDIA_TYPE_UNKNOWN
	AVMEDIA_TYPE_VIDEO      = C.AVMEDIA_TYPE_VIDEO
)

const (
	AV_PKT_DATA_PALETTE                    = C.AV_PKT_DATA_PALETTE
	AV_PKT_DATA_NEW_EXTRADATA              = C.AV_PKT_DATA_NEW_EXTRADATA
	AV_PKT_DATA_PARAM_CHANGE               = C.AV_PKT_DATA_PARAM_CHANGE
	AV_PKT_DATA_H263_MB_INFO               = C.AV_PKT_DATA_H263_MB_INFO
	AV_PKT_DATA_REPLAYGAIN                 = C.AV_PKT_DATA_REPLAYGAIN
	AV_PKT_DATA_DISPLAYMATRIX              = C.AV_PKT_DATA_DISPLAYMATRIX
	AV_PKT_DATA_STEREO3D                   = C.AV_PKT_DATA_STEREO3D
	AV_PKT_DATA_AUDIO_SERVICE_TYPE         = C.AV_PKT_DATA_AUDIO_SERVICE_TYPE
	AV_PKT_DATA_SKIP_SAMPLES               = C.AV_PKT_DATA_SKIP_SAMPLES
	AV_PKT_DATA_JP_DUALMONO                = C.AV_PKT_DATA_JP_DUALMONO
	AV_PKT_DATA_STRINGS_METADATA           = C.AV_PKT_DATA_STRINGS_METADATA
	AV_PKT_DATA_SUBTITLE_POSITION          = C.AV_PKT_DATA_SUBTITLE_POSITION
	AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL   = C.AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL
	AV_PKT_DATA_WEBVTT_IDENTIFIER          = C.AV_PKT_DATA_WEBVTT_IDENTIFIER
	AV_PKT_DATA_WEBVTT_SETTINGS            = C.AV_PKT_DATA_WEBVTT_SETTINGS
	AV_PKT_DATA_METADATA_UPDATE            = C.AV_PKT_DATA_METADATA_UPDATE
	AV_PKT_DATA_MPEGTS_STREAM_ID           = C.AV_PKT_DATA_MPEGTS_STREAM_ID
	AV_PKT_DATA_MASTERING_DISPLAY_METADATA = C.AV_PKT_DATA_MASTERING_DISPLAY_METADATA
	AV_PKT_DATA_CONTENT_LIGHT_LEVEL        = C.AV_PKT_DATA_CONTENT_LIGHT_LEVEL
	AV_PKT_DATA_SPHERICAL                  = C.AV_PKT_DATA_SPHERICAL
	AV_PKT_DATA_A53_CC                     = C.AV_PKT_DATA_A53_CC
)

const (
	FF_THREAD_FRAME = C.FF_THREAD_FRAME
	FF_THREAD_SLICE = C.FF_THREAD_SLICE
)

func (c *Codec) AvCodecGetMaxLowres() int {
	panic("deprecated")
	return 0
	//return int(C.av_codec_get_max_lowres((*C.struct_AVCodec)(c)))
}

//If c is NULL, returns the first registered codec, if c is non-NULL,
func (c *Codec) AvCodecNext() *Codec {
	panic("deprecated")
	return nil
	//return (*Codec)(C.av_codec_next((*C.struct_AVCodec)(c)))
}

//Register all the codecs, parsers and bitstream filters which were enabled at configuration time.
func AvcodecRegisterAll() {
	//for ffmpeg version < 4, not to call deprecated functions
	C.register_codecs()
}

//Register the codec codec and initialize libavcodec.
func (c *Codec) AvcodecRegister() {
	panic("deprecated")
	//C.avcodec_register((*C.struct_AVCodec)(c))
}

//Return a name for the specified profile, if available.
func (c *Codec) AvGetProfileName(p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

//Allocate an Context and set its fields to default values.
func (c *Codec) AvcodecAllocContext3() *Context {
	return (*Context)(unsafe.Pointer(C.avcodec_alloc_context3((*C.struct_AVCodec)(c))))
}

func (c *Codec) AvCodecIsEncoder() int {
	return int(C.av_codec_is_encoder((*C.struct_AVCodec)(c)))
}

func (c *Codec) AvCodecIsDecoder() int {
	return int(C.av_codec_is_decoder((*C.struct_AVCodec)(c)))
}

func (c *Codec) SupportedSamplerates() []int {
	r := make([]int, 0)
	if c.supported_samplerates == nil {
		return r
	}
	size := unsafe.Sizeof(*c.supported_samplerates)
	for i := 0; ; i++ {
		p := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(c.supported_samplerates)) + uintptr(i)*size))
		if p == 0 {
			break
		}
		r = append(r, int(p))
	}
	return r
}

func (c *Codec) SampleFmts() []AvSampleFormat {
	r := make([]AvSampleFormat, 0)
	if c.sample_fmts == nil {
		return r
	}
	size := unsafe.Sizeof(*c.sample_fmts)
	for i := 0; ; i++ {
		p := *(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(c.sample_fmts)) + uintptr(i)*size))
		if p == C.AV_SAMPLE_FMT_NONE {
			break
		}
		r = append(r, AvSampleFormat(p))
	}
	return r
}

func (c *Codec) ChannelLayouts() []uint64 {
	r := make([]uint64, 0)
	if c.channel_layouts == nil {
		return r
	}
	size := unsafe.Sizeof(*c.channel_layouts)
	for i := 0; ; i++ {
		p := *(*C.uint64_t)(unsafe.Pointer(uintptr(unsafe.Pointer(c.channel_layouts)) + uintptr(i)*size))
		if p == 0 {
			break
		}
		r = append(r, uint64(p))
	}
	return r
}

//Same behaviour av_fast_malloc but the buffer has additional FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
func AvFastPaddedMalloc(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

//Return the LIBAvCODEC_VERSION_INT constant.
func AvcodecVersion() uint {
	return uint(C.avcodec_version())
}

//Return the libavcodec build-time configuration.
func AvcodecConfiguration() string {
	return C.GoString(C.avcodec_configuration())

}

//Return the libavcodec license.
func AvcodecLicense() string {
	return C.GoString(C.avcodec_license())
}

//Get the Class for Context.
func AvcodecGetClass() *Class {
	return (*Class)(C.avcodec_get_class())
}

//Get the Class for Frame.
func AvcodecGetFrameClass() *Class {
	return (*Class)(C.avcodec_get_frame_class())
}

//Get the Class for AvSubtitleRect.
func AvcodecGetSubtitleRectClass() *Class {
	return (*Class)(C.avcodec_get_subtitle_rect_class())
}

//Free all allocated data in the given subtitle struct.
func AvsubtitleFree(s *AvSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

//Pack a dictionary for use in side_data.
func AvPacketPackDictionary(d *avutil.Dictionary, s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

//Unpack a dictionary from side_data.
func AvPacketUnpackDictionary(d *uint8, s int, dt **avutil.Dictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt))))
}

//Find a registered decoder with a matching codec ID.
func AvcodecFindDecoder(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

//Find a registered decoder with the specified name.
func AvcodecFindDecoderByName(n string) *Codec {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return (*Codec)(C.avcodec_find_decoder_by_name(cn))
}

//Converts AvChromaLocation to swscale x/y chroma position.
func AvcodecEnumToChromaPos(x, y *int, l AvChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)), (C.enum_AVChromaLocation)(l)))
}

//Converts swscale x/y chroma position to AvChromaLocation.
func AvcodecChromaPosToEnum(x, y int) AvChromaLocation {
	return (AvChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

//Find a registered encoder with a matching codec ID.
func AvcodecFindEncoder(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

//Find a registered encoder with the specified name.
func AvcodecFindEncoderByName(c string) *Codec {
	cc := C.CString(c)
	defer C.free(unsafe.Pointer(cc))
	return (*Codec)(C.avcodec_find_encoder_by_name(cc))
}

//Put a string representing the codec tag codec_tag in buf.
// Deprecated
// func AvGetCodecTagString(b string, bf uintptr, c uint) uintptr {
// 	return uintptr(C.av_get_codec_tag_string(C.CString(b), C.size_t(bf), C.uint(c)))
// }

func AvcodecString(b string, bs int, ctxt *Context, e int) {
	cb := C.CString(b)
	defer C.free(unsafe.Pointer(cb))
	C.avcodec_string(cb, C.int(bs), (*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), C.int(e))
}

//Fill Frame audio data and linesize pointers.
func AvcodecFillAudioFrame(f *Frame, c int, s AvSampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

//Return codec bits per sample.
func AvGetBitsPerSample(c CodecId) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

//Return the PCM codec associated with a sample format.
func AvGetPcmCodec(f AvSampleFormat, b int) CodecId {
	return (CodecId)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(f), C.int(b)))
}

//Return codec bits per sample.
func AvGetExactBitsPerSample(c CodecId) int {
	return int(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(c)))
}

//Same behaviour av_fast_padded_malloc except that buffer will always be 0-initialized after call.
func AvFastPaddedMallocz(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_mallocz(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

//Encode extradata length to a buffer.
func AvXiphlacing(s *string, v uint) uint {
	return uint(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), (C.uint)(v)))
}

//If hwaccel is NULL, returns the first registered hardware accelerator, if hwaccel is non-NULL,
//returns the next registered hardware accelerator after hwaccel, or NULL if hwaccel is the last one.
func (a *AvHWAccel) AvHwaccelNext() *AvHWAccel {
	panic("deprecated")
	return nil
	//return (*AvHWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(a)))
}

//Get the type of the given codec.
func AvcodecGetType(c CodecId) MediaType {
	return (MediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

//Get the name of a codec.
func AvcodecGetName(d CodecId) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(d)))
}

//const Descriptor *avcodec_descriptor_get (enum CodecId id)
func AvcodecDescriptorGet(id CodecId) *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

func (d *Descriptor) Name() string {
	return C.GoString(d.name)
}

//Iterate over all codec descriptors known to libavcodec.
func (d *Descriptor) AvcodecDescriptorNext() *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(d)))
}

func AvcodecDescriptorGetByName(n string) *Descriptor {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return (*Descriptor)(C.avcodec_descriptor_get_by_name(cn))
}

func AvcodecReceivePacket(avctx *Context, avpkt *Packet) int {
	return int(C.avcodec_receive_packet((*C.struct_AVCodecContext)(avctx), (*C.struct_AVPacket)(avpkt)))
}

// AvcodecSendPacket send packet data as input to a decoder.
// See the ffmpeg documentation: https://www.ffmpeg.org/doxygen/trunk/group__lavc__encdec.html
func AvcodecSendPacket(avctx *Context, avpkt *Packet) int {
	return int(C.avcodec_send_packet((*C.struct_AVCodecContext)(avctx), (*C.struct_AVPacket)(avpkt)))
}

// AvcodecReceiveFrame receives a frame from the codec
// See the ffmpeg documentation: https://www.ffmpeg.org/doxygen/trunk/group__lavc__encdec.html
func AvcodecReceiveFrame(avctx *Context, frame *avutil.Frame) int {
	return int(C.avcodec_receive_frame((*C.struct_AVCodecContext)(avctx), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}

func AvcodecSendFrame(avctx *Context, frame *avutil.Frame) int {
	return int(C.avcodec_send_frame((*C.struct_AVCodecContext)(avctx), (*C.struct_AVFrame)(unsafe.Pointer(frame))))
}
