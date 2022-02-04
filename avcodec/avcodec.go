// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package avcodec

/*
#cgo pkg-config: libavformat libavcodec libavutil libswresample
#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>
#include <stdint.h>
#include <string.h>
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <libavutil/frame.h>
void register_codecs(){
#if LIBAVCODEC_VERSION_MAJOR < 58
	  av_register_all();
	  avcodec_register_all();
#endif
}

static inline int avcodec_profile_name_to_int(AVCodec *codec, const char *name) {
	const AVProfile *p;
	for (p = codec->profiles; p != NULL && p->profile != FF_PROFILE_UNKNOWN; p++)
		if (!strcasecmp(p->name, name))
			return p->profile;
	return FF_PROFILE_UNKNOWN;
}


*/
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
	FF_PROFILE_UNKNOWN                               = -99
	FF_PROFILE_RESERVED                              = -100
	FF_PROFILE_AAC_MAIN                              = 0
	FF_PROFILE_AAC_LOW                               = 1
	FF_PROFILE_AAC_SSR                               = 2
	FF_PROFILE_AAC_LTP                               = 3
	FF_PROFILE_AAC_HE                                = 4
	FF_PROFILE_AAC_HE_V2                             = 28
	FF_PROFILE_AAC_LD                                = 22
	FF_PROFILE_AAC_ELD                               = 38
	FF_PROFILE_MPEG2_AAC_LOW                         = 128
	FF_PROFILE_MPEG2_AAC_HE                          = 131
	FF_PROFILE_DNXHD                                 = 0
	FF_PROFILE_DNXHR_LB                              = 1
	FF_PROFILE_DNXHR_SQ                              = 2
	FF_PROFILE_DNXHR_HQ                              = 3
	FF_PROFILE_DNXHR_HQX                             = 4
	FF_PROFILE_DNXHR_444                             = 5
	FF_PROFILE_DTS                                   = 20
	FF_PROFILE_DTS_ES                                = 30
	FF_PROFILE_DTS_96_24                             = 40
	FF_PROFILE_DTS_HD_HRA                            = 50
	FF_PROFILE_DTS_HD_MA                             = 60
	FF_PROFILE_DTS_EXPRESS                           = 70
	FF_PROFILE_MPEG2_422                             = 0
	FF_PROFILE_MPEG2_HIGH                            = 1
	FF_PROFILE_MPEG2_SS                              = 2
	FF_PROFILE_MPEG2_SNR_SCALABLE                    = 3
	FF_PROFILE_MPEG2_MAIN                            = 4
	FF_PROFILE_MPEG2_SIMPLE                          = 5
	FF_PROFILE_H264_CONSTRAINED                      = (1 << 9)  // 8+1; constraint_set1_flag
	FF_PROFILE_H264_INTRA                            = (1 << 11) // 8+3; constraint_set3_flag
	FF_PROFILE_H264_BASELINE                         = 66
	FF_PROFILE_H264_CONSTRAINED_BASELINE             = (66 | FF_PROFILE_H264_CONSTRAINED)
	FF_PROFILE_H264_MAIN                             = 77
	FF_PROFILE_H264_EXTENDED                         = 88
	FF_PROFILE_H264_HIGH                             = 100
	FF_PROFILE_H264_HIGH_10                          = 110
	FF_PROFILE_H264_HIGH_10_INTRA                    = (110 | FF_PROFILE_H264_INTRA)
	FF_PROFILE_H264_MULTIVIEW_HIGH                   = 118
	FF_PROFILE_H264_HIGH_422                         = 122
	FF_PROFILE_H264_HIGH_422_INTRA                   = (122 | FF_PROFILE_H264_INTRA)
	FF_PROFILE_H264_STEREO_HIGH                      = 128
	FF_PROFILE_H264_HIGH_444                         = 144
	FF_PROFILE_H264_HIGH_444_PREDICTIVE              = 244
	FF_PROFILE_H264_HIGH_444_INTRA                   = (244 | FF_PROFILE_H264_INTRA)
	FF_PROFILE_H264_CAVLC_444                        = 44
	FF_PROFILE_VC1_SIMPLE                            = 0
	FF_PROFILE_VC1_MAIN                              = 1
	FF_PROFILE_VC1_COMPLEX                           = 2
	FF_PROFILE_VC1_ADVANCED                          = 3
	FF_PROFILE_MPEG4_SIMPLE                          = 0
	FF_PROFILE_MPEG4_SIMPLE_SCALABLE                 = 1
	FF_PROFILE_MPEG4_CORE                            = 2
	FF_PROFILE_MPEG4_MAIN                            = 3
	FF_PROFILE_MPEG4_N_BIT                           = 4
	FF_PROFILE_MPEG4_SCALABLE_TEXTURE                = 5
	FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION           = 6
	FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE          = 7
	FF_PROFILE_MPEG4_HYBRID                          = 8
	FF_PROFILE_MPEG4_ADVANCED_REAL_TIME              = 9
	FF_PROFILE_MPEG4_CORE_SCALABLE                   = 10
	FF_PROFILE_MPEG4_ADVANCED_CODING                 = 11
	FF_PROFILE_MPEG4_ADVANCED_CORE                   = 12
	FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE       = 13
	FF_PROFILE_MPEG4_SIMPLE_STUDIO                   = 14
	FF_PROFILE_MPEG4_ADVANCED_SIMPLE                 = 15
	FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0        = 1
	FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1        = 2
	FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION       = 32768
	FF_PROFILE_JPEG2000_DCINEMA_2K                   = 3
	FF_PROFILE_JPEG2000_DCINEMA_4K                   = 4
	FF_PROFILE_VP9_0                                 = 0
	FF_PROFILE_VP9_1                                 = 1
	FF_PROFILE_VP9_2                                 = 2
	FF_PROFILE_VP9_3                                 = 3
	FF_PROFILE_HEVC_MAIN                             = 1
	FF_PROFILE_HEVC_MAIN_10                          = 2
	FF_PROFILE_HEVC_MAIN_STILL_PICTURE               = 3
	FF_PROFILE_HEVC_REXT                             = 4
	FF_PROFILE_VVC_MAIN_10                           = 1
	FF_PROFILE_VVC_MAIN_10_444                       = 33
	FF_PROFILE_AV1_MAIN                              = 0
	FF_PROFILE_AV1_HIGH                              = 1
	FF_PROFILE_AV1_PROFESSIONAL                      = 2
	FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT            = 0xc0
	FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT = 0xc1
	FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT         = 0xc2
	FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS                = 0xc3
	FF_PROFILE_MJPEG_JPEG_LS                         = 0xf7
	FF_PROFILE_SBC_MSBC                              = 1
	FF_PROFILE_PRORES_PROXY                          = 0
	FF_PROFILE_PRORES_LT                             = 1
	FF_PROFILE_PRORES_STANDARD                       = 2
	FF_PROFILE_PRORES_HQ                             = 3
	FF_PROFILE_PRORES_4444                           = 4
	FF_PROFILE_PRORES_XQ                             = 5
	FF_PROFILE_ARIB_PROFILE_A                        = 0
	FF_PROFILE_ARIB_PROFILE_C                        = 1
	FF_PROFILE_KLVA_SYNC                             = 0
	FF_PROFILE_KLVA_ASYNC                            = 1
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

// AvCodecNext If c is NULL, returns the first registered codec, if c is non-NULL,
func (c *Codec) AvCodecNext() *Codec {
	panic("deprecated")
	return nil
	//return (*Codec)(C.av_codec_next((*C.struct_AVCodec)(c)))
}

// AvcodecRegisterAll Register all the codecs, parsers and bitstream filters which were enabled at configuration time.
func AvcodecRegisterAll() {
	//for ffmpeg version < 4, not to call deprecated functions
	C.register_codecs()
}

// AvcodecRegister Register the codec codec and initialize libavcodec.
func (c *Codec) AvcodecRegister() {
	panic("deprecated")
	//C.avcodec_register((*C.struct_AVCodec)(c))
}

// AvGetProfileName Return a name for the specified profile, if available.
func (c *Codec) AvGetProfileName(p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

// AvProfileNameToInt Return a int for the specified profile name, if available.
func (c *Codec) AvProfileNameToInt(name string) int {
	return int(C.avcodec_profile_name_to_int((*C.struct_AVCodec)(c), C.CString(name)))
}

// AvcodecAllocContext3 Allocate an Context and set its fields to default values.
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

// AvFastPaddedMalloc Same behaviour av_fast_malloc but the buffer has additional FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
func AvFastPaddedMalloc(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

// AvcodecVersion Return the LIBAvCODEC_VERSION_INT constant.
func AvcodecVersion() uint {
	return uint(C.avcodec_version())
}

// AvcodecConfiguration Return the libavcodec build-time configuration.
func AvcodecConfiguration() string {
	return C.GoString(C.avcodec_configuration())

}

// AvcodecLicense Return the libavcodec license.
func AvcodecLicense() string {
	return C.GoString(C.avcodec_license())
}

// AvcodecGetClass Get the Class for Context.
func AvcodecGetClass() *Class {
	return (*Class)(C.avcodec_get_class())
}

// AvcodecGetFrameClass Get the Class for Frame.
func AvcodecGetFrameClass() *Class {
	return (*Class)(C.avcodec_get_frame_class())
}

// AvcodecGetSubtitleRectClass Get the Class for AvSubtitleRect.
func AvcodecGetSubtitleRectClass() *Class {
	return (*Class)(C.avcodec_get_subtitle_rect_class())
}

// AvsubtitleFree Free all allocated data in the given subtitle struct.
func AvsubtitleFree(s *AvSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

// AvPacketPackDictionary Pack a dictionary for use in side_data.
func AvPacketPackDictionary(d *avutil.Dictionary, s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

// AvPacketUnpackDictionary Unpack a dictionary from side_data.
func AvPacketUnpackDictionary(d *uint8, s int, dt **avutil.Dictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt))))
}

// AvcodecFindDecoder Find a registered decoder with a matching codec ID.
func AvcodecFindDecoder(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

// AvcodecFindDecoderByName Find a registered decoder with the specified name.
func AvcodecFindDecoderByName(n string) *Codec {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return (*Codec)(C.avcodec_find_decoder_by_name(cn))
}

// AvcodecEnumToChromaPos Converts AvChromaLocation to swscale x/y chroma position.
func AvcodecEnumToChromaPos(x, y *int, l AvChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)), (C.enum_AVChromaLocation)(l)))
}

// AvcodecChromaPosToEnum Converts swscale x/y chroma position to AvChromaLocation.
func AvcodecChromaPosToEnum(x, y int) AvChromaLocation {
	return (AvChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

// AvcodecFindEncoder Find a registered encoder with a matching codec ID.
func AvcodecFindEncoder(id CodecId) *Codec {
	return (*Codec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

// AvcodecFindEncoderByName Find a registered encoder with the specified name.
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

// AvcodecFillAudioFrame Fill Frame audio data and line size pointers.
func AvcodecFillAudioFrame(f *Frame, c int, s AvSampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

// AvGetBitsPerSample Return codec bits per sample.
func AvGetBitsPerSample(c CodecId) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

// AvGetPcmCodec Return the PCM codec associated with a sample format.
func AvGetPcmCodec(f AvSampleFormat, b int) CodecId {
	return (CodecId)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(f), C.int(b)))
}

// AvGetExactBitsPerSample Return codec bits per sample.
func AvGetExactBitsPerSample(c CodecId) int {
	return int(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(c)))
}

// AvFastPaddedMallocz Same behaviour av_fast_padded_malloc except that buffer will always be 0-initialized after call.
func AvFastPaddedMallocz(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_mallocz(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

// AvXiphlacing Encode extradata length to a buffer.
func AvXiphlacing(s *string, v uint) uint {
	return uint(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), (C.uint)(v)))
}

// AvHwaccelNext If hwaccel is NULL, returns the first registered hardware accelerator, if hwaccel is non-NULL,
//returns the next registered hardware accelerator after hwaccel, or NULL if hwaccel is the last one.
func (a *AvHWAccel) AvHwaccelNext() *AvHWAccel {
	panic("deprecated")
	return nil
	//return (*AvHWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(a)))
}

// AvcodecGetType Get the type of the given codec.
func AvcodecGetType(c CodecId) MediaType {
	return (MediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

// AvcodecGetName Get the name of a codec.
func AvcodecGetName(d CodecId) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(d)))
}

// AvcodecDescriptorGet const Descriptor *avcodec_descriptor_get (enum CodecId id)
func AvcodecDescriptorGet(id CodecId) *Descriptor {
	return (*Descriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

func (d *Descriptor) Name() string {
	return C.GoString(d.name)
}

// AvcodecDescriptorNext Iterate over all codec descriptors known to libavcodec.
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
