package downloader_test

import (
	pluginGoConsts "github.com/itbasis/go-tools-sdkm/internal/plugins/golang/consts"
	"github.com/itbasis/go-tools-sdkm/internal/plugins/golang/downloader"
	sdkmPlugin "github.com/itbasis/go-tools-sdkm/pkg/plugin"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = ginkgo.Describe(
	"URLForDownload", func() {
		ginkgo.GinkgoRecover()

		ginkgo.DescribeTable(
			"success", func(os, arch, version, wantURL string) {
				var (
					mockController = gomock.NewController(ginkgo.GinkgoT())
					mockBasePlugin = sdkmPlugin.NewMockBasePlugin(mockController)
				)

				gomega.Expect(
					downloader.NewDownloader(os, arch, pluginGoConsts.URLReleases, mockBasePlugin).
						URLForDownload(version),
				).
					To(gomega.Equal(wantURL))
			},
			// stable
			ginkgo.Entry(nil, "darwin", "amd64", "1.22.5", "https://go.dev/dl/go1.22.5.darwin-amd64.tar.gz"),
			ginkgo.Entry(nil, "darwin", "arm64", "1.22.5", "https://go.dev/dl/go1.22.5.darwin-arm64.tar.gz"),
			ginkgo.Entry(nil, "windows", "amd64", "1.22.5", "https://go.dev/dl/go1.22.5.windows-amd64.zip"),
			ginkgo.Entry(nil, "linux", "386", "1.22.5", "https://go.dev/dl/go1.22.5.linux-386.tar.gz"),
			ginkgo.Entry(nil, "linux", "amd64", "1.22.5", "https://go.dev/dl/go1.22.5.linux-amd64.tar.gz"),
			// unstable
			ginkgo.Entry(nil, "darwin", "arm64", "1.23rc2", "https://go.dev/dl/go1.23rc2.darwin-arm64.tar.gz"),
			// archived
			ginkgo.Entry(nil, "darwin", "arm64", "1.21.6", "https://go.dev/dl/go1.21.6.darwin-arm64.tar.gz"),
		)
	},
)
