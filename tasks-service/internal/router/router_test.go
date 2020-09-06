package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/service"
	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/task"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"

	"github.com/golang/mock/gomock"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router"
	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request"
	requestMock "github.com/pedroquerido/sword-challenge/tasks-service/internal/router/request/mock"
	serviceMock "github.com/pedroquerido/sword-challenge/tasks-service/internal/service/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter_CreateTask(t *testing.T) {

	ctlr := gomock.NewController(t)
	svc := serviceMock.NewMockTaskService(ctlr)
	validator := requestMock.NewMockValidator(ctlr)
	testRouter := router.New(svc, validator)

	headerUserID := "x-user-id"
	headerUserRole := "x-user-role"

	t.Run("should get response code 400 - bad request format", func(t *testing.T) {

		reqBody := "bad request format"

		expectResponse := response.NewErrorResponse(http.StatusBadRequest, "bad request")

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any_role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 400 - missing required fields", func(t *testing.T) {

		reqBody := request.CreateTaskRequestBody{
			Summary: "summary",
		}

		expectResponse := response.NewErrorResponse(http.StatusBadRequest, "bad request")

		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(request.ErrorBadRequest)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any_role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 422 - invalid request fields", func(t *testing.T) {

		dateString := "2021-01-01T00:01:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		reqBody := request.CreateTaskRequestBody{
			Summary: "summary",
			Date:    date,
		}

		expectResponse := response.NewErrorResponse(http.StatusUnprocessableEntity, "unprocessable entity")

		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(nil)
		svc.EXPECT().
			Create(gomock.Any(), gomock.Eq(reqBody.Summary), gomock.Eq(reqBody.Date)).
			Times(1).
			Return("", service.ErrorInvalidTask)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any_role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - missing context", func(t *testing.T) {

		dateString := "2020-01-01T00:01:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		reqBody := request.CreateTaskRequestBody{
			Summary: "summary",
			Date:    date,
		}

		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(nil)
		svc.EXPECT().
			Create(gomock.Any(), gomock.Eq(reqBody.Summary), gomock.Eq(reqBody.Date)).
			Times(1).
			Return("", service.ErrorMissingContext)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any_role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - unexpected error", func(t *testing.T) {

		dateString := "2020-01-01T00:01:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		reqBody := request.CreateTaskRequestBody{
			Summary: "summary",
			Date:    date,
		}

		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(nil)
		svc.EXPECT().
			Create(gomock.Any(), gomock.Eq(reqBody.Summary), gomock.Eq(reqBody.Date)).
			Times(1).
			Return("", service.ErrorUnexpectedError)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any_role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 201 with task_id", func(t *testing.T) {

		dateString := "2020-09-06T11:45:00Z"
		date, err := time.Parse(time.RFC3339, dateString)

		reqBody := request.CreateTaskRequestBody{
			Summary: "summary",
			Date:    date,
		}

		svcResponse := "task_id"
		expectResponse := response.NewCreateTaskResponse(http.StatusCreated, "task created", svcResponse)

		validator.EXPECT().
			Validate(gomock.Any()).
			Times(1).
			Return(nil)
		svc.EXPECT().
			Create(gomock.Any(), gomock.Eq(reqBody.Summary), gomock.Eq(reqBody.Date)).
			Times(1).
			Return(svcResponse, nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any_role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		outResponse := response.CreateTaskResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
		assert.Equal(t, expectResponse.Data, outResponse.Data)
	})
}

func TestRouter_ListTasks(t *testing.T) {

	ctlr := gomock.NewController(t)
	svc := serviceMock.NewMockTaskService(ctlr)
	validator := requestMock.NewMockValidator(ctlr)
	testRouter := router.New(svc, validator)

	headerUserID := "x-user-id"
	headerUserRole := "x-user-role"
	headerUserRoleValueManager := "manager"

	t.Run("should get response code 403", func(t *testing.T) {

		svcResponse := service.ErrorUserNotAllowed
		expectResponse := response.NewErrorResponse(http.StatusForbidden, "forbidden")

		svc.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, svcResponse)

		req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - missing context", func(t *testing.T) {

		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, service.ErrorMissingContext)

		req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - unexpected error", func(t *testing.T) {

		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, service.ErrorUnexpectedError)

		req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 200 with tasks", func(t *testing.T) {

		userID := "user_id"
		dateString := "2020-09-06T11:45:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		svcResponse := []*task.Task{
			&task.Task{
				ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
				UserID:  userID,
				Summary: "summary",
				Date:    date,
			},
		}
		expectResponse := response.NewListTasksResponse(http.StatusOK, "tasks retrieved", svcResponse)

		svc.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Times(1).
			Return(svcResponse, nil)

		req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
		req.Header.Set(headerUserID, userID)
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		outResponse := response.ListTasksResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
		assert.Equal(t, expectResponse.Data, outResponse.Data)
	})
	t.Run("should get response code 200 with tasks - with user_id", func(t *testing.T) {

		userID := "user_id"
		dateString := "2020-09-06T11:45:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		svcResponse := []*task.Task{
			&task.Task{
				ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
				UserID:  userID,
				Summary: "summary",
				Date:    date,
			},
		}
		expectResponse := response.NewListTasksResponse(http.StatusOK, "tasks retrieved", svcResponse)

		svc.EXPECT().
			List(gomock.Any(), gomock.Eq(&userID)).
			Times(1).
			Return(svcResponse, nil)

		req, err := http.NewRequest(http.MethodGet, "/tasks?user_id="+userID, nil)
		req.Header.Set(headerUserID, userID)
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		outResponse := response.ListTasksResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
		assert.Equal(t, expectResponse.Data, outResponse.Data)
	})
}

func TestRouter_RetrieveTask(t *testing.T) {

	ctlr := gomock.NewController(t)
	svc := serviceMock.NewMockTaskService(ctlr)
	validator := requestMock.NewMockValidator(ctlr)
	testRouter := router.New(svc, validator)

	headerUserID := "x-user-id"
	headerUserRole := "x-user-role"

	t.Run("should get response code 403", func(t *testing.T) {

		taskID := "task_id"
		expectResponse := response.NewErrorResponse(http.StatusForbidden, "forbidden")

		svc.EXPECT().
			Retrieve(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(nil, service.ErrorUserNotAllowed)

		req, err := http.NewRequest(http.MethodGet, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 404", func(t *testing.T) {

		taskID := "task_id"
		expectResponse := response.NewErrorResponse(http.StatusNotFound, "not found")

		svc.EXPECT().
			Retrieve(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(nil, service.ErrorTaskNotFound)

		req, err := http.NewRequest(http.MethodGet, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - missing context", func(t *testing.T) {

		taskID := "task_id"
		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			Retrieve(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(nil, service.ErrorMissingContext)

		req, err := http.NewRequest(http.MethodGet, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - unexpected error", func(t *testing.T) {

		taskID := "task_id"
		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			Retrieve(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(nil, service.ErrorUnexpectedError)

		req, err := http.NewRequest(http.MethodGet, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 200 with task", func(t *testing.T) {

		userID := "user_id"
		taskID := "task_id"
		dateString := "2020-09-06T11:45:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		svcResponse := &task.Task{
			ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
			UserID:  userID,
			Summary: "summary",
			Date:    date,
		}
		expectResponse := response.NewRetrieveTaskResponse(http.StatusOK, "task retrieved", svcResponse)

		svc.EXPECT().
			Retrieve(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(svcResponse, nil)

		req, err := http.NewRequest(http.MethodGet, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		outResponse := response.RetrieveTaskResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
		assert.Equal(t, expectResponse.Data, outResponse.Data)
	})
}

func TestRouter_UpdateTask(t *testing.T) {

	ctlr := gomock.NewController(t)
	svc := serviceMock.NewMockTaskService(ctlr)
	validator := requestMock.NewMockValidator(ctlr)
	testRouter := router.New(svc, validator)

	headerUserID := "x-user-id"
	headerUserRole := "x-user-role"

	t.Run("should get response code 400 - bad request format", func(t *testing.T) {

		taskID := "task_id"
		reqBody := "bad request format"

		expectResponse := response.NewErrorResponse(http.StatusBadRequest, "bad request")

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 403", func(t *testing.T) {

		taskID := "task_id"
		reqBody := request.UpdateTaskRequestBody{
			Summary: "summary",
		}

		expectResponse := response.NewErrorResponse(http.StatusForbidden, "forbidden")

		svc.EXPECT().
			Update(gomock.Any(), gomock.Eq(taskID), gomock.Eq(&reqBody.Summary), gomock.Any()).
			Times(1).
			Return(service.ErrorUserNotAllowed)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 404", func(t *testing.T) {

		taskID := "task_id"
		reqBody := request.UpdateTaskRequestBody{
			Summary: "summary",
		}

		expectResponse := response.NewErrorResponse(http.StatusNotFound, "not found")

		svc.EXPECT().
			Update(gomock.Any(), gomock.Eq(taskID), gomock.Eq(&reqBody.Summary), gomock.Any()).
			Times(1).
			Return(service.ErrorTaskNotFound)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 422 - invalid request fields", func(t *testing.T) {

		taskID := "task_id"
		dateString := "2021-01-01T00:01:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		reqBody := request.UpdateTaskRequestBody{
			Summary: "summary",
			Date:    date,
		}

		expectResponse := response.NewErrorResponse(http.StatusUnprocessableEntity, "unprocessable entity")

		svc.EXPECT().
			Update(gomock.Any(), gomock.Eq(taskID), gomock.Eq(&reqBody.Summary), gomock.Eq(&reqBody.Date)).
			Times(1).
			Return(service.ErrorInvalidTask)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - missing context", func(t *testing.T) {

		taskID := "task_id"
		dateString := "2021-01-01T00:01:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		reqBody := request.UpdateTaskRequestBody{
			Summary: "summary",
			Date:    date,
		}

		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			Update(gomock.Any(), gomock.Eq(taskID), gomock.Eq(&reqBody.Summary), gomock.Eq(&reqBody.Date)).
			Times(1).
			Return(service.ErrorMissingContext)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - unexpected error", func(t *testing.T) {

		taskID := "task_id"
		dateString := "2021-01-01T00:01:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		reqBody := request.UpdateTaskRequestBody{
			Summary: "summary",
			Date:    date,
		}

		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			Update(gomock.Any(), gomock.Eq(taskID), gomock.Eq(&reqBody.Summary), gomock.Eq(&reqBody.Date)).
			Times(1).
			Return(service.ErrorUnexpectedError)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 200", func(t *testing.T) {

		taskID := "task_id"
		dateString := "2021-01-01T00:01:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		reqBody := request.UpdateTaskRequestBody{
			Summary: "summary",
			Date:    date,
		}

		expectResponse := response.NewUpdateTaskResponse(http.StatusOK, "task updated")

		svc.EXPECT().
			Update(gomock.Any(), gomock.Eq(taskID), gomock.Eq(&reqBody.Summary), gomock.Eq(&reqBody.Date)).
			Times(1).
			Return(nil)

		body, err := json.Marshal(reqBody)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPatch, "/tasks/"+taskID, bytes.NewBuffer(body))
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		outResponse := response.UpdateTaskResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
}

func TestRouter_DeleteTask(t *testing.T) {

	ctlr := gomock.NewController(t)
	svc := serviceMock.NewMockTaskService(ctlr)
	validator := requestMock.NewMockValidator(ctlr)
	testRouter := router.New(svc, validator)

	headerUserID := "x-user-id"
	headerUserRole := "x-user-role"
	headerUserRoleValueManager := "manager"

	t.Run("should get response code 403", func(t *testing.T) {

		taskID := "task_id"
		expectResponse := response.NewErrorResponse(http.StatusForbidden, "forbidden")

		svc.EXPECT().
			Delete(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(service.ErrorUserNotAllowed)

		req, err := http.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 404", func(t *testing.T) {

		taskID := "task_id"
		expectResponse := response.NewErrorResponse(http.StatusNotFound, "not found")

		svc.EXPECT().
			Delete(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(service.ErrorTaskNotFound)

		req, err := http.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - missing context", func(t *testing.T) {

		taskID := "task_id"
		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			Delete(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(service.ErrorMissingContext)

		req, err := http.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - unexpected error", func(t *testing.T) {

		taskID := "task_id"
		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			Delete(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(service.ErrorUnexpectedError)

		req, err := http.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 200", func(t *testing.T) {

		taskID := "task_id"
		expectResponse := response.NewDeleteTaskResponse(http.StatusOK, "task deleted")

		svc.EXPECT().
			Delete(gomock.Any(), gomock.Eq(taskID)).
			Times(1).
			Return(nil)

		req, err := http.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, headerUserRoleValueManager)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		outResponse := response.DeleteTaskResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
}

func TestRouter_ListUserTasks(t *testing.T) {

	ctlr := gomock.NewController(t)
	svc := serviceMock.NewMockTaskService(ctlr)
	validator := requestMock.NewMockValidator(ctlr)
	testRouter := router.New(svc, validator)

	headerUserID := "x-user-id"
	headerUserRole := "x-user-role"

	t.Run("should get response code 403", func(t *testing.T) {

		expectResponse := response.NewErrorResponse(http.StatusForbidden, "forbidden")

		svc.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, service.ErrorUserNotAllowed)

		req, err := http.NewRequest(http.MethodGet, "/users/1/tasks", nil)
		req.Header.Set(headerUserID, "2")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - missing context", func(t *testing.T) {

		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, service.ErrorMissingContext)

		req, err := http.NewRequest(http.MethodGet, "/users/1/tasks", nil)
		req.Header.Set(headerUserID, "1")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 500 - unexpected error", func(t *testing.T) {

		expectResponse := response.NewErrorResponse(http.StatusInternalServerError, "unknown error")

		svc.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil, service.ErrorUnexpectedError)

		req, err := http.NewRequest(http.MethodGet, "/users/1/tasks", nil)
		req.Header.Set(headerUserID, "1")
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		outResponse := response.ErrorResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
	})
	t.Run("should get response code 200 with tasks", func(t *testing.T) {

		userID := "1"
		dateString := "2020-09-06T11:45:00Z"
		date, err := time.Parse(time.RFC3339, dateString)
		require.Nil(t, err)

		svcResponse := []*task.Task{
			&task.Task{
				ID:      "419d535a-ced5-4d5e-8885-49df44d3f5ff",
				UserID:  userID,
				Summary: "summary",
				Date:    date,
			},
		}
		expectResponse := response.NewListUserTasksResponse(http.StatusOK, "user tasks retrieved", svcResponse)

		svc.EXPECT().
			List(gomock.Any(), gomock.Any()).
			Times(1).
			Return(svcResponse, nil)

		req, err := http.NewRequest(http.MethodGet, "/users/"+userID+"/tasks", nil)
		req.Header.Set(headerUserID, userID)
		req.Header.Set(headerUserRole, "any role")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		outResponse := response.ListUserTasksResponse{}
		err = json.NewDecoder(recorder.Body).Decode(&outResponse)
		require.Nil(t, err)
		assert.Equal(t, expectResponse.Code, outResponse.Code)
		assert.Equal(t, expectResponse.Message, outResponse.Message)
		assert.Equal(t, expectResponse.Data, outResponse.Data)
	})
}
