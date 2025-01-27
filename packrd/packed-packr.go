// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "0581941e5555a42d45edd6a86b897e15"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"793124cfbcf9ae17892b017c666be8f1": "1f8b08000000000000ff3cccb14e042110c6f19a798a0985010bb637b1b43d2d2eb617c4d95da20b38ccc618c2bb9bdd906be73fdfaff8f0e517c2d6d025bf917b1b87de01e256320b1a505aa84a4c8b06507a89b2ee1f2ee46daac22461e5e9ecf3dfc4f4b347260d1660de53c02b55b9bd168939d5dbbbff8e9f5ec8083e0ed15d2d36508c4fcf38c6ee42bf462c80ca45ea111e06d03a28e2f3f748eeee5950ec2ef98539b321660b1dfe030000ffff4d5ec109db000000",
		"920fdc8a8285bc442cfb791bcafe45b6": "1f8b08000000000000ff8c51cb4ec33010bcef572c3ed955e5481c41bd8004b78204e2ee868d6bd5b12d67034151fe1dd9e5a11e785c6766673ce364da83b184f38c3a989ef4fd07b02c00ae4f31334a4014d6f17edce936f68d8dbbb1eb8c8f8da510dec48f74f2e3b0b7147e5198f6909b97f33f3c0428806e0c2d6ee955c6c403aeee12bb18068572559fa16f2950361cf31a29e79815ce8068f162834741b9550088ae2b8a42142bfd64bc7b364c525d56fc6c83c1f97a8d9889c71cd05653405c4e0dacbe8a93ac35aabf2843eee2b4357dd950ac5168dd7cad7be37c811ba63e79c33408f5bf50c496a71258f72851d731304d2cd591d30fc452943e625d6b15dceac76cc2d0c5dc53969fdf7102b63ca922fd8e0ccec302f01e0000ffff7aa1c8be1a020000",
		"a61b972dce26544398f591b4e2f715f3": "1f8b08000000000000ff548fc16ae3301086cfd653081d16693172f6bab0a7253d95b484400fa514c51d2ba296e48c4634c1f8dd8b1453da8b0edf7c9aff9fc9f4efc6029f67ae83f1a01f57b02c8c393f45242e59230812b96005638db08e4ef9a8fbe83b1b8f7918cc183b0b215ccb5b4cf1d34a8440fd09bbba65b87608e7ec1004538c0d39f4fc00895e77f02189ff5ea3f441f19935c8fffee3abafaba1186b6ccb01eba8a05f0f13b918d2bc28d6a0dec52d62440988c5c51c8ab8562b3bf6390440a9ea4c3f393a495bcdafaf85ef7390aa6248b5446190f24849deec7b081221e9ffd17b13de52cb37ea3bbf7323a496ff29f670bb6385cf9b97226ecfd98c52c0c5f869044d17122d1ff4ce7828d10bfb0c0000fffffb4a36cd9f010000",
		"a95877cfc2957859690620af727512eb": "1f8b08000000000000fff248cdc9c957482bcacf55a8ae56d0cb4bcc4dd50bc92cc949cdac4a55a8ade502040000ffff8a41cdb420000000",
		"cb9f04739ceb356f960fa53f13efd3f3": "1f8b08000000000000ff3c8d31ae833010447b9f624af8053eca4f957e036bb0426c6bbd5b20cb778f8210ede8bd3785e637ad8cd63025faf0f4b886de9df31eff45634e15210b564e2ca431ad68eda47b777a14bea9aa62b3a239c07bd0b2e0c826a86a21606361d7dd997dd21e1752866ea4c8974ec2b06af4dad9054b33865cb4e2efca8fb7368c60912ce793b09a24a4b8ffeadf000000ffffb2b3ff10d1000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("github.com/gobuffalo/genny/genny/new", "../new/templates")
		b.SetResolver("-name-/-name-.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "920fdc8a8285bc442cfb791bcafe45b6"})
		b.SetResolver("-name-/-name-_test.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "a61b972dce26544398f591b4e2f715f3"})
		b.SetResolver("-name-/options.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "cb9f04739ceb356f960fa53f13efd3f3"})
		b.SetResolver("-name-/options_test.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "793124cfbcf9ae17892b017c666be8f1"})
		b.SetResolver("-name-/templates/example.txt.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "a95877cfc2957859690620af727512eb"})
		}()

	return nil
}()
