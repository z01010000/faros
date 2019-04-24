/*
Copyright 2018 Pusher Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package log contains utilities for fetching a new logger
// when one is not already available.
//
// The Log Handle
//
// This package contains a root logr.Logger Log.  It may be used to
// get a handle to whatever the root logging implementation is.
package log

import (
	"github.com/pusher/faros/pkg/log/klogr"
	logr "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

func init() {
	logr.SetLogger(klogr.New(3))
}

// Log is the base logger used by Faros.
var Log = klogr.New(0)
