package gopocket

import (
	"bytes"
	"encoding/json"
)

// ApiRate contains rate information from the API
type ApiRate struct {
	Error         string `header:"X-Error"`
	ErrorCode     int    `header:"X-Error-Code"`
	UserLimit     int    `header:"X-Limit-User-Limit"`
	UserRemaining int    `header:"X-Limit-User-Remaining"`
	UserReset     int    `header:"X-Limit-User-Reset"`
	KeyLimit      int    `header:"X-Limit-Key-Limit"`
	KeyRemaining  int    `header:"X-Limit-Key-Remaining"`
	KeyReset      int    `header:"X-Limit-Key-Reset"`
}

// AddResponse is the result type when calling Pocket.Add
type AddResponse struct {
	Item   responseItem `json:"item"`
	Status int          `json:"status"`
}

// ModifyResponse is the result type when calling Pocket.Modify
type ModifyResponse struct {
	Results []responseItem `json:"action_results"`
	Status  int            `json:"status"`
}

// RetrieveResponse is the result type when calling Pocket.Retrieve
type RetrieveResponse struct {
	List   retrieveMap `json:"list"`
	Status int         `json:"status"`
}

type responseItem struct {
	ItemId            int       `json:"item_id,string"`
	NormalUrl         string    `json:"normal_url"`
	ResolvedNormalUrl string    `json:"resolved_normal_url"`
	GivenUrl          string    `json:"given_url"`
	ResolvedId        int       `json:"resolved_id,string"`
	ResolvuedUrl      string    `json:"resolved_url"`
	ExtendedItemId    int       `json:"extended_item_id,string"`
	DomainId          int       `json:"domain_id,string"`
	OriginDomainId    int       `json:"origin_domain_id,string"`
	ResponseCode      int       `json:"response_code,string"`
	MimeType          string    `json:"mime_type"`
	ContentLength     int       `json:"content_length,string"`
	Encoding          string    `json:"encoding"`
	DateResolved      string    `json:"date_resolved"`
	DatePublished     string    `json:"date_published"`
	Title             string    `json:"title"`
	Excerpt           string    `json:"excerpt"`
	WordCount         int       `json:"word_count,string"`
	LoginRequired     int       `json:"login_required,string"`
	HasImage          int       `json:"has_image,string"`
	HasVideo          int       `json:"has_video,string"`
	IsIndex           int       `json:"is_index,string"`
	IsArticle         int       `json:"is_article,string"`
	UsedFallback      int       `json:"used_fallback,string"`
	Favorite          int       `json:"favorite,string"`
	Status            int       `json:"status,string"`
	Authors           authorMap `json:"authors"`
	Images            imageMap  `json:"images"`
	Videos            videoMap  `json:"videos"`
}

type retrieveMap map[string]responseItem

func (r *retrieveMap) UnmarshalJSON(b []byte) error {
	if bytes.Compare([]byte("[]"), b) == 0 {
		return nil
	}
	return json.Unmarshal(b, (*map[string]responseItem)(r))
}

// Author contains author information for an item
type Author struct {
	AuthorID int    `json:"author_id,string"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

type authorMap map[string]Author

func (a *authorMap) UnmarshalJSON(b []byte) error {
	if bytes.Compare([]byte("[]"), b) == 0 {
		return nil
	}
	return json.Unmarshal(b, (*map[string]Author)(a))
}

// Image contains image information for an item
type Image struct {
	ItemId  int    `json:"item_id,string"`
	ImageId int    `json:"image_id,string"`
	Src     string `json:"src"`
	Width   int    `json:"width,string"`
	Height  int    `json:"height,string"`
	Credit  string `json:"credit"`
	Caption string `json:"caption"`
}

type imageMap map[string]Image

func (i *imageMap) UnmarshalJSON(b []byte) error {
	if bytes.Compare([]byte("[]"), b) == 0 {
		return nil
	}
	return json.Unmarshal(b, (*map[string]Image)(i))
}

// Video contains video information for an item
type Video struct {
	ItemId  int    `json:"item_id,string"`
	VideoId int    `json:"video_id,string"`
	Src     string `json:"src"`
	Width   int    `json:"width,string"`
	Height  int    `json:"height,string"`
	Type    int    `json:"type,string"`
	Vid     string `json:"vid"`
}

type videoMap map[string]Video

func (v *videoMap) UnmarshalJSON(b []byte) error {
	if bytes.Compare([]byte("[]"), b) == 0 {
		return nil
	}
	return json.Unmarshal(b, (*map[string]Video)(v))
}
