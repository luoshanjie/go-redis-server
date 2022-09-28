package configure

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parse(t *testing.T) {
	asserts := assert.New(t)
	src := "bind 0.0.0.0\n" +
		"port 6399\n"
	p := parse(strings.NewReader(src))
	asserts.NotNil(p)
	asserts.Equal("0.0.0.0", p.Bind)
	asserts.Equal(6399, p.Port)
}
