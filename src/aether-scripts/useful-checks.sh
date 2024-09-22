#!/bin/bash
# Check the type of Resource the System is Virtualised in -- QEMU/VirtualBox/LXD-LXC/Container etc.
sudo dmidecode | grep -i 'manufacture\|product'
sudo dmidecode -s system-product-name
cat /sys/class/dmi/id/product_name
systemd-detect-virt