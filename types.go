package main

type ResponseBodyV2 struct {
	Data struct {
		Order struct {
			ID              string `json:"id"`
			DeliveryMethods []struct {
				ID                  string        `json:"id"`
				DeliveryNumber      interface{}   `json:"deliveryNumber"`
				ServiceID           string        `json:"serviceId"`
				Type                string        `json:"type"`
				Carrier             interface{}   `json:"carrier"`
				TrackingNumber      interface{}   `json:"trackingNumber"`
				Status              string        `json:"status"`
				PastStatuses        []string      `json:"pastStatuses"`
				TrackingLinks       []interface{} `json:"trackingLinks"`
				PartnerTrackingLink interface{}   `json:"partnerTrackingLink"`
				Quantity            int           `json:"quantity"`
				PackageDimensions   interface{}   `json:"packageDimensions"`
				DeliveryDate        struct {
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
				DeliveryAddress struct {
					CompanyName  interface{} `json:"companyName"`
					AddressLine1 interface{} `json:"addressLine1"`
					AddressLine2 interface{} `json:"addressLine2"`
					AddressLine3 interface{} `json:"addressLine3"`
					AddressLine4 interface{} `json:"addressLine4"`
					Postcode     string      `json:"postcode"`
					City         string      `json:"city"`
					Country      string      `json:"country"`
					StreetName   interface{} `json:"streetName"`
					Typename     string      `json:"__typename"`
				} `json:"deliveryAddress"`
				Articles struct {
					Quantity  int    `json:"quantity"`
					Direction string `json:"direction"`
					Any       []struct {
						Name            string      `json:"name"`
						Description     string      `json:"description"`
						Href            string      `json:"href"`
						Quantity        int         `json:"quantity"`
						DecimalQuantity float64     `json:"decimalQuantity"`
						PriceUnitText   interface{} `json:"priceUnitText"`
						ID              string      `json:"id"`
						SplitDelivery   bool        `json:"splitDelivery"`
						Image           struct {
							Small    string `json:"small"`
							Medium   string `json:"medium"`
							Large    string `json:"large"`
							Typename string `json:"__typename"`
						} `json:"image"`
						ReferencedItems []struct {
							Name            string      `json:"name"`
							Description     string      `json:"description"`
							Href            string      `json:"href"`
							Quantity        int         `json:"quantity"`
							DecimalQuantity float64     `json:"decimalQuantity"`
							PriceUnitText   interface{} `json:"priceUnitText"`
							ID              string      `json:"id"`
							SplitDelivery   bool        `json:"splitDelivery"`
							Image           struct {
								Small    string `json:"small"`
								Medium   string `json:"medium"`
								Large    string `json:"large"`
								Typename string `json:"__typename"`
							} `json:"image"`
							Typename string `json:"__typename"`
						} `json:"referencedItems"`
						Typename string `json:"__typename"`
					} `json:"any"`
					Typename string `json:"__typename"`
				} `json:"articles"`
				OpeningHours interface{} `json:"openingHours"`
				IsDeviated   bool        `json:"isDeviated"`
				StatusV2     []struct {
					DeliveryStatusV2 string `json:"deliveryStatusV2"`
					Tense            string `json:"tense"`
				} `json:"statusV2"`
				Actions struct {
					Reschedule struct {
						Doable               bool        `json:"doable"`
						Enabled              bool        `json:"enabled"`
						Token                interface{} `json:"token"`
						PartialActionWarning bool        `json:"partialActionWarning"`
						Typename             string      `json:"__typename"`
					} `json:"reschedule"`
					Typename string `json:"__typename"`
				} `json:"actions"`
				Typename string `json:"__typename"`
			} `json:"deliveryMethods"`
			Typename string `json:"__typename"`
		} `json:"order"`
	} `json:"data"`
}

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
	OrderNumber   string `json:"orderNumber"`
	ReceiptNumber string `json:"receiptNumber"`
	LiteID        string `json:"liteId"`
}

const query = "query orderLookup($orderNumber: String!, $liteId: String) {\n  order(orderNumber: $orderNumber, liteId: $liteId) {\n    status\n    deliveryMethods {\n      status: status2\n      deliveryDate {\n        ...DeliveryDate\n        __typename\n      }\n      type\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DeliveryDate on DeliveryDate {\n  actual {\n    ...DateAndTime\n    __typename\n  }\n  estimatedFrom {\n    ...DateAndTime\n    __typename\n  }\n  estimatedTo {\n    ...DateAndTime\n    __typename\n  }\n  dateTimeRange\n  timeZone\n  __typename\n}\n\nfragment DateAndTime on DateAndTime {\n  time\n  date\n  formattedLocal\n  formattedShortDate\n  formattedLongDate\n  formattedShortDateTime\n  formattedLongDateTime\n  __typename\n}"
const queryV2 = "query DeliveryMethodsOrder($orderNumber: String!, $liteId: String) {\n  order(orderNumber: $orderNumber, liteId: $liteId) {\n    id\n    deliveryMethods {\n      ...DeliveryMethod\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DeliveryMethod on DeliveryMethod {\n  id\n  deliveryNumber\n  serviceId\n  type\n  carrier\n  trackingNumber\n  status: status2\n  pastStatuses\n  trackingLinks\n  partnerTrackingLink\n  quantity\n  packageDimensions {\n    width\n    weightUnit\n    weight\n    volumeUnit\n    volume\n    length\n    height\n    dimensionUnit\n    __typename\n  }\n  deliveryDate {\n    ...DeliveryDate\n    __typename\n  }\n  deliveryAddress {\n    ...Address\n    __typename\n  }\n  articles {\n    ... on AnyDirectionProducts {\n      quantity\n      direction\n      any {\n        ...Article\n        __typename\n      }\n      __typename\n    }\n    ... on ExchangeProducts {\n      inbound {\n        ...Article\n        __typename\n      }\n      outbound {\n        ...Article\n        __typename\n      }\n      inboundQuantity\n      outboundQuantity\n      __typename\n    }\n    __typename\n  }\n  openingHours {\n    dayOfWeek\n    span\n    __typename\n  }\n  isDeviated\n  statusV2 {\n    deliveryStatusV2\n    tense\n    __typename\n  }\n  actions {\n    reschedule {\n      doable\n      enabled\n      token\n      partialActionWarning\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment DeliveryDate on DeliveryDate {\n  actual {\n    ...DateAndTime\n    __typename\n  }\n  estimatedFrom {\n    ...DateAndTime\n    __typename\n  }\n  estimatedTo {\n    ...DateAndTime\n    __typename\n  }\n  dateTimeRange\n  timeZone\n  __typename\n}\n\nfragment DateAndTime on DateAndTime {\n  time\n  date\n  formattedLocal\n  formattedShortDate\n  formattedLongDate\n  formattedShortDateTime\n  formattedLongDateTime\n  __typename\n}\n\nfragment Address on Address {\n  companyName\n  addressLine1\n  addressLine2\n  addressLine3\n  addressLine4\n  postcode\n  city\n  country\n  streetName\n  __typename\n}\n\nfragment Article on Product {\n  name\n  description\n  href\n  quantity\n  decimalQuantity\n  priceUnitText\n  id\n  splitDelivery\n  image {\n    small\n    medium\n    large\n    __typename\n  }\n  referencedItems {\n    ...ReferencedArticle\n    __typename\n  }\n  __typename\n}\n\nfragment ReferencedArticle on Product {\n  name\n  description\n  href\n  quantity\n  decimalQuantity\n  priceUnitText\n  id\n  splitDelivery\n  image {\n    small\n    medium\n    large\n    __typename\n  }\n  __typename\n}"
