// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsFirewallRuleTrafficDirectionType undocumented
type WindowsFirewallRuleTrafficDirectionType int

const (
	// WindowsFirewallRuleTrafficDirectionTypeVNotConfigured undocumented
	WindowsFirewallRuleTrafficDirectionTypeVNotConfigured WindowsFirewallRuleTrafficDirectionType = 0
	// WindowsFirewallRuleTrafficDirectionTypeVOut undocumented
	WindowsFirewallRuleTrafficDirectionTypeVOut WindowsFirewallRuleTrafficDirectionType = 1
	// WindowsFirewallRuleTrafficDirectionTypeVIn undocumented
	WindowsFirewallRuleTrafficDirectionTypeVIn WindowsFirewallRuleTrafficDirectionType = 2
)

// WindowsFirewallRuleTrafficDirectionTypePNotConfigured returns a pointer to WindowsFirewallRuleTrafficDirectionTypeVNotConfigured
func WindowsFirewallRuleTrafficDirectionTypePNotConfigured() *WindowsFirewallRuleTrafficDirectionType {
	v := WindowsFirewallRuleTrafficDirectionTypeVNotConfigured
	return &v
}

// WindowsFirewallRuleTrafficDirectionTypePOut returns a pointer to WindowsFirewallRuleTrafficDirectionTypeVOut
func WindowsFirewallRuleTrafficDirectionTypePOut() *WindowsFirewallRuleTrafficDirectionType {
	v := WindowsFirewallRuleTrafficDirectionTypeVOut
	return &v
}

// WindowsFirewallRuleTrafficDirectionTypePIn returns a pointer to WindowsFirewallRuleTrafficDirectionTypeVIn
func WindowsFirewallRuleTrafficDirectionTypePIn() *WindowsFirewallRuleTrafficDirectionType {
	v := WindowsFirewallRuleTrafficDirectionTypeVIn
	return &v
}
