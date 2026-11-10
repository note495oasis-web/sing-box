package option

import (
	"github.com/sagernet/sing/common/byteformats"
	"github.com/sagernet/sing/common/json/badoption"
)

type HTTP2Options struct {
	IdleTimeout             badoption.Duration       `json:"idle_timeout,omitempty"`
	KeepAlivePeriod         badoption.Duration       `json:"keep_alive_period,omitempty"`
	StreamReceiveWindow     *byteformats.MemoryBytes `json:"stream_receive_window,omitempty"`
	ConnectionReceiveWindow *byteformats.MemoryBytes `json:"connection_receive_window,omitempty"`
	MaxConcurrentStreams    int                      `json:"max_concurrent_streams,omitempty"`
}

type QUICOptions struct {
	HTTP2Options
	InitialPacketSize       int  `json:"initial_packet_size,omitempty"`
	DisablePathMTUDiscovery bool `json:"disable_path_mtu_discovery,omitempty"`
}