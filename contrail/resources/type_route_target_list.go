//
// Automatically generated. DO NOT EDIT.
//

package resources

type RouteTargetList struct {
	RouteTarget []string `json:"route_target,omitempty"`
}

func (obj *RouteTargetList) AddRouteTarget(value string) {
	obj.RouteTarget = append(obj.RouteTarget, value)
}
