package kv

import (
	"github.com/rs/zerolog/log"
	"os"
	"testing"
)

func TestCodec(t *testing.T) {
	defer func() {
		os.Remove("test_codec.0.data")
		os.Remove("test_codec.0.meta")
	}()
	rGen := RecordGenerator{
		DataSizeMin: 1 * 1024,
		DataSizeMax: 100 * 1024,
	}
	rGen.Init()
	fbMgr := FileBlockWriter{
		DataFilenameBase: "test_codec",
		BlockSize:        64 * 1024 * 1024,
		MaxBlockNum:      1,
	}
	records := genRecordsFiles(rGen, fbMgr, true)
	recordsReadBack := ReadRecordsFile("test_codec", 0)
	log.Info().Msgf("len(records) = %d", len(records))
	for i, r := range records {
		rrb := recordsReadBack[i]
		if r.Key != rrb.Key {
			t.Errorf("different key, i: %d\nrecord: %v\nrecordReadBack: %v\n", i, r, rrb)
			break
		}
	}
}
