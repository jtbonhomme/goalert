package retry

import (
	"database/sql"
	"database/sql/driver"
	"net"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type clientErr interface {
	ClientError() bool
}

// IsTemporaryError will determine if an error is temporary, and thus
// the action can/should be retried.
func IsTemporaryError(err error) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(clientErr); ok && e.ClientError() {
		return false
	}
	cause := errors.Cause(err)
	if _, ok := cause.(net.Error); ok {
		return true
	}
	if cause == sql.ErrConnDone {
		return true
	}
	if cause == driver.ErrBadConn {
		return true
	}
	if pqe, ok := cause.(*pq.Error); ok {
		switch pqe.Code.Class() {
		// Allow retry for tx or connection errors:
		// - Class 40 — Transaction Rollback
		// - Class 08 — Connection Exception
		//
		// https://www.postgresql.org/docs/10/static/errcodes-appendix.html
		case "40", "08":
			return true
		}
	}
	return false
}

// DoTempFunc is a simplified version of DoFunc that just returns an error value.
type DoTempFunc func(int) error

// DoTemporaryError will retry as long as the error returned from fn is
// temporary as defined by IsTemporaryError.
func DoTemporaryError(fn DoTempFunc, opts ...Option) error {
	return Do(func(n int) (bool, error) {
		err := fn(n)
		return IsTemporaryError(err), err
	}, opts...)
}