package certificateparser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/patrickmsieber/certificate-exporter/pkg/certificateparser"
)

func TestParser(t *testing.T) {
	t.Parallel()

	res, err := certificateparser.Parse(nil)

	require.Error(t, err)
	assert.Empty(t, res)
}
