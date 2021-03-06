// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SalesInvoice undocumented
type SalesInvoice struct {
	// Entity is the base model of SalesInvoice
	Entity
	// Number undocumented
	Number *string `json:"number,omitempty"`
	// ExternalDocumentNumber undocumented
	ExternalDocumentNumber *string `json:"externalDocumentNumber,omitempty"`
	// InvoiceDate undocumented
	InvoiceDate *time.Time `json:"invoiceDate,omitempty"`
	// DueDate undocumented
	DueDate *time.Time `json:"dueDate,omitempty"`
	// CustomerPurchaseOrderReference undocumented
	CustomerPurchaseOrderReference *string `json:"customerPurchaseOrderReference,omitempty"`
	// CustomerID undocumented
	CustomerID *UUID `json:"customerId,omitempty"`
	// CustomerNumber undocumented
	CustomerNumber *string `json:"customerNumber,omitempty"`
	// CustomerName undocumented
	CustomerName *string `json:"customerName,omitempty"`
	// BillToName undocumented
	BillToName *string `json:"billToName,omitempty"`
	// BillToCustomerID undocumented
	BillToCustomerID *UUID `json:"billToCustomerId,omitempty"`
	// BillToCustomerNumber undocumented
	BillToCustomerNumber *string `json:"billToCustomerNumber,omitempty"`
	// ShipToName undocumented
	ShipToName *string `json:"shipToName,omitempty"`
	// ShipToContact undocumented
	ShipToContact *string `json:"shipToContact,omitempty"`
	// SellingPostalAddress undocumented
	SellingPostalAddress *PostalAddressType `json:"sellingPostalAddress,omitempty"`
	// BillingPostalAddress undocumented
	BillingPostalAddress *PostalAddressType `json:"billingPostalAddress,omitempty"`
	// ShippingPostalAddress undocumented
	ShippingPostalAddress *PostalAddressType `json:"shippingPostalAddress,omitempty"`
	// CurrencyID undocumented
	CurrencyID *UUID `json:"currencyId,omitempty"`
	// CurrencyCode undocumented
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// OrderID undocumented
	OrderID *UUID `json:"orderId,omitempty"`
	// OrderNumber undocumented
	OrderNumber *string `json:"orderNumber,omitempty"`
	// PaymentTermsID undocumented
	PaymentTermsID *UUID `json:"paymentTermsId,omitempty"`
	// ShipmentMethodID undocumented
	ShipmentMethodID *UUID `json:"shipmentMethodId,omitempty"`
	// Salesperson undocumented
	Salesperson *string `json:"salesperson,omitempty"`
	// PricesIncludeTax undocumented
	PricesIncludeTax *bool `json:"pricesIncludeTax,omitempty"`
	// DiscountAmount undocumented
	DiscountAmount *int `json:"discountAmount,omitempty"`
	// DiscountAppliedBeforeTax undocumented
	DiscountAppliedBeforeTax *bool `json:"discountAppliedBeforeTax,omitempty"`
	// TotalAmountExcludingTax undocumented
	TotalAmountExcludingTax *int `json:"totalAmountExcludingTax,omitempty"`
	// TotalTaxAmount undocumented
	TotalTaxAmount *int `json:"totalTaxAmount,omitempty"`
	// TotalAmountIncludingTax undocumented
	TotalAmountIncludingTax *int `json:"totalAmountIncludingTax,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// PhoneNumber undocumented
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// SalesInvoiceLines undocumented
	SalesInvoiceLines []SalesInvoiceLine `json:"salesInvoiceLines,omitempty"`
	// Customer undocumented
	Customer *Customer `json:"customer,omitempty"`
	// Currency undocumented
	Currency *Currency `json:"currency,omitempty"`
	// PaymentTerm undocumented
	PaymentTerm *PaymentTerm `json:"paymentTerm,omitempty"`
	// ShipmentMethod undocumented
	ShipmentMethod *ShipmentMethod `json:"shipmentMethod,omitempty"`
}
