package rtcrtpencodingparameters

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpfecparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtprtxparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcssrcrange"
)

type RTCRtpEncodingParameters struct {
	active                *bool
	codecPayloadType      *byte
	dependencyEncodingIds *[]string
	encodingId            *string
	fec                   *rtcrtpfecparameters.RTCRtpFecParameters
	framerateScale        *float32
	maxBitrate            *float32
	maxFramerate          *uint
	minQuality            *float32
	priority              *float32
	resolutionScale       *float32
	rtx                   *rtcrtprtxparameters.RTCRtpRtxParameters
	ssrc                  *uint
	ssrcRange             *rtcssrcrange.RTCSsrcRange
}