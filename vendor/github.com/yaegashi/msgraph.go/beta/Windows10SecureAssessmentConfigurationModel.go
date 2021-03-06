// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows10SecureAssessmentConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the secureAssessment resource.
type Windows10SecureAssessmentConfiguration struct {
	// DeviceConfiguration is the base model of Windows10SecureAssessmentConfiguration
	DeviceConfiguration
	// LaunchURI Url link to an assessment that's automatically loaded when the secure assessment browser is launched. It has to be a valid Url (http[s]://msdn.microsoft.com/).
	LaunchURI *string `json:"launchUri,omitempty"`
	// ConfigurationAccount The account used to configure the Windows device for taking the test. The user can be a domain account (domain\user), an AAD account (username@tenant.com) or a local account (username).
	ConfigurationAccount *string `json:"configurationAccount,omitempty"`
	// ConfigurationAccountType The account type used to by ConfigurationAccount.
	ConfigurationAccountType *SecureAssessmentAccountType `json:"configurationAccountType,omitempty"`
	// AllowPrinting Indicates whether or not to allow the app from printing during the test.
	AllowPrinting *bool `json:"allowPrinting,omitempty"`
	// AllowScreenCapture Indicates whether or not to allow screen capture capability during a test.
	AllowScreenCapture *bool `json:"allowScreenCapture,omitempty"`
	// AllowTextSuggestion Indicates whether or not to allow text suggestions during the test.
	AllowTextSuggestion *bool `json:"allowTextSuggestion,omitempty"`
	// LocalGuestAccountName Specifies the display text for the local guest account shown on the sign-in screen. Typically is the name of an assessment. When the user clicks the local guest account on the sign-in screen, an assessment app is launched with a specified assessment URL. Secure assessments can only be configured with local guest account sign-in on devices running Windows 10, version 1903 or later. Important notice: this property must be set with assessmentAppUserModelID in order to make the local guest account sign-in experience work properly for secure assessments.
	LocalGuestAccountName *string `json:"localGuestAccountName,omitempty"`
	// AssessmentAppUserModelID Specifies the application user model ID of the assessment app launched when a user signs in to a secure assessment with a local guest account. Important notice: this property must be set with localGuestAccountName in order to make the local guest account sign-in experience work properly for secure assessments.
	AssessmentAppUserModelID *string `json:"assessmentAppUserModelId,omitempty"`
}
