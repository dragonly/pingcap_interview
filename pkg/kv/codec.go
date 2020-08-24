package kv

import (
	"encoding/binary"
	"fmt"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
)

type RecordGenerator struct {
	DataSizeMin int // Record.Data 字段最小长度
	DataSizeMax int // Record.Data 字段最大长度
}

func (g *RecordGenerator) Generate() Record {
	dataLen := g.DataSizeMin + rand.Int()%(g.DataSizeMax-g.DataSizeMin) // 这不是一个真正的均匀分布，不过在这个场景下影响不大
	record := Record{
		Key:  rand.Int63(),
		Data: make([]byte, dataLen),
	}
	if n, err := rand.Read(record.Data); err != nil || n != dataLen {
		panic(fmt.Sprintf("err: %s, n: %d", err, n))
	}
	return record
}

type FileBlockManager struct {
	DataFilenameBase string // data 文件名，后缀会添加 meta 或 block index
	BlockSize        int    // 文件分块大小，单位 byte
	MaxBlockNum      int    // 最大块数

	fData              *os.File // data 文件指针
	fMeta              *os.File // metadata 文件指针
	blockIndex         int      // 当前 block
	dataFileBytesWrote int      // 已经写入 data 文件的 byte 数
}

func (m *FileBlockManager) rotateFiles() {
	if m.fMeta != nil {
		dataFileBytesRemaining := m.BlockSize - m.dataFileBytesWrote
		pad := make([]byte, dataFileBytesRemaining)
		if n, err := m.fData.Write(pad); err != nil || n != dataFileBytesRemaining {
			panic(fmt.Sprintf("err: %s, n: %d", err, n))
		}
		if err := m.fMeta.Sync(); err != nil {
			panic(err)
		}
		if err := m.fMeta.Close(); err != nil {
			panic(err)
		}
		if err := m.fData.Sync(); err != nil {
			panic(err)
		}
		if err := m.fData.Close(); err != nil {
			panic(err)
		}
		m.blockIndex++
		m.dataFileBytesWrote = 0
	}
	if m.blockIndex < m.MaxBlockNum {
		var err error
		if m.fMeta, err = os.Create(fmt.Sprintf("%s.%d.meta", m.DataFilenameBase, m.blockIndex)); err != nil {
			panic(err)
		}
		if m.fData, err = os.Create(fmt.Sprintf("%s.%d.data", m.DataFilenameBase, m.blockIndex)); err != nil {
			panic(err)
		}
	}
}

func (m *FileBlockManager) writeFiles(record Record) {
	keyBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(keyBytes, uint64(record.Key))
	if n, err := m.fData.Write(keyBytes); err != nil || n != 8 {
		panic(fmt.Sprintf("err: %s, n: %d", err, n))
	}
	if n, err := m.fData.Write(record.Data); err != nil || n != len(record.Data) {
		panic(fmt.Sprintf("err: %s, n: %d", err, n))
	}
	lenStr := fmt.Sprintf("%d\n", len(record.Data))
	if n, err := m.fMeta.WriteString(lenStr); err != nil || n != len(lenStr) {
		panic(fmt.Sprintf("err: %s, n: %d", err, n))
	}
	m.dataFileBytesWrote += 8 + len(record.Data)
}

func (m *FileBlockManager) write(record Record) bool {
	if m.fMeta == nil {
		m.rotateFiles()
	}

	dataFileBytesRemaining := m.BlockSize - m.dataFileBytesWrote
	if dataFileBytesRemaining >= 8 /*Key*/ +len(record.Data) { // 当前 block 还有容量
		m.writeFiles(record)
		return true
	} else if m.blockIndex < m.MaxBlockNum-1 { // 还有新 block 容量
		log.Info().
			Int("remaining bytes", dataFileBytesRemaining).
			Int("data bytes", 8+len(record.Data)).
			Msgf("no capacity in current block")
		m.rotateFiles()
		m.writeFiles(record)
		return true
	} else { // MaxBlockNum 已经写满
		return false
	}
}

// genRecordsFiles 生成分块的 record 文件，为了简化处理，暂时将跨当前文件 block 边缘的 record 放入下一个 block，
// 并将前一个 block 结尾 pad 成 0 字节
func genRecordsFiles(rGen RecordGenerator, fbMgr FileBlockManager) {
	for {
		record := rGen.Generate()
		if !fbMgr.write(record) {
			break
		}
	}
}

func GenRecordsFiles() {
	rGen := RecordGenerator{
		DataSizeMin: 1 * 1024,
		DataSizeMax: 100 * 1024,
	}
	fbMgr := FileBlockManager{
		DataFilenameBase: "data/test",
		BlockSize:        64 * 1024 * 1024,
		MaxBlockNum:      3,
	}
	genRecordsFiles(rGen, fbMgr)
}
