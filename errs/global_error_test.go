package errs

import (
	"net/http"
	"testing"
)

func Test_should_return_newNotFoundError(t *testing.T) {

	//Arrange/Act-> setting up test condition
	appError := NewNotFoundError("Not found Exception")

	//Assert-> Validate expected
	if appError.Message != "Not found Exception" {
		t.Error("Invalid Message Type")
	}

	if appError.Code != http.StatusNotFound {
		t.Error("Invalid Invalid Error Code")
	}

}

func Test_should_return_newUnexpectedError(t *testing.T) {

	//Arrange/Act-> setting up test condition
	appError := NewUnexpectedError("Un expected error")

	//Assert-> Validate expected
	if appError.Message != "Un expected error" {
		t.Error("Un expected error")
	}

	if appError.Code != http.StatusInternalServerError {
		t.Error("Invalid Invalid Error Code")
	}

}

func Test_should_return_NewValidationError(t *testing.T) {

	//Arrange/Act-> setting up test condition
	appError := NewValidationError("Unprocessable Entity")

	//Assert-> Validate expected
	if appError.Message != "Unprocessable Entity" {
		t.Error("Failed while testing NewValidationError message")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Failed while testing NewValidationError error code")
	}

}
