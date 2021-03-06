package civo

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceCivoSnapshots_basic(t *testing.T) {
	datasourceName := "data.civo_snapshot.foobar"
	name := acctest.RandomWithPrefix("snapshot-test")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCivoSnapshotsConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "name", name),
					resource.TestCheckResourceAttrSet(datasourceName, "hostname"),
					resource.TestCheckResourceAttrSet(datasourceName, "size_gb"),
				),
			},
		},
	})
}

func testAccDataSourceCivoSnapshotsConfig(name string) string {
	return fmt.Sprintf(`
resource "civo_instance" "vm" {
	hostname = "instance-%s"
}

resource "civo_snapshot" "foobar" {
	name = "%s"
	instance_id = civo_instance.vm.id
}

data "civo_snapshot" "foobar" {
	name = civo_snapshot.foobar.name
}
`, name, name)
}
