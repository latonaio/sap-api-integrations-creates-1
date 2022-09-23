package requests

type HeaderItem struct {
	Header
	ToItem `json:"to_XXXXXXXXXXXXItem"`
}

type ToItem struct {
	Results []Item `json:"results"`
}
