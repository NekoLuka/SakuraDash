# SakuraDash
A simple dashboard that you configure with a simple yaml file, and pulls icons from the selfhst/icons project.

## Disclaimer
The current background of the site is generated using AI. 
This is to be replaced with commissioned art sometime in the future.

## TODO
- Make the config file location configurable
- Implement icon caching
- Implement logging
- Publish a docker image
- Create overrideable static files for customisation

## Config
Adapt the following example to your own needs
```yaml
title: SakuraDash
categories:
  - name: "Infrastructure"
    items:
      - name: "OPNSense"
        icon: "https://cdn.jsdelivr.net/gh/selfhst/icons/webp/opnsense.webp"
        url: "http://opnsense.lan"
      - name: "Proxmox"
        icon: "https://cdn.jsdelivr.net/gh/selfhst/icons/webp/proxmox.webp"
        url: "http://proxmox.lan:8006"
```

## Build
Clone the git repo and run the following command:  
`docker build -t sakuradash .`

## Running
Run the following command after you build the image:  
`docker run -p 8080:8080 -v ./config.yaml:/config.yaml sakuradash`

## Screenshot


## Mentions
This project makes use of w3.css
