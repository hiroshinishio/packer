package version

import (
	"github.com/hashicorp/packer/packer-plugin-sdk/version"
	packerVersion "github.com/hashicorp/packer/version"
)

var ChefClientPluginVersion *version.PluginVersion

func init() {
	ChefClientPluginVersion = version.InitializePluginVersion(
		packerVersion.Version, packerVersion.VersionPrerelease)
}
