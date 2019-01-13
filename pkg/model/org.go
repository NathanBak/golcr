package golcr

import (
	"encoding/json"
	"log"
)

// Org describes a organization in a unit.
type Org struct {
	Name     string        `json:"name"`
	SubOrgID int64         `json:"subOrgId"`
	Children []Org         `json:"children,omitempty"`
	Callings []Calling     `json:"callings,omitempty"`
	Members  []ClassMember `json:"members,omitempty"`

	UnitNumber                   int64   `json:"unitNumber"`
	OrgTypeIds                   []int64 `json:"orgTypeIds"`
	DefaultOrgTypeIds            []int64 `json:"defaultOrgTypeIds,omitempty"`
	UserCanEditCallings          bool    `json:"userCanEditCallings"`
	IsClass                      bool    `json:"isClass,omitempty"`
	IsRealClass                  bool    `json:"isRealClass,omitempty"`
	IsSplit                      bool    `json:"isSplit,omitempty"`
	ClassGroup                   string  `json:"classGroup,omitempty"`
	FirstTypeID                  int64   `json:"firstTypeId,omitempty"`
	IsCombined                   bool    `json:"isCombined,omitempty"`
	Gender                       string  `json:"gender,omitempty"`
	ParentName                   string  `json:"parentName,omitempty"`
	UnitType                     int64   `json:"unitType,omitempty"`
	ParentSubOrgID               int64   `json:"parentSubOrgId,omitempty"`
	CustomOrgName                string  `json:"customOrgName,omitempty"`
	DefaultOrgName               string  `json:"defaultOrgName,omitempty"`
	RelatedSubOrgName            string  `json:"relatedSubOrgName,omitempty"`
	Priesthood                   bool    `json:"priesthood,omitempty"`
	AdultClass                   bool    `json:"adultClass,omitempty"`
	CanHaveMembers               bool    `json:"canHaveMembers,omitempty"`
	UserCanEditClassAssignments  bool    `json:"userCanEditClassAssignments,omitempty"`
	UserCanSeeBasicInfo          bool    `json:"userCanSeeBasicInfo,omitempty"`
	MultipleAllowed              bool    `json:"multipleAllowed,omitempty"`
	NameEditable                 bool    `json:"nameEditable,omitempty"`
	Resetable                    bool    `json:"resetable,omitempty"`
	Required                     bool    `json:"required,omitempty"`
	CustomPositionsAllowed       bool    `json:"customPositionsAllowed,omitempty"`
	Missionaries                 bool    `json:"missionaries,omitempty"`
	HighCouncilOrDistrictCouncil bool    `json:"highCouncilOrDistrictCouncil,omitempty"`
	Room                         string  `json:"room,omitempty"`
	SplitOrCombinedClassNotice   string  `json:"splitOrCombinedClassNotice,omitempty"`
	DisplayOrder                 int64   `json:"displayOrder,omitempty"`
	Instance                     int64   `json:"instance,omitempty"`
	TotalInstances               int64   `json:"totalInstances,omitempty"`
	ClassesCollapsed             bool    `json:"classesCollapsed,omitempty"`
	EnableVT                     bool    `json:"enableVT,omitempty"`
	SupportsAttendanceRoll       bool    `json:"supportsAttendanceRoll,omitempty"`
	EnableHT                     bool    `json:"enableHT,omitempty"`
	Other                        bool    `json:"other,omitempty"`
	HasAnyAllowMultiple          bool    `json:"hasAnyAllowMultiple,omitempty"`
	FirstOrgType                 string  `json:"firstOrgType,omitempty"`
	HasCustomCallings            bool    `json:"hasCustomCallings,omitempty"`
}

func (o *Org) String() string {
	bytes, err := json.Marshal(o)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
