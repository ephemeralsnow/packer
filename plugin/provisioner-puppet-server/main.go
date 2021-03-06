package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/mitchellh/packer/provisioner/puppet-server"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(puppetserver.Provisioner))
	server.Serve()
}
