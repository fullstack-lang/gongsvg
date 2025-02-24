package nggithubcomfullstacklanggongsvg

import "embed"

// Projects is the export of the angular source code. This allows for go vendoring to import ng sources
// and align go source and ng source of the gong stack
//
//go:embed projects package.json
var Projects embed.FS
