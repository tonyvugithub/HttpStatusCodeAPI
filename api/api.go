package api

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/tonyvugithub/statusCodeApi/helpers"
	"github.com/tonyvugithub/statusCodeApi/customFileServer"
)

func extractDelayQuery(r *http.Request) (int, time.Duration, error) {
	//Extract "delay" query value if any
	delays, isProvided := r.URL.Query()["delay"]
	var err error = nil
	var val int = -1
	var timeUnit time.Duration = time.Second
	b := regexp.MustCompile(`^[1-9][0-9]*$`)
	s := regexp.MustCompile(`^[1-9][0-9]*s$`)
	ms := regexp.MustCompile(`^[1-9][0-9]*ms$`)

	//Check if the delay query is provided and extract value
	if isProvided {
		delay := delays[0]
		if len(delay) > 0 {
			if b.MatchString(delay) {
				val, _ = strconv.Atoi(delay)
			} else if s.MatchString(delay) {
				val, _ = strconv.Atoi(delay[:len(delay)-1])
			} else if ms.MatchString(delay) {
				val, _ = strconv.Atoi(delay[:len(delay)-2])
				timeUnit = time.Millisecond
			} else {
				err = errors.New("Invalid input for delay query")
			}
		} else {
			err = errors.New("delay query needs to have a value")
		}
	}
	return val, timeUnit, err
}

//Request handler
func getStatusCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	vars := mux.Vars(r)
	code := vars["code"]
	delayVal, timeUnit, err := extractDelayQuery(r)
	if err == nil {
		time.Sleep(time.Duration(delayVal) * timeUnit)
	}
	switch code {
	case "100":
		w.WriteHeader(http.StatusContinue)
		w.Write([]byte(`{"status": "100 Continue"}`))
	case "101":
		w.WriteHeader(http.StatusSwitchingProtocols)
		w.Write([]byte(`{"status": "101 Switching Protocols "}`))
	case "102":
		w.WriteHeader(http.StatusProcessing)
		w.Write([]byte(`{"status": "102 Processing "}`))
	case "103":
		w.WriteHeader(http.StatusEarlyHints)
		w.Write([]byte(`{"status": "103 Early Hints "}`))
	case "200":
	    //w.WriteHeader(http.StatusEarlyHints)
		customFileServer.ServeFile(w,r,"static/index.html", 403)
	case "201":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status": "201 Created"}`))
	case "202":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"status": "202 Accepted"}`))
	case "203":
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		w.Write([]byte(`{"status": "203 Non Authorative Info"}`))
	case "204":
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"status": "204 No Content"}`))
	case "205":
		w.WriteHeader(http.StatusResetContent)
		w.Write([]byte(`{"status": "205 Reset Content"}`))
	case "206":
		w.WriteHeader(http.StatusPartialContent)
		w.Write([]byte(`{"status": "206 Partial Content"}`))
	case "207":
		w.WriteHeader(http.StatusMultiStatus)
		w.Write([]byte(`{"status": "207 Multi Status"}`))
	case "208":
		w.WriteHeader(http.StatusAlreadyReported)
		w.Write([]byte(`{"status": "208 Already Reported"}`))
	case "226":
		w.WriteHeader(http.StatusIMUsed)
		w.Write([]byte(`{"status": "226 IM Used"}`))
	case "300":
		w.WriteHeader(http.StatusMultipleChoices)
		w.Write([]byte(`{"status": "300 Multiple Choices"}`))
	case "301":
		w.WriteHeader(http.StatusMovedPermanently)
		w.Write([]byte(`{"status": "301 Moved Permanently"}`))
	case "302":
		w.WriteHeader(http.StatusFound)
		w.Write([]byte(`{"status": "302 Found"}`))
	case "303":
		w.WriteHeader(http.StatusSeeOther)
		w.Write([]byte(`{"status": "303 See Other"}`))
	case "304":
		w.WriteHeader(http.StatusNotModified)
		w.Write([]byte(`{"status": "304 Not Modified"}`))
	case "305":
		w.WriteHeader(http.StatusUseProxy)
		w.Write([]byte(`{"status": "305 Use Proxy"}`))
	case "306":
		w.WriteHeader(306)
		w.Write([]byte(`{"status": "306 Unused"}`))
	case "307":
		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Write([]byte(`{"status": "307 Temporary Redirect"}`))
	case "308":
		w.WriteHeader(http.StatusPermanentRedirect)
		w.Write([]byte(`{"status": "308 Permanent Redirect"}`))
	case "400":
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status": "400 Bad Request"}`))
	case "401":
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"status": "401 Unauthorized"}`))
	case "402":
		w.WriteHeader(http.StatusPaymentRequired)
		w.Write([]byte(`{"status": "402 Payment Required"}`))
	case "403":
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"status": "403 Forbidden"}`))
	case "404":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status": "404 Not Found"}`))
	case "405":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"status": "405 Method Not Allowed"}`))
	case "406":
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(`{"status": "406 Not Acceptable"}`))
	case "407":
		w.WriteHeader(http.StatusProxyAuthRequired)
		w.Write([]byte(`{"status": "407 Proxy Authentication Required"}`))
	case "408":
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte(`{"status": "408 Request Timeout"}`))
	case "409":
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"status": "409 Conflict"}`))
	case "410":
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(`{"status": "410 Gone"}`))
	case "411":
		w.WriteHeader(http.StatusLengthRequired)
		w.Write([]byte(`{"status": "411 Length Required"}`))
	case "412":
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte(`{"status": "412 Precondition Failed"}`))
	case "413":
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		w.Write([]byte(`{"status": "413 Request Entity Too Large"}`))
	case "414":
		w.WriteHeader(http.StatusRequestURITooLong)
		w.Write([]byte(`{"status": "414 Request URI Too Long"}`))
	case "415":
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(`{"status": "415 Unsupported Media Type"}`))
	case "416":
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
		w.Write([]byte(`{"status": "416 Requested Range Not Satisfiable"}`))
	case "417":
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(`{"status": "417 Expectation Failed"}`))
	case "418":
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte(`{"status": "418 Teapot"}`))
	case "421":
		w.WriteHeader(http.StatusMisdirectedRequest)
		w.Write([]byte(`{"status": "421 Misdirected Request"}`))
	case "422":
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"status": "422 Unprocessable Entity"}`))
	case "423":
		w.WriteHeader(http.StatusLocked)
		w.Write([]byte(`{"status": "423 Locked"}`))
	case "424":
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte(`{"status": "424 Failed Dependency"}`))
	case "426":
		w.WriteHeader(http.StatusUpgradeRequired)
		w.Write([]byte(`{"status": "426 Upgrade Required"}`))
	case "428":
		w.WriteHeader(http.StatusPreconditionRequired)
		w.Write([]byte(`{"status": "428 Precondition Failed"}`))
	case "429":
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"status": "429 Too Many Requests"}`))
	case "431":
		w.WriteHeader(http.StatusRequestHeaderFieldsTooLarge)
		w.Write([]byte(`{"status": "431 Request Header Fields Too Large"}`))
	case "451":
		w.WriteHeader(http.StatusUnavailableForLegalReasons)
		w.Write([]byte(`{"status": "451 Unavailable For Legal Reasons"}`))
	case "500":
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status": "500 Internal Server Error"}`))
	case "501":
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"status": "501 Not Implemented"}`))
	case "502":
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(`{"status": "502 Bad Gateway"}`))
	case "503":
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"status": "503 Service Unavailable"}`))
	case "504":
		w.WriteHeader(http.StatusGatewayTimeout)
		w.Write([]byte(`{"status": "504 Gateway Timeout"}`))
	case "505":
		w.WriteHeader(http.StatusHTTPVersionNotSupported)
		w.Write([]byte(`{"status": "505 HTTP Version Not Supported"}`))
	case "506":
		w.WriteHeader(http.StatusVariantAlsoNegotiates)
		w.Write([]byte(`{"status": "506 Variant Also Negotiates"}`))
	case "507":
		w.WriteHeader(http.StatusInsufficientStorage)
		w.Write([]byte(`{"status": "507 Insufficient Storage"}`))
	case "508":
		w.WriteHeader(http.StatusLoopDetected)
		w.Write([]byte(`{"status": "508 Loop Detected"}`))
	case "509":
		w.WriteHeader(509)
		w.Write([]byte(`{"status": "509 Unassigned"}`))
	case "510":
		w.WriteHeader(http.StatusNotExtended)
		w.Write([]byte(`{"status": "510 Not Extended"}`))
	case "511":
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		w.Write([]byte(`{"status": "511 Network Authentication Required"}`))
	}
}

func HandleRequests() {
	//Create a mux router object to extract params
	router := mux.NewRouter().StrictSlash(true)

	//Set up request handler
	router.HandleFunc("/status/{code}", getStatusCode)
	//router.Use(mux.CORSMethodMiddleware(router))
	//Listen for request and log error if any
	log.Fatal(http.ListenAndServe(helpers.Port(), router))
}
