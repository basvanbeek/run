// Copyright (c) Bas van Beek 2024.
// Copyright (c) Tetrate, Inc 2022.
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

package test_test

import (
	"runtime"
	"sync"
	"testing"

	"github.com/basvanbeek/telemetry"

	"github.com/basvanbeek/run"
	"github.com/basvanbeek/run/pkg/test"
)

// TestIRQService test if IRQs returns a valid error for deliberate termination.
func TestIRQService(t *testing.T) {
	if ps := runtime.GOMAXPROCS(runtime.NumCPU()); ps < 3 {
		t.Skipf("GOMAXPROCS not sufficient for test: %d", ps)
	}
	var (
		g    = &run.Group{Name: "test", Logger: telemetry.NoopLogger()}
		irqs = test.NewIRQService(func() {})
	)

	g.Register(irqs)
	if err := g.RunConfig(); err != nil {
		t.Fatalf("configuring run.Group: %v", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	t.Run("primary thread", func(t *testing.T) {
		t.Parallel()
		defer wg.Done()

		if err := g.Run(); err != nil {
			t.Fatalf("server exit: %v", err)
		}
	})

	// Try to close the run group on primary thread from the secondary thread.
	t.Run("secondary thread", func(t *testing.T) {
		t.Parallel()
		defer func() {
			_ = irqs.Close()
			wg.Done()
		}()
	})

	t.Run("waiter thread", func(t *testing.T) {
		t.Parallel()
		wg.Wait()
	})
}
