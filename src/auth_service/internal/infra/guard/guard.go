package guard

import (
	"context"
	"encoding/json"

	"github.com/1layar/merasa/backend/src/auth_service/internal/app/appconfig"
	openfga "github.com/openfga/go-sdk"
	"github.com/openfga/go-sdk/client"
)

func New(config *appconfig.Config) *client.OpenFgaClient {
	fgaClient, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl: config.FgaApiUrl,
	})

	if err != nil {
		panic(err)
	}

	respStores, err := fgaClient.ListStores(context.Background()).Execute()
	if err != nil {
		panic(err)
	}

	hasStore := false
	for _, store := range respStores.GetStores() {
		if store.GetName() == config.FgaStore {
			err = fgaClient.SetStoreId(store.GetId())
			if err != nil {
				panic(err)
			}

			hasStore = true
			break
		}
	}

	if !hasStore {
		resp, err := fgaClient.CreateStore(context.Background()).Body(client.ClientCreateStoreRequest{
			Name: config.FgaStore,
		}).Execute()

		if err != nil {
			panic(err)
		}
		err = fgaClient.SetStoreId(resp.GetId())
		if err != nil {
			panic(err)
		}
	}

	respModels, err := fgaClient.ReadLatestAuthorizationModel(context.Background()).Execute()

	if err != nil {
		panic(err)
	}

	var body openfga.WriteAuthorizationModelRequest

	schema, err := loadModel()

	if err != nil || respModels.AuthorizationModel == nil {
		if err := json.Unmarshal(schema, &body); err != nil {
			panic(err)
		}

		data, err := fgaClient.WriteAuthorizationModel(context.Background()).
			Body(body).
			Execute()

		if err != nil {
			panic(err)
		}

		err = fgaClient.SetAuthorizationModelId(data.AuthorizationModelId)

		if err != nil {
			panic(err)
		}
	} else {
		authModelId := respModels.GetAuthorizationModel().Id
		err = fgaClient.SetAuthorizationModelId(authModelId)

		if err != nil {
			panic(err)
		}
	}

	return fgaClient
}
