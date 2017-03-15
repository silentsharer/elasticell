// Copyright 2016 DeepFabric, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"fmt"

	"github.com/deepfabric/elasticell/pkg/redis"
	"github.com/fagongzi/goetty"
)

func (s *Server) startRedisAPIServer() error {
	select {
	case <-s.stopC:
		return s.stopRedisAPIServer()
	default:
		return s.s.Start(s.doConnection)
	}
}

func (s *Server) stopRedisAPIServer() error {
	// TODO: 考虑一致性问题
	s.s.Stop()
	return nil
}

func (s *Server) doConnection(session goetty.IOSession) error {
	for {
		req, err := session.Read()
		if err != nil {
			return err
		}

		cmd, _ := req.(*redis.Command)
		fmt.Printf("cmd: %s", cmd.Cmd)
		for _, arg := range cmd.Args {
			fmt.Printf(" %s", string(arg))
		}
		fmt.Println("")

		// session.Write(redis.StatusReply("OK"))

		return nil
	}
}