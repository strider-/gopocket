package gopocket

import "testing"

func createTestPocket() *Pocket {
	return &Pocket{key: "0000-2885", token: "IIDX20-TRICORO", post: mockPost}
}

func mockPost(url string, requestModel interface{}, result interface{}) (rate *ApiRate, err error) {
	return &ApiRate{}, nil
}

func TestPocketAdd(t *testing.T) {
	pocket := createTestPocket()
	pocket.Add("", "", []string{})
	// test stub
}

func TestPocketModify(t *testing.T) {
	pocket := createTestPocket()
	batch := NewBatch()
	pocket.Modify(batch)
	// test stub
}

func TestPocketRetrieve(t *testing.T) {
	pocket := createTestPocket()
	opts := NewOptions()
	pocket.Retrieve(opts)
	// test stub
}
