# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "puppetlabs/centos-6.5-64-puppet"

  config.vm.network "public_network", bridge: 'en0: Wi-Fi (AirPort)'

  if Vagrant.has_plugin?("vagrant-cachier")
    config.cache.scope = :box
  end

  config.vm.provision :shell, inline: "yum -y install http://ftp.riken.jp/Linux/fedora/epel/6/i386/epel-release-6-8.noarch.rpm"
  config.vm.provision :shell, inline: "yum -y install golang ruby-devel git rpm-build nss-mdns"
  config.vm.provision :shell, inline: "gem install bundler"
end
