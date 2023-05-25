package reporting

import (
	"errors"
	"testing"

	reportstore "github.com/amin-mir/reporting/reportstore"
	mockreportstore "github.com/amin-mir/reporting/reportstore/mock"
	mockuuid "github.com/amin-mir/reporting/uuid/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ReportManagerSuite struct {
	suite.Suite
	*require.Assertions

	controller        *gomock.Controller
	mockReportStore   *mockreportstore.MockStore
	mockUUIDGenerator *mockuuid.MockGenerator

	manager *ReportManager
}

func TestReportManagerSuite(t *testing.T) {
	suite.Run(t, new(ReportManagerSuite))
}

func (s *ReportManagerSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.controller = gomock.NewController(s.T())
	s.mockReportStore = mockreportstore.NewMockStore(s.controller)
	s.mockUUIDGenerator = mockuuid.NewMockGenerator(s.controller)
	s.manager = NewReportManager(s.mockUUIDGenerator, s.mockReportStore)

}

func (s *ReportManagerSuite) TearDownTest() {
	s.controller.Finish()
}

func (s *ReportManagerSuite) TestCreateReport() {
	reportID := "reportid"
	UserID := "userid"
	title := "tittle"

	s.mockUUIDGenerator.EXPECT().Generate().Return(reportID).Times(1)
	s.mockReportStore.EXPECT().CreateReport(gomock.Eq(reportstore.CreateReportRequest{
		ReportID: reportID,
		UserID:   UserID,
		Status:   reportstore.ReportStatusPending.String(),
		Title:    title,
	})).Return(nil).Times(1)

	res, err := s.manager.CreateReport((CreateReportRequest{
		UserID: UserID,
		Title:  title,
	}))
	s.NoError(err)
	expectedResponse := CreateReportResponse{
		ReportID: reportID,
	}
	s.Equal(expectedResponse, res)
}

func (s *ReportManagerSuite) TestCreateReportError() {
	s.mockUUIDGenerator.EXPECT().Generate().Return("reportid").Times(1)
	err := errors.New("error")
	s.mockReportStore.EXPECT().CreateReport((gomock.Any())).Return(err).Times(1)
	_, e := s.manager.CreateReport(CreateReportRequest{})
	s.Equal(err, e)
}
