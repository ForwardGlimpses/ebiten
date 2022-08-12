// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build nintendosdk
// +build nintendosdk

package audio

import (
	"io"

	"github.com/hajimehoshi/oto/v2/mux"

	"github.com/hajimehoshi/ebiten/v2/internal/nintendosdk"
)

func newContext(sampleRate, channelCount, bitDepthInBytes int) (context, chan struct{}, error) {
	ready := make(chan struct{})
	close(ready)

	c := &contextProxy{mux.New(sampleRate, channelCount, bitDepthInBytes)}
	nintendosdk.OpenAudio(sampleRate, channelCount, c.mux.ReadFloat32s)
	return c, ready, nil
}

type contextProxy struct {
	mux *mux.Mux
}

// NewPlayer implements context.
func (c *contextProxy) NewPlayer(r io.Reader) player {
	return c.mux.NewPlayer(r)
}

func (c *contextProxy) Suspend() error {
	// Do nothing so far.
	return nil
}

func (c *contextProxy) Resume() error {
	// Do nothing so far.
	return nil
}

func (c *contextProxy) Err() error {
	return nil
}
