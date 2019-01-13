package golcr

import (
	"encoding/json"
	"log"
)

// Member contains information about an individual member.
type Member struct {
	NameListPreferredLocal   string `json:"nameListPreferredLocal"`
	NameFamilyPreferredLocal string `json:"nameFamilyPreferredLocal"`
	NameGivenPreferredLocal  string `json:"nameGivenPreferredLocal"`

	Age                   int64  `json:"age"`
	BirthDateDisplayLocal string `json:"birthDateDisplayLocal"`
	BirthDateSort         string `json:"birthDateSort"`
	Sex                   string `json:"sex"`
	LegacyCmisID          int64  `json:"legacyCmisId"`
	PersonUUID            string `json:"personUuid"`

	Convert                  bool   `json:"convert,omitempty"`
	Email                    string `json:"email,omitempty"`
	IsAdult                  bool   `json:"isAdult,omitempty"`
	IsHead                   bool   `json:"isHead,omitempty"`
	IsOutOfUnitMember        bool   `json:"isOutOfUnitMember,omitempty"`
	IsProspectiveElder       bool   `json:"isProspectiveElder,omitempty"`
	IsSingleAdult            bool   `json:"isSingleAdult,omitempty"`
	IsSpouse                 bool   `json:"isSpouse,omitempty"`
	IsYoungSingleAdult       bool   `json:"isYoungSingleAdult,omitempty"`
	PriesthoodTeacherOrAbove bool   `json:"priesthoodTeacherOrAbove,omitempty"`
}

func (m *Member) String() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
