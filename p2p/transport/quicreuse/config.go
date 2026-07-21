package quicreuse

import (
	"time"

	"github.com/quic-go/quic-go"
)

var quicConfig = &quic.Config{
	MaxIncomingStreams:         256,
	MaxIncomingUniStreams:      5,              // allow some unidirectional streams, in case we speak WebTransport
	MaxStreamReceiveWindow:     32 * (1 << 20), // 32 MB
	MaxConnectionReceiveWindow: 64 * (1 << 20), // 64 MB
	KeepAlivePeriod:            15 * time.Second,
	Versions:                   []quic.Version{quic.Version1},
	// We don't use datagrams (yet), but this is necessary for WebTransport
	EnableDatagrams: true,
	// Required for WebTransport
	EnableStreamResetPartialDelivery: true,
	InitialStreamReceiveWindow:       32 * (1 << 20),
	InitialConnectionReceiveWindow:   64 * (1 << 20),
}
