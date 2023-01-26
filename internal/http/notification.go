package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/palavrapasse/damn/pkg/entity"
	"github.com/palavrapasse/import/internal/logging"
)

const MaxAttemptsNotify = 5
const WaitingSecondsBetweenAttemptsNotify = 3

func NotifyNewLeak(leakId entity.AutoGenKey, subscribeServiceURL string) error {
	logging.Aspirador.Info(fmt.Sprintf("Starting notification of new leak %d", leakId))

	postBody, err := json.Marshal(map[string]int64{
		"leakId": int64(leakId),
	})

	if err != nil {
		return err
	}

	responseBody := bytes.NewBuffer(postBody)

	attempt := 1
	var resp *http.Response

	for attempt <= MaxAttemptsNotify {
		resp, err = http.Post(subscribeServiceURL, "application/json", responseBody)

		if err != nil {
			logging.Aspirador.Error(fmt.Sprintf("Error occured: '%s'. Trying again (done %d attempts)", err, attempt))
		} else if resp.StatusCode == http.StatusNoContent {
			break
		} else {
			logging.Aspirador.Warning(fmt.Sprintf("Expected %d status but received %d status. Trying again (done %d attempts)", http.StatusNoContent, resp.StatusCode, attempt))
		}

		logging.Aspirador.Trace(fmt.Sprintf("Waiting %d seconds...", WaitingSecondsBetweenAttemptsNotify))
		time.Sleep(WaitingSecondsBetweenAttemptsNotify * time.Second)

		attempt++
	}

	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	logging.Aspirador.Info("Successful notification of new leak")

	return nil
}
