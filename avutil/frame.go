// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/frame.h>
//#include <stdlib.h>
import "C"
import "unsafe"

type (
	AvBuffer            C.struct_AVBuffer
	AvBufferRef         C.struct_AVBufferRef
	AvBufferPool        C.struct_AVBufferPool
	Frame               C.struct_AVFrame
	AvFrameSideData     C.struct_AVFrameSideData
	AvFrameSideDataType C.enum_AVFrameSideDataType
)

func (f *Frame) AvprivFrameGetMetadatap() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(f.metadata))
}

func AvFrameGetBestEffortTimestamp(f *Frame) int64 {
	panic("deprecated")
	return 0
	//return int64(C.av_frame_get_best_effort_timestamp((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

func AvFrameSetQpTable(f *Frame, b *AvBufferRef, s, q int) int {
	panic("deprecated")
	return 0
	//return int(C.av_frame_set_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.struct_AVBufferRef)(unsafe.Pointer(b)), C.int(s), C.int(q)))
}

func AvFrameGetQpTable(f *Frame, s, t *int) int8 {
	panic("deprecated")
	return 0
	//return int8(*C.av_frame_get_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.int)(unsafe.Pointer(s)), (*C.int)(unsafe.Pointer(t))))
}

// AvFrameAlloc Allocate an Frame and set its fields to default values.
func AvFrameAlloc() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_alloc()))
}

// AvFrameFree Free the frame and any dynamically allocated objects in it, e.g.
func AvFrameFree(f *Frame) {
	var ptr *C.struct_AVFrame = (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_free(&ptr)
}

// AvFrameGetBuffer Allocate new buffer(s) for audio or video data.
func AvFrameGetBuffer(f *Frame, a int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(a)))
}

// AvFrameRef Setup a new reference to the data described by an given frame.
func AvFrameRef(d, s *Frame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

// AvFrameClone Create a new frame that references the same data as src.
func AvFrameClone(f *Frame) *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

// AvFrameUnref Unreference all the buffers referenced by frame and reset the frame fields.
func AvFrameUnref(f *Frame) {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

// AvFrameMoveRef Move everythnig contained in src to dst and reset src.
func AvFrameMoveRef(d, s *Frame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s)))
}

// AvFrameIsWritable Check if the frame data is writable.
func AvFrameIsWritable(f *Frame) int {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// AvFrameMakeWritable Ensure that the frame data is writable, avoiding data copy if possible.
func AvFrameMakeWritable(f *Frame) int {
	return int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

// AvFrameCopyProps Copy only "metadata" fields from src to dst.
func AvFrameCopyProps(d, s *Frame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

// AvFrameGetPlaneBuffer Get the buffer reference a given data plane is stored in.
func AvFrameGetPlaneBuffer(f *Frame, p int) *AvBufferRef {
	return (*AvBufferRef)(unsafe.Pointer(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p))))
}

// AvFrameNewSideData Add a new side data to a frame.
func AvFrameNewSideData(f *Frame, d AvFrameSideDataType, s int) *AvFrameSideData {
	return (*AvFrameSideData)(unsafe.Pointer(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s))))
}

func AvFrameGetSideData(f *Frame, t AvFrameSideDataType) *AvFrameSideData {
	return (*AvFrameSideData)(unsafe.Pointer(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t))))
}

// //static int get_video_buffer (Frame *frame, int align)
// func GetVideoBuffer(f *Frame, a int) int {
// 	return int(C.get_video_buffer(f, C.int(a)))
// }

// //static int get_audio_buffer (Frame *frame, int align)
// func GetAudioBuffer(f *Frame, a int) int {
// 	return C.get_audio_buffer(f, C.int(a))
// }

// //static void get_frame_defaults (Frame *frame)
// func GetFrameDefaults(f *Frame) {
// 	C.get_frame_defaults(*C.struct_AVFrame(f))
// }
