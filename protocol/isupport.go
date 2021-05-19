package protocol

func CreateISupport(server, nick string) Line {
	return NewLine(
		server,
		"005",
		[]string{
			"AWAYLEN=64",
			"CASEMAPPING=ascii",
			"CHANTYPES=#",
			"CHANLIMIT=#:",
			"CHANNELLEN=16",
		},
	)
}
