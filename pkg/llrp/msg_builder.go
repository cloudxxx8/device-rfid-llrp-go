//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package llrp

import (
	"fmt"
	"io"
	"sync"
)

type fieldEncoder interface {
	EncodeFields(w io.Writer) error
}

type paramHeader struct {
	ParamType
	sz   uint16
	data fieldEncoder
	subs []paramHeader
}

func encodeParams(w io.Writer, headers ...paramHeader) error {
	for _, h := range headers {
		if h.ParamType.IsTV() {
			if n, err := w.Write([]byte{byte(h.ParamType | 0x80)}); err != nil {
				return fmt.Errorf("failed to write TV header for %v: %w", h.ParamType, err)
			} else if n < 1 {
				return fmt.Errorf("short write for %v: %d < 1", h.ParamType, n)
			}
		} else {
			if n, err := w.Write([]byte{
				byte(h.ParamType >> 8), byte(h.ParamType & 0xff),
				byte(h.sz >> 8), byte(h.sz & 0xff),
			}); err != nil {
				return fmt.Errorf("failed to write parameter header: %w", err)
			} else if n < 4 {
				return fmt.Errorf("short write: %d < 4", n)
			}
		}

		if err := h.data.EncodeFields(w); err != nil {
			return err
		}

		if err := encodeParams(w, h.subs...); err != nil {
			return err
		}
	}

	return nil
}

// msgWriter isn't for general use at the moment,
// but is handy for certain basic tests.
type msgWriter struct {
	w      io.Writer // target output
	header Header
	mu     sync.Mutex
}

func newMsgWriter(w io.Writer, version VersionNum) *msgWriter {
	return &msgWriter{
		w: w,
		header: Header{
			version: version,
		},
	}
}

func (mw *msgWriter) Write(mid messageID, out Outgoing) error {
	mw.mu.Lock()
	defer mw.mu.Unlock()

	data, err := out.MarshalBinary()
	if err != nil {
		return err
	}

	if uint32(len(data)) > maxPayloadSz {
		return fmt.Errorf("outgoing payload size (%d) "+
			"is larger than that permitted by LLRP", len(data))
	}

	mw.header.payloadLen = uint32(len(data))
	mw.header.typ = out.Type()
	mw.header.id = mid
	if _, err := mw.header.WriteTo(mw.w); err != nil {
		return err
	}

	if _, err = mw.w.Write(data); err != nil {
		return fmt.Errorf("failed to write payload %v: %w", mw.header, err)
	}
	return nil
}
