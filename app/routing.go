package app

import (
	rPayment "tubespbbo/domains/payment/router"
	rUser "tubespbbo/domains/user/router"
)

func mapURLs() {
	rUser.RouteUsers(router)
	rPayment.RoutePaymentMethods(router)
	rPayment.RoutePayments(router)
}
