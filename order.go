package myflyingbox

// Order is the struct used to place an order.
type Order struct {

	// Used in placing the order.
	OfferID        string  `json:"offer_id,omitempty"`
	ThermalLabels  bool    `json:"thermal_labels,omitempty"`
	InsureShipment bool    `json:"insure_shipment,omitempty"`
	Shipper        Shipper `json:"shipper"`

	Parcels []Parcel `json:"parcels,omitempty"`

	// Used when retrieving the order\
	ID        string    `json:"id,omitempty"`
	Recipient Recipient `json:"recipient,omitempty"`
}
