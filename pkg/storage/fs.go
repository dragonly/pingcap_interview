package storage

import (
	"encoding/binary"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadRecordsFile(dataFilenameBase string, blockIndex int64) []Record {
	metadataFilename := fmt.Sprintf("%s.%d.meta", dataFilenameBase, blockIndex)
	dataFilename := fmt.Sprintf("%s.%d.data", dataFilenameBase, blockIndex)
	var err error
	metadataBytes, err := ioutil.ReadFile(metadataFilename)
	if err != nil {
		panic(err)
	}
	dataBytes, err := ioutil.ReadFile(dataFilename)
	if err != nil {
		panic(err)
	}
	metadataStr := string(metadataBytes)
	metadataStr = metadataStr[:len(metadataStr)-1] // remove last empty line
	metadataStrArr := strings.Split(metadataStr, "\n")
	recordLenArr := make([]int, len(metadataStrArr))
	for i, recordLenStr := range metadataStrArr {
		recordLenArr[i], err = strconv.Atoi(recordLenStr)
		if err != nil {
			panic(err)
		}
	}
	bytesRead := 0
	records := make([]Record, len(recordLenArr))
	for i, recordLen := range recordLenArr {
		keyBytes := dataBytes[bytesRead : bytesRead+8]
		records[i].Key = int64(binary.LittleEndian.Uint64(keyBytes))
		records[i].Data = dataBytes[bytesRead+8 : bytesRead+8+recordLen]
		bytesRead += 8 + recordLen
	}
	return records
}

// genRecordsFiles 生成分块的 record 文件，为了简化处理，暂时将跨当前文件 block 边缘的 record 放入下一个 block，
// 并将前一个 block 结尾 pad 成 0 字节
func genRecordsFiles(rGen RecordGenerator, fbMgr FileBlockWriter, debug bool) []Record {
	var records []Record
	existingKeys := make(map[int64]struct{}, 1024)
	for {
		record := rGen.Generate()
		if _, exist := existingKeys[record.Key]; exist {
			continue
		}
		existingKeys[record.Key] = struct{}{}
		if !fbMgr.write(record) {
			break
		}
		if debug {
			records = append(records, record)
		}
	}
	return records
}

func GenRecordsFiles(startBlockIndex int) {
	maxKey := viper.GetInt64("cluster.data.record.maxKey")
	dataSizeMin := viper.GetInt("cluster.data.record.dataSizeMin")
	dataSizeMax := viper.GetInt("cluster.data.record.dataSizeMax")

	dataFilenameBase := viper.GetString("cluster.data.file.path")
	blockSize := viper.GetInt("cluster.data.file.blockSize")
	blockNum := viper.GetInt("cluster.data.file.blockNum")
	log.Info().
		Dict("record generator", zerolog.Dict().
			Int64("maxKey", maxKey).
			Int("dataSizeMin", dataSizeMin).
			Int("dataSizeMax", dataSizeMax)).
		Dict("file writer", zerolog.Dict().
			Str("dataFilenameBase", dataFilenameBase).
			Int("blockSize", blockSize).
			Int("blockNum", blockNum)).
		Msg("generating records files with parameters")
	rGen := RecordGenerator{
		MaxKey:      maxKey,
		DataSizeMin: dataSizeMin,
		DataSizeMax: dataSizeMax,
	}
	rGen.Init()
	fbMgr := FileBlockWriter{
		DataFilenameBase: dataFilenameBase,
		BlockSize:        blockSize,
		BlockNum:         blockNum,
		BlockIndex:       startBlockIndex,
	}
	genRecordsFiles(rGen, fbMgr, false)
}
