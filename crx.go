package crx3lite

import (
	"encoding/binary"
	"errors"

	"github.com/jda/go-crx3/pb"
	"google.golang.org/protobuf/proto"
)

var (
	ErrUnsupportedFileFormat = errors.New("crx3: unsupported file format")
	ErrExtensionNotSpecified = errors.New("crx3: extension id not specified")
)

func IsCRX(buf []byte) bool {
	if string(buf[0:4]) != "Cr24" {
		return false
	}
	if binary.LittleEndian.Uint32(buf[4:8]) != 3 {
		return false
	}
	return true
}

// Unpack unpacks chrome extension into some directory.
func Unpack(crx []byte) (zip []byte, err error) {
	if !IsCRX(crx) {
		return nil, ErrUnsupportedFileFormat
	}

	var (
		headerSize = binary.LittleEndian.Uint32(crx[8:12])
		metaSize   = uint32(12)
		v          = crx[metaSize : headerSize+metaSize]
		header     pb.CrxFileHeader
		signedData pb.SignedData
	)

	if err := proto.Unmarshal(v, &header); err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(header.SignedHeaderData, &signedData); err != nil {
		return nil, err
	}

	if len(signedData.CrxId) != 16 {
		return nil, ErrUnsupportedFileFormat
	}

	zip = crx[len(v)+int(metaSize):]

	return zip, nil
}
