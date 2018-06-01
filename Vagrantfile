# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.provider "virtualbox" do |v|
  	v.memory = 8192
	v.cpus = 3
	v.customize ["modifyvm", :id, "--vram", "128"]
  end
  config.vm.box_check_update = false
  config.vm.define "development" do |node|
    node.vm.box = "ubuntu/trusty64"
    node.vm.synced_folder "/workspace", "/workspace", create: true
    node.vm.network :forwarded_port, guest: 3000, host: 3000, host_ip: "127.0.0.1"
    node.vm.network :forwarded_port, guest: 8080, host: 8080, host_ip: "127.0.0.1"
    node.vm.network :forwarded_port, guest: 80, host: 8888, host_ip: "127.0.0.1"
    node.vm.network :forwarded_port, guest: 4546, host: 4546, host_ip: "127.0.0.1"
    node.vm.network :forwarded_port, guest: 2525, host: 2525, host_ip: "127.0.0.1"
    node.vm.network :forwarded_port, guest: 1521, host: 1521, host_ip: "127.0.0.1"
    node.vm.network :forwarded_port, guest: 9000, host: 9000, host_ip: "127.0.0.1"
	node.vm.network :forwarded_port, guest: 8500, host: 8500, host_ip: "127.0.0.1"

    node.vm.provision "bootstrap", type: "shell" do |s|
      s.privileged = false
    	s.path = "../provision.sh"
    	s.args = [
    				"#{ENV['ZIRAAT_USER']}",
    				"#{ENV['ZIRAAT_PASS']}",
    			 ]
    end
    node.vm.provision "ansible", type: "shell" do |s|
      s.privileged = false
    	s.path = "../ansible.sh"
      s.args = [
            "development",
            "#{ENV['ZIRAAT_USER']}", 
            "#{ENV['ZIRAAT_PASS']}",
          ]
    end
  end
end
