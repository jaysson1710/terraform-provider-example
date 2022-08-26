terraform {
  required_providers {
    example = {
      version = "= 0.2.7"
      source  = "terraform-example.com/exampleprovider/example"
    }
  }
}


resource "example_server" "my-server-name" {
	env = "1"
	resource="env"
}

output "name" {
  value = example_server.my-server-name.name
}