package api

import (
	"net/http"

	"github.com/stripe/stripe-go/v74"
	"go.uber.org/zap"

	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/price"
)

func (s *Server) handleCheckoutSession(w http.ResponseWriter, r *http.Request) {
	log := s.logger
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	lookup_key := r.PostFormValue("lookup_key")

	params := &stripe.PriceListParams{
		LookupKeys: stripe.StringSlice([]string{
			lookup_key,
		}),
	}

	i := price.List(params)
	if i == nil {
		log.Error("Error updating subscription:")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var price *stripe.Price
	for i.Next() {
		p := i.Price()
		price = p
	}
	checkoutParams := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String(price.ID),
				Quantity: stripe.Int64(1),
			},
		},
	}

	sess, err := session.New(checkoutParams)
	if err != nil {
		log.Error("session", zap.Error(err))
	}

	http.Redirect(w, r, sess.URL, http.StatusSeeOther)
}
