package gopocket

// Batch allows for multiple operations to be sent to the API in a single request.
type Batch struct {
	actions []interface{}
}

// ModifyAction specifies an action that will update an article
type ModifyAction string

// TagAction specified an action that will update article tags
type TagAction string

const (
	// ACTION_ARCHIVE will archive an item
	ACTION_ARCHIVE ModifyAction = "archive"
	// ACTION_READD will re-add an item to a users queue
	ACTION_READD ModifyAction = "readd"
	// ACTION_FAVORITE will mark an item as favorited
	ACTION_FAVORITE ModifyAction = "favorite"
	// ACTION_UNFAVORITE will mark an item as unfavorited
	ACTION_UNFAVORITE ModifyAction = "unfavorite"
	// ACTION_DELETE will delete an item
	ACTION_DELETE ModifyAction = "delete"
	// TAGS_ADD will add one or more tags to an item
	TAGS_ADD TagAction = "tags_add"
	// TAGS_REMOVE will remove one or more tags from an item
	TAGS_REMOVE TagAction = "tags_remove"
	// TAGS_REPLACE will replace all of the tags for an item with one of more tags
	TAGS_REPLACE TagAction = "tags_replace"
	// TAGS_CLEAR will remove all tags from an item.
	TAGS_CLEAR TagAction = "tags_clear"
)

func (b *Batch) push(op interface{}) {
	b.actions = append(b.actions, op)
}

// Count returns the current number of actions that have been added to the batch.
func (b *Batch) Count() int {
	return len(b.actions)
}

// Clear will remove every action from the batch.
func (b *Batch) Clear() {
	b.actions = nil
}

// Add will add a Url to a users pocket queue. Title will be used as a fallback if the API couldn't accurately parse the
// title of the article, sending an empty string is acceptable.
func (b *Batch) Add(url, title string, tags []string) {
	b.push(struct {
		Action string   `json:"action"`
		Url    string   `json:"url"`
		Title  string   `json:"title,omitempty"`
		Tags   []string `json:"tags"`
	}{
		"add", url, title, tags,
	})
}

// Action will perform the specified ModifyAction on the given item id.
func (b *Batch) Action(action ModifyAction, itemId int) {
	b.push(struct {
		Action string `json:"action"`
		ItemId int    `json:"item_id"`
	}{
		string(action), itemId,
	})
}

// Tag will perform the specified TagAction on the given item id.  When using TAGS_CLEAR,
// use an empty array/slice for the tags parameter.
func (b *Batch) Tag(action TagAction, itemId int, tags []string) {
	b.push(struct {
		Action string   `json:"action"`
		ItemId int      `json:"item_id"`
		tags   []string `json:"tags,omitempty"`
	}{
		string(action), itemId, tags,
	})
}

// RenameTag will rename a tag, affecting all items with this tag.
func (b *Batch) RenameTag(itemId int, oldTag, newTag string) {
	b.push(struct {
		Action string `json:"action"`
		ItemId int    `json:"item_id"`
		OldTag string `json:"old_tag"`
		NewTag string `json:"new_tag"`
	}{
		"tag_rename", itemId, oldTag, newTag,
	})
}
