package gopocket

import "testing"

func createTestPocket() *Pocket {
	return &Pocket{key: "0000-2885", token: "IIDX20-TRICORO", post: mockPost}
}

func mockPost(url string, requestModel interface{}, result interface{}) (rate *ApiRate, err error) {
	return &ApiRate{}, nil
}

func TestPocketAdd(t *testing.T) {
	// test stub
}

func TestPocketModify(t *testing.T) {
	// test stub
}

func TestPocketRetrieve(t *testing.T) {
	// test stub
}
