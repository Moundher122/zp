package process

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/cilium/ebpf/ringbuf"
)

func ReadFromRingBuf(rd *ringbuf.Reader) (*Process, error) {
	record, err := rd.Read()
	if err != nil {
		// ring buffer closed or interrupted
		if errors.Is(err, ringbuf.ErrClosed) {
			return nil, err
		}
		return nil, err
	}
	var p Process
	reader := bytes.NewReader(record.RawSample)
	if err := binary.Read(reader, binary.LittleEndian, &p); err != nil {
		return nil, err
	}
	return &p, nil
}
