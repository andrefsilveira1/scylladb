package reporting

type CreateReportRequest struct {
	UserID string
	Title  string
}

type CreateReportResponse struct {
	ReportID string
}
