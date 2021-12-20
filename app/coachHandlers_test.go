package app

import (
	"k/golang/gamematic/dto"
	"k/golang/gamematic/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var router *mux.Router
var ch CoachHandlers
var mocservice *service.MockCoachService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCoachService(ctrl)
	ch = CoachHandlers{service: mockService}
	router = mux.NewRouter()
	router.HandleFunc("/coaches", ch.GetAllCoaches)
	return func() {
		router = nil
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
	}

}

func Test_should_return_all_Coaches_200(t *testing.T) {

	//Arrange
	teardown := setup(t)
	defer teardown()

	dummyCoaches := []dto.CoachResponse{
		{Id: "1001", FirstName: "Ole", LastName: "Gunnar", Gender: "M", PhoneNumber: "781-112-7922", EmailAddress: "OleG@yahoo.com",
			AddressResponse: dto.AddressResponse{Address1: "44 Hickery Lyne", Address2: "", City: "Burlington", State: "MA", Zipcode: "01803"}, Role: "Head Coach", Team: "Man United"},
	}

	mocservice.EXPECT().GetAllCoaches().Return(dummyCoaches, nil)

	request, _ := http.NewRequest(http.MethodGet, "/coaches", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//Assert

	if recorder.Code != http.StatusOK {
		t.Error("Failed while Testing the status Code")
	}

}
