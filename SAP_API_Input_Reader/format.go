package sap_api_input_reader

import (
	"sap-api-integrations-xxxxxx-xxxxxx-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeaderItem() *requests.HeaderItem {
	data := sdc.XXXXX
	results := make([]requests.Item, 0, len(data.XXXXXItem))

	header := sdc.ConvertToHeader()

	for i := range data.XXXXXItem {
		results = append(results, *sdc.ConvertToItem(i))
	}

	return &requests.HeaderItem{
		Header: *header,
		ToItem: requests.ToItem{
			Results: results,
		},
	}
}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.XXXXX
	return &requests.Header{
		XXXXX: data.XXXXX,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.Item {
	dataXXXXX := sdc.XXXXX
	data := sdc.XXXXX.XXXXXItem[num]
	return &requests.Item{
		XXXXX: dataXXXXX.XXXXX,
		XXXXX: data.XXXXX,
	}
}
