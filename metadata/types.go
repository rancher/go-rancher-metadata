package metadata

const (
	VERSION1 = "2015-07-25"
	VERSION2 = "2015-12-19"
	VERSION3 = "2016-07-29"
)

/*
LATEST
*/
type Stack struct {
	StackBase
	Services []Service `json:"services"`
}

type Service struct {
	ServiceBase
	Containers []Container `json:"containers"`
}

type Container struct {
	Name                     string            `json:"name"`
	UUID                     string            `json:"uuid"`
	PrimaryIp                string            `json:"primary_ip"`
	Ips                      []string          `json:"ips"`
	Ports                    []string          `json:"ports"`
	ServiceName              string            `json:"service_name"`
	StackName                string            `json:"stack_name"`
	StackUUID                string            `json:"stack_uuid"`
	ServiceUUID              string            `json:"service_uuid"`
	Labels                   map[string]string `json:"labels"`
	CreateIndex              int               `json:"create_index"`
	HostUUID                 string            `json:"host_uuid"`
	Hostname                 string            `json:"hostname"`
	HealthState              string            `json:"health_state"`
	StartCount               int               `json:"start_count"`
	ServiceIndex             string            `json:"service_index"`
	State                    string            `json:"state"`
	ExternalId               string            `json:"external_id"`
	PrimaryMacAddress        string            `json:"primary_mac_address"`
	MemoryReservation        int64             `json:"memory_reservation"`
	MilliCPUReservation      int64             `json:"milli_cpu_reservation"`
	NetworkUUID              string            `json:"network_uuid"`
	NetworkFromContainerUUID string            `json:"network_from_container_uuid"`
	System                   bool              `json:"system"`
	Dns                      []string          `json:"dns"`
	DnsSearch                []string          `json:"dns_search"`
	HealthCheckHosts         []string          `json:"health_check_hosts"`
	EnvironmentUUID          string            `json:"environment_uuid"`
	Links                    map[string]string `json:"links"`
}

type Network struct {
	Name      string                 `json:"name"`
	UUID      string                 `json:"uuid"`
	Metadata  map[string]interface{} `json:"metadata"`
	HostPorts bool                   `json:"host_ports"`
	Default   bool                   `json:"is_default"`
}

type HealthCheck struct {
	HealthyThreshold   int    `json:"healthy_threshold"`
	Interval           int    `json:"interval"`
	Port               int    `json:"port"`
	RequestLine        string `json:"request_line"`
	ResponseTimeout    int    `json:"response_timeout"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
}

type Host struct {
	Name            string            `json:"name"`
	AgentIP         string            `json:"agent_ip"`
	HostId          int               `json:"host_id"`
	Labels          map[string]string `json:"labels"`
	UUID            string            `json:"uuid"`
	Hostname        string            `json:"hostname"`
	Memory          int64             `json:"memory"`
	MilliCPU        int64             `json:"milli_cpu"`
	LocalStorageMb  int64             `json:"local_storage_mb"`
	EnvironmentUUID string            `json:"environment_uuid"`
}

type PortRule struct {
	SourcePort  int    `json:"source_port"`
	Protocol    string `json:"protocol"`
	Path        string `json:"path"`
	Hostname    string `json:"hostname"`
	Service     string `json:"service"`
	TargetPort  int    `json:"target_port"`
	Priority    int    `json:"priority"`
	BackendName string `json:"backend_name"`
	Selector    string `json:"selector"`
}

type LBConfig struct {
	Certs            []string           `json:"certs"`
	DefaultCert      string             `json:"default_cert"`
	PortRules        []PortRule         `json:"port_rules"`
	Config           string             `json:"config"`
	StickinessPolicy LBStickinessPolicy `json:"stickiness_policy"`
}

type LBStickinessPolicy struct {
	Name     string `json:"name"`
	Cookie   string `json:"cookie"`
	Domain   string `json:"domain"`
	Indirect bool   `json:"indirect"`
	Nocache  bool   `json:"nocache"`
	Postonly bool   `json:"postonly"`
	Mode     string `json:"mode"`
}

/*
VERSION 1
*/

type StackVersion1 struct {
	StackBase
	Services []string `json:"services"`
}

type ServiceVersion1 struct {
	ServiceBase
	Containers []string `json:"containers"`
}

/*
COMMON BASE
*/
type StackBase struct {
	EnvironmentName string `json:"environment_name"`
	EnvironmentUUID string `json:"environment_uuid"`
	Name            string `json:"name"`
	UUID            string `json:"uuid"`
	System          bool   `json:"system"`
}

type ServiceBase struct {
	Name               string                 `json:"name"`
	UUID               string                 `json:"uuid"`
	StackName          string                 `json:"stack_name"`
	StackUUID          string                 `json:"stack_uuid"`
	Kind               string                 `json:"kind"`
	Hostname           string                 `json:"hostname"`
	Vip                string                 `json:"vip"`
	CreateIndex        int                    `json:"create_index"`
	ExternalIps        []string               `json:"external_ips"`
	Sidekicks          []string               `json:"sidekicks"`
	Links              map[string]string      `json:"links"`
	Ports              []string               `json:"ports"`
	Labels             map[string]string      `json:"labels"`
	Metadata           map[string]interface{} `json:"metadata"`
	Scale              int                    `json:"scale"`
	Fqdn               string                 `json:"fqdn"`
	Expose             []string               `json:"expose"`
	HealthCheck        HealthCheck            `json:"health_check"`
	System             bool                   `json:"system"`
	LBConfig           LBConfig               `json:"lb_config"`
	PrimaryServiceName string                 `json:"primary_service_name"`
	EnvironmentUUID    string                 `json:"environment_uuid"`
	Token              string                 `json:"token"`
}
