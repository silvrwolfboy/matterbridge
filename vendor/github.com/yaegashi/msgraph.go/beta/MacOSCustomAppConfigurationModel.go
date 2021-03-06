// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOSCustomAppConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the macOSCustomAppConfiguration resource.
type MacOSCustomAppConfiguration struct {
	// DeviceConfiguration is the base model of MacOSCustomAppConfiguration
	DeviceConfiguration
	// BundleID Bundle id for targeting.
	BundleID *string `json:"bundleId,omitempty"`
	// FileName Configuration file name (*.plist | *.xml).
	FileName *string `json:"fileName,omitempty"`
	// ConfigurationXML Configuration xml. (UTF8 encoded byte array)
	ConfigurationXML *Binary `json:"configurationXml,omitempty"`
}
