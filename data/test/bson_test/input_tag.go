// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mytype

type MyType struct {
	local  int `bson:"Local"`
	Local2 int `bson:"Local1"`
}
