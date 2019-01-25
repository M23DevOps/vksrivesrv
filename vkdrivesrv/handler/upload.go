package handler

import (
    "net/http"
    _"fmt"
            
    "../vkservice"
    "../db"
    "../config"
)

const _10Mb = (1 << 20)*10

//UploadImage is used to upload image on server
var UploadImage = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
    //traceHTTPRequest(request)

    vLabel := request.URL.Query().Get("label")
    vSession := request.URL.Query().Get("session")

    request.ParseMultipartForm(_10Mb)
    file, _, err := request.FormFile("file")
    if err != nil {
        responseMessage(GeneralResponse{5, err.Error()}, response, request)
        return
    }

    defer file.Close()

    var albumId string
    if vSession != "" {
        albumId = config.Config.Bot.PrivateAlbumId
    } else
    {
        albumId = config.Config.Bot.PublicAlbumId
    }

    urlImage, err := vkservice.UploadImage(file, albumId)
    if err != nil {
        responseMessage(GeneralResponse{5, err.Error()}, response, request)
        return
    }

    if vSession != "" {
        _, session, err := ValidateSession(vSession, request.UserAgent())
        if err != nil {
            responseMessage(GeneralResponse{4, err.Error()}, response, request)
            return
        }

        err = db.InsertImage(session.UserId, 0, urlImage, "Label")
        if err != nil {
            responseMessage(GeneralResponse{5, err.Error()}, response, request)
            return
        }
    }
    responseMessage(UploadResponse{GeneralResponse{0,""}, Image{urlImage, 0, vLabel}}, response, request)
    return
})