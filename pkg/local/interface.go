package local

import "github.com/dragonly/pingcap_interview/pkg/kv"

type TopNSolver func(r []kv.Record, n int) []kv.Record
