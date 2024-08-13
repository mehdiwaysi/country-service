package handlers

import (
 
)

func GetAllCountriesHandler(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	request := new(dto.CreateEmailRequest)
	err := lambda.UnmarshalRequest(e, true, request, true)
	if err != nil {
		return lambda.Error(err)
	}

	f := factory.NewFactory(request.TenantAlias)
	srv := service.NewEmailService(f)
	response, err := srv.CreateEmail(ctx, *request)
	if err != nil {
		return lambda.Error(err)
	}

	return lambda.MarshalResponse(http.StatusOK, nil, response)
}
