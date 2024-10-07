package player

import "errors"

type Player interface {
	Init() error

	Play(list *MediaList) error
	PlayNext() error
	AdvanceBySeconds(seconds int) int

	Next() string
	Current() string

	Shutdown() error
}

var ErrNoMoreMedia = errors.New("no more media in the list")
var ErrPlayerNotInitialized = errors.New("player wasn't initialized")
