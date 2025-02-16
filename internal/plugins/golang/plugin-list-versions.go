package golang

import (
	"context"
	"strings"

	sdkmSDKVersion "github.com/itbasis/go-tools-sdkm/pkg/sdk-version"
)

func (receiver *goPlugin) ListAllVersions(ctx context.Context, rebuildCache bool) ([]sdkmSDKVersion.SDKVersion, error) {
	return receiver.sdkVersions.AllVersions(ctx, rebuildCache) //nolint:wrapcheck // TODO
}

func (receiver *goPlugin) ListAllVersionsByPrefix(ctx context.Context, rebuildCache bool, prefix string) ([]sdkmSDKVersion.SDKVersion, error) {
	var allVersions, err = receiver.sdkVersions.AllVersions(ctx, rebuildCache)

	if err != nil {
		return nil, err //nolint:wrapcheck // TODO
	}

	if prefix == "" {
		return allVersions, nil
	}

	var versions []sdkmSDKVersion.SDKVersion

	for _, sdkVersion := range allVersions {
		if strings.HasPrefix(sdkVersion.ID, prefix) {
			receiver.enrichSDKVersion(&sdkVersion)

			versions = append(versions, sdkVersion)
		}
	}

	return versions, nil
}
