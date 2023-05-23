package server

import (
	"errors"
)

type State struct {
	CurrentTerm Term
	State ServerState
}

type ServerState string

const (
  LeaderState ServerState = "leader"
  FollowerState ServerState = "follower"
  CandidateState ServerState = "candidate"
)

type (
  Leader ServerState
  Follower ServerState
  Candidate ServerState
)

func ServerStateSwitch[R any](fieldType ServerState,
	leader func(Leader) (R, error),
	follower func(Follower) (R, error),
	candidate func(Candidate) (R, error),
) (res R, err error) {
  switch fieldType {
  case LeaderState:
    if leader != nil {
      return leader("")
    }
  case FollowerState:
    if follower != nil {
      return follower("")
    }
  case CandidateState:
    if candidate != nil {
      return candidate("")
    }
  default:
	  return res, errors.New("unsupported server state")
  }

  // If we get here, it's because we provided a nil function for a
  // type of custom field, implying we don't want to handle it.
  return res, nil
}

type Term struct {
	Id    uint
	Begin int32
	End   int32
}
