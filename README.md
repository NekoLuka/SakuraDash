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
### Authelia specif feature
If you put the dashboard behind authelia with groups configured,
you can forward the header `Remote-Groups` from authelia to sakuradash.
This unlocks functionality to show certain categories only to users who belong to the configured group.
Unlock this functionality by adding `view_group` to a categories entry like so:
```yaml
...
categories:
  - name: ...
    view_group: "jellyfin_users"
    items:
...
```
This makes this particular category only visible to users in the `jellyfin_users` group.

## Build
Clone the git repo and run the following command:  
`docker build -t nekoluka/sakuradash .`

## Pull
`docker pull nekoluka/sakuradash`

## Running
Run the following command after you build the image:  
`docker run -p 8080:8080 -v ./config.yaml:/config.yaml nekoluka/sakuradash`

## Screenshot
<img width="1469" alt="image" src="https://github.com/user-attachments/assets/453caf90-ca54-4fad-896e-d002c0b66460" />


## Mentions
This project makes use of w3.css
