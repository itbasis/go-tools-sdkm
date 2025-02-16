package base

import (
	"context"
	"log/slog"

	itbasisCoreOption "github.com/itbasis/go-tools-core/option"
	sdkmLog "github.com/itbasis/go-tools-sdkm/internal/log"
	sdkmSdk "github.com/itbasis/go-tools-sdkm/internal/sdk"
)

const _optionSdkDirKey itbasisCoreOption.Key = "option-sdk-dir"

func WithDefaultSdkDir() itbasisCoreOption.Option[basePlugin] {
	return &_optionSdkDir{}
}

func WithCustomSdkDir(sdkDir string) itbasisCoreOption.Option[basePlugin] {
	return &_optionSdkDir{dir: sdkDir}
}

type _optionSdkDir struct {
	dir string
}

func (r *_optionSdkDir) Key() itbasisCoreOption.Key { return _optionSdkDirKey }

func (r *_optionSdkDir) Apply(_ context.Context, cmp *basePlugin) error {
	slog.Debug("apply SDK directory option", sdkmLog.SlogAttrRootDir(r.dir))

	if r.dir != "" {
		cmp.sdkDir = r.dir

		return nil
	}

	cmp.sdkDir = sdkmSdk.GetDefaultSdkRoot()

	return nil
}
