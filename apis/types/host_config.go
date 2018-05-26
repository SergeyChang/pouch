// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostConfig Container configuration that depends on the host we are running on
// swagger:model HostConfig

type HostConfig struct {

	// Automatically remove the container when the container's process exits. This has no effect if `RestartPolicy` is set.
	AutoRemove bool `json:"AutoRemove,omitempty"`

	// A list of volume bindings for this container. Each volume binding is a string in one of these forms:
	//
	// - `host-src:container-dest` to bind-mount a host path into the container. Both `host-src`, and `container-dest` must be an _absolute_ path.
	// - `host-src:container-dest:ro` to make the bind mount read-only inside the container. Both `host-src`, and `container-dest` must be an _absolute_ path.
	// - `volume-name:container-dest` to bind-mount a volume managed by a volume driver into the container. `container-dest` must be an _absolute_ path.
	// - `volume-name:container-dest:ro` to mount the volume read-only inside the container.  `container-dest` must be an _absolute_ path.
	//
	Binds []string `json:"Binds"`

	// A list of kernel capabilities to add to the container.
	CapAdd []string `json:"CapAdd"`

	// A list of kernel capabilities to drop from the container.
	CapDrop []string `json:"CapDrop"`

	// Cgroup to use for the container.
	Cgroup string `json:"Cgroup,omitempty"`

	// Initial console size, as an `[height, width]` array. (Windows only)
	// Max Items: 2
	// Min Items: 2
	ConsoleSize []*int64 `json:"ConsoleSize"`

	// Path to a file where the container ID is written
	ContainerIDFile string `json:"ContainerIDFile,omitempty"`

	// A list of DNS servers for the container to use.
	DNS []string `json:"Dns"`

	// A list of DNS options.
	DNSOptions []string `json:"DnsOptions"`

	// A list of DNS search domains.
	DNSSearch []string `json:"DnsSearch"`

	// Whether to enable lxcfs.
	EnableLxcfs bool `json:"EnableLxcfs,omitempty"`

	// A list of hostnames/IP mappings to add to the container's `/etc/hosts` file. Specified in the form `["hostname:IP"]`.
	//
	ExtraHosts []string `json:"ExtraHosts"`

	// A list of additional groups that the container process will run as.
	GroupAdd []string `json:"GroupAdd"`

	// Initial script executed in container. The script will be executed before entrypoint or command
	InitScript string `json:"InitScript,omitempty"`

	// IPC sharing mode for the container. Possible values are:
	// - `"none"`: own private IPC namespace, with /dev/shm not mounted
	// - `"private"`: own private IPC namespace
	// - `"shareable"`: own private IPC namespace, with a possibility to share it with other containers
	// - `"container:<name|id>"`: join another (shareable) container's IPC namespace
	// - `"host"`: use the host system's IPC namespace
	// If not specified, daemon default is used, which can either be `"private"`
	// or `"shareable"`, depending on daemon version and configuration.
	//
	IpcMode string `json:"IpcMode,omitempty"`

	// Isolation technology of the container. (Windows only)
	Isolation string `json:"Isolation,omitempty"`

	// A list of links for the container in the form `container_name:alias`.
	Links []string `json:"Links"`

	// The logging configuration for this container
	LogConfig *LogConfig `json:"LogConfig,omitempty"`

	// Network mode to use for this container. Supported standard values are: `bridge`, `host`, `none`, and `container:<name|id>`. Any other value is taken as a custom network's name to which this container should connect to.
	NetworkMode string `json:"NetworkMode,omitempty"`

	// An integer value containing the score given to the container in order to tune OOM killer preferences.
	// The range is in [-1000, 1000].
	//
	// Maximum: 1000
	// Minimum: -1000
	OomScoreAdj int64 `json:"OomScoreAdj,omitempty"`

	// Set the PID (Process) Namespace mode for the container. It can be either:
	// - `"container:<name|id>"`: joins another container's PID namespace
	// - `"host"`: use the host's PID namespace inside the container
	//
	PidMode string `json:"PidMode,omitempty"`

	// A map of exposed container ports and the host port they should map to.
	PortBindings PortMap `json:"PortBindings,omitempty"`

	// Gives the container full access to the host.
	Privileged bool `json:"Privileged,omitempty"`

	// Allocates a random host port for all of a container's exposed ports.
	PublishAllPorts bool `json:"PublishAllPorts,omitempty"`

	// Mount the container's root filesystem as read only.
	ReadonlyRootfs bool `json:"ReadonlyRootfs,omitempty"`

	// Restart policy to be used to manage the container
	RestartPolicy *RestartPolicy `json:"RestartPolicy,omitempty"`

	// Whether to start container in rich container mode. (default false)
	Rich bool `json:"Rich,omitempty"`

	// Choose one rich container mode.(default dumb-init)
	RichMode string `json:"RichMode,omitempty"`

	// Runtime to use with this container.
	Runtime string `json:"Runtime,omitempty"`

	// A list of string values to customize labels for MLS systems, such as SELinux.
	SecurityOpt []string `json:"SecurityOpt"`

	// Size of `/dev/shm` in bytes. If omitted, the system uses 64MB.
	// Minimum: 0
	ShmSize *int64 `json:"ShmSize,omitempty"`

	// Storage driver options for this container, in the form `{"size": "120G"}`.
	//
	StorageOpt map[string]string `json:"StorageOpt,omitempty"`

	// A list of kernel parameters (sysctls) to set in the container. For example: `{"net.ipv4.ip_forward": "1"}`
	//
	Sysctls map[string]string `json:"Sysctls,omitempty"`

	// A map of container directories which should be replaced by tmpfs mounts, and their corresponding mount options. For example: `{ "/run": "rw,noexec,nosuid,size=65536k" }`.
	//
	Tmpfs map[string]string `json:"Tmpfs,omitempty"`

	// UTS namespace to use for the container.
	UTSMode string `json:"UTSMode,omitempty"`

	// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	UsernsMode string `json:"UsernsMode,omitempty"`

	// Driver that this container uses to mount volumes.
	VolumeDriver string `json:"VolumeDriver,omitempty"`

	// A list of volumes to inherit from another container, specified in the form `<container name>[:<ro|rw>]`.
	VolumesFrom []string `json:"VolumesFrom"`

	Resources
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HostConfig) UnmarshalJSON(raw []byte) error {

	var data struct {
		AutoRemove bool `json:"AutoRemove,omitempty"`

		Binds []string `json:"Binds,omitempty"`

		CapAdd []string `json:"CapAdd,omitempty"`

		CapDrop []string `json:"CapDrop,omitempty"`

		Cgroup string `json:"Cgroup,omitempty"`

		ConsoleSize []*int64 `json:"ConsoleSize,omitempty"`

		ContainerIDFile string `json:"ContainerIDFile,omitempty"`

		DNS []string `json:"Dns,omitempty"`

		DNSOptions []string `json:"DnsOptions,omitempty"`

		DNSSearch []string `json:"DnsSearch,omitempty"`

		EnableLxcfs bool `json:"EnableLxcfs,omitempty"`

		ExtraHosts []string `json:"ExtraHosts,omitempty"`

		GroupAdd []string `json:"GroupAdd,omitempty"`

		InitScript string `json:"InitScript,omitempty"`

		IpcMode string `json:"IpcMode,omitempty"`

		Isolation string `json:"Isolation,omitempty"`

		Links []string `json:"Links,omitempty"`

		LogConfig *LogConfig `json:"LogConfig,omitempty"`

		NetworkMode string `json:"NetworkMode,omitempty"`

		OomScoreAdj int64 `json:"OomScoreAdj,omitempty"`

		PidMode string `json:"PidMode,omitempty"`

		PortBindings PortMap `json:"PortBindings,omitempty"`

		Privileged bool `json:"Privileged,omitempty"`

		PublishAllPorts bool `json:"PublishAllPorts,omitempty"`

		ReadonlyRootfs bool `json:"ReadonlyRootfs,omitempty"`

		RestartPolicy *RestartPolicy `json:"RestartPolicy,omitempty"`

		Rich bool `json:"Rich,omitempty"`

		RichMode string `json:"RichMode,omitempty"`

		Runtime string `json:"Runtime,omitempty"`

		SecurityOpt []string `json:"SecurityOpt,omitempty"`

		ShmSize *int64 `json:"ShmSize,omitempty"`

		StorageOpt map[string]string `json:"StorageOpt,omitempty"`

		Sysctls map[string]string `json:"Sysctls,omitempty"`

		Tmpfs map[string]string `json:"Tmpfs,omitempty"`

		UTSMode string `json:"UTSMode,omitempty"`

		UsernsMode string `json:"UsernsMode,omitempty"`

		VolumeDriver string `json:"VolumeDriver,omitempty"`

		VolumesFrom []string `json:"VolumesFrom,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.AutoRemove = data.AutoRemove

	m.Binds = data.Binds

	m.CapAdd = data.CapAdd

	m.CapDrop = data.CapDrop

	m.Cgroup = data.Cgroup

	m.ConsoleSize = data.ConsoleSize

	m.ContainerIDFile = data.ContainerIDFile

	m.DNS = data.DNS

	m.DNSOptions = data.DNSOptions

	m.DNSSearch = data.DNSSearch

	m.EnableLxcfs = data.EnableLxcfs

	m.ExtraHosts = data.ExtraHosts

	m.GroupAdd = data.GroupAdd

	m.InitScript = data.InitScript

	m.IpcMode = data.IpcMode

	m.Isolation = data.Isolation

	m.Links = data.Links

	m.LogConfig = data.LogConfig

	m.NetworkMode = data.NetworkMode

	m.OomScoreAdj = data.OomScoreAdj

	m.PidMode = data.PidMode

	m.PortBindings = data.PortBindings

	m.Privileged = data.Privileged

	m.PublishAllPorts = data.PublishAllPorts

	m.ReadonlyRootfs = data.ReadonlyRootfs

	m.RestartPolicy = data.RestartPolicy

	m.Rich = data.Rich

	m.RichMode = data.RichMode

	m.Runtime = data.Runtime

	m.SecurityOpt = data.SecurityOpt

	m.ShmSize = data.ShmSize

	m.StorageOpt = data.StorageOpt

	m.Sysctls = data.Sysctls

	m.Tmpfs = data.Tmpfs

	m.UTSMode = data.UTSMode

	m.UsernsMode = data.UsernsMode

	m.VolumeDriver = data.VolumeDriver

	m.VolumesFrom = data.VolumesFrom

	var aO1 Resources
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Resources = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HostConfig) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	var data struct {
		AutoRemove bool `json:"AutoRemove,omitempty"`

		Binds []string `json:"Binds,omitempty"`

		CapAdd []string `json:"CapAdd,omitempty"`

		CapDrop []string `json:"CapDrop,omitempty"`

		Cgroup string `json:"Cgroup,omitempty"`

		ConsoleSize []*int64 `json:"ConsoleSize,omitempty"`

		ContainerIDFile string `json:"ContainerIDFile,omitempty"`

		DNS []string `json:"Dns,omitempty"`

		DNSOptions []string `json:"DnsOptions,omitempty"`

		DNSSearch []string `json:"DnsSearch,omitempty"`

		EnableLxcfs bool `json:"EnableLxcfs,omitempty"`

		ExtraHosts []string `json:"ExtraHosts,omitempty"`

		GroupAdd []string `json:"GroupAdd,omitempty"`

		InitScript string `json:"InitScript,omitempty"`

		IpcMode string `json:"IpcMode,omitempty"`

		Isolation string `json:"Isolation,omitempty"`

		Links []string `json:"Links,omitempty"`

		LogConfig *LogConfig `json:"LogConfig,omitempty"`

		NetworkMode string `json:"NetworkMode,omitempty"`

		OomScoreAdj int64 `json:"OomScoreAdj,omitempty"`

		PidMode string `json:"PidMode,omitempty"`

		PortBindings PortMap `json:"PortBindings,omitempty"`

		Privileged bool `json:"Privileged,omitempty"`

		PublishAllPorts bool `json:"PublishAllPorts,omitempty"`

		ReadonlyRootfs bool `json:"ReadonlyRootfs,omitempty"`

		RestartPolicy *RestartPolicy `json:"RestartPolicy,omitempty"`

		Rich bool `json:"Rich,omitempty"`

		RichMode string `json:"RichMode,omitempty"`

		Runtime string `json:"Runtime,omitempty"`

		SecurityOpt []string `json:"SecurityOpt,omitempty"`

		ShmSize *int64 `json:"ShmSize,omitempty"`

		StorageOpt map[string]string `json:"StorageOpt,omitempty"`

		Sysctls map[string]string `json:"Sysctls,omitempty"`

		Tmpfs map[string]string `json:"Tmpfs,omitempty"`

		UTSMode string `json:"UTSMode,omitempty"`

		UsernsMode string `json:"UsernsMode,omitempty"`

		VolumeDriver string `json:"VolumeDriver,omitempty"`

		VolumesFrom []string `json:"VolumesFrom,omitempty"`
	}

	data.AutoRemove = m.AutoRemove

	data.Binds = m.Binds

	data.CapAdd = m.CapAdd

	data.CapDrop = m.CapDrop

	data.Cgroup = m.Cgroup

	data.ConsoleSize = m.ConsoleSize

	data.ContainerIDFile = m.ContainerIDFile

	data.DNS = m.DNS

	data.DNSOptions = m.DNSOptions

	data.DNSSearch = m.DNSSearch

	data.EnableLxcfs = m.EnableLxcfs

	data.ExtraHosts = m.ExtraHosts

	data.GroupAdd = m.GroupAdd

	data.InitScript = m.InitScript

	data.IpcMode = m.IpcMode

	data.Isolation = m.Isolation

	data.Links = m.Links

	data.LogConfig = m.LogConfig

	data.NetworkMode = m.NetworkMode

	data.OomScoreAdj = m.OomScoreAdj

	data.PidMode = m.PidMode

	data.PortBindings = m.PortBindings

	data.Privileged = m.Privileged

	data.PublishAllPorts = m.PublishAllPorts

	data.ReadonlyRootfs = m.ReadonlyRootfs

	data.RestartPolicy = m.RestartPolicy

	data.Rich = m.Rich

	data.RichMode = m.RichMode

	data.Runtime = m.Runtime

	data.SecurityOpt = m.SecurityOpt

	data.ShmSize = m.ShmSize

	data.StorageOpt = m.StorageOpt

	data.Sysctls = m.Sysctls

	data.Tmpfs = m.Tmpfs

	data.UTSMode = m.UTSMode

	data.UsernsMode = m.UsernsMode

	data.VolumeDriver = m.VolumeDriver

	data.VolumesFrom = m.VolumesFrom

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	aO1, err := swag.WriteJSON(m.Resources)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this host config
func (m *HostConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBinds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapDrop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsoleSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsolation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOomScoreAdj(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestartPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRichMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityOpt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShmSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumesFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.Resources.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostConfig) validateBinds(formats strfmt.Registry) error {

	if swag.IsZero(m.Binds) { // not required
		return nil
	}

	return nil
}

func (m *HostConfig) validateCapAdd(formats strfmt.Registry) error {

	if swag.IsZero(m.CapAdd) { // not required
		return nil
	}

	return nil
}

func (m *HostConfig) validateCapDrop(formats strfmt.Registry) error {

	if swag.IsZero(m.CapDrop) { // not required
		return nil
	}

	return nil
}

func (m *HostConfig) validateConsoleSize(formats strfmt.Registry) error {

	if swag.IsZero(m.ConsoleSize) { // not required
		return nil
	}

	iConsoleSizeSize := int64(len(m.ConsoleSize))

	if err := validate.MinItems("ConsoleSize", "body", iConsoleSizeSize, 2); err != nil {
		return err
	}

	if err := validate.MaxItems("ConsoleSize", "body", iConsoleSizeSize, 2); err != nil {
		return err
	}

	for i := 0; i < len(m.ConsoleSize); i++ {

		if swag.IsZero(m.ConsoleSize[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("ConsoleSize"+"."+strconv.Itoa(i), "body", int64(*m.ConsoleSize[i]), 0, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *HostConfig) validateDNS(formats strfmt.Registry) error {

	if swag.IsZero(m.DNS) { // not required
		return nil
	}

	return nil
}

func (m *HostConfig) validateDNSOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSOptions) { // not required
		return nil
	}

	return nil
}

func (m *HostConfig) validateDNSSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSSearch) { // not required
		return nil
	}

	return nil
}

func (m *HostConfig) validateExtraHosts(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtraHosts) { // not required
		return nil
	}

	return nil
}

func (m *HostConfig) validateGroupAdd(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupAdd) { // not required
		return nil
	}

	return nil
}

var hostConfigTypeIsolationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","process","hyperv"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostConfigTypeIsolationPropEnum = append(hostConfigTypeIsolationPropEnum, v)
	}
}

// property enum
func (m *HostConfig) validateIsolationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hostConfigTypeIsolationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HostConfig) validateIsolation(formats strfmt.Registry) error {

	if swag.IsZero(m.Isolation) { // not required
		return nil
	}

	// value enum
	if err := m.validateIsolationEnum("Isolation", "body", m.Isolation); err != nil {
		return err
	}

	return nil
}

func (m *HostConfig) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	return nil
}

func (m *HostConfig) validateLogConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.LogConfig) { // not required
		return nil
	}

	if m.LogConfig != nil {

		if err := m.LogConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LogConfig")
			}
			return err
		}
	}

	return nil
}

func (m *HostConfig) validateOomScoreAdj(formats strfmt.Registry) error {

	if swag.IsZero(m.OomScoreAdj) { // not required
		return nil
	}

	if err := validate.MinimumInt("OomScoreAdj", "body", int64(m.OomScoreAdj), -1000, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("OomScoreAdj", "body", int64(m.OomScoreAdj), 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *HostConfig) validateRestartPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.RestartPolicy) { // not required
		return nil
	}

	if m.RestartPolicy != nil {

		if err := m.RestartPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RestartPolicy")
			}
			return err
		}
	}

	return nil
}

var hostConfigTypeRichModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dumb-init","sbin-init","systemd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostConfigTypeRichModePropEnum = append(hostConfigTypeRichModePropEnum, v)
	}
}

// property enum
func (m *HostConfig) validateRichModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hostConfigTypeRichModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HostConfig) validateRichMode(formats strfmt.Registry) error {

	if swag.IsZero(m.RichMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRichModeEnum("RichMode", "body", m.RichMode); err != nil {
		return err
	}

	return nil
}

func (m *HostConfig) validateSecurityOpt(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityOpt) { // not required
		return nil
	}

	return nil
}

func (m *HostConfig) validateShmSize(formats strfmt.Registry) error {

	if swag.IsZero(m.ShmSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("ShmSize", "body", int64(*m.ShmSize), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *HostConfig) validateVolumesFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumesFrom) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostConfig) UnmarshalBinary(b []byte) error {
	var res HostConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
