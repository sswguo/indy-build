package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDir(t *testing.T) {
	assert.Equal(t, "partyline", getGitDir("https://github.com/abc/partyline"))
	assert.Equal(t, "partyline", getGitDir("https://github.com/abc/partyline.git"))
	assert.Equal(t, "partyline", getGitDir("git@github.com:abc/partyline"))
	assert.Equal(t, "partyline", getGitDir("git@github.com:abc/partyline.git"))
	assert.Equal(t, "partyline", getGitDir("git+ssh://abc.com/abc/partyline"))
	assert.Equal(t, "partyline", getGitDir("git+ssh://abc.com/abc/partyline.git"))
}
