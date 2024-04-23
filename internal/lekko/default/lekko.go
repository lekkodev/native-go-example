package lekkodefault

// Controls behavior for access to a protected endpoint
func getCanAccess(endpoint string, isUserAdmin bool) bool {
	if isUserAdmin == true {
		return true
	} else if endpoint == "/protected" {
		return false
	}
	return true
}
