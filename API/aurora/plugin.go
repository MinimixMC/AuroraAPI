package aurora

type PluginInterface interface {
	Init()
	Plugin() *Plugin
}

type PluginVersion struct {
	Version        string       // The Version tag (e.g. v1.0.0-ALPHA)
	VersionSupport map[int]bool // What versions of aurora does this plugin support
}

// Check if a given version of the plugin is supported
func (pv *PluginVersion) CheckVersionSupported(version int) bool {
	// Retrieve the value from the map for the given version
	supported, ok := pv.VersionSupport[version]
	if !ok {
		// If the version is not in the map, return false
		return false
	}
	// Return the value from the map for the given version
	return supported
}

type Plugin struct {
	Name    string        // Name of the plugin
	Author  []string      // Author(s) of the plugin
	Version PluginVersion // Plugin version info

	MainFunc func(p *Plugin)
}

func (p *Plugin) Init() {
	p.MainFunc(p)
}

func (p *Plugin) Plugin() *Plugin {
	return p
}
