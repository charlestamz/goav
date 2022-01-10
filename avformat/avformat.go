// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avformat provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package avformat

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavutil/rational.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"unsafe"

	"github.com/charlestamz/goav/avutil"
)

type (
	AvProbeData                C.struct_AVProbeData
	InputFormat                C.struct_AVInputFormat
	OutputFormat               C.struct_AVOutputFormat
	Context                    C.struct_AVFormatContext
	Frame                      C.struct_AVFrame
	CodecContext               C.struct_AVCodecContext
	AvIndexEntry               C.struct_AVIndexEntry
	Stream                     C.struct_AVStream
	AvProgram                  C.struct_AVProgram
	AvChapter                  C.struct_AVChapter
	AvPacketList               C.struct_AVPacketList
	Packet                     C.struct_AVPacket
	CodecParserContext         C.struct_AVCodecParserContext
	AvIOContext                C.struct_AVIOContext
	AvCodec                    C.struct_AVCodec
	AvCodecTag                 C.struct_AVCodecTag
	Class                      C.struct_AVClass
	AvFormatInternal           C.struct_AVFormatInternal
	AvIOInterruptCB            C.struct_AVIOInterruptCB
	AvPacketSideData           C.struct_AVPacketSideData
	FFFrac                     C.struct_FFFrac
	AvStreamParseType          C.enum_AVStreamParseType
	AvDiscard                  C.enum_AVDiscard
	MediaType                  C.enum_AVMediaType
	AvDurationEstimationMethod C.enum_AVDurationEstimationMethod
	AvPacketSideDataType       C.enum_AVPacketSideDataType
	CodecId                    C.enum_AVCodecID
)

const (
	MAX_ARRAY_SIZE       = 1<<29 - 1
	AVFMT_NOFILE         = C.AVFMT_NOFILE
	AVFMT_GLOBALHEADER   = C.AVFMT_GLOBALHEADER
	AVFMT_FLAG_CUSTOM_IO = C.AVFMT_FLAG_CUSTOM_IO
	AVIO_FLAG_WRITE      = C.AVIO_FLAG_WRITE
	FF_FDEBUG_TS         = C.FF_FDEBUG_TS
)

const (
	AVSEEK_FLAG_ANY      = C.AVSEEK_FLAG_ANY
	AVSEEK_FLAG_BACKWARD = C.AVSEEK_FLAG_BACKWARD
	AVSEEK_FLAG_BYTE     = C.AVSEEK_FLAG_BYTE
	AVSEEK_FLAG_FRAME    = C.AVSEEK_FLAG_FRAME
)

type File C.FILE

//Allocate and read the payload of a packet and initialize its fields with default values.
func (ctxt *AvIOContext) AvGetPacket(pkt *Packet, s int) int {
	return int(C.av_get_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(s)))
}

//Read data and append it to the current content of the Packet.
func (ctxt *AvIOContext) AvAppendPacket(pkt *Packet, s int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(s)))
}

func (f *InputFormat) AvRegisterInputFormat() {
	panic("deprecated")
	//C.av_register_input_format((*C.struct_AVInputFormat)(f))
}

func (f *OutputFormat) AvRegisterOutputFormat() {
	panic("deprecated")
	//C.av_register_output_format((*C.struct_AVOutputFormat)(f))
}

//If f is NULL, returns the first registered input format, if f is non-NULL, returns the next registered input format after f or NULL if f is the last one.
func (f *InputFormat) AvIformatNext() *InputFormat {
	panic("deprecated")
	return nil
	//return (*InputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

//If f is NULL, returns the first registered output format, if f is non-NULL, returns the next registered output format after f or NULL if f is the last one.
func (f *OutputFormat) AvOformatNext() *OutputFormat {
	panic("deprecated")
	return nil
	//return (*OutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

func (f *OutputFormat) Flags() int {
	return int(f.flags)
}

//Return the LIBAvFORMAT_VERSION_INT constant.
func AvformatVersion() uint {
	return uint(C.avformat_version())
}

//Return the libavformat build-time configuration.
func AvformatConfiguration() string {
	return C.GoString(C.avformat_configuration())
}

//Return the libavformat license.
func AvformatLicense() string {
	return C.GoString(C.avformat_license())
}

//Do global initialization of network components.
func AvformatNetworkInit() int {
	return int(C.avformat_network_init())
}

//Undo the initialization done by avformat_network_init.
func AvformatNetworkDeinit() int {
	return int(C.avformat_network_deinit())
}

//Allocate an Context.
func AvformatAllocContext() *Context {
	return (*Context)(C.avformat_alloc_context())
}

//Get the Class for Context.
func AvformatGetClass() *Class {
	return (*Class)(C.avformat_get_class())
}

//Get side information from stream.
func (s *Stream) AvStreamGetSideData(t AvPacketSideDataType, z int) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(s), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&z))))
}

//Allocate an Context for an output format.
func AvformatAllocOutputContext2(ctx **Context, o *OutputFormat, fo, fi string) int {
	cfo := (*C.char)(nil)
	if len(fo) > 0 {
		cfo = C.CString(fo)
		defer C.free(unsafe.Pointer(cfo))
	}
	cfi := C.CString(fi)
	defer C.free(unsafe.Pointer(cfi))
	return int(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(ctx)), (*C.struct_AVOutputFormat)(o), cfo, cfi))
}

//Find InputFormat based on the short name of the input format.
func AvFindInputFormat(s string) *InputFormat {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return (*InputFormat)(C.av_find_input_format(cs))
}

//Guess the file format.
func AvProbeInputFormat(pd *AvProbeData, i int) *InputFormat {
	return (*InputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), C.int(i)))
}

//Guess the file format.
func AvProbeInputFormat2(pd *AvProbeData, o int, sm *int) *InputFormat {
	return (*InputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sm))))
}

//Guess the file format.
func AvProbeInputFormat3(pd *AvProbeData, o int, sl *int) *InputFormat {
	return (*InputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sl))))
}

//Probe a bytestream to determine the input format.
func AvProbeInputBuffer2(pb *AvIOContext, f **InputFormat, fi string, l int, o, m uint) int {
	cfi := C.CString(fi)
	defer C.free(unsafe.Pointer(cfi))
	return int(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), cfi, unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//Like av_probe_input_buffer2() but returns 0 on success.
func AvProbeInputBuffer(pb *AvIOContext, f **InputFormat, fi string, l int, o, m uint) int {
	cfi := C.CString(fi)
	defer C.free(unsafe.Pointer(cfi))
	return int(C.av_probe_input_buffer((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), cfi, unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//Create and initialize a AVIOContext for accessing the resource indicated by url.
func AvIOOpen(pb **AvIOContext, fi string, flags int) int {
	cfi := C.CString(fi)
	defer C.free(unsafe.Pointer(cfi))
	return int(C.avio_open((**C.struct_AVIOContext)(unsafe.Pointer(pb)), cfi, C.int(flags)))
}

//Force flushing of buffered data.
func AvIOFlush(pb *AvIOContext) {
	C.avio_flush((*C.struct_AVIOContext)(unsafe.Pointer(pb)))
}

//Close the resource accessed by the AVIOContext *s, free it and set the pointer pointing to it to NULL.
func AvIOClosep(pb **AvIOContext) int {
	return int(C.avio_closep((**C.struct_AVIOContext)(unsafe.Pointer(pb))))
}

//Open an input stream and read the header.
func AvformatOpenInput(ps **Context, fi string, fmt *InputFormat, d **avutil.Dictionary) int {
	cfi := C.CString(fi)
	defer C.free(unsafe.Pointer(cfi))
	return int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)), cfi, (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
func AvGuessFormat(sn, f, mt string) *OutputFormat {
	csn := C.CString(sn)
	defer C.free(unsafe.Pointer(csn))
	cf := C.CString(f)
	defer C.free(unsafe.Pointer(cf))
	cmt := C.CString(mt)
	defer C.free(unsafe.Pointer(cmt))
	return (*OutputFormat)(C.av_guess_format(csn, cf, cmt))
}

//Guess the codec ID based upon muxer and filename.
func AvGuessCodec(fmt *OutputFormat, sn, f, mt string, t MediaType) CodecId {
	csn := C.CString(sn)
	defer C.free(unsafe.Pointer(csn))
	cf := C.CString(f)
	defer C.free(unsafe.Pointer(cf))
	cmt := C.CString(mt)
	defer C.free(unsafe.Pointer(cmt))
	return (CodecId)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt), csn, cf, cmt, (C.enum_AVMediaType)(t)))
}

//Send a nice hexadecimal dump of a buffer to the specified file stream.
func AvHexDump(f *File, b *uint8, s int) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(b), C.int(s))
}

//Send a nice hexadecimal dump of a buffer to the log.
func AvHexDumpLog(a, l int, b *uint8, s int) {
	C.av_hex_dump_log(unsafe.Pointer(&a), C.int(l), (*C.uint8_t)(b), C.int(s))
}

//Send a nice dump of a packet to the specified file stream.
func AvPktDump2(f *File, pkt *Packet, dp int, st *Stream) {
	C.av_pkt_dump2((*C.FILE)(f), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//Send a nice dump of a packet to the log.
func AvPktDumpLog2(a int, l int, pkt *Packet, dp int, st *Stream) {
	C.av_pkt_dump_log2(unsafe.Pointer(&a), C.int(l), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//enum CodecId av_codec_get_id (const struct AvCodecTag *const *tags, unsigned int tag)
//Get the CodecId for the given codec tag tag.
func AvCodecGetId(t **AvCodecTag, tag uint) CodecId {
	return (CodecId)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(t)), C.uint(tag)))
}

//Get the codec tag for the given codec id id.
func AvCodecGetTag(t **AvCodecTag, id CodecId) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id)))
}

//Get the codec tag for the given codec id.
func AvCodecGetTag2(t **AvCodecTag, id CodecId, tag *uint) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

//Get the index for a specific timestamp.
func AvIndexSearchTimestamp(st *Stream, t int64, f int) int {
	return int(C.av_index_search_timestamp((*C.struct_AVStream)(st), C.int64_t(t), C.int(f)))
}

//Add an index entry into a sorted list.
func AvAddIndexEntry(st *Stream, pos, t, int64, s, d, f int) int {
	return int(C.av_add_index_entry((*C.struct_AVStream)(st), C.int64_t(pos), C.int64_t(t), C.int(s), C.int(d), C.int(f)))
}

//Split a URL string into components.
func AvUrlSplit(p string, ps int, a string, as int, h string, hs int, pp *int, path string, psize int, url string) {
	cp := C.CString(p)
	defer C.free(unsafe.Pointer(cp))
	ca := C.CString(a)
	defer C.free(unsafe.Pointer(ca))
	ch := C.CString(h)
	defer C.free(unsafe.Pointer(ch))
	cpa := C.CString(path)
	defer C.free(unsafe.Pointer(cpa))
	cu := C.CString(url)
	defer C.free(unsafe.Pointer(cu))
	C.av_url_split(cp, C.int(ps), ca, C.int(as), ch, C.int(hs), (*C.int)(unsafe.Pointer(pp)), cpa, C.int(psize), cu)
}

//int av_get_frame_filename (char *buf, int buf_size, const char *path, int number)
//Return in 'buf' the path with 'd' replaced by a number.
func AvGetFrameFilename(b string, bs int, pa string, n int) int {
	cb := C.CString(b)
	defer C.free(unsafe.Pointer(cb))
	cpa := C.CString(pa)
	defer C.free(unsafe.Pointer(cpa))
	return int(C.av_get_frame_filename(cb, C.int(bs), cpa, C.int(n)))
}

//Check whether filename actually is a numbered sequence generator.
func AvFilenameNumberTest(f string) int {
	cf := C.CString(f)
	defer C.free(unsafe.Pointer(cf))
	return int(C.av_filename_number_test(cf))
}

//Generate an SDP for an RTP session.
func AvSdpCreate(ac **Context, nf int, b string, s int) int {
	cb := C.CString(b)
	defer C.free(unsafe.Pointer(cb))
	return int(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(ac)), C.int(nf), cb, C.int(s)))
}

//int av_match_ext (const char *filename, const char *extensions)
//Return a positive value if the given filename has one of the given extensions, 0 otherwise.
func AvMatchExt(f, e string) int {
	cf := C.CString(f)
	defer C.free(unsafe.Pointer(cf))
	ce := C.CString(e)
	defer C.free(unsafe.Pointer(ce))
	return int(C.av_match_ext(cf, ce))
}

//Test if the given container can store a codec.
func AvformatQueryCodec(o *OutputFormat, cd CodecId, sc int) int {
	return int(C.avformat_query_codec((*C.struct_AVOutputFormat)(o), (C.enum_AVCodecID)(cd), C.int(sc)))
}

func AvformatGetRiffVideoTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_video_tags())
}

//struct AvCodecTag * avformat_get_riff_audio_tags (void)
func AvformatGetRiffAudioTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_audio_tags())
}

func AvformatGetMovVideoTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_video_tags())
}

func AvformatGetMovAudioTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_audio_tags())
}
