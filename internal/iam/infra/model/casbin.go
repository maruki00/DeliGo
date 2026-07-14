package model

// CREATE TABLE casbin_rule (
//     id    BIGSERIAL PRIMARY KEY, -- Auto-incrementing ID (internal use only)
//     ptype VARCHAR(100) NOT NULL, -- 'p' for standard policy, 'g' for group/role policy
//     v0    VARCHAR(100),          -- Matches field 1 of the rule (e.g., subject/user)
//     v1    VARCHAR(100),          -- Matches field 2 of the rule (e.g., object/role)
//     v2    VARCHAR(100),          -- Matches field 3 of the rule (e.g., action)
//     v3    VARCHAR(100),          -- Optional / Domain (tenant)
//     v4    VARCHAR(100),          -- Optional
//     v5    VARCHAR(100),          -- Optional
//     CONSTRAINT unique_casbin_key UNIQUE (ptype, v0, v1, v2, v3, v4, v5)
// );

const (
	POLICY_TYPE_STANDARD = "p"
	POLICY_TYPE_GROUP    = "g"
)

type CasbinRule struct {
	Id    int     `json:"id"`
	PType string  `json:"ptype"`
	V0    string  `json:"v0"`
	V1    string  `json:"v1"`
	V2    string  `json:"v2"`
	V3    *string `json:"v3"` // (optional)
	V4    *string `json:"v4"` // (optional)
	V5    *string `json:"v5"` // (optional)
}

type Policy struct {
	ID         int     `json:"id"`
	PType      string  `json:"ptype"`
	Role       string  `json:"role"`        // Maps to v0
	Resource   string  `json:"resource"`    // Maps to v1
	Action     string  `json:"action"`      // Maps to v2
	Effect     *string `json:"effect"`      // Maps to v3 (optional)
	IpSource   *string `json:"ip_source"`   // Maps to v4 (optional)
	TimeWindow *string `json:"time_window"` // Maps to v5 (optional)
}

type GroupPolicy struct {
	User   string `json:"user"`   // Maps to v0
	Group  string `json:"group"`  // Maps to v1
	Domain string `json:"domain"` // Maps to v2 (optional)
}

func (_this *GroupPolicy) Map2Casbin() CasbinRule {
	return CasbinRule{
		PType: POLICY_TYPE_GROUP,
		V0:    _this.User,
		V1:    _this.Group,
		V2:    _this.Domain,
	}
}

func (_this *Policy) Map2Casbin() CasbinRule {
	return CasbinRule{
		PType: POLICY_TYPE_STANDARD,
		V0:    _this.Role,
		V1:    _this.Resource,
		V2:    _this.Action,
		V3:    _this.Effect,
		V4:    _this.IpSource,
		V5:    _this.TimeWindow,
	}
}
