package gopocket

import "time"

type Options struct {
	dict map[string]interface{}
}

type ArticleState string
type ContentType string
type SortOption string
type DetailType string

const (
	STATE_UNREAD    ArticleState = "unread"
	STATE_ALL       ArticleState = "all"
	STATE_ARCHIVED  ArticleState = "archived"
	CONTENT_ARTICLE ContentType  = "article"
	CONTENT_VIDEO   ContentType  = "video"
	CONTENT_IMAGE   ContentType  = "image"
	SORT_NEWEST     SortOption   = "newest"
	SORT_OLDEST     SortOption   = "oldest"
	SORT_TITLE      SortOption   = "title"
	SORT_SITE       SortOption   = "site"
	DETAIL_SIMPLE   DetailType   = "simple"
	DETAIL_COMPLETE DetailType   = "complete"
)

func (o *Options) Clear() {
	o.dict = make(map[string]interface{})
}

func (o *Options) State(state ArticleState) {
	o.dict["state"] = state
}

func (o *Options) Favorited() {
	o.dict["favorite"] = 1
}

func (o *Options) Unfavorited() {
	o.dict["favorite"] = 0
}

func (o *Options) Tag(tag string) {
	if tag != "" {
		o.dict["tag"] = tag
	} else {
		o.dict["tag"] = "_untagged_"
	}
}

func (o *Options) ContentType(content ContentType) {
	o.dict["contentType"] = content
}

func (o *Options) Sort(sort SortOption) {
	o.dict["sort"] = sort
}

func (o *Options) Detail(detail DetailType) {
	o.dict["detailType"] = detail
}

func (o *Options) Search(term string) {
	o.dict["search"] = term
}

func (o *Options) Domain(domain string) {
	o.dict["domain"] = domain
}

func (o *Options) Since(time time.Time) {
	o.dict["since"] = time.Unix()
}

func (o *Options) Limit(count, offset int) {
	o.dict["count"] = count
	o.dict["offset"] = offset
}
