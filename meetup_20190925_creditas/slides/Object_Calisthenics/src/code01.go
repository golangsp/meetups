import "strings"

//show code1 OMIT
func hasPermission(roles []Role, statusFlows []StatusFlow, currentStatus, newStatus string) bool {
	for _, role := range roles {
		for _, statusFlow := range statusFlows {
			if strings.ToLower(role.Name) == strings.ToLower(statusFlow.Role) {
				if len(statusFlow.Flows) == 0 {
					return true
				}
				for _, permission := range statusFlow.Flows {
					if permission.From == currentStatus {
						for _, newStatuses := range permission.To {
							if newStatuses == newStatus {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

//end show code1 OMIT

//show code2 OMIT
func (u *User) HasPermission(fromStatus, toStatus string) bool {
	for _, role := range u.roles {
		return role.hasPermission(fromStatus, toStatus)
	}
	return false
}

//end show code2 OMIT

//show code3 OMIT
func (r *Role) hasPermission(fromStatus, toStatus string) bool {
	flows := Flows()
	flow := matchedFlow(flows, r.Name)
	if flow != nil {
		return true
	}
	return flow.IsTransitionAllowed(fromStatus, toStatus)
}

//end show code3 OMIT
//show code4 OMIT
func matchedFlow(statusFlows []StatusFlow, name string) *StatusFlow {
	for _, statusFlow := range statusFlows {
		if strings.EqualFold(name, statusFlow.Role) {
			return &statusFlow
		}
	}
	return nil
}

//end show code4 OMIT
//show code5 OMIT
func (f *StatusFlow) IsTransitionAllowed(from, to string) bool {
	if len(f.Flows) == 0 {
		return true
	}
	permission := f.PermissionForTransition(from)
	for _, new := range permission.To {
		if new == to {
			return true
		}
	}
	return false
}

//end show code5 OMIT
//show code6 OMIT
func (f *StatusFlow) PermissionForTransition(name string) *Permission {
	for _, permission := range f.Flows {
		if permission.From == name {
			return &permission
		}
	}
	return nil
}

//end show code6 OMIT