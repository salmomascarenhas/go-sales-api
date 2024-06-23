package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner1 struct {
	BaseUrl string
}

// Partner1ReservationRequest represents the request to the partner1 reservation endpoint
type Partner1ReservationRequest struct {
	Spots []string `json:"spots"`
	TicketKind string `json:"ticket_kind"`
	Email string `json:"email"`
}

// Partner1ReservationResponse represents the response from the partner1 reservation endpoint
type Partner1ReservationResponse struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Spot string `json:"spot"`
	TicketKind string `json:"ticket_kind"`
	Status string `json:"status"`
	EventID string `json:"event_id"`
}

func (p *Partner1) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	partnerReq := Partner1ReservationRequest{
		Spots: req.Spots,
		TicketKind: req.TicketKind,
		Email: req.Email,
	}

	body, err := json.Marshal(partnerReq)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/events/%s/reserve", p.BaseUrl, req.EventID)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	client:= &http.Client{}
	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()
	if httpReq.Response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpReq.Response.StatusCode)
	}

	var partnerRes []Partner1ReservationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&partnerRes); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerRes))
	for i, pR := range partnerRes {
		responses[i] = ReservationResponse{
			ID: pR.ID,
			Spot: pR.Spot,	
			Status: pR.Status,
		}
	}
	return responses, nil
}	