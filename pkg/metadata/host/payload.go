// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package host

import (
	hostMetadataUtils "github.com/DataDog/datadog-agent/comp/metadata/host/utils"
)

type systemStats struct {
	CPUCores  int32     `json:"cpuCores"`
	Machine   string    `json:"machine"`
	Platform  string    `json:"platform"`
	Pythonv   string    `json:"pythonV"`
	Processor string    `json:"processor"`
	Macver    osVersion `json:"macV"`
	Nixver    osVersion `json:"nixV"`
	Fbsdver   osVersion `json:"fbsdV"`
	Winver    osVersion `json:"winV"`
}

// NetworkMeta is metadata about the host's network
type NetworkMeta struct {
	ID         string `json:"network-id"`
	PublicIPv4 string `json:"public-ipv4,omitempty"`
}

// LogsMeta is metadata about the host's logs agent
type LogsMeta struct {
	Transport            string `json:"transport"`
	AutoMultilineEnabled bool   `json:"auto_multi_line_detection_enabled"`
}

// InstallMethod is metadata about the agent's installation
type InstallMethod struct {
	Tool             *string `json:"tool"`
	ToolVersion      string  `json:"tool_version"`
	InstallerVersion *string `json:"installer_version"`
}

// ProxyMeta is metatdata about the proxy configuration
type ProxyMeta struct {
	NoProxyNonexactMatch              bool `json:"no-proxy-nonexact-match"`
	ProxyBehaviorChanged              bool `json:"proxy-behavior-changed"`
	NoProxyNonexactMatchExplicitlySet bool `json:"no-proxy-nonexact-match-explicitly-set"`
}

// OtlpMeta is metadata about the otlp pipeline
type OtlpMeta struct {
	Enabled bool `json:"enabled"`
}

// Payload handles the JSON unmarshalling of the metadata payload
type Payload struct {
	Os            string                  `json:"os"`
	AgentFlavor   string                  `json:"agent-flavor"`
	PythonVersion string                  `json:"python"`
	SystemStats   *systemStats            `json:"systemStats"`
	Meta          *hostMetadataUtils.Meta `json:"meta"`
	HostTags      *hostMetadataUtils.Tags `json:"host-tags"`
	ContainerMeta map[string]string       `json:"container-meta,omitempty"`
	NetworkMeta   *NetworkMeta            `json:"network"`
	LogsMeta      *LogsMeta               `json:"logs"`
	InstallMethod *InstallMethod          `json:"install-method"`
	ProxyMeta     *ProxyMeta              `json:"proxy-info"`
	OtlpMeta      *OtlpMeta               `json:"otlp"`
}
