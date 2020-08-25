package storage

import (
	"bytes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math"
	"os"
	"testing"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.With().Caller().Logger()
}

func TestCodec(t *testing.T) {
	defer func() {
		os.Remove("test_codec.0.data")
		os.Remove("test_codec.0.meta")
	}()
	rGen := RecordGenerator{
		MaxKey:      math.MaxInt64,
		DataSizeMin: 1 * 1024,
		DataSizeMax: 100 * 1024,
	}
	rGen.Init()
	fbMgr := FileBlockWriter{
		DataFilenameBase: "test_codec",
		BlockSize:        64 * 1024 * 1024,
		BlockNum:         1,
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
		if !bytes.Equal(r.Data, rrb.Data) {
			t.Errorf("different data, i: %d\n", i)
			break
		}
	}
}
