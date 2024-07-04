#!/usr/bin/env bash

echo "pve (local) pve build16"
. .env.pve
./proxrm --id 999

echo "-----"

unset PROXMOX_TOKEN

echo "pvj-p (prod is dev) vpc-router0.pvj-p-1.test.dev.ctf"
. .env.pvj-p
./proxrm --id 999

echo "-----"

unset PROXMOX_PASSWORD

echo "pvj-i (dev is prod) vpn.prosversusjoes.net"
. .env.pvj-i
./proxrm --id 999

unset PROXMOX_PASSWORD
# EOF
