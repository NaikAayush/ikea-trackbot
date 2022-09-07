package main

type ResponseBody struct {
	Data struct {
		Order struct {
			Status          string `json:"status"`
			DeliveryMethods []struct {
				Status       string `json:"status"`
				DeliveryDate struct {
					Actual        interface{} `json:"actual"`
					EstimatedFrom struct {
						Time                   string `json:"time"`
						Date                   string `json:"date"`
						FormattedLocal         string `json:"formattedLocal"`
						FormattedShortDate     string `json:"formattedShortDate"`
						FormattedLongDate      string `json:"formattedLongDate"`
						FormattedShortDateTime string `json:"formattedShortDateTime"`
						FormattedLongDateTime  string `json:"formattedLongDateTime"`
						Typename               string `json:"__typename"`
					} `json:"estimatedFrom"`
					EstimatedTo struct {
						Time                   string `json:"time"`
						Date                   string `json:"date"`
						FormattedLocal         string `json:"formattedLocal"`
						FormattedShortDate     string `json:"formattedShortDate"`
						FormattedLongDate      string `json:"formattedLongDate"`
						FormattedShortDateTime string `json:"formattedShortDateTime"`
						FormattedLongDateTime  string `json:"formattedLongDateTime"`
						Typename               string `json:"__typename"`
					} `json:"estimatedTo"`
					DateTimeRange string `json:"dateTimeRange"`
					TimeZone      string `json:"timeZone"`
					Typename      string `json:"__typename"`
				} `json:"deliveryDate"`
				Type     string `json:"type"`
				Typename string `json:"__typename"`
			} `json:"deliveryMethods"`
			Typename string `json:"__typename"`
		} `json:"order"`
	} `json:"data"`
}

type Payload struct {
	OperationName string    `json:"operationName"`
	Variables     Variables `json:"variables"`
	Query         string    `json:"query"`
}
type Variables struct {
	OrderNumber string `json:"orderNumber"`
	LiteID      string `json:"liteId"`
}

const query = "query orderLookup($orderNumber: String!, $liteId: String) {\n  order(orderNumber: $orderNumber, liteId: $liteId) {\n    status\n    deliveryMethods {\n      status: status2\n      deliveryDate {\n        ...DeliveryDate\n        __typename\n      }\n      type\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DeliveryDate on DeliveryDate {\n  actual {\n    ...DateAndTime\n    __typename\n  }\n  estimatedFrom {\n    ...DateAndTime\n    __typename\n  }\n  estimatedTo {\n    ...DateAndTime\n    __typename\n  }\n  dateTimeRange\n  timeZone\n  __typename\n}\n\nfragment DateAndTime on DateAndTime {\n  time\n  date\n  formattedLocal\n  formattedShortDate\n  formattedLongDate\n  formattedShortDateTime\n  formattedLongDateTime\n  __typename\n}"
