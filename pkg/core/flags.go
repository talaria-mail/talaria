package core

type Flags uint

const (
	SeenFlag Flags = 1 << iota
	AnsweredFlag
	FlaggedFlag
	DeletedFlag
	DraftFlag
	RecentFlag
)
