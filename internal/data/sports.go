package data

type Sport struct {
	Key          string `json:"amiericanfootball_ncaaf"`
	Group        string `json:"group"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Active       bool   `json:"active"`
	HasOutrights bool   `json:"has_outrights"`
}
