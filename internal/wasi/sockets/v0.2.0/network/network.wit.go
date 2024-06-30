// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package network represents the imported interface "wasi:sockets/network@0.2.0".
package network

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Network represents the imported resource "wasi:sockets/network@0.2.0#network".
//
// An opaque resource that represents access to (a subset of) the network.
// This enables context-based security for networking.
// There is no need for this to map 1:1 to a physical network interface.
//
//	resource network
type Network cm.Resource

// ResourceDrop represents the imported resource-drop for resource "network".
//
// Drops a resource handle.
//
//go:nosplit
func (self Network) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetworkResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasi:sockets/network@0.2.0 [resource-drop]network
//go:noescape
func wasmimport_NetworkResourceDrop(self0 uint32)

// ErrorCode represents the enum "wasi:sockets/network@0.2.0#error-code".
//
// Error codes.
//
// In theory, every API can return any error code.
// In practice, API's typically only return the errors documented per API
// combined with a couple of errors that are always possible:
// - `unknown`
// - `access-denied`
// - `not-supported`
// - `out-of-memory`
// - `concurrency-conflict`
//
// See each individual API for what the POSIX equivalents are. They sometimes differ
// per API.
//
//	enum error-code {
//		unknown,
//		access-denied,
//		not-supported,
//		invalid-argument,
//		out-of-memory,
//		timeout,
//		concurrency-conflict,
//		not-in-progress,
//		would-block,
//		invalid-state,
//		new-socket-limit,
//		address-not-bindable,
//		address-in-use,
//		remote-unreachable,
//		connection-refused,
//		connection-reset,
//		connection-aborted,
//		datagram-too-large,
//		name-unresolvable,
//		temporary-resolver-failure,
//		permanent-resolver-failure
//	}
type ErrorCode uint8

const (
	// Unknown error
	ErrorCodeUnknown ErrorCode = iota

	// Access denied.
	//
	// POSIX equivalent: EACCES, EPERM
	ErrorCodeAccessDenied

	// The operation is not supported.
	//
	// POSIX equivalent: EOPNOTSUPP
	ErrorCodeNotSupported

	// One of the arguments is invalid.
	//
	// POSIX equivalent: EINVAL
	ErrorCodeInvalidArgument

	// Not enough memory to complete the operation.
	//
	// POSIX equivalent: ENOMEM, ENOBUFS, EAI_MEMORY
	ErrorCodeOutOfMemory

	// The operation timed out before it could finish completely.
	ErrorCodeTimeout

	// This operation is incompatible with another asynchronous operation that is already
	// in progress.
	//
	// POSIX equivalent: EALREADY
	ErrorCodeConcurrencyConflict

	// Trying to finish an asynchronous operation that:
	// - has not been started yet, or:
	// - was already finished by a previous `finish-*` call.
	//
	// Note: this is scheduled to be removed when `future`s are natively supported.
	ErrorCodeNotInProgress

	// The operation has been aborted because it could not be completed immediately.
	//
	// Note: this is scheduled to be removed when `future`s are natively supported.
	ErrorCodeWouldBlock

	// The operation is not valid in the socket's current state.
	ErrorCodeInvalidState

	// A new socket resource could not be created because of a system limit.
	ErrorCodeNewSocketLimit

	// A bind operation failed because the provided address is not an address that the
	// `network` can bind to.
	ErrorCodeAddressNotBindable

	// A bind operation failed because the provided address is already in use or because
	// there are no ephemeral ports available.
	ErrorCodeAddressInUse

	// The remote address is not reachable
	ErrorCodeRemoteUnreachable

	// The TCP connection was forcefully rejected
	ErrorCodeConnectionRefused

	// The TCP connection was reset.
	ErrorCodeConnectionReset

	// A TCP connection was aborted.
	ErrorCodeConnectionAborted

	// The size of a datagram sent to a UDP socket exceeded the maximum
	// supported size.
	ErrorCodeDatagramTooLarge

	// Name does not exist or has no suitable associated IP addresses.
	ErrorCodeNameUnresolvable

	// A temporary failure in name resolution occurred.
	ErrorCodeTemporaryResolverFailure

	// A permanent failure in name resolution occurred.
	ErrorCodePermanentResolverFailure
)

// IPAddressFamily represents the enum "wasi:sockets/network@0.2.0#ip-address-family".
//
//	enum ip-address-family {
//		ipv4,
//		ipv6
//	}
type IPAddressFamily uint8

const (
	// Similar to `AF_INET` in POSIX.
	IPAddressFamilyIPv4 IPAddressFamily = iota

	// Similar to `AF_INET6` in POSIX.
	IPAddressFamilyIPv6
)

// IPv4Address represents the tuple "wasi:sockets/network@0.2.0#ipv4-address".
//
//	type ipv4-address = tuple<u8, u8, u8, u8>
type IPv4Address [4]uint8

// IPv6Address represents the tuple "wasi:sockets/network@0.2.0#ipv6-address".
//
//	type ipv6-address = tuple<u16, u16, u16, u16, u16, u16, u16, u16>
type IPv6Address [8]uint16

// IPAddress represents the variant "wasi:sockets/network@0.2.0#ip-address".
//
//	variant ip-address {
//		ipv4(ipv4-address),
//		ipv6(ipv6-address),
//	}
type IPAddress cm.Variant[uint8, IPv6Address, IPv6Address]

// IPAddressIPv4 returns a [IPAddress] of case "ipv4".
func IPAddressIPv4(data IPv4Address) IPAddress {
	return cm.New[IPAddress](0, data)
}

// IPv4 returns a non-nil *[IPv4Address] if [IPAddress] represents the variant case "ipv4".
func (self *IPAddress) IPv4() *IPv4Address {
	return cm.Case[IPv4Address](self, 0)
}

// IPAddressIPv6 returns a [IPAddress] of case "ipv6".
func IPAddressIPv6(data IPv6Address) IPAddress {
	return cm.New[IPAddress](1, data)
}

// IPv6 returns a non-nil *[IPv6Address] if [IPAddress] represents the variant case "ipv6".
func (self *IPAddress) IPv6() *IPv6Address {
	return cm.Case[IPv6Address](self, 1)
}

// IPv4SocketAddress represents the record "wasi:sockets/network@0.2.0#ipv4-socket-address".
//
//	record ipv4-socket-address {
//		port: u16,
//		address: ipv4-address,
//	}
type IPv4SocketAddress struct {
	// sin_port
	Port uint16

	// sin_addr
	Address IPv4Address
}

// IPv6SocketAddress represents the record "wasi:sockets/network@0.2.0#ipv6-socket-address".
//
//	record ipv6-socket-address {
//		port: u16,
//		flow-info: u32,
//		address: ipv6-address,
//		scope-id: u32,
//	}
type IPv6SocketAddress struct {
	// sin6_port
	Port uint16

	// sin6_flowinfo
	FlowInfo uint32

	// sin6_addr
	Address IPv6Address

	// sin6_scope_id
	ScopeID uint32
}

// IPSocketAddress represents the variant "wasi:sockets/network@0.2.0#ip-socket-address".
//
//	variant ip-socket-address {
//		ipv4(ipv4-socket-address),
//		ipv6(ipv6-socket-address),
//	}
type IPSocketAddress cm.Variant[uint8, IPv6SocketAddress, IPv6SocketAddress]

// IPSocketAddressIPv4 returns a [IPSocketAddress] of case "ipv4".
func IPSocketAddressIPv4(data IPv4SocketAddress) IPSocketAddress {
	return cm.New[IPSocketAddress](0, data)
}

// IPv4 returns a non-nil *[IPv4SocketAddress] if [IPSocketAddress] represents the variant case "ipv4".
func (self *IPSocketAddress) IPv4() *IPv4SocketAddress {
	return cm.Case[IPv4SocketAddress](self, 0)
}

// IPSocketAddressIPv6 returns a [IPSocketAddress] of case "ipv6".
func IPSocketAddressIPv6(data IPv6SocketAddress) IPSocketAddress {
	return cm.New[IPSocketAddress](1, data)
}

// IPv6 returns a non-nil *[IPv6SocketAddress] if [IPSocketAddress] represents the variant case "ipv6".
func (self *IPSocketAddress) IPv6() *IPv6SocketAddress {
	return cm.Case[IPv6SocketAddress](self, 1)
}
