// Copyright (c) 2017, Roland Kammerer
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//Package errorlog implements a concurrency safe log of errors.
package errorlog

import "sync"

// ErrorLog is the struct that tracks a list of errors.
type ErrorLog struct {
	errs []error
	sync.Mutex
}

// NewErrorLog returns a new ErrorLog.
func NewErrorLog() *ErrorLog {
	return &ErrorLog{}
}

// Len returns the length (number of errors) in the ErrorLog.
func (e *ErrorLog) Len() int {
	e.Lock()
	defer e.Unlock()
	return len(e.errs)
}

// Errs returns a copy of the slice of errors written to the ErrorLog.
func (e *ErrorLog) Errs() []error {
	e.Lock()
	defer e.Unlock()

	// don't leak internal slice
	cpy := make([]error, len(e.errs))
	copy(cpy, e.errs)
	return cpy
}

// Append appends an error to the ErrorLog.
func (e *ErrorLog) Append(errs ...error) {
	e.Lock()
	e.errs = append(e.errs, errs...)
	e.Unlock()
}
