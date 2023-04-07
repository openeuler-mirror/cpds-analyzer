package detector

import (
	"errors"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func SendRuleUpdatedRequset(detecotrHost string, detectorPort int) error {
	urlStr := fmt.Sprintf("http://%s:%d/api/v1/rule_updated", detecotrHost, detectorPort)
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := jsoniter.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	} else if data["status"] != float64(200) {
		return errors.New("cannot send rule updated requset to detector")
	}

	return nil
}
