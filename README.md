# tag-metainfo-xml

tool to update appstream manifests with release data using git tags

## install

```bash
git clone https://github.com/dada513/tag-metainfo-xml
cd tag-metainfo-xml
go build
# to install non-portable - copy tag-metainfo-xml to another location like /usr/bin
```

## usage

```
❯ ./tag-metainfo-xml
Usage: ./tag-metainfo-xml <user> <repo> <filename>
```

## api auth

use this if you face ratelimits or need to access private repo  
set PAT env variable to your personal access token
