package models

type CV struct {
	Name         string       `json:"name"`
	Surname      string       `json:"surname"`
	Contacts     Contacts     `json:"contacts"`
	Summary      []string     `json:"summary"`
	Experience   []Experience `json:"experience"`
	Faq          []Faq        `json:"faq"`
	Compensation Compensation `json:"compensation"`
}
type Contacts struct {
	Skype    string `json:"skype"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Linkedin string `json:"linkedin"`
	Github   string `json:"github"`
}
type Experience struct {
	StartDate        string        `json:"startDate"`
	EndDate          string        `json:"endDate"`
	Role             string        `json:"role"`
	Project          string        `json:"project"`
	Company          string        `json:"company"`
	ProjectSite      string        `json:"projectSite"`
	FrontendSize     int           `json:"frontendSize"`
	BackendSize      int           `json:"backendSize"`
	Stack            []string      `json:"stack"`
	Responsibilities []interface{} `json:"responsibilities"`
	QaSize           int           `json:"qaSize,omitempty"`
}
type Question struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
}
type Answer struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
}
type Faq struct {
	Question []Question `json:"question"`
	Answer   []Answer   `json:"answer"`
	Company  string     `json:"company"`
	Project  string     `json:"project"`
}
type Hourly struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
type Monthly struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
type Compensation struct {
	Hourly  Hourly  `json:"hourly"`
	Monthly Monthly `json:"monthly"`
}
