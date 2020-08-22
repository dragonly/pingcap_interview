# pingcap 面试小作业

原题如下

> 内存中的行列结构的数据集，存在主键 k，求 TopN 算法
> 
> 上述题目在多核环境下的优化
> 
> 数据集大小为 1TB，分布规律未知。存储在某存储服务上，以 get(min_k, max_k) 接口获取数据，求多台服务器的计算方案


# 题目理解与疑问

首先尝试形式化定义原题，确保在实现前与题目原义对齐。这里使用 golang 作为实现语言。

## 约定

内存中的行列数据结构以行为数据记录单元，可以使用键值对的形式表示，其中 key 为主键，data 为该行数据。由于需要求 TopN，因此所有行的集合形成一个*全序集*。暂且不考虑行数据的具体存储编码，将逻辑结构定义为

```go
type Record struct {
	Key   int    // 主键，排序字段
	Data  []byte // 数据
}
```

## 题目形式化定义

在以上约定的前提下，这里尝试将原题定义得更形式化，便于消除歧义。

### 题目
内存中有数据集 `records []Record`，两个 `Record` `R1` 和 `R2` 的排序等价于 `rank` 字段的排序，即 `R1 > R2` 当且仅当 `R1.key > R2.key`，每个 record 的主键唯一。基于以上定义，完成下列三道题目：
1. 设计单核情况下，查询 `records` 中最小的 `N` 个值的 TopN 算法。
2. 在多核情况下，优化 1 中的算法。
3. 考虑分布式情况：数据集大小为 1TB，主键顺序随机，存储在存储服务 `Store` 上。用户以 `get(n, min_k, max_k)` 接口获取主键范围在 `[min_k, max_k]` 内的 TopN 数据，求多台计算节点、一个存储服务、一个请求服务的计算方案。
