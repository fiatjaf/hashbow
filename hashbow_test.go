package hashbow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashbow(t *testing.T) {
	assert.Equal(t,
		"#8b3c61",
		Hashbow("Olivia Wilde"),
	)
	assert.Equal(t,
		"#613c3c",
		Hashbow("Olivias Wilde"),
	)
	assert.Equal(t,
		"#ae613c",
		Hashbow(21),
	)
	assert.Equal(t,
		"#3c61d8",
		Hashbow(22),
	)
	assert.Equal(t,
		"#3c6134",
		Hashbow(23),
	)
	assert.Equal(t,
		"#61343c",
		Hashbow(200023),
	)
}
