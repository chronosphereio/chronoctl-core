package pagination

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockListFn(p Page) ([]string, string, error) {
	switch p.Token {
	case "":
		return []string{"item1", "item2"}, "token1", nil
	case "token1":
		return []string{"item3", "item4"}, "", nil
	default:
		return nil, "", errors.New("unexpected token")
	}
}

func mockListFnWithError(p Page) ([]string, string, error) {
	if p.Token == "" {
		return nil, "", errors.New("mock error")
	}
	return nil, "", nil
}

func TestListWithPagination(t *testing.T) {
	page := Page{Size: 4, Token: ""}

	// Call List with mock function
	items, token, err := List(page, mockListFn)

	assert.Nil(t, err)
	assert.Equal(t, []string{"item1", "item2", "item3", "item4"}, items)
	assert.Equal(t, "", token) // no next token since results are exhausted
}

func TestListExhaustedResults(t *testing.T) {
	page := Page{Size: 2, Token: ""}

	// Call List with mock function
	items, token, err := List(page, mockListFn)

	assert.Nil(t, err)
	assert.Equal(t, []string{"item1", "item2"}, items)
	assert.Equal(t, "token1", token) // should have a next token
}

func TestListWithError(t *testing.T) {
	page := Page{Size: 2, Token: ""}

	// Call List with mock function that returns an error
	items, token, err := List(page, mockListFnWithError)

	assert.NotNil(t, err)
	assert.Nil(t, items)
	assert.Equal(t, "", token) // no next token since an error occurred
}

func TestListWithEmptyResult(t *testing.T) {
	mockEmptyListFn := func(p Page) ([]string, string, error) {
		return []string{}, "", nil
	}

	page := Page{Size: 2, Token: ""}

	// Call List with mock function that returns an empty list
	items, token, err := List(page, mockEmptyListFn)

	assert.Nil(t, err)
	assert.Empty(t, items)
	assert.Equal(t, "", token) // no next token since no items are returned
}

func TestListExactPageSize(t *testing.T) {
	mockExactPageSizeFn := func(p Page) ([]string, string, error) {
		return []string{"item1", "item2"}, "", nil
	}

	page := Page{Size: 2, Token: ""}

	// Call List with mock function that returns exactly the page size
	items, token, err := List(page, mockExactPageSizeFn)

	assert.Nil(t, err)
	assert.Equal(t, []string{"item1", "item2"}, items)
	assert.Equal(t, "", token) // no next token since results are exhausted
}

func TestListZeroPageSize(t *testing.T) {
	mockExactPageSizeFn := func(p Page) ([]string, string, error) {
		if p.Token == "t1" {
			return []string{"item3", "item4"}, "", nil
		}
		return []string{"item1", "item2"}, "t1", nil
	}

	page := Page{}

	// Call List with mock function that returns exactly the page size
	items, token, err := List(page, mockExactPageSizeFn)

	assert.Nil(t, err)
	assert.Equal(t, []string{"item1", "item2", "item3", "item4"}, items)
	assert.Equal(t, "", token) // no next token since results are exhausted
}
