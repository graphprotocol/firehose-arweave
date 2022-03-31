package codec

import (
	"fmt"

	pbcodec "github.com/ChainSafe/firehose-arweave/pb/cs/arweave/codec/v1"
	"github.com/streamingfast/bstream"
	pbbstream "github.com/streamingfast/pbgo/sf/bstream/v1"
	"google.golang.org/protobuf/proto"
)

func BlockFromProto(b *pbcodec.Block) (*bstream.Block, error) {
	content, err := proto.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal to binary form: %s", err)
	}

	block := &bstream.Block{
		Id:             b.Hash(),
		Number:         b.Height(),
		PreviousId:     b.PreviousBlock(),
		Timestamp:      b.Timestamp(),
		LibNum:         b.Height() - 1,
		PayloadKind:    pbbstream.Protocol_UNKNOWN,
		PayloadVersion: 1,
	}
	return bstream.GetBlockPayloadSetter(block, content)
}
