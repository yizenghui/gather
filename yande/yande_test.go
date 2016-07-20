// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package yande_test

import (
	"fmt"
	//	"log"
	"testing"

	"github.com/yizenghui/gather/yande"
)

func Test_NewQuerier(t *testing.T) {
	//	que := yande.NewQuerier("", "1", "10")
	que := yande.Querier{"", 1, 10}

	fmt.Println(que)
}

func Test_GetPost(t *testing.T) {
	que := yande.Querier{"", 1, 10}
	//	que := yande.NewQuerier("sex", "1", "10")
	fmt.Println(que)
	data := yande.GetPost(que)
	fmt.Println(data)

}
