package main

import "middleware-chain/middleware"

func main() {
	logging := &middleware.LoggingHandler{}
	auth := &middleware.AuthHandler{}
	authz := &middleware.AuthorizationHandler{}
	business := &middleware.BusinessHandler{}

	logging.SetNext(auth)
	auth.SetNext(authz)
	authz.SetNext(business)

	logging.Handle("valid-request")
	logging.Handle("invalid")
}
