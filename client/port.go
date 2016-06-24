// Copyright (C) 2016 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"strconv"

	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/ishidawataru/opennsl-server/client/proto/port"
	"github.com/ishidawataru/opennsl-server/client/proto/portservice"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var portClient portservice.PortClient

func FormatPortConfig(c *port.PortConfig) string {
	s := bytes.NewBuffer(make([]byte, 0, 64))
	s.WriteString("      ")
	for i := 0; i < PBMP_WORD_MAX; i++ {
		s.WriteString("v               v               ")
	}
	s.WriteString("\nFe:   ")
	s.WriteString(PBMP(c.Fe).String())
	s.WriteString("\nGe:   ")
	s.WriteString(PBMP(c.Ge).String())
	s.WriteString("\nXe:   ")
	s.WriteString(PBMP(c.Xe).String())
	s.WriteString("\nCe:   ")
	s.WriteString(PBMP(c.Ce).String())
	s.WriteString("\nE:    ")
	s.WriteString(PBMP(c.E).String())
	s.WriteString("\nHg:   ")
	s.WriteString(PBMP(c.Hg).String())
	s.WriteString("\nPort: ")
	s.WriteString(PBMP(c.Port).String())
	s.WriteString("\nCPU:  ")
	s.WriteString(PBMP(c.Cpu).String())
	s.WriteString("\nAll:  ")
	s.WriteString(PBMP(c.All).String())
	return s.String()
}

func NewPortCmd() *cobra.Command {
	portCmd := &cobra.Command{
		Use: "port",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			conn, err := connGrpc()
			if err != nil {
				log.Fatal(err)
			}
			portClient = portservice.NewPortClient(conn)
		},
	}

	initCmd := &cobra.Command{
		Use: "init",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := portClient.Init(context.Background(), &port.InitRequest{})
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	clearCmd := &cobra.Command{
		Use: "clear",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := portClient.Clear(context.Background(), &port.ClearRequest{})
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	// probe
	// detach

	getConfig := &cobra.Command{
		Use: "config",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := portClient.GetConfig(context.Background(), &port.GetConfigRequest{})
			if err != nil {
				log.Fatal(err)
			}
			log.Info("reponse:\n", FormatPortConfig(res.Config))
		},
	}

	getPortName := &cobra.Command{
		Use: "name",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				log.Fatal("usage: name <port>")
			}
			i, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			res, err := portClient.GetPortName(context.Background(), &port.GetPortNameRequest{
				Port: int64(i),
			})
			log.Info("response:", res)
		},
	}

	portCmd.AddCommand(initCmd, clearCmd, getConfig, getPortName)
	return portCmd
}
