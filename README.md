     AAAAAAAAAAAA     AAAAAAAAAAAA       AAAAAAAAAA     NNNN        NNN    AAAAAAAAAAA     AAAAA         AAAA
     AAAAAAAAAAAAAA  AAAAAAAAAAAAAAA   AAAAAAAAAAAAAA  NNNNNN      NNNNNN AAAAAAAAAAAAAA   AAAAAAA     AAAAAAA
    AAAAAAAAAAAAAAAA AAAAAAAAAAAAAAAA  AAAAAAAAAAAAAAA  NNNNNN    NNNNNN  AAAAAAAAAAAAAAA  AAAAAAAA   AAAAAAAA
    AAAAAAAAAAAAAAAA AAAAAAAAAAAAAAAA AAAAAAAAAAAAAAAA   NNNNNN  NNNNNN   AAAAAAAAAAAAAAA  AAAAAAAA   AAAAAAAA
    AAAAAAAAAAAAAAAA AAAAAAAAAAAAAAAA AAAAAAAAAAAAAAAA    NNNNNNNNNNNNN   AAAAAAAAAAAAAAAA AAAAAAAAA AAAAAAAAA
    AAAAA       AAAA AAAAA      AAAAA AAAAA      AAAAA     NNNNNNNNNNN    AAAAA      AAAAA AAAA AAAA AAAAAAAAA
    AAAAA       AAAA AAAAA      AAAAA AAAAA      AAAAA     NNNNNNNNNN     AAAAA      AAAAA AAAA AAAAAAAAA AAAA
    AAAAA      AAAAA AAAAAAAAAAAAAAAA AAAAA      AAAAA      NNNNNNNN      AAAAAAAAAAAAAAAA AAAA AAAAAAAAA AAAA
    AAAAAAAAAAAAAAAA AAAAAAAAAAAAAAAA AAAAA      AAAAA      NNNNNNNN      AAAAAAAAAAAAAAA  AAAA  AAAAAAA  AAAA
    AAAAAAAAAAAAAAAA AAAAAAAAAAAAAAA  AAAAA      AAAAA     NNNNNNNNN      AAAAAAAAAAAAAAA  AAAA  AAAAAAA  AAAA
    AAAAAAAAAAAAAAA  AAAAAAAAAAAAAAA  AAAAA      AAAAA    NNNNNNNNNNN     AAAAAAAAAAAAAA   AAAA   AAAAA   AAAA
    AAAAAAAAAAAAAA   AAAAA   AAAAA    AAAAAAAAAAAAAAAA   NNNNNNNNNNNNN    AAAAA   AAAAA    AAAA   AAAAA   AAAA
    AAAAA            AAAAA    AAAAA   AAAAAAAAAAAAAAAA   NNNNNN  NNNNNN   AAAAA    AAAAA   AAAA    AAAA   AAAA
    AAAAA            AAAAA    AAAAAA   AAAAAAAAAAAAAAA  NNNNNN    NNNNNN  AAAAA    AAAAAA  AAAA     AA    AAAA
    AAAA             AAAA      AAAAA   AAAAAAAAAAAAAA  NNNNNN      NNNNNN AAAA      AAAAA  AAAA            AAA
    AAA              AAA         AAA     AAAAAAAAAAA     NNN        NNNN  AAA         AAA  AA               AA

# proxrm

[“If I have seen further, it is by standing on the shoulders of giants.”](https://en.wikipedia.org/wiki/Standing_on_the_shoulders_of_giants)  -- Sir Isaac Newton

[![Build Status](https://github.com/CTFfactory/proxrm/actions/workflows/cd.yml/badge.svg?branch=master)](https://github.com/CTFfactory/proxrm/actions/workflows/cd.yml)

Remove Packer build template/vm artifacts from Proxmox with this easy to use tool.

## .env
```
export PROXMOX_URL=http://10.31.33.7:8006/api2/json
export PROXMOX_USERNAME=user@node!tokenid
export PROXMOX_PASSWORD=Password
export PROXMOX_TOKEN=c2c42a66-2866-3795-2238-ef60a11544c0
```

### TODO List:

 - [ ] Github Actions
   - [ ] actionlint
   - [ ] yamllint
   - [x] get go-proxmox package
   - [x] format
   - [x] vet
   - [x] tidy modules
   - [x] download modules
   - [x] verify modules
   - [x] lint
   - [x] staticcheck
   - [x] gocyclo
     - [x] top 10
     - [x] over 5
     - [x] average
     - [x] over 2 average
   - [x] ineffassign
   - [x] vulncheck
 - [ ] error handling
   - [x] functions return error
   - [ ] parse proxmox errors
