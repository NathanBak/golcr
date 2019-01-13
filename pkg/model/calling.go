package golcr

// Calling describes a calling held by a member.
type Calling struct {
	// ActiveDate is the date the member started the calling.
	ActiveDate string `json:"activeDate"`

	// MemberID is the ID of the member holding the calling.
	MemberID int64 `json:"memberId"`

	// MemberName is the name of the member holding the calling.
	MemberName string `json:"memberName"`

	// Position is the description of the calling.
	Position string `json:"position"`
}
