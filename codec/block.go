package codec

import (
	"encoding/hex"
	"fmt"

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

	block := &bstream.Block{
		Id:             hex.EncodeToString(b.Hash),
		Number:         b.Height,
		PreviousId:     hex.EncodeToString(b.PreviousBlock),
		Timestamp:      b.Timestamp.AsTime(),
		LibNum:         libNum,
		PayloadKind:    pbbstream.Protocol_UNKNOWN,
		PayloadVersion: 1,
	}
	return bstream.GetBlockPayloadSetter(block, content)
}
