package grpc

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// InstrumentIncomingRequestAndStartTxn instruments and starts a new transaction on each incoming Message.
func (s *Server) InstrumentIncomingRequestAndStartTxn(ctx context.Context, msg string) *newrelic.Transaction {
	txn := newrelic.FromContext(ctx)

	segment := &newrelic.Segment{
		Name:      msg,
		StartTime: txn.StartSegmentNow(),
	}

	defer segment.End()
	s.logger.Info(fmt.Sprintf("Message received: %s\n", msg))

	return txn
}
