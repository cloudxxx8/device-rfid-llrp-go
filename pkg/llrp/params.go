//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

//go:generate python3 generate_param_code.py -i messages.yaml -s generated_structs.go -t binary_test.go -m generated_marshal.go -u generated_unmarshal.go -e generated_encoder.go
//go:generate stringer -type=ParamType,ConnectionAttemptEventType,StatusCode,AirProtocolIDType

package llrp

import (
	"strconv"
)

const (
	paramInvalid = ParamType(0)

	// TV-encoded params

	ParamAntennaID                 = ParamType(1)
	ParamFirstSeenUTC              = ParamType(2)
	ParamFirstSeenUptime           = ParamType(3)
	ParamLastSeenUTC               = ParamType(4)
	ParamLastSeenUptime            = ParamType(5)
	ParamPeakRSSI                  = ParamType(6)
	ParamChannelIndex              = ParamType(7)
	ParamTagSeenCount              = ParamType(8)
	ParamROSpecID                  = ParamType(9)
	ParamInventoryParameterSpecID  = ParamType(10)
	ParamC1G2CRC                   = ParamType(11)
	ParamC1G2PC                    = ParamType(12)
	ParamEPC96                     = ParamType(13)
	ParamSpecIndex                 = ParamType(14)
	ParamClientRequestOpSpecResult = ParamType(15)
	ParamAccessSpecID              = ParamType(16)
	ParamOpSpecID                  = ParamType(17)
	ParamC1G2SingulationDetails    = ParamType(18)
	ParamC1G2XPCW1                 = ParamType(19)
	ParamC1G2XPCW2                 = ParamType(20)

	// TLV encoded parameters

	ParamUTCTimestamp                      = ParamType(128)
	ParamUptime                            = ParamType(129)
	ParamGeneralDeviceCapabilities         = ParamType(137)
	ParamReceiveSensitivityTableEntry      = ParamType(139)
	ParamPerAntennaAirProtocol             = ParamType(140)
	ParamGPIOCapabilities                  = ParamType(141)
	ParamLLRPCapabilities                  = ParamType(142)
	ParamRegulatoryCapabilities            = ParamType(143)
	ParamUHFBandCapabilities               = ParamType(144)
	ParamTransmitPowerLevelTableEntry      = ParamType(145)
	ParamFrequencyInformation              = ParamType(146)
	ParamFrequencyHopTable                 = ParamType(147)
	ParamFixedFrequencyTable               = ParamType(148)
	ParamPerAntennaReceiveSensitivityRange = ParamType(149)
	ParamROSpec                            = ParamType(177)
	ParamROBoundarySpec                    = ParamType(178)
	ParamROSpecStartTrigger                = ParamType(179)
	ParamPeriodicTriggerValue              = ParamType(180)
	ParamGPITriggerValue                   = ParamType(181)
	ParamROSpecStopTrigger                 = ParamType(182)
	ParamAISpec                            = ParamType(183)
	ParamAISpecStopTrigger                 = ParamType(184)
	ParamTagObservationTrigger             = ParamType(185)
	ParamInventoryParameterSpec            = ParamType(186)
	ParamRFSurveySpec                      = ParamType(187)
	ParamRFSurveySpecStopTrigger           = ParamType(188)
	ParamAccessSpec                        = ParamType(207)
	ParamAccessSpecStopTrigger             = ParamType(208)
	ParamAccessCommand                     = ParamType(209)
	ParamClientRequestOpSpec               = ParamType(210)
	ParamClientRequestResponse             = ParamType(211)
	ParamLLRPConfigurationStateValue       = ParamType(217)
	ParamIdentification                    = ParamType(218)
	ParamGPOWriteData                      = ParamType(219)
	ParamKeepAliveSpec                     = ParamType(220)
	ParamAntennaProperties                 = ParamType(221)
	ParamAntennaConfiguration              = ParamType(222)
	ParamRFReceiver                        = ParamType(223)
	ParamRFTransmitter                     = ParamType(224)
	ParamGPIPortCurrentState               = ParamType(225)
	ParamEventsAndReports                  = ParamType(226)
	ParamROReportSpec                      = ParamType(237)
	ParamTagReportContentSelector          = ParamType(238)
	ParamAccessReportSpec                  = ParamType(239)
	ParamTagReportData                     = ParamType(240)
	ParamEPCData                           = ParamType(241)
	ParamRFSurveyReportData                = ParamType(242)
	ParamFrequencyRSSILevelEntry           = ParamType(243)
	ParamReaderEventNotificationSpec       = ParamType(244)
	ParamEventNotificationState            = ParamType(245)
	ParamReaderEventNotificationData       = ParamType(246)
	ParamHoppingEvent                      = ParamType(247)
	ParamGPIEvent                          = ParamType(248)
	ParamROSpecEvent                       = ParamType(249)
	ParamReportBufferLevelWarningEvent     = ParamType(250)
	ParamReportBufferOverflowErrorEvent    = ParamType(251)
	ParamReaderExceptionEvent              = ParamType(252)
	ParamRFSurveyEvent                     = ParamType(253)
	ParamAISpecEvent                       = ParamType(254)
	ParamAntennaEvent                      = ParamType(255)
	ParamConnectionAttemptEvent            = ParamType(256)
	ParamConnectionCloseEvent              = ParamType(257)
	ParamLLRPStatus                        = ParamType(287)
	ParamFieldError                        = ParamType(288)
	ParamParameterError                    = ParamType(289)
	ParamLoopSpec                          = ParamType(355)
	ParamSpecLoopEvent                     = ParamType(356)
	ParamCustom                            = ParamType(1023)

	ParamC1G2LLRPCapabilities                        = ParamType(327)
	ParamUHFC1G2RFModeTable                          = ParamType(328)
	ParamUHFC1G2RFModeTableEntry                     = ParamType(329)
	ParamC1G2InventoryCommand                        = ParamType(330)
	ParamC1G2Filter                                  = ParamType(331)
	ParamC1G2TagInventoryMask                        = ParamType(332)
	ParamC1G2TagInventoryStateAwareFilterAction      = ParamType(333)
	ParamC1G2TagInventoryStateUnawareFilterAction    = ParamType(334)
	ParamC1G2RFControl                               = ParamType(335)
	ParamC1G2SingulationControl                      = ParamType(336)
	ParamC1G2TagInventoryStateAwareSingulationAction = ParamType(337)
	ParamC1G2TagSpec                                 = ParamType(338)
	ParamC1G2TargetTag                               = ParamType(339)
	ParamC1G2Read                                    = ParamType(341)
	ParamC1G2Write                                   = ParamType(342)
	ParamC1G2Kill                                    = ParamType(343)
	ParamC1G2Lock                                    = ParamType(344)
	ParamC1G2LockPayload                             = ParamType(345)
	ParamC1G2BlockErase                              = ParamType(346)
	ParamC1G2BlockWrite                              = ParamType(347)
	ParamC1G2EPCMemorySelector                       = ParamType(348)
	ParamC1G2ReadOpSpecResult                        = ParamType(349)
	ParamC1G2WriteOpSpecResult                       = ParamType(350)
	ParamC1G2KillOpSpecResult                        = ParamType(351)
	ParamC1G2LockOpSpecResult                        = ParamType(352)
	ParamC1G2BlockEraseOpSpecResult                  = ParamType(353)
	ParamC1G2BlockWriteOpSpecResult                  = ParamType(354)
	ParamC1G2Recommission                            = ParamType(357)
	ParamC1G2BlockPermalock                          = ParamType(358)
	ParamC1G2GetBlockPermalockStatus                 = ParamType(359)
	ParamC1G2RecommissionOpSpecResult                = ParamType(360)
	ParamC1G2BlockPermalockOpSpecResult              = ParamType(361)
	ParamC1G2GetBlockPermalockStatusOpSpecResult     = ParamType(362)

	ParamMaximumReceiveSensitivity     = ParamType(363)
	ParamRFSurveyFrequencyCapabilities = ParamType(365)

	paramResvStart = ParamType(900)
	paramResvEnd   = ParamType(999)
)

// IsTV returns true if the ParamType is TV-encoded.
// TV-encoded parameters have specific lengths which must be looked up.
func (pt ParamType) IsTV() bool {
	return pt != 0 && pt <= 127
}

// IsTLV returns true if the ParamType is TLV-encoded.
// TLV-encoded parameters include their length in their headers.
func (pt ParamType) IsTLV() bool {
	return pt >= 128 && pt <= 2047
}

// IsValid returns true if the ParamType is within the valid LLRP Parameter range
// and not one of the reserved parameter types (900-999)
func (pt ParamType) IsValid() bool {
	return 0 < pt && pt <= 2047 && !(paramResvStart <= pt && pt <= paramResvEnd)
}

const (
	statusMsgStart    = StatusMsgParamError
	statusMsgEnd      = StatusMsgMsgUnexpected
	statusParamStart  = StatusParamParamError
	statusParamEnd    = StatusParamParamUnsupported
	statusFieldStart  = StatusFieldInvalid
	statusFieldEnd    = StatusFieldOutOfRange
	statusDeviceStart = StatusDeviceError
	statusDeviceEnd   = StatusDeviceError
)

func (sc StatusCode) isMsgStatus() bool {
	return statusMsgStart <= sc && sc <= statusMsgEnd
}

func (sc StatusCode) isParamStatus() bool {
	return statusParamStart <= sc && sc <= statusParamEnd
}

func (sc StatusCode) isFieldStatus() bool {
	return statusFieldStart <= sc && sc <= statusFieldEnd
}

func (sc StatusCode) isDeviceStatus() bool {
	return statusDeviceStart <= sc && sc <= statusDeviceEnd
}

func (sc StatusCode) defaultText() string {
	switch {
	case sc == StatusSuccess:
		return "success"
	case sc.isMsgStatus():
		return statusMsgErrs[sc-statusMsgStart]
	case sc.isParamStatus():
		return statusParamErrs[sc-statusParamStart]
	case sc.isFieldStatus():
		return statusFieldErrs[sc-statusFieldStart]
	case sc.isDeviceStatus():
		return statusDeviceErrs[sc-statusDeviceStart]
	}

	return "unknown LLRP status code " + strconv.Itoa(int(sc))
}

var statusMsgErrs = [...]string{
	StatusMsgParamError - statusMsgStart:       "message parameter error",
	StatusMsgFieldError - statusMsgStart:       "message field error",
	StatusMsgParamUnexpected - statusMsgStart:  "message parameter unexpected",
	StatusMsgParamMissing - statusMsgStart:     "message parameter missing",
	StatusMsgParamDuplicate - statusMsgStart:   "message has duplicate parameter",
	StatusMsgParamOverflow - statusMsgStart:    "message parameter instances exceed Reader max",
	StatusMsgFieldOverflow - statusMsgStart:    "message field instances exceed Reader max",
	StatusMsgParamUnknown - statusMsgStart:     "message parameter unknown",
	StatusMsgFieldUnknown - statusMsgStart:     "message field unknown",
	StatusMsgMsgUnsupported - statusMsgStart:   "message type unsupported",
	StatusMsgVerUnsupported - statusMsgStart:   "LLRP version not supported",
	StatusMsgParamUnsupported - statusMsgStart: "message parameter unsupported",
	StatusMsgMsgUnexpected - statusMsgStart:    "unexpected message type",
}
var _ = statusMsgErrs[statusMsgEnd-statusMsgStart] // compile error == missing message

var statusParamErrs = [...]string{
	StatusParamParamError - statusParamStart:       "error in sub-parameter",
	StatusParamFieldError - statusParamStart:       "error in parameter field",
	StatusParamParamUnexpected - statusParamStart:  "unexpected sub-parameter",
	StatusParamParamMissing - statusParamStart:     "missing required sub-parameter",
	StatusParamParamDuplicate - statusParamStart:   "duplicate sub-parameter",
	StatusParamParamOverflow - statusParamStart:    "sub-parameter instances exceed Reader max",
	StatusParamFieldOverflow - statusParamStart:    "parameter field instances exceeds Reader max",
	StatusParamParamUnknown - statusParamStart:     "unknown sub-parameter",
	StatusParamFieldUnknown - statusParamStart:     "unknown parameter field",
	StatusParamParamUnsupported - statusParamStart: "unsupported sub-parameter",
}
var _ = statusParamErrs[statusParamEnd-statusParamStart] // compile error == missing message

var statusFieldErrs = [...]string{
	StatusFieldInvalid - statusFieldStart:    "invalid value",
	StatusFieldOutOfRange - statusFieldStart: "value out of range",
}
var _ = statusFieldErrs[statusFieldEnd-statusFieldStart] // compile error == missing message

var statusDeviceErrs = [...]string{
	StatusDeviceError - statusDeviceEnd: "the Reader encountered a problem unrelated to the message",
}
var _ = statusDeviceErrs[statusDeviceEnd-statusDeviceStart] // compile error == missing message

func (fe FieldError) Error() string {
	return fe.ErrorCode.defaultText() + " at index " + strconv.Itoa(int(fe.FieldIndex))
}

// Error constructs a string from the parameter's error.
func (pe *ParameterError) Error() string {
	msg := pe.ErrorCode.defaultText() + " " + pe.ParameterType.String()
	if pe.ParameterType.IsValid() {
		msg += " (type " + strconv.Itoa(int(pe.ParameterType)) + ")"
	}

	if pe.FieldError != nil {
		msg += ": " + pe.FieldError.Error()
	}

	if pe.ParameterError != nil {
		msg += ": " + pe.ParameterError.Error()
	}

	return msg
}

// StatusError is an LLRPStatus that implements the Error interface.
// This is a different type than a regular LLRPStatus
// because the latter can also represent success.
type StatusError LLRPStatus

// Error implements the error interface for a StatusError.
//
// It returns the ErrorDescription set by the Reader,
// and appends to it
func (se *StatusError) Error() string {
	msg := se.Status.defaultText()
	if se.ErrorDescription != "" {
		msg += ": " + se.ErrorDescription
	}

	if se.FieldError != nil {
		msg += ": " + se.FieldError.Error()
	}

	if se.ParameterError != nil {
		msg += ": " + se.ParameterError.Error()
	}

	if msg == "" {
		msg = "unknown error"
	}

	return msg
}

// Err returns an error represented by this LLRPStatus, if any.
// If the Status is Success, this returns nil.
func (ls *LLRPStatus) Err() error {
	if ls.Status == StatusSuccess {
		return nil
	}

	se := StatusError(*ls)
	return &se
}
