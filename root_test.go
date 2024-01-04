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
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixPath(t *testing.T) {
	assert := assert.New(t)
	items := []struct {
		Input  string
		Output string
	}{
		{"/foo/bar/baz", "/foo/bar/baz"},
		{"/foo/bar/baz/", "/foo/bar/baz"},
		{"/foo/bar/baz.js", "/foo/bar/baz.js"},
		{"/foo/bar/../baz", "/foo/baz"},
		{"/foo/bar/./baz", "/foo/bar/baz"},
	}
	for _, item := range items {
		prefix = item.Input
		r := PrefixPath()
		assert.Equal(item.Output, r)
	}
}

func TestLogPath(t *testing.T) {
	assert := assert.New(t)
	items := []struct {
		Input  string
		Output []string
	}{
		{"/foo/bar/baz", []string{"/foo/bar/baz", "var", "log"}},
		{"/foo/bar/baz/", []string{"/foo/bar/baz", "var", "log"}},
		{"/foo/bar/../baz", []string{"/foo/baz", "var", "log"}},
		{"/foo/bar/./baz", []string{"/foo/bar/baz", "var", "log"}},
	}

	for _, item := range items {
		prefix = item.Input
		lp := LogPath()
		expected := filepath.Join(item.Output...)
		assert.Equal(expected, lp)
	}
}

func TestConfigFile(t *testing.T) {
	assert := assert.New(t)
	items := []struct {
		Input  string
		Output []string
	}{
		{"/foo/bar/baz", []string{"/foo/bar/baz", "etc", "conf", "Hotfile"}},
		{"/foo/bar/baz/", []string{"/foo/bar/baz", "etc", "conf", "Hotfile"}},
		{"/foo/bar/../baz", []string{"/foo/baz", "etc", "conf", "Hotfile"}},
		{"/foo/bar/./baz", []string{"/foo/bar/baz", "etc", "conf", "Hotfile"}},
	}

	for _, item := range items {
		prefix = item.Input
		lp := ConfigFile()
		expected := filepath.Join(item.Output...)
		assert.Equal(expected, lp)
	}
}
