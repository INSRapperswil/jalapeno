package arango

type VRF struct {
	Hash                          string `json:"hash"` // hash of kafka message to verify if changes need to be written into arango
	AfAfName                      string `json:"af/af_name"`
	AfRouteTargetAfName           string `json:"af/route_target/af_name"`
	AfRouteTargetRouteTargetType  string `json:"af/route_target/route_target_type"`
	AfRouteTargetRouteTargetValue string `json:"af/route_target/route_target_value"`
	AfRouteTargetSafName          string `json:"af/route_target/saf_name"`
	AfSafName                     string `json:"af/saf_name"`
	InterfaceInterfaceName        string `json:"interface/interface_name"`
	IsBigVrf                      string `json:"is_big_vrf"`
	RouteDistinguisher            string `json:"route_distinguisher"`
	VrfNameXr                     string `json:"vrf_name_xr"`
}
