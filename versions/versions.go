package main

import (
	"log"

	"github.com/charlestamz/goav/avcodec"
	"github.com/charlestamz/goav/avdevice"
	"github.com/charlestamz/goav/avfilter"
	"github.com/charlestamz/goav/avutil"
	"github.com/charlestamz/goav/swresample"
	"github.com/charlestamz/goav/swscale"
)

func main() {
	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())
}
