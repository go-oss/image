// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jpeg

const blockSize = 64 // A DCT block is 8x8.

type block [blockSize]int32
