#!/usr/bin/env bash

echo "pve (local) pve test-999"
. .env.pve
./proxrm --id 999

echo "-----"

unset PROXMOX_TOKEN

echo "pvj-p (prod is dev) test-999"
. .env.pvj-p
./proxrm --id 999

echo "-----"

unset PROXMOX_PASSWORD

echo "pvj-i (dev is prod) test-999"
. .env.pvj-i
./proxrm --id 999

unset PROXMOX_PASSWORD
# EOF
