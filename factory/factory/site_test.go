package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFacebook(t *testing.T) {
	fb := facebook{baseSocialNetworkSite{}}
	fb.createProfile()
	s := fb.getSections()
	assert.Equal(t, 2, len(s))
	assert.Contains(t, s, albumSection{})
	assert.Contains(t, s, personalInfoSection{})
	assert.Contains(t, []string{"this is album section", "this is personal info section"}, s[0].description())
	assert.Contains(t, []string{"this is album section", "this is personal info section"}, s[1].description())
}

func TestLinkedin(t *testing.T) {
	l := linkedin{baseSocialNetworkSite{}}
	l.createProfile()
	s := l.getSections()
	assert.Equal(t, 2, len(s))
	assert.Contains(t, s, achievementSection{})
	assert.Contains(t, s, personalInfoSection{})
	assert.Contains(t, []string{"this is achievement section",
		"this is personal info section"}, s[0].description())
	assert.Contains(t, []string{"this is achievement section",
		"this is personal info section"}, s[1].description())
}
