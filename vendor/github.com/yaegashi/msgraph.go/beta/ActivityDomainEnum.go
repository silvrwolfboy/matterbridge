// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ActivityDomain undocumented
type ActivityDomain int

const (
	// ActivityDomainVUnknown undocumented
	ActivityDomainVUnknown ActivityDomain = 0
	// ActivityDomainVWork undocumented
	ActivityDomainVWork ActivityDomain = 1
	// ActivityDomainVPersonal undocumented
	ActivityDomainVPersonal ActivityDomain = 2
	// ActivityDomainVUnrestricted undocumented
	ActivityDomainVUnrestricted ActivityDomain = 3
)

// ActivityDomainPUnknown returns a pointer to ActivityDomainVUnknown
func ActivityDomainPUnknown() *ActivityDomain {
	v := ActivityDomainVUnknown
	return &v
}

// ActivityDomainPWork returns a pointer to ActivityDomainVWork
func ActivityDomainPWork() *ActivityDomain {
	v := ActivityDomainVWork
	return &v
}

// ActivityDomainPPersonal returns a pointer to ActivityDomainVPersonal
func ActivityDomainPPersonal() *ActivityDomain {
	v := ActivityDomainVPersonal
	return &v
}

// ActivityDomainPUnrestricted returns a pointer to ActivityDomainVUnrestricted
func ActivityDomainPUnrestricted() *ActivityDomain {
	v := ActivityDomainVUnrestricted
	return &v
}
