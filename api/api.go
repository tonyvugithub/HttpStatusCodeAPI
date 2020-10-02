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

	file := "static/index.html"

	w.Header().Set("Content-Type", "text/html")
	vars := mux.Vars(r)
	code := vars["code"]
	delayVal, timeUnit, err := extractDelayQuery(r)
	if err == nil {
		time.Sleep(time.Duration(delayVal) * timeUnit)
	}
	switch code {
	case "100":
		customFileServer.ServeFile(w,r,file, http.StatusContinue) 
	case "101":
		customFileServer.ServeFile(w,r,file, http.StatusSwitchingProtocols) 
	case "102":
		customFileServer.ServeFile(w,r,file, http.StatusProcessing) 
	case "103":
		customFileServer.ServeFile(w,r,file, http.StatusEarlyHints) 
	case "200":
		customFileServer.ServeFile(w,r,file, http.StatusOK) // here is the change
	case "201":
		customFileServer.ServeFile(w,r,file, http.StatusCreated)
	case "202":
		customFileServer.ServeFile(w,r,file, http.StatusAccepted)
	case "203":
		customFileServer.ServeFile(w,r,file, http.StatusNonAuthoritativeInfo)
	case "204":
		customFileServer.ServeFile(w,r,file, http.StatusNoContent)
	case "205":
		customFileServer.ServeFile(w,r,file, http.StatusResetContent)
	case "206":
		customFileServer.ServeFile(w,r,file, http.StatusPartialContent)
	case "207":
		customFileServer.ServeFile(w,r,file, http.StatusMultiStatus)
	case "208":
		customFileServer.ServeFile(w,r,file, http.StatusAlreadyReported)
	case "226":
		customFileServer.ServeFile(w,r,file, http.StatusIMUsed)
	case "300":
		customFileServer.ServeFile(w,r,file, http.StatusMultipleChoices)
	case "301":
		customFileServer.ServeFile(w,r,file, http.StatusMovedPermanently)
	case "302":
		customFileServer.ServeFile(w,r,file, http.StatusFound)
	case "303":
		customFileServer.ServeFile(w,r,file, http.StatusSeeOther)
	case "304":
		customFileServer.ServeFile(w,r,file, http.StatusNotModified)
	case "305":
		customFileServer.ServeFile(w,r,file, http.StatusUseProxy)
	case "306":
		customFileServer.ServeFile(w,r,file, 306)
	case "307":
		customFileServer.ServeFile(w,r,file, http.StatusTemporaryRedirect)
	case "308":
		customFileServer.ServeFile(w,r,file, http.StatusPermanentRedirect)
	case "400":
		customFileServer.ServeFile(w,r,file, http.StatusBadRequest)
	case "401":
		customFileServer.ServeFile(w,r,file, http.StatusUnauthorized)
	case "402":
		customFileServer.ServeFile(w,r,file, http.StatusPaymentRequired)
	case "403":
		customFileServer.ServeFile(w,r,file, http.StatusForbidden)
	case "404":
		customFileServer.ServeFile(w,r,file, http.StatusNotFound)
	case "405":
		customFileServer.ServeFile(w,r,file, http.StatusMethodNotAllowed)
	case "406":
		customFileServer.ServeFile(w,r,file, http.StatusNotAcceptable)
	case "407":
		customFileServer.ServeFile(w,r,file, http.StatusProxyAuthRequired)
	case "408":
		customFileServer.ServeFile(w,r,file, http.StatusRequestTimeout)
	case "409":
		customFileServer.ServeFile(w,r,file, http.StatusConflict)
	case "410":
		customFileServer.ServeFile(w,r,file, http.StatusGone)
	case "411":
		customFileServer.ServeFile(w,r,file, http.StatusLengthRequired)
	case "412":
		customFileServer.ServeFile(w,r,file, http.StatusPreconditionFailed)
	case "413":
		customFileServer.ServeFile(w,r,file, http.StatusRequestEntityTooLarge)
	case "414":
		customFileServer.ServeFile(w,r,file, http.StatusRequestURITooLong)
	case "415":
		customFileServer.ServeFile(w,r,file, http.StatusUnsupportedMediaType)
	case "416":
		customFileServer.ServeFile(w,r,file, http.StatusRequestedRangeNotSatisfiable)
	case "417":
		customFileServer.ServeFile(w,r,file, http.StatusExpectationFailed)
	case "418":
		customFileServer.ServeFile(w,r,file, http.StatusTeapot)
	case "421":
		customFileServer.ServeFile(w,r,file, http.StatusMisdirectedRequest)
	case "422":
		customFileServer.ServeFile(w,r,file, http.StatusUnprocessableEntity)
	case "423":
		customFileServer.ServeFile(w,r,file, http.StatusLocked)
	case "424":
		customFileServer.ServeFile(w,r,file, http.StatusFailedDependency)
	case "426":
		customFileServer.ServeFile(w,r,file, http.StatusUpgradeRequired)
	case "428":
		customFileServer.ServeFile(w,r,file, http.StatusPreconditionRequired)
	case "429":
		customFileServer.ServeFile(w,r,file, http.StatusTooManyRequests)
	case "431":
		customFileServer.ServeFile(w,r,file, http.StatusRequestHeaderFieldsTooLarge)
	case "451":
		customFileServer.ServeFile(w,r,file, http.StatusUnavailableForLegalReasons)
	case "500":
		customFileServer.ServeFile(w,r,file, http.StatusInternalServerError)
	case "501":
		customFileServer.ServeFile(w,r,file, http.StatusNotImplemented)
	case "502":
		customFileServer.ServeFile(w,r,file, http.StatusBadGateway)
	case "503":
		customFileServer.ServeFile(w,r,file, http.StatusServiceUnavailable)
	case "504":
		customFileServer.ServeFile(w,r,file, http.StatusGatewayTimeout)
	case "505":
		customFileServer.ServeFile(w,r,file, http.StatusHTTPVersionNotSupported)
	case "506":
		customFileServer.ServeFile(w,r,file, http.StatusVariantAlsoNegotiates)
	case "507":
		customFileServer.ServeFile(w,r,file, http.StatusInsufficientStorage)
	case "508":
		customFileServer.ServeFile(w,r,file, http.StatusLoopDetected)
	case "509":
		customFileServer.ServeFile(w,r,file, 509)
	case "510":
		customFileServer.ServeFile(w,r,file, http.StatusNotExtended)
	case "511":
		customFileServer.ServeFile(w,r,file, http.StatusNetworkAuthenticationRequired)
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
