// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	_ "embed"
	// "os"
	"path/filepath"

	"github.com/siderolabs/go-copy/copy"
	"github.com/siderolabs/talos/pkg/machinery/overlay"
	"github.com/siderolabs/talos/pkg/machinery/overlay/adapter"
)

const (
	dtb       = "rockchip/rk3328-roc-cc.dtb"
)

func main() {
	adapter.Execute(&RocCcRk3328Installer{})
}

type RocCcRk3328Installer struct{}

type RocCcRk3328ExtraOptions struct {
	Console    []string `json:"console"`
	ConfigFile string   `json:"configFile"`
}

func (i *RocCcRk3328Installer) GetOptions(extra RocCcRk3328ExtraOptions) (overlay.Options, error) {
	kernelArgs := []string{
		"console=tty0",
		"sysctl.kernel.kexec_load_disabled=1",
		"talos.dashboard.disabled=1",
	}

	kernelArgs = append(kernelArgs, extra.Console...)

	return overlay.Options{
		Name:       "roc-cc-rk3328",
		KernelArgs: kernelArgs,
	}, nil
}

func (i *RocCcRk3328Installer) Install(options overlay.InstallOptions[RocCcRk3328ExtraOptions]) error {
	// allows to copy a directory from the overlay to the target
	// err := copy.Dir(filepath.Join(options.ArtifactsPath, "arm64/firmware/boot"), filepath.Join(options.MountPrefix, "/boot/EFI"))
	// if err != nil {
	// 	return err
	// }

	// allows to copy a file from the overlay to the target
	err := copy.File(filepath.Join(options.ArtifactsPath, "arm64/u-boot/RocCcRk3328/u-boot.bin"), filepath.Join(options.MountPrefix, "/boot/EFI/u-boot.bin"))
	if err != nil {
		return err
	}

	if options.ExtraOptions.ConfigFile != "" {
		// do something with the config file
	}

	return nil
	// var f *os.File
	// err = f.Sync()
	// if err != nil {
	// 	return err
	// }

	// src := filepath.Join(options.ArtifactsPath, "arm64/dtb", dtb)
	// dst := filepath.Join(options.MountPrefix, "/boot/EFI/dtb", dtb)

	// err = os.MkdirAll(filepath.Dir(dst), 0o600)
	// if err != nil {
	// 	return err
	// }

	// return copy.File(src, dst)
}
