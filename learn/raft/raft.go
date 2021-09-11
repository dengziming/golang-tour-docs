package raft

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/rpc"
	"time"
)

type node struct {
	connect bool
	address string // 节点连接地址
}

type state int

const (
	Follower state = iota + 1
	Candidate
	Leader
)

type Raft struct {
	me int // 自己的 id
	nodes map[int]*node

	state state // 节点状态

	Term int

	votedFor int

	voteCount int // voters

	// heartbeat channel
	heartbeatC chan bool

	// to leader channel
	toLeaderC  chan bool



	// 日志条目集合
	log []LogEntry

	// 被提交的最大索引
	commitIndex int
	// 被应用到状态机的最大索引
	lastApplied int

	// 保存需要发送给每个节点的下一个条目索引
	nextIndex []int
	// 保存已经复制给每个节点日志的最高索引
	matchIndex []int

}

func newNode(address string) *node {
	node := &node{address: address}
	return node
}

func (rf *Raft) start() {
	// 启动 raft server 设置初识状态
	rf.state = Follower
	rf.Term = 0
	rf.votedFor = -1
	rf.heartbeatC = make(chan bool)
	rf.toLeaderC = make(chan bool)

	//
	rand.Seed(time.Now().UnixNano())

	for {
		switch rf.state {
		case Follower:
			select {
			case <- rf.heartbeatC: // 收到 heartbeat 请求
				log.Printf("follower-%d recived heartbeat\n", rf.me)

			case <- time.After(time.Duration(rand.Intn(500-300)+300) * time.Millisecond):
				log.Printf("follower-%d timeout\n", rf.me)
				rf.state = Candidate
			}

		case Leader:
			rf.broadcastHeartbeat()
			time.Sleep(100 * time.Millisecond)

		case Candidate:
			fmt.Printf("Node: %d, I'm candidate\n", rf.me)
			rf.Term++

			rf.votedFor = rf.me
			rf.voteCount = 1

			// 发送投票请求
			go rf.broadcastRequestVote()

			select {
			case <-time.After(time.Duration(rand.Intn(5000-300)+300) * time.Millisecond):
				rf.state = Follower

			case <- rf.toLeaderC:
				fmt.Printf("Node: %d, I'm leader\n", rf.me)
				rf.state = Leader

				// 初始化 peers 的 nextIndex 和 matchIndex
				rf.nextIndex = make([]int, len(rf.nodes))
				rf.matchIndex = make([]int, len(rf.nodes))

				// 为每个节点初始化 nextIndex 和 matchIndex
				// 这里不考虑leader重新选举的情况
				for i := range rf.nodes {
					rf.nextIndex[i] = 1
					rf.matchIndex[i] = 0
				}

				// 模拟客户端每3s发送一条command
				go func() {
					i := 0
					for {
						i++
						// leader收到客户端的command，封装成LogEntry并append 到 log
						rf.log = append(rf.log, LogEntry{rf.Term, i, fmt.Sprintf("user send : %d", i)})
						time.Sleep(10 * time.Second)
					}
				}()

			}
		}
	}
}

type VoteRequest struct {
	Term int
	CandidateID int
}

type VoteResponse struct {
	Term int
	VoteGranted bool
}

func (rf *Raft) broadcastRequestVote() {
	var request = VoteRequest {
		Term: rf.Term,
		CandidateID: rf.me,
	}

	for i := range rf.nodes {
		go func(i int) {
			var response VoteResponse
			rf.sendRequestVote(i, request, &response)
		}(i)
	}
}

func (rf *Raft) sendRequestVote(serverID int, args VoteRequest, reply *VoteResponse) {
	client, err := rpc.DialHTTP("tcp", rf.nodes[serverID].address)
	if err != nil {
		log.Fatal("rpc error", err)
	}
	defer client.Close()
	client.Call("Raft.RequestVote", args, reply)

	if reply.Term > rf.Term {
		rf.Term = reply.Term
		rf.state = Follower
		rf.votedFor = -1
		return
	}

	if reply.VoteGranted {
		// 票数 + 1
		rf.voteCount++
		// 获取票数大于集群的一半（N/2+1）
		if rf.voteCount > len(rf.nodes)/2+1 {
			// 发送通知给 toLeaderC channel
			rf.toLeaderC <- true
		}
	}

}

func (rf *Raft) RequestVote(args VoteRequest, reply *VoteResponse) error {

	if args.Term < rf.Term {
		reply.Term = rf.Term
		reply.VoteGranted = false
		return nil
	}

	// TODO 如果当前没有投票而且 log at least up to date with candidate
	if rf.votedFor == -1 {
		rf.Term = args.Term
		rf.votedFor = args.CandidateID
		reply.Term = rf.Term
		reply.VoteGranted = true
		return nil
	}

	reply.Term = rf.Term
	reply.VoteGranted = false
	return nil
}

type HeartbeatRequest struct {
	Term     int
	LeaderID int

	PrevLogIndex int
	PrevLogTerm int
	Entries []LogEntry
	LeaderCommit int
}

type HeartbeatResponse struct {
	Term int
	Success bool

	// 如果 Follower Index小于 Leader Index， 会告诉 Leader 下次开始发送的索引位置
	NextIndex int
}

// 广播 heartbeat
func (rf *Raft) broadcastHeartbeat() {
	// 遍历所有节点
	for i := range rf.nodes {
		// request 参数
		request := HeartbeatRequest{
			Term:     rf.Term,
			LeaderID: rf.me,
			LeaderCommit: rf.commitIndex,
		}

		prevLogIndex := rf.nextIndex[i] - 1
		if rf.getLastIndex() > prevLogIndex {
			// 提取 preLogIndex 之后的entry，发送给 follower
			request.PrevLogTerm = prevLogIndex
			request.PrevLogTerm = rf.log[prevLogIndex].LogTerm
			request.Entries = rf.log[prevLogIndex:]
		}

		go func(i int, args HeartbeatRequest) {
			var reply HeartbeatResponse
			// 向某一个节点发送 heartbeat
			rf.sendHeartbeat(i, args, &reply)
		}(i, request)
	}
}


// 向某一个节点发送 heartbeat
func (rf *Raft) sendHeartbeat(serverID int, request HeartbeatRequest, response *HeartbeatResponse) {
	// 创建 RPC client
	client, err := rpc.DialHTTP("tcp", rf.nodes[serverID].address)
	if err != nil {
		log.Fatal("rpc error:", err)
	}

	defer client.Close()
	// 调用 follower 节点 Heartbeat 方法
	client.Call("Raft.Heartbeat", request, response)

	// 如果follower节点term大于leader节点term
	// leader节点过时，leader转变为follower状态
	if response.Term > rf.Term {
		rf.Term = response.Term
		rf.state = Follower
		rf.votedFor = -1
	}

	if response.Success {
		if response.NextIndex > 0 {
			rf.nextIndex[serverID] = response.NextIndex
			rf.matchIndex[serverID] = rf.nextIndex[serverID] - 1
		}

		// TODO
		// 如果大于半数节点同步成功
		// 1. 更新 leader 节点的 commitIndex
		// 2. 返回给客户端
		// 3. 应用状态就机
		// 4. 通知 Followers Entry 已提交
	} else {
		if response.Term > rf.Term {
			rf.Term = response.Term
			rf.state = Follower
			rf.votedFor = -1
		}
	}

}

// Heartbeat rpc method
func (rf *Raft) Heartbeat(request HeartbeatRequest, response *HeartbeatResponse) error {

	// 如果leader节点term小于follower节点，不做处理并返回
	if request.Term < rf.Term {
		response.Term = rf.Term
		return nil
	}

	// 如果leader节点term大于follower节点
	// 说明 follower 过时，重置follower节点term
	if request.Term > rf.Term {
		rf.Term = request.Term
		rf.state = Follower
		rf.votedFor = -1
	}

	// 将当前 follower 节点term返回给leader
	response.Term = rf.Term

	// 心跳成功，将消息发给 heartbeatC 通道
	rf.heartbeatC <- true
	if len(request.Entries) == 0 {
		response.Success = true
		response.Term = rf.Term
		return nil
	}

	if request.PrevLogIndex > rf.getLastIndex() {
		response.Success = false
		response.Term = rf.Term
		response.NextIndex = rf.getLastIndex() + 1
		return nil
	}

	// TODO 处理 diverge 的情况

	// append leader 发送 log 到自己 log
	rf.log = append(rf.log, request.Entries...)
	rf.commitIndex = int(math.Min(float64(rf.getLastIndex()), float64(request.LeaderCommit)))

	// 返回 term 和 NextIndex
	response.Success = true
	response.Term = rf.Term
	response.NextIndex = rf.getLastIndex() + 1

	return nil
}

func (rf *Raft) getLastIndex() int {
	rlen := len(rf.log)
	if rlen == 0 {
		return 0
	}
	return rf.log[rlen-1].LogIndex
}

func (rf *Raft) getLastTerm() int {
	rlen := len(rf.log)
	if rlen == 0 {
		return 0
	}
	return rf.log[rlen-1].LogTerm
}



// LogEntry struct
type LogEntry struct {
	// 当前term
	LogTerm  int

	// 当前 index
	LogIndex int

	// 真正的消息体
	LogCMD   interface{}
}
