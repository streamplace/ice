// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build windows

package netutil

import (
	"errors"
	"syscall"
)

// Go's syscall.EADDRNOTAVAIL is an invented POSIX-compat constant that does not
// match the raw Winsock errno returned by the kernel, so we check both.
const wsaeaddrnotavail syscall.Errno = 10049

const wsaeaddrinuse syscall.Errno = 10048

// IsAddrInUse reports whether err indicates that a local port is already bound.
func IsAddrInUse(err error) bool {
	return errors.Is(err, syscall.EADDRINUSE) || errors.Is(err, wsaeaddrinuse)
}

// IsAddrUnavailable reports whether err indicates that the address
// is unavailable (as opposed to a specific port being busy).
func IsAddrUnavailable(err error) bool {
	var errno syscall.Errno
	if errors.As(err, &errno) {
		return errno == syscall.EADDRNOTAVAIL || errno == wsaeaddrnotavail
	}

	return false
}
