// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package netutil

import (
	"errors"
	"fmt"
	"runtime"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

var errAddressInUseText = errors.New("address already in use")

func TestIsAddrInUse(t *testing.T) {
	for _, tt := range []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"in use", syscall.EADDRINUSE, true},
		{"wrapped", fmt.Errorf("bind: %w", syscall.EADDRINUSE), true},
		{"unavailable", syscall.EADDRNOTAVAIL, false},
		{"permission", syscall.EACCES, false},
		{"text only", errAddressInUseText, false},
		{"winsock", syscall.Errno(10048), runtime.GOOS == "windows"},
		{"wrapped winsock", fmt.Errorf("bind winsock: %w", syscall.Errno(10048)), runtime.GOOS == "windows"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, IsAddrInUse(tt.err))
		})
	}
}
