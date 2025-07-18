package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

func Channel5(c http.Client) Result {
	timestamp := time.Now().Unix()
	resp, err := GET(c, "https://cassie.channel5.com/api/v2/live_media/my5desktopng/C5.json?timestamp="+strconv.FormatInt(timestamp, 10)+"&auth=0_rZDiY0hp_TNcDyk2uD-Kl40HqDbXs7hOawxyqPnbI")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}

	var res struct {
		Code string `json:"code"`
	}

	if err := json.Unmarshal(b, &res); err != nil {
		if err.Error() == `invalid character '<' looking for beginning of value` {
			return Result{Status: StatusBanned}
		}
		return Result{Status: StatusFailed, Err: err}
	}
	switch res.Code {
	case "3000":
		return Result{Status: StatusNo}
	case "3001":
		return Result{Status: StatusNo, Info: "Proxy Detected"}
	case "4003":
		return Result{Status: StatusOK}
	default:
		return Result{Status: StatusUnexpected}
	}
}
