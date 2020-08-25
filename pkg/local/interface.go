package local

import "github.com/dragonly/pingcap_interview/pkg/storage"

type TopNSolver func(records []storage.Record, topN int) []storage.Record
