package gopocket

import "time"

// Options type holds all the current filter options for a retrieval request.
type Options struct {
	dict map[string]interface{}
}

// ArticleState represents the state of an article
type ArticleState string

// ContentType represents the type of article
type ContentType string

// SortOption represents the sort method for returned items
type SortOption string

// DetailType represents the level of detail returned by the API
type DetailType string

const (
	// STATE_UNREAD specifies only unread items (default)
	STATE_UNREAD ArticleState = "unread"
	// STATE_ALL specifies all items
	STATE_ALL ArticleState = "all"
	// STATE_ARCHIVED specifies only archived items
	STATE_ARCHIVED ArticleState = "archived"
	// CONTENT_ARTICLE specifies only items that are articles
	CONTENT_ARTICLE ContentType = "article"
	// CONTENT_VIDEO specifies only items that are videos
	CONTENT_VIDEO ContentType = "video"
	// CONTENT_IMAGE specifies only items that are images
	CONTENT_IMAGE ContentType = "image"
	// SORT_NEWEST specifies items will be returned newest to oldest
	SORT_NEWEST SortOption = "newest"
	// SORT_OLDEST specifies items will be returned oldest to newest
	SORT_OLDEST SortOption = "oldest"
	// SORT_TITLE specifies items will be returned in order of title, alphabetically
	SORT_TITLE SortOption = "title"
	// SORT_SITE specifies items will be returned in order of url, alphabetically
	SORT_SITE SortOption = "site"
	// DETAIL_SIMPLE specifies only titles and urls will be returned.
	DETAIL_SIMPLE DetailType = "simple"
	// DETAIL_COMPLETE specifies all item data (tags, images, authors, ect) will be returned.
	DETAIL_COMPLETE DetailType = "complete"
)

// Clear resets all options
func (o *Options) Clear() {
	o.dict = make(map[string]interface{})
}

// State adds a filter for article state
func (o *Options) State(state ArticleState) *Options {
	o.dict["state"] = state
	return o
}

// Favorited adds a filter for only favorited items. Mutually exclusive with Unfavorited().
func (o *Options) Favorited() *Options {
	o.dict["favorite"] = 1
	return o
}

// Unfavorited adds a filter for only unfavorited items. Multually exclusive with Favorited().
func (o *Options) Unfavorited() *Options {
	o.dict["favorite"] = 0
	return o
}

// Tag adds a filter for items containing this tag
func (o *Options) Tag(tag string) *Options {
	if tag != "" {
		o.dict["tag"] = tag
	} else {
		o.dict["tag"] = "_untagged_"
	}
	return o
}

// ContentType adds a filter for the items of a specific content
func (o *Options) ContentType(content ContentType) *Options {
	o.dict["contentType"] = content
	return o
}

// Sort specifies the order to return results in
func (o *Options) Sort(sort SortOption) *Options {
	o.dict["sort"] = sort
	return o
}

// Detail specifies how much information should be returned
func (o *Options) Detail(detail DetailType) *Options {
	o.dict["detailType"] = detail
	return o
}

// Search adds a filter for items whose title or url contains the given term
func (o *Options) Search(term string) *Options {
	o.dict["search"] = term
	return o
}

// Domain adds a filter for items from a specific domain
func (o *Options) Domain(domain string) *Options {
	o.dict["domain"] = domain
	return o
}

// Since adds a filter for items modified since the given time
func (o *Options) Since(time time.Time) *Options {
	o.dict["since"] = time.Unix()
	return o
}

// Limit specifies a count of items to return, starting from a given offset.
func (o *Options) Limit(count, offset int) *Options {
	o.dict["count"] = count
	o.dict["offset"] = offset
	return o
}
