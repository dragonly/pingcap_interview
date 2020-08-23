package local

import "github.com/dragonly/pingcap_interview/pkg/kv"

type TopNSolver func(records []kv.Record, topN int) []kv.Record
