package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"github.com/stripe/stripe-go/v78/refund"
)

func init() {
	godotenv.Load(".env")
	stripe.Key = os.Getenv("STRIPE_KEY")

	if stripe.Key == "" {
		log.Fatal("Stripe Key is not found in the environment")
	}
}

type CreateIntentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func handlerCreateIntent(w http.ResponseWriter, r *http.Request) {
	var req CreateIntentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request Payload")
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(req.Currency),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	result, err := paymentintent.New(params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, `Failed to create payment intent`)
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

func handlerCaptureIntent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	params := &stripe.PaymentIntentCaptureParams{}
	pi, err := paymentintent.Capture(id, params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to capture payment intent")
		return
	}

	respondWithJSON(w, http.StatusOK, pi)
}

func handlerCreateRefund(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	params := &stripe.RefundParams{
		PaymentIntent: stripe.String(id),
	}

	ref, err := refund.New(params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create refund")
		return
	}

	respondWithJSON(w, http.StatusOK, ref)
}

func handlerGetIntents(w http.ResponseWriter, r *http.Request) {
	params := &stripe.PaymentIntentListParams{}
	i := paymentintent.List(params)

	var intents []*stripe.PaymentIntent
	for i.Next() {
		intents = append(intents, i.PaymentIntent())
	}

	if err := i.Err(); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve payment intents")
		return
	}

	respondWithJSON(w, http.StatusOK, intents)
}
