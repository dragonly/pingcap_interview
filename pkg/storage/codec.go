package storage

import (
	"encoding/binary"
	"fmt"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"time"
)

type RecordGenerator struct {
	MaxKey      int64 // Record.Key 最大值
	DataSizeMin int   // Record.Data 字段最小长度
	DataSizeMax int   // Record.Data 字段最大长度
}

func (g *RecordGenerator) Init() {
	rand.Seed(time.Now().Unix())
	//rand.Seed(0)
}

func (g *RecordGenerator) Generate() Record {
	var dataLen int
	if g.DataSizeMax-g.DataSizeMin == 0 {
		dataLen = g.DataSizeMin
	} else {
		dataLen = g.DataSizeMin + rand.Int()%(g.DataSizeMax-g.DataSizeMin) // 这不是一个真正的均匀分布，不过在这个场景下影响不大
	}
	record := Record{
		Key:  rand.Int63() % g.MaxKey,
		Data: make([]byte, dataLen),
	}
	if n, err := rand.Read(record.Data); err != nil || n != dataLen {
		panic(fmt.Sprintf("err: %s, n: %d", err, n))
	}
	return record
}

type FileBlockWriter struct {
	DataFilenameBase string // data 文件名，后缀会添加 meta 或 block index
	BlockSize        int    // 文件分块大小，单位 byte
	BlockNum         int    // 最大块数
	BlockIndex       int    // 当前 block

	fData              *os.File // data 文件指针
	fMeta              *os.File // metadata 文件指针
	dataFileBytesWrote int      // 已经写入 data 文件的 byte 数
}

func (m *FileBlockWriter) rotateFiles() {
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
		m.BlockIndex++
		m.dataFileBytesWrote = 0
	}
	if m.BlockIndex < m.BlockNum {
		var err error
		if m.fMeta, err = os.Create(fmt.Sprintf("%s.%d.meta", m.DataFilenameBase, m.BlockIndex)); err != nil {
			panic(err)
		}
		if m.fData, err = os.Create(fmt.Sprintf("%s.%d.data", m.DataFilenameBase, m.BlockIndex)); err != nil {
			panic(err)
		}
	}
}

func (m *FileBlockWriter) writeFiles(record Record) {
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

func (m *FileBlockWriter) write(record Record) bool {
	if m.fMeta == nil {
		m.rotateFiles()
	}

	dataFileBytesRemaining := m.BlockSize - m.dataFileBytesWrote
	if dataFileBytesRemaining >= 8 /*Key*/ +len(record.Data) { // 当前 block 还有容量
		m.writeFiles(record)
		return true
	} else if m.BlockIndex < m.BlockNum-1 { // 还有新 block 容量
		log.Info().
			Int("remaining bytes", dataFileBytesRemaining).
			Int("data bytes", 8+len(record.Data)).
			Msgf("no capacity in current block")
		m.rotateFiles()
		m.writeFiles(record)
		return true
	} else { // BlockNum 已经写满
		return false
	}
}

// GenRecords 生成 Key 唯一且随机的包含 n 个 Record 的数组，data 数据随机
func GenRecords(n int, maxKey int64) []Record {
	rGen := RecordGenerator{
		MaxKey:      maxKey,
		DataSizeMin: 0,
		DataSizeMax: 0,
	}
	rGen.Init()
	var records []Record
	existingKeys := make(map[int64]struct{}, 1024)
	log.Debug().Msg("generating records")
	for {
		record := rGen.Generate()
		if _, exist := existingKeys[record.Key]; exist {
			//fmt.Println(record.Key)
			continue
		}
		existingKeys[record.Key] = struct{}{}
		records = append(records, record)
		if len(records) == n {
			break
		}
	}
	return records
}
