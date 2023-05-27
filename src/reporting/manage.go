package reporting

import (
	"src/reportstore"

	"src/uuid"
)

type ReportManager struct {
	uuidGenerator uuid.Generator
	store         reportstore.Store
}

func NewReportManager(gen uuid.Generator, store reportstore.Store) *ReportManager {
	return &ReportManager{
		uuidGenerator: gen,
		store:         store,
	}
}

func (m *ReportManager) CreateReport(request CreateReportRequest) (response CreateReportResponse, err error) {
	reportID := m.uuidGenerator.Generate()
	r := reportstore.CreateReportRequest{
		ReportID: reportID,
		UserID:   reportstore.ReportStatusPending.String(),
		Title:    request.Title,
	}
	err = m.store.CreateReport(r)

	response.ReportID = r.ReportID
	return

}
