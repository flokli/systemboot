package bls

import ()

type Entry struct {
	Title        string   `key:"title"`
	Version      string   `key:"version"`
	MachineID    string   `key:"machine-id"`
	Linux        string   `key:"linux"`
	Initrd       string   `key:"initrd"`
	EFI          string   `key:"efi"`
	Options      []string `key:"options"`
	DeviceTree   string   `key:"devicetree"`
	Architecture string   `key:"architecture"`
}

// IsValid verifies at least linux or an efi key are set
func (entry *Entry) IsValid() bool {
	if entry.Linux != "" || entry.EFI != "" {
		//TODO: think about that
		return true
	}
	return false
}
