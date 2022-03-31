package codec

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	pbcodec "github.com/ChainSafe/firehose-arweave/pb/cs/arweave/codec/v1"
	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

// ConsoleReader is what reads the `geth` output directly. It builds
// up some LogEntry objects. See `LogReader to read those entries .
type ConsoleReader struct {
	lines chan string
	close func()

	ctx  *parseCtx
	done chan interface{}
}

func NewConsoleReader(lines chan string, rpcUrl string) (*ConsoleReader, error) {
	l := &ConsoleReader{
		lines: lines,
		close: func() {},
		done:  make(chan interface{}),
	}
	return l, nil
}

//todo: WTF?
func (r *ConsoleReader) Done() <-chan interface{} {
	return r.done
}

func (r *ConsoleReader) Close() {
	r.close()
}

type parsingStats struct {
	startAt  time.Time
	blockNum uint64
	data     map[string]int
}

func newParsingStats(block uint64) *parsingStats {
	return &parsingStats{
		startAt:  time.Now(),
		blockNum: block,
		data:     map[string]int{},
	}
}

func (s *parsingStats) log() {
	zlog.Info("mindreader block stats",
		zap.Uint64("block_num", s.blockNum),
		zap.Int64("duration", int64(time.Since(s.startAt))),
		zap.Reflect("stats", s.data),
	)
}

func (s *parsingStats) inc(key string) {
	if s == nil {
		return
	}
	k := strings.ToLower(key)
	value := s.data[k]
	value++
	s.data[k] = value
}

type parseCtx struct {
	currentBlock *pbcodec.Block

	stats *parsingStats
}

func newContext(height uint64) *parseCtx {
	return &parseCtx{
		currentBlock: &pbcodec.Block{
			Height: height,
			Txs:    []*pbcodec.Transaction{},
		},
	}

}

func (r *ConsoleReader) Read() (out interface{}, err error) {
	return r.next()
}

const (
	LogPrefix = "DMLOG"
	LogBlock  = "BLOCK"
)

func (r *ConsoleReader) next() (out interface{}, err error) {
	for line := range r.lines {
		if !strings.HasPrefix(line, LogPrefix) {
			continue
		}

		tokens := strings.Split(line[len(LogPrefix)+1:], " ")
		if len(tokens) < 2 {
			return nil, fmt.Errorf("invalid log line format: %s", line)
		}

		switch tokens[0] {
		case LogBlock:
			return r.block(tokens[1:])
		default:
			if tracer.Enabled() {
				zlog.Debug("skipping unknown deep mind log line", zap.String("line", line))
			}
			continue
		}

		if err != nil {
			chunks := strings.SplitN(line, " ", 2)
			return nil, fmt.Errorf("%s: %s (line %q)", chunks[0], err, line)
		}
	}

	zlog.Info("lines channel has been closed")
	return nil, io.EOF
}

func (r *ConsoleReader) ProcessData(reader io.Reader) error {
	scanner := r.buildScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		r.lines <- line
	}

	if scanner.Err() == nil {
		close(r.lines)
		return io.EOF
	}

	return scanner.Err()
}

func (r *ConsoleReader) buildScanner(reader io.Reader) *bufio.Scanner {
	buf := make([]byte, 50*1024*1024)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(buf, 50*1024*1024)

	return scanner
}

// Format:
// DMLOG BLOCK <HEIGHT> <ENCODED_BLOCK>
func (r *ConsoleReader) block(params []string) (*pbcodec.Block, error) {
	if err := validateChunk(params, 2); err != nil {
		return nil, fmt.Errorf("invalid log line length: %w", err)
	}

	// <HEIGHT>
	//
	// parse block height
	blockHeight, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid block num: %w", err)
	}

	r.ctx = newContext(blockHeight)

	// <ENCODED_BLOCK>
	//
	// hex decode block
	bytes, err := hex.DecodeString(params[1])
	if err != nil {
		return nil, fmt.Errorf("invalid encoded block: %w", err)
	}

	// decode bytes to Block
	err = proto.Unmarshal(bytes, r.ctx.currentBlock)

	if blockHeight != r.ctx.currentBlock.Height {
		return nil, fmt.Errorf("block height does not match active block height, got block height %d but current is block height %d", blockHeight, r.ctx.currentBlock.Height)
	}

	if err != nil {
		return nil, fmt.Errorf("invalid encoded block: %w", err)
	}

	// logging
	zlog.Info("console reader read block",
		zap.Uint64("height", r.ctx.currentBlock.Height),
		zap.String("hash", base64url.Encode(r.ctx.currentBlock.Hash)),
		zap.String("prev_hash", base64url.Encode(r.ctx.currentBlock.PreviousBlock)),
		zap.Int("trx_count", len(r.ctx.currentBlock.Txs)),
	)

	return r.ctx.currentBlock, nil
}

func validateChunk(params []string, count int) error {
	if len(params) != count {
		return fmt.Errorf("%d fields required but found %d", count, len(params))
	}
	return nil
}
