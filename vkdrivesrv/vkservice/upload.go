package vkservice

import (
	_"fmt"
	"io/ioutil"
	"net/http"
	_"os"
	_"strings"
	"bytes"
	"mime/multipart"
	_"path/filepath"
	"io"
	"errors"
	"strconv"

	"github.com/bitly/go-simplejson"

	"../config"
)


//UploadImg upload image on vk.com and return URL of image
func UploadImage(imgFile io.Reader, strAlbumId string) (string, error){

	token := config.Config.Bot.Token
	albumId := strAlbumId

	url, err := getVkUploadServer(albumId, token)
	if err != nil {
		return "", err
	}

	server, photosList, hash, err := uploadImageOnVk(url, imgFile)
	if err != nil {
		return "", err
	}

	imageURL, err := saveImageOnVk(albumId, server, photosList, hash, token)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func getVkUploadServer(strAlbumId, strToken string) (string, error){
	resp, err := http.Get("https://api.vk.com/method/photos.getUploadServer?album_id=" + strAlbumId + "&access_token=" + strToken + "&v=5.67")
	if err != nil {
		return "", errors.New("Can't get upload Server")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(body)
	url, _ := json.Get("response").Get("upload_url").String()

	return url, nil
}

func uploadImageOnVk(strURL string, imgFile io.Reader) (string, string, string, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("photo", "img.jpg")
	if err != nil {
        return "","","", errors.New("Can't create form file")
	}

	_, err = io.Copy(part, imgFile)
	if err != nil {
        return "","","", errors.New("Can't copy file")
    }
	writer.Close()

	resp, err := http.DefaultClient.Post(strURL, writer.FormDataContentType(), body)
	if err != nil {
        return "","","", errors.New("Can't send POST request")
	}
	defer resp.Body.Close()
	
	resBody, _ := ioutil.ReadAll(resp.Body)

	json, _ := simplejson.NewJson(resBody)

	photosList,_ := json.Get("photos_list").String()
	if photosList == "[]" {
		return "","","", errors.New("Can't upload file on VK")
	}

	server, _ := json.Get("server").Int()
	hash, _ := json.Get("hash").String()

	return strconv.Itoa(server), photosList, hash, nil
}

func saveImageOnVk(strAlbumId, strServer, strPhotosList, strHash, strToken string) (string, error) {
	resp, err := http.Get("https://api.vk.com/method/photos.save?album_id=" + strAlbumId + "&server=" + strServer + "&photos_list=" + strPhotosList + "&hash=" + strHash + "&access_token=" + strToken + "&v=5.67")
	if err != nil {
        return "", errors.New("Can't send request for saving photo")
	}

	resbody, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(resbody)
	
	errCode := 0
	errCode, _ = json.Get("error").Get("error_code").Int()
	errMsg, _ := json.Get("error").Get("error_msg").String()

	if errCode != 0 {
		return "", errors.New("Vk Error: " + strconv.Itoa(errCode) + ". Error Message: " + errMsg)
	}

	imageURL, _ := json.Get("response").GetIndex(0).Get("photo_2560").String()
	if imageURL != "" {return imageURL, nil}

	imageURL, _ = json.Get("response").GetIndex(0).Get("photo_1280").String()
	if imageURL != "" {return imageURL, nil}

	imageURL, _ = json.Get("response").GetIndex(0).Get("photo_807").String()
	if imageURL != "" {return imageURL, nil}

	imageURL, _ = json.Get("response").GetIndex(0).Get("photo_604").String()
	if imageURL != "" {return imageURL, nil}

	imageURL, _ = json.Get("response").GetIndex(0).Get("photo_130").String()
	if imageURL != "" {return imageURL, nil}

	imageURL, _ = json.Get("response").GetIndex(0).Get("photo_75").String()

	return imageURL, nil
}
