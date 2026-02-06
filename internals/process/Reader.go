package process

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
	"zp/internals/config"

	"github.com/cilium/ebpf/ringbuf"
	"github.com/dgraph-io/badger/v4"
)

func ReadFromRingBuf(rd *ringbuf.Reader, db *badger.DB) (*Process, error) {
	record, err := rd.Read()
	if err != nil {
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
	key := strconv.Itoa(int(p.Port))
	value := strconv.Itoa(int(p.Pid))
	config.AddToDb(db, []byte(key), []byte(value))
	return &p, nil
}
