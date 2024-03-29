![Go version](https://img.shields.io/github/go-mod/go-version/ravmda/go-bots?color=green)
![License](https://img.shields.io/github/license/ravmda/go-bots)
[![Go Report Card](https://goreportcard.com/badge/github.com/RavMda/go-bots)](https://goreportcard.com/report/github.com/RavMda/go-bots)

# Go-Bots

Go-Bots was created to help server owners stress-test their servers with bots

**Be aware that many anti-bot measures will successfully block this tool**


## Features
- Convenient configuration via YAML
- Register/Login on the server
- Change protocol to the one you prefer [(check out this list)](https://wiki.vg/Protocol_version_numbers)
- Spam chat with custom phrases
- "Smart" cooldown system
- ~~It can simulate jumps, arm swings and head rotations~~
- Custom packet spam *(not tested, better not use it)*
- Low RAM and CPU usage compared to some other solutions

  ![uh](https://cdn.discordapp.com/attachments/744430106067599362/815245304345133076/ezgif-5-e19a83f1263e.gif)
## Known Issues
- Bots can stuck in blocks when jumping because there are no physics involved
- Packet spam can freeze on Windows systems
- Older protocol versions may not work
- ![help...](https://i.imgur.com/GAfisZH.png)

## Usage

1. Download latest artifact [here](https://github.com/RavMda/go-bots/actions) or compile the binary yourself

1. Create "proxies.txt" and add SOCKS4 proxies there in this format:
    ```
   host:port
   host:port
   ...
   ```


3. Change values that will correspond to you in "config.yml"
4. Run it!


## Build


```
git clone https://github.com/RavMda/go-bots.git
cd ./go-bots/
go build -ldflags "-s -w" -o ./go-bots
```

## Preventive Measures
[[2LS] AntiBot](https://www.spigotmc.org/resources/2ls-antibot-the-ultimate-antibot-plugin.62847/)

## License
[GPL-3.0 License](https://choosealicense.com/licenses/gpl-3.0/)
