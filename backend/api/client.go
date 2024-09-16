package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type client struct {
	logger  *slog.Logger
	address string
}

type Client interface {
	SubmitInput(input Input)
}

func NewClient(
	domain string,
	port int,
	logger *slog.Logger,
) (Client, error) {
	// Add http:// if not provided in the address
	address := fmt.Sprintf("%s:%d", domain, port)
	if !strings.HasPrefix(address, "http://") {
		address = fmt.Sprintf("http://%s", address)
	}
	newClient := &client{
		logger:  logger,
		address: address,
	}

	return newClient, nil
}

func (c *client) SubmitInput(input Input) {
	c.logger.Info("submitting input")

	inputJson, err := json.Marshal(input)
	if err != nil {
		c.logger.Error(
			"failed to convert input to json",
			slog.Int("AdCampaignId", input.AdCampaignId),
			slog.Int("CustomerId", input.CustomerId),
			slog.String("GameName", input.GameName),
			slog.String("ImageName", input.ImageName),
			slog.Bool("ValidAccount", input.ValidAccount),
			slog.String("error", err.Error()),
		)
		return
	}

	response, err := http.Post(
		fmt.Sprintf(`%s/submit/input`, c.address),
		"application/json",
		bytes.NewBuffer(inputJson),
	)
	if err != nil {
		c.logger.Error(
			"failed send http post request",
			slog.Int("AdCampaignId", input.AdCampaignId),
			slog.Int("CustomerId", input.CustomerId),
			slog.String("GameName", input.GameName),
			slog.String("ImageName", input.ImageName),
			slog.Bool("ValidAccount", input.ValidAccount),
			slog.String("error", err.Error()),
		)
		return
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		c.logger.Error(
			"failed convert response body to string",
			slog.Int("AdCampaignId", input.AdCampaignId),
			slog.Int("CustomerId", input.CustomerId),
			slog.String("GameName", input.GameName),
			slog.String("ImageName", input.ImageName),
			slog.Bool("ValidAccount", input.ValidAccount),
			slog.String("error", err.Error()),
		)
		return
	}

	c.logger.Info(
		"successfully sent post request",
		slog.String("response", string(responseBody)),
	)
}

// Takes a json file, validates the json in the file, and returns the json as input struct and error
// Note - This was not made apart of the client to make unit testing easier
func ValidateInputFile(file string) (Input, error) {
	var input Input

	jsonData, err := os.ReadFile(file)
	if err != nil {
		return input, err
	}

	err = json.Unmarshal(jsonData, &input)
	if err != nil {
		// Design Decision: I could have returned an empty input but wanted to use this for my unit testing.
		// It's also best practice to always check for error, so if someone want to ignore the error, they could still send the input we received
		return input, err
	}

	return input, nil
}

// Takes a raw json provided as a cli arg, validates the json, and returns the json as input struct and error
// Note - This was not made apart of the client to make unit testing easier
func ValidateInputRawJson(jsonString string) (Input, error) {
	var input Input

	err := json.Unmarshal([]byte(jsonString), &input)
	if err != nil {
		// Design Decision: I could have returned an empty input but wanted to use this for my unit testing.
		// It's also best practice to always check for error, so if someone want to ignore the error, they could still send the input we received
		return input, err
	}

	return input, nil
}
