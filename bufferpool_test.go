// Copyright 2013 The Bufferpool Authors. All rights reserved.
// Use of this source code is governed by the BSD 2-Clause license,
// which can be found in the LICENSE file.

package bufferpool_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/pushrax/bufferpool"
)

func TestTakeFromEmpty(t *testing.T) {
	bp := bufferpool.New(1, 1)
	poolBuf := bp.Take()
	if !bytes.Equal(poolBuf.Bytes(), []byte("")) {
		t.Fatalf("Buffer from empty bufferpool was allocated incorrectly.")
	}
}

func TestTakeFromFilled(t *testing.T) {
	bp := bufferpool.New(1, 1)
	bp.Give(bytes.NewBuffer([]byte("X")))
	reusedBuf := bp.Take()
	if !bytes.Equal(reusedBuf.Bytes(), []byte("")) {
		t.Fatalf("Buffer from filled bufferpool was recycled incorrectly.")
	}
}

func ExampleNew() {
	bp := bufferpool.New(10, 255)

	dogBuffer := bp.Take()
	dogBuffer.writeString("Dog!")
	bp.Give(dogBuffer)

	catBuffer := bp.Take() // dogBuffer is reused and reset.
	catBuffer.WriteString("Cat!")

	fmt.Println(catBuffer)
	// Output:
	// Cat!
}
