package transport

import (
	"context"

	cm "BGP-Golang-TRPL3A/Framework/git/order/common"
	"BGP-Golang-TRPL3A/Framework/git/order/services"

	log "github.com/Sirupsen/logrus"

	"github.com/go-kit/kit/endpoint"
)

func invalidRequest() cm.Message {
	return cm.Message{
			Result: &cm.Result{
				Code:   99,
				Remark: "Invalid Request",
			},
		
	}
}

func OrderEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if req, ok := request.(cm.Message); ok {
			return svc.OrderHandler(ctx, req), nil
		}
		log.WithField("Error", request).Info("Request in in unkwon format")
		return invalidRequest(), nil
	}
}


