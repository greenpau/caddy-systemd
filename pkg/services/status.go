// Copyright 2024 Paul Greenberg greenpau@outlook.com
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

package services

import (
	"encoding/json"
	"strings"
)

type StatusKind int

const (
	UnknownStatus StatusKind = iota
	PendingStatus
	FailureStatus
	SuccessStatus
)

// Status represent the last recorded status of a service.
type Status struct {
	Current     StatusKind `json:"current,omitempty"`
	ServiceName string     `json:"service_name,omitempty"`
	Error       error      `json:"error,omitempty"`
}

// NewStatus creates Status instance.
func NewStatus(name string, kind StatusKind) *Status {
	st := &Status{
		ServiceName: name,
		Current:     kind,
	}
	return st
}

func (k StatusKind) String() string {
	return [...]string{"Unknown", "Pending", "Failure", "Success"}[k]
}

func (k StatusKind) EnumIndex() int {
	return int(k)
}

func (k StatusKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.ToLower(k.String()))
}
