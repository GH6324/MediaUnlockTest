package mediaunlocktest

import (
	"net/http"
)

func PrettyDerbyJP(c http.Client) Result {
	for i := 0; i < 3; i++ {
		//resp, err := GET_Dalvik(c, "https://api-umamusume.cygames.jp/")
		resp, err := GET(c, "https://api-umamusume.cygames.jp/",
		    H{"user-agent", UA_Dalvik},
		    H{"connection", "keep-alive"},
		)
		if err != nil {
			return Result{Status: StatusNetworkErr, Err: err}
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 404:
			return Result{Status: StatusOK}
		}
	}
	return Result{Status: StatusNo}
}
