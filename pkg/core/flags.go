package core

type Flags uint

const (
	FlagSeen Flags = 1 << iota
	FlagAnswered
	FlagFlagged
	FlagDeleted
	FlagDraft
	FlagRecent

	AllFlags = FlagSeen | FlagAnswered | FlagFlagged | FlagDeleted | FlagDraft | FlagRecent
)

var flagToStr map[Flags]string = map[Flags]string{
	FlagSeen:     `\Seen`,
	FlagAnswered: `\Answered`,
	FlagFlagged:  `\Flagged`,
	FlagDeleted:  `\Deleted`,
	FlagDraft:    `\Draft`,
	FlagRecent:   `\Recent`,
}

// Has checks if a set of flags contains a flag
func (fs Flags) Has(flag Flags) bool {
	return fs&flag != 0
}

// Stings converts a set of flags to a list of strings
func (fs Flags) Strings() []string {
	var r []string
	for flag, str := range flagToStr {
		if fs.Has(flag) {
			r = append(r, str)
		}
	}
	return r
}
