package assets

import "embed"

//go:embed russian5.txt
var Assets embed.FS

//go:embed html/*
var HTML embed.FS

//go:embed js/*
var Static embed.FS
