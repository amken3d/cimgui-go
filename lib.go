//go:build required
// +build required

package imgui

import (
	_ "github.com/amken3d/cimgui-go/lib/linux/x64"
	_ "github.com/amken3d/cimgui-go/lib/macos/arm64"
	_ "github.com/amken3d/cimgui-go/lib/macos/x64"
	_ "github.com/amken3d/cimgui-go/lib/windows/x64"
)
