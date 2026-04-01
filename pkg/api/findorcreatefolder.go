package api

import (
	"context"
	"strings"

	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-openapi-client-go/client"
	"github.com/grafana/grafana-openapi-client-go/client/folders"
	"github.com/grafana/grafana-openapi-client-go/models"
)

type Folder struct {
	UID     string
	Name    string
	Folders []Folder
}

func CreateFolders(ctx context.Context, client *client.GrafanaHTTPAPI, folders []Folder, parentUID string) error {
	for _, folder := range folders {
		if _, err := FindOrCreateFolder(ctx, client, folder.Name, folder.UID, parentUID); err != nil {
			return err
		}
		if err := CreateFolders(ctx, client, folder.Folders, folder.UID); err != nil {
			return err
		}
	}
	return nil
}

func FindOrCreateFolder(ctx context.Context, client *client.GrafanaHTTPAPI, name, uid, parentUID string) (string, error) {
	// FIXME: this doesn't handle pagination.
	// It will misbehave if the target Grafana instance has >1000 folders.
	getParams := folders.NewGetFoldersParams().WithContext(ctx)
	if parentUID != "" {
		getParams.SetParentUID(util.ToPtr(parentUID))
	}
	response, err := client.Folders.GetFolders(getParams)
	if err != nil {
		return "", err
	}

	for _, folder := range response.Payload {
		if strings.EqualFold(folder.UID, uid) {
			return folder.UID, nil
		}
	}

	// The folder doesn't exist: create it.
	createResponse, err := client.Folders.CreateFolder(&models.CreateFolderCommand{
		UID:       uid,
		Title:     name,
		ParentUID: parentUID,
	})
	if err != nil {
		return "", err
	}

	return createResponse.Payload.UID, nil
}
