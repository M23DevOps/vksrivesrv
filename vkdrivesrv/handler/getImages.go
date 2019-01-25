package handler

import (
    "net/http"
    _"fmt"
            
    "../db"
)

//GetImages is used to get all saved images
var GetImages = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
    //traceHTTPRequest(request)

    vSession := request.URL.Query().Get("session")
    if vSession == "" {
        responseMessage(GeneralResponse{4, "No session was passed"}, response, request)
        return
    }

    _, session, err := ValidateSession(vSession, request.UserAgent())
    if err != nil {
        responseMessage(GeneralResponse{4, err.Error()}, response, request)
        return
    }

    images, err  := db.SelectImagesByUser(session.UserId)
    if err != nil {
        responseMessage(GeneralResponse{6, err.Error()}, response, request)
        return
    }

    resImages := make([]Image,0)
    for _, i := range images {
        resImages = append(resImages, Image{i.URL, i.FolderId, i.Label})
    }

    responseMessage(GetImagesResponse{GeneralResponse{0,""}, resImages}, response, request)
    return
})