//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package llrp

import (
	"encoding/binary"
	"testing"
)

func TestMsgReader_readerEventNotification(t *testing.T) {
	pConnAttempt := []byte{
		0x1, 0x0, // ConnectionAttemptEvent
		0x0, 0x6, // size; 4 byte header + 2 byte field
		0x0, 0x4, // 4 == anotherConnAttempted
	}

	pTimestamp := []byte{
		0x0, 128, // UTCTimestamp
		0x0, 12, // header + 8 bytes of microseconds
		0x00, 0x02, 0x4b, 0xd4, 0xc0, 0x03, 0x1a, 0x00, // June 25, 1990, 11:18AM EST
	}

	eventData := append(pTimestamp, pConnAttempt...)
	pReaderEvent := append([]byte{
		0x0, 246, // ReaderEventNotificationParameter
		0x0, 0x0, // length (set below)
	}, eventData...)
	binary.BigEndian.PutUint16(pReaderEvent[2:4], uint16(len(pReaderEvent)))

	ren := ReaderEventNotification{}
	if err := ren.UnmarshalBinary(pReaderEvent); err != nil {
		t.Fatalf("%+v", err)
	}

	nd := ren.ReaderEventNotificationData
	expTS := UTCTimestamp(646327080000000)
	if nd.UTCTimestamp != expTS {
		t.Errorf("utc timestamp mismatch: %+v != %+v", nd.UTCTimestamp, expTS)
	}
	if nd.ConnectionAttemptEvent == nil ||
		ConnectionAttemptEventType(*nd.ConnectionAttemptEvent) != ConnAttemptedAgain {
		t.Errorf("expected ConnAttemptedAgain, but got %+v", nd.ConnectionAttemptEvent)
	}
}

func BenchmarkReaderEventNotification_UnmarshalBinary(b *testing.B) {
	data := []byte{
		0x0, 246, // ReaderEventNotificationParameter
		0x0, 22, // length

		0x0, 128, // UTCTimestamp
		0x0, 12, // header + 8 bytes of microseconds
		0x00, 0x02, 0x4b, 0xd4, 0xc0, 0x03, 0x1a, 0x00, // June 25, 1990, 11:18AM EST

		0x1, 0x0, // ConnectionAttemptEvent
		0x0, 0x6, // size; 4 byte header + 2 byte field
		0x0, 0x3, // failedReasonUnknown
	}

	b.ReportAllocs()
	b.ResetTimer()
	ren := ReaderEventNotification{}
	for i := 0; i < b.N; i++ {
		if err := ren.UnmarshalBinary(data); err != nil {
			b.Fatalf("%+v", err)
		}
	}
}

func TestReaderEventNotification_UnmarshalBinary(t *testing.T) {
	data := []byte{
		0x0, 246, // ReaderEventNotificationParameter
		0x0, 22, // length

		0x0, 128, // UTCTimestamp
		0x0, 12, // header + 8 bytes of microseconds
		0x00, 0x02, 0x4b, 0xd4, 0xc0, 0x03, 0x1a, 0x00, // June 25, 1990, 11:18AM EST

		0x1, 0x0, // ConnectionAttemptEvent
		0x0, 0x6, // size; 4 byte header + 2 byte field
		0x0, 0x3, // failedReasonUnknown
	}

	ren := ReaderEventNotification{}
	if err := ren.UnmarshalBinary(data); err != nil {
		t.Fatalf("%+v", err)
	}

	cae := ConnectionAttemptEvent(ConnFailedReasonUnknown)
	exp := ReaderEventNotification{
		ReaderEventNotificationData: ReaderEventNotificationData{
			UTCTimestamp:           UTCTimestamp(0x00024bd4c0031a00),
			ConnectionAttemptEvent: &cae,
		},
	}

	if exp.ReaderEventNotificationData.UTCTimestamp != ren.ReaderEventNotificationData.UTCTimestamp {
		t.Fatalf("expected timestamp %+v; got %+v",
			exp.ReaderEventNotificationData.UTCTimestamp,
			ren.ReaderEventNotificationData.UTCTimestamp)
	}

	if nil == ren.ReaderEventNotificationData.ConnectionAttemptEvent {
		t.Fatalf("expected non-nil timestamp and connection attempt; got %+v", ren.ReaderEventNotificationData.ConnectionAttemptEvent)
	}

	if ConnFailedReasonUnknown != ConnectionAttemptEventType(*ren.ReaderEventNotificationData.ConnectionAttemptEvent) {
		t.Fatalf("expected ConnFailedReasonUnknown; got %+v", ren.ReaderEventNotificationData.ConnectionAttemptEvent)
	}
}
