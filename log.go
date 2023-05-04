package lattice_common

import (
	"io"
	"zkjg.com/lattice/common/rlp"
)

type Log struct {
	Address Address `json:"address"  gencodec:"required"` // address of the contract that generated the event
	Topics  []Hash  `json:"topics"   gencodec:"required"` // list of topics provided by the contract.
	Data    []byte  `json:"data"     gencodec:"required"` // supplied by the contract, usually ABI-encoded
	// index of the log in the block
	Index uint `json:"logIndex" gencodec:"required"`

	TBlockHash   Hash   `json:"tblockHash"`
	DBlockNumber uint64 `json:"dblockNumber"`

	Removed bool `json:"removed"`

	DataHex string `json:"dataHex"`
}

type rlpLog struct {
	Address Address
	Topics  []Hash
	Data    []byte
}

// EncodeRLP implements rlp.Encoder.
func (l *Log) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, rlpLog{Address: l.Address, Topics: l.Topics, Data: l.Data})
}

// DecodeRLP implements rlp.Decoder.
func (l *Log) DecodeRLP(s *rlp.Stream) error {
	var dec rlpLog
	err := s.Decode(&dec)
	if err == nil {
		l.Address, l.Topics, l.Data = dec.Address, dec.Topics, dec.Data
	}
	return err
}

type rlpStorageLog rlpLog

type LogForStorage Log

// EncodeRLP implements rlp.Encoder.
func (l *LogForStorage) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, rlpStorageLog{
		Address: l.Address,
		Topics:  l.Topics,
		Data:    l.Data,
	})
}

// DecodeRLP implements rlp.Decoder.
func (l *LogForStorage) DecodeRLP(s *rlp.Stream) error {
	blob, err := s.Raw()
	if err != nil {
		return err
	}
	var dec rlpStorageLog
	err = rlp.DecodeBytes(blob, &dec)
	if err != nil {
		return err
	}

	*l = LogForStorage{
		Address: dec.Address,
		Topics:  dec.Topics,
		Data:    dec.Data,
	}
	return nil
}
