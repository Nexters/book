package auth_test

import (
	"testing"

	"github.com/nexters/book/app/auth"
	"github.com/stretchr/testify/assert"
)

func TestParseTokenSuccessCase(t *testing.T) {
	b := auth.NewBearerAuth()
	token, err := b.ParseToken("Bearer 61016a5d-30d0-46d1-ad92-f0ea6e1adc71")
	assert.NoError(t, err)
	assert.Equal(t, "61016a5d-30d0-46d1-ad92-f0ea6e1adc71", token)
}

func TestParseTokenFailCase(t *testing.T) {
	b := auth.NewBearerAuth()
	_, err := b.ParseToken("Bearerabcd123-124")
	assert.Error(t, err)
}
