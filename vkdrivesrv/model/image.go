package model

import "time"

//Image is table for images
type Image struct {
	Id			int
	UserId		int
	URL			string
	Date		time.Time `sql:"DEFAULT:now()"`
	FolderId	int
	Label		string
	IsDelete	int
}
