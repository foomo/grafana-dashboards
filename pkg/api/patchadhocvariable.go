package api

import (
	"context"
	"fmt"

	"github.com/grafana/grafana-openapi-client-go/client"
)

func PatchAdHocVariable(ctx context.Context, client *client.GrafanaHTTPAPI, uid string) error {
	response, err := client.Dashboards.GetDashboardByUID(uid)
	if err != nil {
		return err
	}

	fmt.Println(response.Payload)
	return nil
}
