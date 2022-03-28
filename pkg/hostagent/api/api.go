package api

type Info struct {
	SSHLocalPort int `json:"sshLocalPort,omitempty"`
}

type ReloadMounts struct {
	SSHLocalPort int `json:"sshLocalPort,omitempty"`
}
