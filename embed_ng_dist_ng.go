package gongsvg

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongsvg/dist/ng-github.com-fullstack-lang-gongsvg
var NgDistNg embed.FS
