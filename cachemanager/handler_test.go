// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cachemanager

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransformKey(t *testing.T) {
	// Initialize
	h := handler{
		prefix: "prefix_",
	}

	// Assert
	assert.Equal(t, "prefix_test", h.buildKey("test"))
}

func TestTransformTtl(t *testing.T) {
	// Initialize
	h := handler{
		ttl: time.Duration(5),
	}

	// Assert
	assert.Equal(t, time.Duration(5), h.buildTTL(time.Duration(-1)))
	assert.Equal(t, time.Duration(3), h.buildTTL(time.Duration(3)))
}

func TestSerialize(t *testing.T) {
	// Initialize
	d := []string{
		"test1",
		"test2",
	}
	h := handler{}

	// Encode
	c, e := h.serialize(d)

	// Assert
	assert.NoError(t, e)
	assert.Equal(t, "\f\xff\x81\x02\x01\x02\xff\x82\x00\x01\f\x00\x00\x10\xff\x82\x00\x02\x05test1\x05test2", string(c))

	// Decode
	de := []string{}
	e = h.unserialize(c, &de)

	// Assert
	assert.NoError(t, e)
	assert.Equal(t, d, de)
}
