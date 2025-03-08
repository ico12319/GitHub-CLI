package commits

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Commits struct {
	Comms []Commit
}

func parseCommitsInfo(response *http.Response) (*Commits, error) {
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not ok response code %v", response.StatusCode)
	}
	var com []Commit
	err := json.NewDecoder(response.Body).Decode(&com)
	if err != nil {
		return nil, err
	}
	return &Commits{Comms: com}, nil
}

func NewCommits(response *http.Response) (*Commits, error) {
	comm, err := parseCommitsInfo(response)
	if err != nil {
		return nil, err
	}
	return &Commits{Comms: comm.Comms}, nil
}
