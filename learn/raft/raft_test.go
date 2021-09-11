package raft

import (
	"flag"
	"fmt"
	"strings"
	"testing"
)

func TestRaftVote(t *testing.T)  {

	port := flag.String("port", ":9091", "rpc listen port")
	cluster := flag.String("cluster", "127.0.0.1:9091", "comma sep")
	id := flag.Int("id", 1, "node ID")

	// 参数解析
	flag.Parse()
	clusters := strings.Split(*cluster, ",")
	ns := make(map[int]*node)
	for k, v := range clusters {
		ns[k] = newNode(v)
	}

	// 创建节点
	raft := &Raft{}
	raft.me = *id
	raft.nodes = ns

	// TODO 监听rpc
	fmt.Println(port)
	// raft.rpc(*port)
	// 开启 raft
	raft.start()

	select {}
}
