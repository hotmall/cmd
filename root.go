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
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	ProcName      string // ProcName represents the service name
	Version       string // Version define software version
	CommitHash    string // CommitHash represents the Git commit hash at built time
	BuildDate     string // BuildDate represents the date when this tool was built
	GoVersion     string // GoVersion represents go version
	prefix        string // prefix path
	defaultPrefix string // default prefix path
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     ProcName,
	Version: Version,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("[cmd.init] getwd fail, err = %v\n", err)
	}
	defaultPrefix = wd
	rootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", defaultPrefix, "set prefix path")
}

func PrefixPath() string {
	p, err := filepath.Abs(prefix)
	if err != nil {
		log.Printf("[cmd.PrefixPath] get prefix path fail, prefix=%s, err=%v\n", prefix, err)
		p = filepath.Clean(prefix)
	}
	return p
}

func LogPath() string {
	return filepath.Join(PrefixPath(), "var", "log")
}

func ConfigFile() string {
	return filepath.Join(PrefixPath(), "etc", "conf", "Hotfile")
}
