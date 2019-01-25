package handler

type GeneralResponse struct {
	ErrorCode   	int 	`json:"err_code"`
	ErrorMessage    string 	`json:"err_msg"`
}

type LoginResponse struct {
	GeneralResponse
	FirstName		string `json:"first_name"`
	LastName		string `json:"last_name"`
	PhotoURL		string `json:"photo_url"`
	Session			string `json:"session"`
}

type UploadResponse struct {
	GeneralResponse
	Image
}

type GetImagesResponse struct {
	GeneralResponse
	Images []Image `json:"images"`
}

type Image struct {
	ImageURL	string 	`json:"image_url"`
	FolderId	int 	`json:"folder_id"`
	Label		string 	`json:"label"`
}

