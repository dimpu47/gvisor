// Copyright 2018 Google Inc.
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

package boot

import "fmt"

// PlatformType tells which platform to use.
type PlatformType int

const (
	// Ptrace runs the sandbox with the ptrace platform.
	PlatformPtrace PlatformType = iota

	// KVM runs the sandbox with the KVM platform.
	PlatformKVM
)

// MakePlatformType converts type from string.
func MakePlatformType(s string) (PlatformType, error) {
	switch s {
	case "ptrace":
		return PlatformPtrace, nil
	case "kvm":
		return PlatformKVM, nil
	default:
		return 0, fmt.Errorf("invalid platform type %q", s)
	}
}

func (p PlatformType) String() string {
	switch p {
	case PlatformPtrace:
		return "ptrace"
	case PlatformKVM:
		return "kvm"
	default:
		return fmt.Sprintf("unknown(%d)", p)
	}
}

// FileAccessType tells how the filesystem is accessed.
type FileAccessType int

const (
	// FileAccessProxy sends IO requests to a Gofer process that validates the
	// requests and forwards them to the host.
	FileAccessProxy FileAccessType = iota

	// FileAccessDirect connects the sandbox directly to the host filesystem.
	FileAccessDirect
)

// MakeFileAccessType converts type from string.
func MakeFileAccessType(s string) (FileAccessType, error) {
	switch s {
	case "proxy":
		return FileAccessProxy, nil
	case "direct":
		return FileAccessDirect, nil
	default:
		return 0, fmt.Errorf("invalid file access type %q", s)
	}
}

func (f FileAccessType) String() string {
	switch f {
	case FileAccessProxy:
		return "proxy"
	case FileAccessDirect:
		return "direct"
	default:
		return fmt.Sprintf("unknown(%d)", f)
	}
}

// NetworkType tells which network stack to use.
type NetworkType int

const (
	// NetworkSandbox uses internal network stack, isolated from the host.
	NetworkSandbox NetworkType = iota

	// NetworkHost redirects network related syscalls to the host network.
	NetworkHost

	// NetworkNone sets up just loopback using netstack.
	NetworkNone
)

// MakeNetworkType converts type from string.
func MakeNetworkType(s string) (NetworkType, error) {
	switch s {
	case "sandbox":
		return NetworkSandbox, nil
	case "host":
		return NetworkHost, nil
	case "none":
		return NetworkNone, nil
	default:
		return 0, fmt.Errorf("invalid network type %q", s)
	}
}

func (n NetworkType) String() string {
	switch n {
	case NetworkSandbox:
		return "sandbox"
	case NetworkHost:
		return "host"
	case NetworkNone:
		return "none"
	default:
		return fmt.Sprintf("unknown(%d)", n)
	}
}

// Config holds configuration that is not part of the runtime spec.
type Config struct {
	// RootDir is the runtime root directory.
	RootDir string

	// FileAccess indicates how the filesystem is accessed.
	FileAccess FileAccessType

	// Overlay is whether to wrap the root filesystem in an overlay.
	Overlay bool

	// Network indicates what type of network to use.
	Network NetworkType

	// LogPackets indicates that all network packets should be logged.
	LogPackets bool

	// Platform is the platform to run on.
	Platform PlatformType

	// Strace indicates that strace should be enabled.
	Strace bool

	// StraceSyscalls is the set of syscalls to trace.  If StraceEnable is
	// true and this list is empty, then all syscalls will be traced.
	StraceSyscalls []string

	// StraceLogSize is the max size of data blobs to display.
	StraceLogSize uint

	// DisableSeccomp indicates whether seccomp syscall filters should be
	// disabled. Pardon the double negation, but default to enabled is important.
	DisableSeccomp bool
}