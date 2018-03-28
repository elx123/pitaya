package services

import (
	"fmt"

	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/examples/demo/cluster_protobuf/protos"
)

// ConnectorRemote is a remote that will receive rpc's
type ConnectorRemote struct {
	component.Base
}

// Connector struct
type Connector struct {
	component.Base
}

<<<<<<< HEAD
// SessionData is the session data struct
type SessionData struct {
	Data map[string]interface{} `json:"data"`
}

// GetSessionData gets the session data
func (c *Connector) GetSessionData(s *session.Session, data *SessionData) (map[string]interface{}, error) {
	return s.GetData(), nil
}

// SetSessionData sets the session data
func (c *Connector) SetSessionData(s *session.Session, data *SessionData) (string, error) {
	err := s.SetData(data.Data)
	if err != nil {
		return "", err
	}
	return "success", nil
}

// NotifySessionData sets the session data
func (c *Connector) NotifySessionData(s *session.Session, data *SessionData) {
	err := s.SetData(data.Data)
	if err != nil {
		fmt.Println("got error on notify", err)
	}
}

// RemoteFunc is a function that will be called remotelly
func (c *ConnectorRemote) RemoteFunc(message []byte) (*protos.Response, error) {
	fmt.Printf("received a remote call with this message: %s\n", message)
	return &protos.Response{
		Msg: string(message),
	}, nil
}