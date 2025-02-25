package ng

import "embed"

// Projects is the export of the angular source code. This allows for go vendoring to import ng sources
// and align go source and ng source of the gong stack
//
// when this package is imported and vendored,
// it will import in the vendor dir all the files in the current directory and the projects directory
//
//go:embed projects
var Projects embed.FS
