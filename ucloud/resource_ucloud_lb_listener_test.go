package ucloud

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
)

func TestAccUCloudLBListener_basic(t *testing.T) {
	var lbSet ulb.ULBSet
	var vserverSet ulb.ULBVServerSet
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		IDRefreshName: "ucloud_lb_listener.foo",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckLBListenerDestroy,

		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLBListenerConfig,

				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBExists("ucloud_lb.foo", &lbSet),
					testAccCheckLBListenerExists("ucloud_lb_listener.foo", &lbSet, &vserverSet),
					testAccCheckLBListenerAttributes(&vserverSet),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "protocol", "HTTPS"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "method", "Source"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "name", "testAcc"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "idle_timeout", "80"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "persistence_type", "ServerInsert"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "health_check_type", "Path"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "path", "/l"),
				),
			},

			resource.TestStep{
				Config: testAccLBListenerConfigTwo,

				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBExists("ucloud_lb.foo", &lbSet),
					testAccCheckLBListenerExists("ucloud_lb_listener.foo", &lbSet, &vserverSet),
					testAccCheckLBListenerAttributes(&vserverSet),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "protocol", "HTTP"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "method", "Roundrobin"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "name", "testAccTwo"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "idle_timeout", "100"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "persistence_type", "None"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "health_check_type", "Port"),
					resource.TestCheckResourceAttr("ucloud_lb_listener.foo", "domain", "www.ucloud.cn"),
				),
			},
		},
	})
}

func testAccCheckLBListenerExists(n string, lbSet *ulb.ULBSet, vserverSet *ulb.ULBVServerSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("LBListener id is empty")
		}

		client := testAccProvider.Meta().(*UCloudClient)
		ptr, err := client.describeVServerById(lbSet.ULBId, rs.Primary.ID)

		log.Printf("[INFO] LBListener id %#v", rs.Primary.ID)

		if err != nil {
			return err
		}

		*vserverSet = *ptr
		return nil
	}
}

func testAccCheckLBListenerAttributes(vserverSet *ulb.ULBVServerSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if vserverSet.VServerId == "" {
			return fmt.Errorf("LBListener id is empty")
		}
		return nil
	}
}

func testAccCheckLBListenerDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ucloud_lb_listener" {
			continue
		}

		client := testAccProvider.Meta().(*UCloudClient)
		d, err := client.describeVServerById(rs.Primary.Attributes["load_balancer_id"], rs.Primary.ID)

		// Verify the error is what we want
		if err != nil {
			if isNotFoundError(err) {
				continue
			}
			return err
		}

		if d.VServerId != "" {
			return fmt.Errorf("LBListener still exist")
		}
	}

	return nil
}

const testAccLBListenerConfig = `
resource "ucloud_lb" "foo" {
}

resource "ucloud_lb_listener" "foo" {
	load_balancer_id = "${ucloud_lb.foo.id}"
	protocol = "HTTPS"
	method = "Source"
	name = "testAcc"
	idle_timeout = 80
	persistence_type = "ServerInsert"
	health_check_type = "Path"
	path = "/l"

}
`
const testAccLBListenerConfigTwo = `
resource "ucloud_lb" "foo" {
}

resource "ucloud_lb_listener" "foo" {
	load_balancer_id = "${ucloud_lb.foo.id}"
	protocol = "HTTP"
	method = "Roundrobin"
	name = "testAccTwo"
	idle_timeout = 100
	persistence_type = "None"
	health_check_type = "Port"
	domain = "www.ucloud.cn"
}
`
