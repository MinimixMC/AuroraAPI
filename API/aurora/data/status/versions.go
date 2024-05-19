package status

// Protocol version is the key
var Version map[int]VersionInfo = map[int]VersionInfo{
	766: {
		Name:     "1.20.5-1.20.6",
		Protocol: 766,
	},
	765: {
		Name:     "1.20.3-1.20.4",
		Protocol: 765,
	},
	764: {
		Name:     "1.20.2",
		Protocol: 764,
	},
	763: {
		Name:     "1.20-1.20.1",
		Protocol: 763,
	},
}
