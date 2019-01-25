package handler

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	//"strconv"
	"errors"
	"strings"

	"../db"
	"../model"
	"../crypt"
)

//StatusHandler is used for check server state
var StatusHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	//traceHTTPRequest(request)
	fmt.Fprintf(response, "Server is alive!")
})

//VerifySession is used to verify token for give access to Angular's modules 
var VerifySession = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	//traceHTTPRequest(request)
	
	session := request.URL.Query().Get("session")
	
	err := crypt.JWTVerify(session)
	if err != nil {
		responseMessage(GeneralResponse{2, err.Error()}, response, request)
        return
	} 
	responseMessage(GeneralResponse{0, ""}, response, request)
    return
})

func responseMessage(data interface{}, response http.ResponseWriter, request *http.Request) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

func traceHTTPRequest(r *http.Request) {
	fmt.Println(r.Method, r.URL, r.Proto)
    fmt.Println("Host: ", r.Host)
    for name, headers := range r.Header {
        name = strings.ToLower(name)
        for _, h := range headers {
            fmt.Println(name, h)
        }
    }
        if r.Method == "POST" {
        r.ParseForm()
        fmt.Println(r.Form.Encode())
	}
}

//ValidateSession is used to validate session and return claim and session objects
func ValidateSession(strSession, strUserAgent string) (crypt.JWTClaim, model.Session, error) {
	var claim crypt.JWTClaim
	var session model.Session

	err := crypt.JWTVerify(strSession)
	if err != nil {
		return claim, session, err
	}

	parts := strings.Split(strSession, ".")

	jsonBytes, err := crypt.DecodeSegment(parts[0])
	if err != nil {
		return claim, session, errors.New("Can't Decode Claim")
	}

	err = json.Unmarshal(jsonBytes, &claim)
	if err != nil {
		return claim, session, errors.New("Can't Unmarshal Claim")
	}

	session, err = db.SelectSession(claim.Token, strUserAgent)
	if err != nil {
		return claim, session, errors.New("No such session")
	}

	return claim, session, nil
}
