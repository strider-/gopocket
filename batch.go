package gopocket

type Batch struct {
	actions []interface{}
}

type ModifyAction string
type TagAction string

const (
	ACTION_ARCHIVE    ModifyAction = "archive"
	ACTION_READD      ModifyAction = "readd"
	ACTION_FAVORITE   ModifyAction = "favorite"
	ACTION_UNFAVORITE ModifyAction = "unfavorite"
	ACTION_DELETE     ModifyAction = "delete"
	TAGS_ADD          TagAction    = "tags_add"
	TAGS_REMOVE       TagAction    = "tags_remove"
	TAGS_REPLACE      TagAction    = "tags_replace"
	TAGS_CLEAR        TagAction    = "tags_clear"
)

func (b *Batch) push(op interface{}) {
	b.actions = append(b.actions, op)
}

func (b *Batch) Count() int {
	return len(b.actions)
}

func (b *Batch) Clear() {
	b.actions = nil
}

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

func (b *Batch) Action(action ModifyAction, itemId int) {
	b.push(struct {
		Action string `json:"action"`
		ItemId int    `json:"item_id"`
	}{
		string(action), itemId,
	})
}

func (b *Batch) Tag(action TagAction, itemId int, tags []string) {
	b.push(struct {
		Action string   `json:"action"`
		ItemId int      `json:"item_id"`
		tags   []string `json:"tags,omitempty"`
	}{
		string(action), itemId, tags,
	})
}

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
