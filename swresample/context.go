// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package swresample

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"
import (
	"unsafe"
)

// SwrInit Initialize context after user parameters have been set.
func (s *Context) SwrInit() int {
	return int(C.swr_init((*C.struct_SwrContext)(s)))
}

// SwrIsInitialized Check whether an swr context has been initialized or not.
func (s *Context) SwrIsInitialized() int {
	return int(C.swr_is_initialized((*C.struct_SwrContext)(s)))
}

// SwrAllocSetOpts Allocate Context if needed and set/reset common parameters.
func (s *Context) SwrAllocSetOpts(ocl int64, osf AvSampleFormat, osr int, icl int64, isf AvSampleFormat, isr, lo, lc int) *Context {
	return (*Context)(C.swr_alloc_set_opts((*C.struct_SwrContext)(s), C.int64_t(ocl), (C.enum_AVSampleFormat)(osf), C.int(osr), C.int64_t(icl), (C.enum_AVSampleFormat)(isf), C.int(isr), C.int(lo), unsafe.Pointer(&lc)))
}

// SwrFree Context destructor functions. Free the given Context and set the pointer to NULL.
func (s *Context) SwrFree() {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(&s)))
}

// SwrClose Closes the context so that swr_is_initialized() returns 0.
func (s *Context) SwrClose() {
	C.swr_close((*C.struct_SwrContext)(s))
}

//SwrGetOutSamples
//Find an upper bound on the number of samples that the next swr_convert
//call will output, if called with in_samples of input samples. This
//depends on the internal state, and anything changing the internal state
//(like further swr_convert() calls) will may change the number of samples
//swr_get_out_samples() returns for the same number of input samples.
//@param in_samples    number of input samples.
//@note any call to swr_inject_silence(), swr_convert(), swr_next_pts() or swr_set_compensation() invalidates this limit
//@note it is recommended to pass the correct available buffer size
//to all functions like swr_convert() even if swr_get_out_samples()
//indicates that less would be used.//@returns an upper bound on the number of samples that the next swr_convert
//will output or a negative value to indicate an error
//
func (s *Context) SwrGetOutSamples(sampleCount int) int {
	return int(C.swr_get_out_samples((*C.struct_SwrContext)(s), C.int(sampleCount)))
}

// SwrConvert Core conversion functions. Convert audio
func (s *Context) SwrConvert(out **uint8, oc int, in **uint8, ic int) int {
	return int(C.swr_convert((*C.struct_SwrContext)(s), (**C.uint8_t)(unsafe.Pointer(out)), C.int(oc), (**C.uint8_t)(unsafe.Pointer(in)), C.int(ic)))
}

// SwrNextPts Convert the next timestamp from input to output timestamps are in 1/(in_sample_rate * out_sample_rate) units.
func (s *Context) SwrNextPts(pts int64) int64 {
	return int64(C.swr_next_pts((*C.struct_SwrContext)(s), C.int64_t(pts)))
}

// SwrSetCompensation Low-level option setting functions
//These functons provide a means to set low-level options that is not possible with the AvOption API.
//Activate resampling compensation ("soft" compensation).
func (s *Context) SwrSetCompensation(sd, cd int) int {
	return int(C.swr_set_compensation((*C.struct_SwrContext)(s), C.int(sd), C.int(cd)))
}

// SwrSetChannelMapping Set a customized input channel mapping.
func (s *Context) SwrSetChannelMapping(cm *int) int {
	return int(C.swr_set_channel_mapping((*C.struct_SwrContext)(s), (*C.int)(unsafe.Pointer(cm))))
}

// SwrSetMatrix Set a customized remix matrix.
func (s *Context) SwrSetMatrix(m *int, t int) int {
	return int(C.swr_set_matrix((*C.struct_SwrContext)(s), (*C.double)(unsafe.Pointer(m)), C.int(t)))
}

// SwrDropOutput Sample handling functions. Drops the specified number of output samples.
func (s *Context) SwrDropOutput(c int) int {
	return int(C.swr_drop_output((*C.struct_SwrContext)(s), C.int(c)))
}

// SwrInjectSilence Injects the specified number of silence samples.
func (s *Context) SwrInjectSilence(c int) int {
	return int(C.swr_inject_silence((*C.struct_SwrContext)(s), C.int(c)))
}

// SwrGetDelay Gets the delay the next input sample will experience relative to the next output sample.
func (s *Context) SwrGetDelay(b int64) int64 {
	return int64(C.swr_get_delay((*C.struct_SwrContext)(s), C.int64_t(b)))
}

// SwrConvertFrame Frame based API. Convert the samples in the input Frame and write them to the output Frame.
func (s *Context) SwrConvertFrame(o, i *Frame) int {
	return int(C.swr_convert_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}

// SwrConfigFrame Configure or reconfigure the Context using the information provided by the AvFrames.
func (s *Context) SwrConfigFrame(o, i *Frame) int {
	return int(C.swr_config_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}
