# letscloud-cli

The LetsCloud Command Line is an interface for interacting with LetsCloud Infrastructure.

```
NAME:
   letscloud - manage your LetsCloud resources from your terminal

USAGE:
   letscloud [command]

VERSION:
   v1.2.0

COMMANDS:
   api-key    Show or Set your API Key
   locations  Show All Locations
   plans      Show All Plans by Location
   images     Show All Images by Location
   ssh-key    Manage your SSH Keys
   snapshot   Manage your snapshots
   instance   Manage your instances
   profile    Show your Profile Info
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

Use "letscloud [command] --help" for more information about a command.
```

- [Installing `letscloud`](#installing-letscloud)
   - [Downloading a Release from GitHub](#downloading-a-release-from-github)
   - [Building the Development Version from Source](#building-the-development-version-from-source)
   - [Using Docker](#using-docker)
- [Getting Started](#getting-started)
- [Examples](#examples)
   - [Create a New SSH Key](#create-a-new-ssh-key)
   - [List all instances](#list-all-instances)
   - [Create new instance](#create-new-instance)
   - [List all available locations](#list-all-available-locations)
   - [List all available images by location](#list-all-available-images-by-location)
   - [List all available plans by location](#list-all-available-plans-by-location)

## Installing `letscloud`

You can download pre-built binaries for Linux, macOS, and Windows on
the [releases page](https://github.com/letscloud-community/letscloud-cli/releases).

### Downloading a Release from GitHub

Find the appropriate archive for your operating system and architecture and download the archive from your browser or copy its URL and retrieve it to your home directory with `wget` or `curl`.

In this example we used the `version` `1.2.0`, but you have to change for the latest version or what version you want.

Example using `wget`:

```
cd ~
wget https://github.com/letscloud-community/letscloud-cli/releases/download/v1.2.0/letscloud-1.2.0-linux-amd64.tar.gz
```

Or using `curl`

```
cd ~
curl -OL https://github.com/letscloud-community/letscloud-cli/releases/download/v1.2.0/letscloud-1.2.0-linux-amd64.tar.gz
```

Extract the binary:

```
tar xf ~/letscloud-1.2.0-linux-amd64.tar.gz
```

Move the letscloud binary to somewhere in your path. For example, on GNU/Linux and OS X systems:

```
sudo mv ~/letscloud /usr/local/bin
```

### Building the Development Version from Source

If you have a Go environment configured, you can install the development version from the command line.

```
go install github.com/letscloud-community/letscloud-cli/cmd/letscloud@latest
```

Another way to build from source is clone the repository:

```
git clone https://github.com/letscloud-community/letscloud-cli.git
cd letscloud-cli
make build
```

### Using Docker

You can also find the image on [Docker Hub](https://hub.docker.com/r/letscloudcommunity/letscloud-cli)

Get the letscloud-cli for Docker
```
docker pull letscloudcommunity/letscloud-cli:latest
```
#### Using letscloud-cli with Docker
When using letscloud-cli with Docker, you must pass the environment variable with your API KEY.

> Note: Replace `<YOUR-API-KEY>` with your actual LetsCloud API key.

Example: Running a basic command to return your profile:
```
docker run --rm -it \
  --env=LETSCLOUD_API_KEY='<YOUR-API-KEY>' \
  letscloudcommunity/letscloud-cli profile
```
Example: Running a command to list your instances:
```
docker run --rm -it \
  --env=LETSCLOUD_API_KEY='<YOUR-API-KEY>' \
  letscloudcommunity/letscloud-cli instance list
```

## Getting Started

1. Visit the LetsCloud Account at [my.letscloud.io](https://my.letscloud.io),

2. Go to [API tab](https://my.letscloud.io/profile/client-api), and press `Enable API` button.

3. You'll get your API Key. Copy it and save to a safe place.

4. Run the following command to set your API Key for the first time.

   `letscloud api-key set 'your_api_key'`
   **remember to use single quotes in your api key**  
   
   Alternatively, you can export the API Key as an environment variable in linux OS.
   
   `export LETSCLOUD_API_KEY='your_api_key'`
    
5. That's it. Now you can use the program.

   For example, you can run the `profile` command to see your profile info and your balance.
   
   `letscloud profile`
    
## Examples

### Create a New SSH Key
If you do not provide a `public key`, we will generate a new key and the `private key` will be shown.  
**Save this key as it is not stored**
```
letscloud ssh-key create --title=my-ssh-key
```

### List all instances
```
letscloud instance list
```

### Create new instance
```
letscloud instance create \
   --location <location-slug> \
   --plan <plan-slug> \
   --image <image-slug> \
   --hostname <host-name> \
   --label <label> \
   --password <password>
```
**The password must:**
- Be between 12 and 32 characters.
- Contain at least one uppercase letter.
- Contain at least one lowercase letter.
- Contain at least one digit.
- Contain at least one special character.

Example:
```
letscloud instance create \
   --location MIA1 \
   --plan 1vcpu-1gb-10ssd \
   --image ubuntu-24.04-x86_64 \
   --hostname test-hostname.com \
   --label test-cli \
   --password exampleExample@123
```

Example one-line:
```
letscloud instance create --location MIA1 --plan 1vcpu-1gb-10ssd --image ubuntu-24.04-x86_64 --hostname test-hostname.com --label test-cli --password exampleExample@123
```

or you can create an instance using your SSH key:
```
letscloud ssh-key list
```
```
letscloud instance create \
   --location MIA1 \
   --plan 1vcpu-1gb-10ssd \
   --image ubuntu-24.04-x86_64 \
   --hostname test-hostname.com \
   --label test-cli-key \
   --ssh <ssh-slug>
```

### List all available locations
```
letscloud locations
```

### List all available images by location
```
letscloud images <location-slug>
```

### List all available plans by location
```
letscloud plans <location-slug>
```

## License

MIT license