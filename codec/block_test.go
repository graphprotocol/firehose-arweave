package codec

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	pbcodec "github.com/ChainSafe/firehose-arweave/pb/sf/arweave/type/v1"
	"github.com/streamingfast/bstream"
	pbbstream "github.com/streamingfast/pbgo/sf/bstream/v1"

	"google.golang.org/protobuf/proto"
)

const BLOCK_19 = "080112307bbf0f04a8c9fba14ecc332db142290db9a0fc73a772e1cd199c5cfcdbd17deab4744bb805d21675cfc7e955469346fc1a5d000000010001000100000000010100010000010000010000000000000100010101000000010001000101010101010101000101010000010101010100010101000100010001000001000000010100000000000001010000010001000000223048b492556e4f1fa1e3a269985c7116a44090c914827fa65b0c64ad6e9292fc9adbf23edf2b546564ccf4174ccdd02d212a0608f8e1ebd805320608c7dfebd8053a02323940134a300000000545cc45b1d1d38c439cd638c2c307d2bcc1abb8aa0f0fbdf443831c896ea471ae66d8945561308bc22150a4f952006220f03a97f90891194b1563e51e4cb85214fc52de338e9b1d75d791a1c74742580b6a2b346e51526471486b4e3369626a515178624a46317a365f5f474576363569396a30517475736265617661457a0130820101308a010130aa01030a0131"

func TestDecode(t *testing.T) {
	bytes, _ := hex.DecodeString(BLOCK_19)

	b := &pbcodec.Block{}
	proto.Unmarshal(bytes, b)

	block := &bstream.Block{
		Id:             hex.EncodeToString(b.Hash),
		Number:         b.Height,
		PreviousId:     hex.EncodeToString(b.PreviousBlock),
		Timestamp:      b.Timestamp.AsTime(),
		LibNum:         b.Height - 1,
		PayloadKind:    pbbstream.Protocol_UNKNOWN,
		PayloadVersion: 1,
	}

	j, _ := json.Marshal(block)
	t.Logf("%s", string(j))
}
