package handler

import (
	_"fmt"
	"net/http"
    "encoding/json"
	_"errors"    
	
	"../db"
)


//Authorize is used to authorize user in application
var Logout = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
    //traceHTTPRequest(request)

    session := request.URL.Query().Get("session")

	err := logoutUser(session, request.UserAgent())
	
    if err != nil {
        json.NewEncoder(response).Encode(GeneralResponse{3, err.Error()})
    } else {
        json.NewEncoder(response).Encode(GeneralResponse{0, ""})
    }
})


func logoutUser(strSession, strUserAgent string) (error){

	claim, _, err := ValidateSession(strSession, strUserAgent)
	if err != nil {
		return err		
	}

	db.DeleteSession(claim.Token, strUserAgent)

    return nil
}