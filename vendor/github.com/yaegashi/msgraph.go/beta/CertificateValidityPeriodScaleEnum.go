// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CertificateValidityPeriodScale undocumented
type CertificateValidityPeriodScale int

const (
	// CertificateValidityPeriodScaleVDays undocumented
	CertificateValidityPeriodScaleVDays CertificateValidityPeriodScale = 0
	// CertificateValidityPeriodScaleVMonths undocumented
	CertificateValidityPeriodScaleVMonths CertificateValidityPeriodScale = 1
	// CertificateValidityPeriodScaleVYears undocumented
	CertificateValidityPeriodScaleVYears CertificateValidityPeriodScale = 2
)

// CertificateValidityPeriodScalePDays returns a pointer to CertificateValidityPeriodScaleVDays
func CertificateValidityPeriodScalePDays() *CertificateValidityPeriodScale {
	v := CertificateValidityPeriodScaleVDays
	return &v
}

// CertificateValidityPeriodScalePMonths returns a pointer to CertificateValidityPeriodScaleVMonths
func CertificateValidityPeriodScalePMonths() *CertificateValidityPeriodScale {
	v := CertificateValidityPeriodScaleVMonths
	return &v
}

// CertificateValidityPeriodScalePYears returns a pointer to CertificateValidityPeriodScaleVYears
func CertificateValidityPeriodScalePYears() *CertificateValidityPeriodScale {
	v := CertificateValidityPeriodScaleVYears
	return &v
}
