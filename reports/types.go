package reports

import "time"

type ReportStatus int

const (
	ReportStatusUnknown ReportStatus = iota
	ReportStatusPending
	ReportStatusReviewing
	ReportStatusResolved
)

func (s ReportStatus) String() string {
	switch s {
	case ReportStatusPending:
		return "Pending"
	case ReportStatusReviewing:
		return "Reviewing"
	case ReportStatusResolved:
		return "Resolved"
	default:
		return ""
	}
}

type Report struct {
	ReportID  string
	UserID    string
	Resolver  string
	Reviewers []string
	CreatedAt time.Time
	UpdateAt  time.Time
	Status    ReportStatus
	Title     string
}

type Message struct {
	ReportID  string
	MessageID string
	Body      string
	CreatedAt time.Time
}
