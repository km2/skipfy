package skipper

import (
	"strings"

	"github.com/km2/skipfy/internal/model"
)

type ContainsSkipper struct {
	Substr string
}

func (s *ContainsSkipper) IsSkip(track *model.Track) bool {
	return strings.Contains(strings.ToLower(track.Title), s.Substr)
}
