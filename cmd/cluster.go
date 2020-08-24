/*
Copyright © 2020 Li Yilong

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/dragonly/pingcap_interview/pkg/cluster"

	"github.com/spf13/cobra"
)

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cluster called")
	},
}

var startMapperCmd = &cobra.Command{
	Use:   "startMapper",
	Short: "启动 mapper 服务",
	Long: `start mapper server
mapper will listen on port 2333 to receive get-top-n request, with key range [min_k, max_k]
the calculation is done on each mapper for blocks of data on shared storage`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cluster startMapperCmd called")
		cluster.StartServer()
	},
}

var (
	pMinKey *int64
	pMaxKey *int64
)

var getTopNKeysInRangeCmd = &cobra.Command{
	Use:   "getTopNKeysInRange",
	Short: "get top n keys in range",
	Long: `从提供的 [min, max] 范围内，找到 key 最小的前 n 个记录
该计算过程采用 map-reduce 模型进行，当前进程会将所有数据分块调度给 mapper 节点进行计算，并在 reduce 过程完成后，将结果返回`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cluster getTopNKeysInRangeCmd called")
		fmt.Printf("minKey=%d, maxKey=%d\n", *pMinKey, *pMaxKey)
		cluster.GetTopNKeysInRange(*pMinKey, *pMaxKey)
	},
}

func init() {
	rootCmd.AddCommand(clusterCmd)

	clusterCmd.AddCommand(startMapperCmd)
	clusterCmd.AddCommand(getTopNKeysInRangeCmd)

	pMinKey = getTopNKeysInRangeCmd.Flags().Int64("minKey", -1, "min key, inclusive")
	pMaxKey = getTopNKeysInRangeCmd.Flags().Int64("maxKey", -1, "max key, inclusive")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
