package middleware

import (
	"time"
	"context"
	cm "BGP-Golang-TRPL3A/Framework/git/order/common"
	"BGP-Golang-TRPL3A/Framework/git/order/services"
	log "github.com/Sirupsen/logrus"
)

func BasicMiddleware() services.ServiceMiddleware {
	return func(next services.PaymentServices) services.PaymentServices {
		return BasicMiddlewareStruct{next}
	}
}

type BasicMiddlewareStruct struct {
	services.PaymentServices
}

func (mw BasicMiddlewareStruct) OrderHandler(ctx context.Context, request cm.Message) cm.Message {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("OrderHandler ends")
	}(time.Now())
	log.WithField("request", request).Info("OrderHandler begins")
	return mw.PaymentServices.OrderHandler(ctx, request)
}

// add middleware CustomerHandler and ProductHandler
func (mw BasicMiddlewareStruct) CustomerHandler(ctx context.Context, request cm.Customer) cm.Customer {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("CustomerHandler ends")
	}(time.Now())
	log.WithField("request", request).Info("CustomerHandler begins")
	return mw.PaymentServices.CustomerHandler(ctx, request)
}

func (mw BasicMiddlewareStruct) ProductHandler(ctx context.Context, request cm.Product) cm.Product {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("ProductHandler ends")
	}(time.Now())
	log.WithField("request", request).Info("ProductHandler begins")
	return mw.PaymentServices.ProductHandler(ctx, request)	
}
