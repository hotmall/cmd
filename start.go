// Copyright © 2024 The Hot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

const (
	defaultPort = 53720
)

var (
	port int
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start one stopped service",
	Run: func(cmd *cobra.Command, args []string) {
		writeProcID(LogPath())

		// Set up channel on which to send signal notifications.
		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, os.Interrupt)

		addr := fmt.Sprintf("localhost:%d", port)
		server := &http.Server{
			Addr:    addr,
			Handler: nil,
		}

		log.Printf("[cmd.Run] starting server, listen on %s\n", addr)
		go func() {
			err := server.ListenAndServe()
			if err != nil {
				log.Fatalf("[cmd.Run] listen and serve fail, addr=%s, err=%v\n", addr, err)
			}
		}()

		// Block until interrupt signal is received.
		quitSignal := <-quitChan
		log.Println("[cmd.Run] Get signal:", quitSignal)
		server.Close()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntVar(&port, "port", defaultPort, "set the service listening port")
}

// pidFile return procid file
func pidFile(logPath string) (pidFile string) {
	if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
		log.Fatalf("[cmd.pidFile] make all dir fail, logPath=%s, err=%v\n", logPath, err)
	}
	pidFile = filepath.Join(logPath, ProcName+".pid")
	return
}

// writeProcID write procID into xx.pid file
func writeProcID(logPath string) {
	pidFile := pidFile(logPath)
	pid := strconv.Itoa(os.Getpid())
	err := os.WriteFile(pidFile, []byte(pid), 0644)
	if err != nil {
		log.Fatalf("[cmd.writeProcID] write pid file fail, pidFile=%s, err=%v\n", pidFile, err)
	}
}

// readProcID read procID from xx.pid file
func readProcID(logPath string) int {
	pidFile := pidFile(logPath)

	content, err := os.ReadFile(pidFile)
	if err != nil {
		log.Fatalf("[cmd.readProcID] read pid file fail, pidFile=%s, err=%v\n", pidFile, err)
	}

	procID, err := strconv.Atoi(string(content))
	if err != nil {
		log.Fatalf("[cmd.readProcID] strconv procID fail, strProcID=%s, err=%v\n", string(content), err)
	}
	return procID
}

// exit send a SIGINT(2) signal to the process
func exit(logPath string) {
	pid := readProcID(logPath)
	p, err := os.FindProcess(pid)
	if err != nil {
		log.Fatalf("[cmd.exit] find process fail, pid=%d, err=%v\n", pid, err)
	}

	if err = p.Signal(os.Interrupt); err != nil {
		log.Fatalf("[cmd.exit] send interrupt signal fail, pid=%d, err=%v\n", pid, err)
	}
}

func Port() int {
	return port
}
