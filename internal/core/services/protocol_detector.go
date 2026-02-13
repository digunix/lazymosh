// Copyright 2025.
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

package services

import (
	"os/exec"
	"sync"
)

// ProtocolDetector detects the availability of network protocols like mosh.
type ProtocolDetector struct {
	moshAvailable *bool
	mu            sync.RWMutex
}

// NewProtocolDetector creates a new ProtocolDetector instance.
func NewProtocolDetector() *ProtocolDetector {
	return &ProtocolDetector{}
}

// IsMoshAvailable checks if the mosh binary is available in the system PATH.
// The result is cached after the first check for performance.
func (pd *ProtocolDetector) IsMoshAvailable() bool {
	pd.mu.RLock()
	if pd.moshAvailable != nil {
		cached := *pd.moshAvailable
		pd.mu.RUnlock()
		return cached
	}
	pd.mu.RUnlock()

	// Check if mosh is in PATH (cross-platform)
	_, err := exec.LookPath("mosh")
	available := err == nil

	pd.mu.Lock()
	pd.moshAvailable = &available
	pd.mu.Unlock()

	return available
}
