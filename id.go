package crx3lite

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/jda/go-crx3/pb"

	"google.golang.org/protobuf/proto"
)

const symbols = "abcdefghijklmnopqrstuvwxyz"

// ID returns the extension id.
func ID(crx []byte) (id string, err error) {
	if !IsCRX(crx) {
		return id, ErrUnsupportedFileFormat
	}
	var (
		headerSize = binary.LittleEndian.Uint32(crx[8:12])
		metaSize   = uint32(12)
		v          = crx[metaSize : headerSize+metaSize]
		header     pb.CrxFileHeader
		signedData pb.SignedData
	)

	if err := proto.Unmarshal(v, &header); err != nil {
		return id, err
	}
	if err := proto.Unmarshal(header.SignedHeaderData, &signedData); err != nil {
		return id, err
	}

	idx := strIDx()
	sid := fmt.Sprintf("%x", signedData.CrxId[:16])
	buf := bytes.NewBuffer(nil)
	for _, char := range sid {
		index := idx[char]
		buf.WriteString(string(symbols[index]))
	}
	return buf.String(), nil
}

func strIDx() map[rune]int {
	index := make(map[rune]int)
	src := "0123456789abcdef"
	for i, char := range src {
		index[char] = i
	}
	return index
}
