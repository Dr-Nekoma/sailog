package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "sailog/proto"
	raftServer "sailog/server/pkgs"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedRaftServer
	state raftServer.State
}

// RequestVote implements voting system to decide leader
func (s *server) RequestVote(ctx context.Context, in *pb.RequestVoteMessage) (*pb.ReplyVoteMessage, error) {
	log.Printf("Judging candidate %v...", in.CandidateId)

	// TODO: Next step is to use the global state to judge if the candidate should be accepted as leader

	return &pb.ReplyVoteMessage{
		Term: 1,
		VoteGranted: true}, nil
}

func main() {
	// Parse flags
	flag.Parse()

	// Create server instance
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	var initialState raftServer.State
	initialState.CurrentTerm = raftServer.Term{1, 0, 0}
	initialState.State = raftServer.FollowerState
	
	// Set initial state
	pb.RegisterRaftServer(s, &server{state: initialState})
	
	// Listening to port
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
