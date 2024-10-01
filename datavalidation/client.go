package datavalidation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	cl *http.Client

	pvt string
	evt string
}

func NewClient(phoneValidationToken, emailValidationToken string) *Client {
	return &Client{
		cl: &http.Client{
			Timeout: 5 * time.Second,
		},
		pvt: phoneValidationToken,
		evt: emailValidationToken,
	}
}

func (c *Client) PhoneValidation(phone string) (PhoneResponse, error) {
	url := fmt.Sprintf("https://phonevalidation.abstractapi.com/v1/?api_key=%s&phone=%s", c.pvt, phone)

	resp, err := c.cl.Get(url)
	if err != nil {
		return PhoneResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PhoneResponse{}, err
	}
	defer resp.Body.Close()

	var validPhone PhoneResponse
	if err = json.Unmarshal(body, &validPhone); err != nil {
		return PhoneResponse{}, err
	}

	return validPhone, nil
}

func (c *Client) EmailValidation(email string) (EmailResponse, error) {
	url := fmt.Sprintf("https://emailvalidation.abstractapi.com/v1/?api_key=%s&email=%s", c.evt, email)

	resp, err := c.cl.Get(url)
	if err != nil {
		return EmailResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return EmailResponse{}, err
	}
	defer resp.Body.Close()

	var emailValid EmailResponse
	if err = json.Unmarshal(body, &emailValid); err != nil {
		return EmailResponse{}, err
	}

	return emailValid, nil
}
