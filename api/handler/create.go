package handler

import (
	"net/http"

	"github.com/rancher/norman/types"
)

func CreateHandler(apiContext *types.APIContext) error {
	var err error

	data, err := ParseAndValidateBody(apiContext)
	if err != nil {
		return err
	}

	store := apiContext.Schema.Store
	if store != nil {
		data, err = store.Create(apiContext, apiContext.Schema, data)
		if err != nil {
			return err
		}
	}

	apiContext.WriteResponse(http.StatusCreated, data)
	return nil
}
