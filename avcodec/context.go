// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"github.com/charlestamz/goav/avutil"
	"unsafe"
)

func (ctxt *Context) AvCodecGetPktTimebase() Rational {
	panic("deprecated")
	return Rational{}
	//return (Rational)(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))))
}

func (ctxt *Context) AvCodecSetPktTimebase(r Rational) {
	panic("deprecated")
	//C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (C.struct_AVRational)(r))
}

func (ctxt *Context) AvCodecGetCodecDescriptor() *Descriptor {
	panic("deprecated")
	return nil
	//return (*Descriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))))
}

func (ctxt *Context) AvCodecSetCodecDescriptor(d *Descriptor) {
	panic("deprecated")
	//C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVCodecDescriptor)(d))
}

func (ctxt *Context) AvCodecGetLowres() int {
	panic("deprecated")
	return 0
	//return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))))
}

func (ctxt *Context) AvCodecSetLowres(i int) {
	panic("deprecated")
	//C.av_codec_set_lowres((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), C.int(i))
}

func (ctxt *Context) AvCodecGetSeekPreroll() int {
	panic("deprecated")
	return 0
	//return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))))
}

func (ctxt *Context) AvCodecSetSeekPreroll(i int) {
	panic("deprecated")
	//C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), C.int(i))
}

func (ctxt *Context) AvCodecGetChromaIntraMatrix() *uint16 {
	panic("deprecated")
	return nil
	//return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))))
}

func (ctxt *Context) AvCodecSetChromaIntraMatrix(t *uint16) {
	panic("deprecated")
	//C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.uint16_t)(t))
}

// AvcodecFreeContext Free the codec context and everything associated with it and write NULL to the provided pointer.
func AvcodecFreeContext(ctxt *Context) {
	var ptr *C.struct_AVCodecContext = (*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))
	C.avcodec_free_context(&ptr)
}

// AvcodecGetContextDefaults3 Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (ctxt *Context) AvcodecGetContextDefaults3(c *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVCodec)(c)))
}

//Copy the settings of the source Context into the destination Context.
// func (ctxt *Context) AvcodecCopyContext(ctxt2 *Context) int {
// 	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVCodecContext)(unsafe.Pointer(ctxt2))))
// }

// AvcodecOpen2 Initialize the Context to use the given Codec
func (ctxt *Context) AvcodecOpen2(c *Codec, d **avutil.Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

// AvcodecClose Close a given Context and free all the data associated with it (but not the Context itself).
func (ctxt *Context) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))))
}

// AvcodecDefaultGetBuffer2 The default callback for Context.get_buffer2().
func (ctxt *Context) AvcodecDefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVFrame)(f), C.int(l)))
}

// AvcodecAlignDimensions Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctxt *Context) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

// AvcodecAlignDimensions2 Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctxt *Context) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

// AvcodecDecodeAudio4 Decode the audio frame of size avpkt->size from avpkt->data into frame.
func (ctxt *Context) AvcodecDecodeAudio4(f *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

// AvcodecDecodeVideo2 Decode the video frame of size avpkt->size from avpkt->data into picture.
func (ctxt *Context) AvcodecDecodeVideo2(p *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVFrame)(p), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

func (ctxt *Context) ReceivePacket(a *Packet) int {
	return AvcodecReceivePacket(ctxt, a)
}

// SendPacket sends a packet to the context for decoding
// OO form of AvcodecSendPacket
func (ctxt *Context) SendPacket(a *Packet) int {
	return AvcodecSendPacket(ctxt, a)
}

// ReceiveFrame receives a decoded from from a context
// OO form of AvcodecReceiveFrame
func (ctxt *Context) ReceiveFrame(f *avutil.Frame) int {
	return AvcodecReceiveFrame(ctxt, f)
}

func (ctxt *Context) SendFrame(f *avutil.Frame) int {
	return AvcodecSendFrame(ctxt, f)
}

// AvcodecDecodeSubtitle2 Decode a subtitle message.
func (ctxt *Context) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

// AvcodecEncodeAudio2 Encode a frame of audio.
func (ctxt *Context) AvcodecEncodeAudio2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

// AvcodecEncodeVideo2 Encode a frame of video
func (ctxt *Context) AvcodecEncodeVideo2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

func (ctxt *Context) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (ctxt *Context) AvcodecDefaultGetFormat(f *avutil.PixelFormat) avutil.PixelFormat {
	return (avutil.PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (*C.enum_AVPixelFormat)(f)))
}

// AvcodecFlushBuffers Reset the internal decoder state / flush internal buffers.
func (ctxt *Context) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)))
}

// AvGetAudioFrameDuration Return audio frame duration.
func (ctxt *Context) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), C.int(f)))
}

func (ctxt *Context) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(unsafe.Pointer(ctxt))))
}

// AvParserParse2 Parse a packet.
func (ctxt *Context) AvParserParse2(ctxtp *ParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func (ctxt *Context) AvParserChange(ctxtp *ParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(unsafe.Pointer(ctxt)), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

func AvParserInit(c int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

func AvParserClose(ctxtp *ParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

func (p *Parser) AvParserNext() *Parser {
	panic("deprecated")
	return nil
	//return (*Parser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

func (p *Parser) AvRegisterCodecParser() {
	panic("deprecated")
	//C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}
