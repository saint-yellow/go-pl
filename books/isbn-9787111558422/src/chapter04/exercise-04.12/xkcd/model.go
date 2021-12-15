package xkcd

import (
	"encoding/json"
	"fmt"
)

type ComicResponse struct {
    Month      string `json:"month"`
    Num        int    `json:"num"`
    Link       string `json:"link"`
    Year       string `json:"year"`
    News       string `json:"news"`
    SafeTitle  string `json:"safe_title"`
    Transcript string `json:"transcript"`
    Alt        string `json:"alt"`
    Img        string `json:"img"`
    Title      string `json:"title"`
    Day        string `json:"day"`
}


type Comic struct {
    Title       string `json:"title"`
    Number      int    `json:"number"`
    Date        string `json:"date"`
    Description string `json:"description"`
    Image       string `json:"image"`
}

func (cr ComicResponse) FormattedDate() string {
	return fmt.Sprintf("%s-%s-%s", cr.Year, cr.Month, cr.Day)
}

func (cr ComicResponse) Comic() Comic {
	return Comic{
		Title: cr.Title,
		Number: cr.Num,
		Description: cr.Alt,
		Date: cr.FormattedDate(),
		Image: cr.Img,
	}
}

func (c Comic) PrettyString() string {
	return fmt.Sprintf(
		"Title: %s\nComic No: %d\nDate: %s\nDescription: %s\nImage: %s\n",
		c.Title, c.Number, c.Date, c.Description, c.Image)
}

func (c Comic) JSON() string {
	if content, err := json.MarshalIndent(c, "", "    "); err != nil {
		return ""
	} else {
		return string(content)
	}
}

type ComicNumber int