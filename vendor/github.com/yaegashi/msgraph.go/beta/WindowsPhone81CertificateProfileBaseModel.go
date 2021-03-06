// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPhone81CertificateProfileBase Base Windows Phone 8.1+ certificate profile.
type WindowsPhone81CertificateProfileBase struct {
	// DeviceConfiguration is the base model of WindowsPhone81CertificateProfileBase
	DeviceConfiguration
	// RenewalThresholdPercentage Certificate renewal threshold percentage.
	RenewalThresholdPercentage *int `json:"renewalThresholdPercentage,omitempty"`
	// KeyStorageProvider Key Storage Provider (KSP).
	KeyStorageProvider *KeyStorageProviderOption `json:"keyStorageProvider,omitempty"`
	// SubjectNameFormat Certificate Subject Name Format.
	SubjectNameFormat *SubjectNameFormat `json:"subjectNameFormat,omitempty"`
	// SubjectAlternativeNameType Certificate Subject Alternative Name Type.
	SubjectAlternativeNameType *SubjectAlternativeNameType `json:"subjectAlternativeNameType,omitempty"`
	// CertificateValidityPeriodValue Value for the Certificate Validtiy Period.
	CertificateValidityPeriodValue *int `json:"certificateValidityPeriodValue,omitempty"`
	// CertificateValidityPeriodScale Scale for the Certificate Validity Period.
	CertificateValidityPeriodScale *CertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	// ExtendedKeyUsages Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
	ExtendedKeyUsages []ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`
}
