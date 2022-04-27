package codec

import (
	"encoding/hex"
	"fmt"
	"time"

	pbcodec "github.com/ChainSafe/firehose-arweave/pb/sf/arweave/type/v1"
	"github.com/streamingfast/bstream"
	pbbstream "github.com/streamingfast/pbgo/sf/bstream/v1"
	"google.golang.org/protobuf/proto"
)

const CONFIRMS uint64 = 20

func BlockFromProto(b *pbcodec.Block) (*bstream.Block, error) {
	content, err := proto.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal to binary form: %s", err)
	}

	var libNum uint64
	if b.Height > CONFIRMS {
		libNum = b.Height - CONFIRMS
	} else {
		libNum = 0
	}

	var previousId string
	if b.Height != 0 {
		previousId = hex.EncodeToString(b.PreviousBlock)
	} else {
		var empty_hash [64]byte
		previousId = hex.EncodeToString(empty_hash[:])
	}

	block := &bstream.Block{
		Id:             hex.EncodeToString(b.IndepHash),
		Number:         b.Height,
		PreviousId:     previousId,
		Timestamp:      time.UnixMilli(int64(b.Timestamp)),
		LibNum:         libNum,
		PayloadKind:    pbbstream.Protocol_UNKNOWN,
		PayloadVersion: 1,
	}
	return bstream.GetBlockPayloadSetter(block, content)
}
