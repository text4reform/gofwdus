package fwdus

type CallLegislatorArgs struct {
	BioguideID string `json:"bioguide_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Zip        string `json:"zip"`
	Phone      string `json:"phone"`
}

type CallRequest struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	LegislatorID int    `json:"legislator_id"`
	Phone        string `json:"json"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type SearchLegislatorArgs struct {
	Zip      string `json:"zip"`
	District int    `json:"district"`
	State    string `json:"state"`
	Party    string `json:"party"`
}

type Legislator struct {
	ID                 int    `json:"id"`
	FirstName          string `json:"firstname"`
	MiddleName         string `json:"middlename"`
	LastName           string `json:"lastname"`
	Party              string `json:"party"`
	State              string `json:"state"`
	District           int    `json:"district,string"`
	InOffice           int    `json:"in_office,string"`
	Gender             string `json:"gender"`
	Phone              string `json:"phone"`
	Website            string `json:"website"`
	BioguideID         string `json:"bioguide_id"`
	TwitterID          string `json:"twitter_id"`
	ActionPoints       int    `json:"action_points"`
	TargetLevel        string `json:"target_level"`
	Photo              string `json:"photo"`
	Chamber            string `json:"chamber"`
	OverallStance      string `json:"overall_stance"`
	ImmigrationStances []struct {
		Issue    string `json:"issue"`
		Position string `json:"position"`
	} `json:"immigration_stances"`
	Rating int `json:"rating"`
}

func (l *Legislator) FullName() string {
	return l.FirstName + " " + l.LastName
}

type CreateLetterArgs struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	Writing       string `json:"writing"`
	LegislatorID  int    `json:"legislator_id"`
	Shareable     bool   `json:"shareable"`
}

type Letter struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	Writing       string `json:"writing"`
	LegislatorID  int    `json:"legislator_id"`
	Shareable     bool   `json:"shareable"`
	Sent          int    `json:"sent"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type Response struct {
	CallRequest CallRequest         `json:"call_request"`
	Legislators []Legislator        `json:"legislators"`
	Letter      Letter              `json:"letter"`
	Error       string              `json:"error"`
	Errors      map[string][]string `json:"errors"`
}
