// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by "gen.bash" from internal/trace/v2; DO NOT EDIT.nn//go:build go1.21
package event

// SchedReqs is a set of constraints on what the scheduling
// context must look like.
type SchedReqs struct {
	Thread    Constraint
	Proc      Constraint
	Goroutine Constraint
}

// Constraint represents a various presence requirements.
type Constraint uint8

const (
	MustNotHave Constraint = iota
	MayHave
	MustHave
)

// UserGoReqs is a common requirement among events that are running
// or are close to running user code.
var UserGoReqs = SchedReqs{Thread: MustHave, Proc: MustHave, Goroutine: MustHave}
