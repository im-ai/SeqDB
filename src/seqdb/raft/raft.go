package raft

type RequestVoteArgs struct {
	term int
	candidateId int
	lastLogIndex int
	lastLogTerm int
}

type RequestVoteReply struct {
	currentTerm int
	granted bool
}

type Raft interface {
	RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) error
	AppendEntries(term int, leaderId int, prevLogIndex int, prevLogTerm int, entries string, leaderCommit int) (currentTerm int, success bool)
}

type RaftServer struct {
	myId int

	currentTerm int // latest term server has seen
	votedFor int // candidateId that received vote in current term (or null if none)
	// log[] // log entries; each entry contains command for state machine, and term when entry  was received by leader (first index is 1)
	commitIndex int //index of highest log entry known to be committed
	lastApplied int //index of highest log entry applied to state machine

	// Reset after election:
	nextIndex []int //for each server, index of the next log entry to send to that server
	matchIndex []int //for each server, index of highest log entry known to be replicated on server
}

func NewRaftServer() *RaftServer {
	return &RaftServer{
		currentTerm:0,
		votedFor:-1,
		commitIndex:0,
		lastApplied:0,
		nextIndex:make([]int, 0),
		matchIndex:make([]int, 0),
	}
}

