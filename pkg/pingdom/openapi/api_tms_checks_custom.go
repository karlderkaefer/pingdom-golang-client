package openapi

import (
	"context"
	"errors"
	"net/http"
)

type APIAction string

const (
	APIActionCreate APIAction = "created"
	APIActionUpdate APIAction = "updated"
	APIActionError  APIAction = "error"
)

func (a *TMSChecksAPIService) UpsertCheck(ctx context.Context, check CheckWithoutID) (APIAction, *CheckGetWithID, *http.Response, error) {
	if check.Name == "" {
		return APIActionError, nil, nil, errors.New("Check name is required")
	}
	id, resp, err := a.GetTransactionCheckIdkByName(ctx, check.Name)
	if err != nil {
		return APIActionError, nil, resp, err
	}
	if id != 0 {
		updated, resp, err := a.ModifyCheck(ctx, id).CheckWithoutIDPUT(*check.AsPut()).Execute()
		result := &CheckGetWithID{
			CheckWithoutIDGET: *updated,
			ID:                id,
		}
		return APIActionUpdate, result, resp, err
	}

	created, resp, err := a.AddCheck(ctx).CheckWithoutID(check).Execute()
	if err != nil {
		return APIActionError, nil, resp, err
	}
	result, resp, err := a.TestGetTransactionCheck(ctx, *created.Id)
	if err != nil {
		return APIActionError, nil, resp, err
	}
	return APIActionCreate, result, resp, nil
}

func (a *TMSChecksAPIService) GetTransactionCheckIdkByName(ctx context.Context, name string) (int64, *http.Response, error) {
	checks, resp, err := a.GetAllChecksExecute(ApiGetAllChecksRequest{
		ApiService: a,
		ctx:        ctx,
		type_:      PtrString("script"),
	})
	if err != nil {
		return 0, resp, err
	}
	for _, check := range checks.Checks {
		if check.Name != nil && *check.Name == name {
			return *check.Id, resp, nil
		}
	}
	return 0, resp, nil
}

func (a *TMSChecksAPIService) DeleteCheckByName(ctx context.Context, name string) (*DeleteCheck200Response, *http.Response, error) {
	id, resp, err := a.GetTransactionCheckIdkByName(ctx, name)
	if err != nil {
		return nil, resp, err
	}
	if id == 0 {
		return nil, resp, nil
	}
	return a.DeleteCheck(ctx, id).Execute()
}

func (a *TMSChecksAPIService) TestGetTransactionCheck(ctx context.Context, id int64) (*CheckGetWithID, *http.Response, error) {
	check, resp, err := a.GetCheck(ctx, id).Execute()
	if err != nil {
		return nil, resp, err
	}
	result := &CheckGetWithID{
		CheckWithoutIDGET: *check,
		ID:                id,
	}
	return result, resp, nil
}
