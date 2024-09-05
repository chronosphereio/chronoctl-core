package pagination

import "github.com/chronosphereio/chronoctl/src/cmd/pkg/pagination"

// Page is a list page.
type Page struct {
	Size  int64
	Token string
}

// List invokes the listFn repeatedly accumulating results per page until p.Size is met or results are exhausted.
func List[T any](p Page,
	listFn func(p Page) (items []T, token string, err error),
) ([]T, string, error) {
	var (
		items     []T
		pageSize  = p.Size
		nextToken = p.Token
	)

	// loop until we've fetched enough results
	for {
		current, token, err := listFn(Page{Size: pageSize, Token: nextToken})
		if err != nil {
			return nil, "", err
		}

		nextToken = token
		items = append(items, current...)

		if p.Size > 0 && len(items) >= int(p.Size) {
			return items, nextToken, nil
		}

		// no more results to fetch
		if nextToken == "" {
			return items, nextToken, nil
		}

		pageSize = int64(pagination.CalculatePageSize(pagination.Calculation{
			GotItems:    len(items),
			MaxItems:    int(p.Size),
			MaxPageSize: len(current),
		}))
	}
}
