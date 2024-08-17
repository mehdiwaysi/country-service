package factory

import (
	"context"
	"encoding/json"
	"os"

	"github.com/revotech-group/go-aws/params"
	"github.com/revotech-group/go-lib/db"
)

type connectionProvider struct{}

func NewConnectionProvider() db.ConnectionProvider {
	return &connectionProvider{}
}

func getConnectionInfoByParam(_ context.Context, paramName string) (*db.ConnectionInfo, error) {

	paramValue, err := params.GetParameter(paramName)
	if err != nil {
		return nil, err
	}

	conInfo := &db.ConnectionInfo{}
	if err := json.Unmarshal([]byte(paramValue), conInfo); err != nil {
		return nil, err
	}

	return conInfo, nil
}

func (c *connectionProvider) GetConnectionInfo(ctx context.Context, tenantAlias string) (*db.ConnectionInfo, error) {
	paramName := "/" + os.Getenv("APP") + "/" + os.Getenv("STAGE") + "/platform/mongodb"
	return getConnectionInfoByParam(ctx, paramName)
}
