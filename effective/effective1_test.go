package effective

import (
	"fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	"testing"
)


type Message struct {
	// Type MessageType `protobuf:"varint,1,opt,name=type,enum=raftpb.MessageType" json:"type"`
	To   uint64      `protobuf:"varint,2,opt,name=to" json:"to"`
	From uint64      `protobuf:"varint,3,opt,name=from" json:"from"`
	Term uint64      `protobuf:"varint,4,opt,name=term" json:"term"`
	// logTerm is generally used for appending Raft logs to followers. For example,
	// (type=MsgApp,index=100,logTerm=5) means leader appends entries starting at
	// index=101, and the term of entry at index 100 is 5.
	// (type=MsgAppResp,reject=true,index=100,logTerm=5) means follower rejects some
	// entries from its leader as it already has an entry with term 5 at index 100.
	LogTerm    uint64   `protobuf:"varint,5,opt,name=logTerm" json:"logTerm"`
	Index      uint64   `protobuf:"varint,6,opt,name=index" json:"index"`
	// Entries    []Entry  `protobuf:"bytes,7,rep,name=entries" json:"entries"`
	Commit     uint64   `protobuf:"varint,8,opt,name=commit" json:"commit"`
	// Snapshot   Snapshot `protobuf:"bytes,9,opt,name=snapshot" json:"snapshot"`
	Reject     bool     `protobuf:"varint,10,opt,name=reject" json:"reject"`
	RejectHint uint64   `protobuf:"varint,11,opt,name=rejectHint" json:"rejectHint"`
	Context    []byte   `protobuf:"bytes,12,opt,name=context" json:"context,omitempty"`
}


func TestMessage(t *testing.T) {
	message := Message{
		To: 10,
	}

	fmt.Println(message)
}