# letscloud: Command-line interface for LetsCloud

`letscloud` is a command-line interface for interacting with LetsCloud Infrastructure.

```
$ letscloud instance list
IDENTIFIER      LABEL           IPv4            OS                      STATUS
rstpgafocoznzx  my-new-label-1    38.143.68.93    CentOS 6.8 x86          running
xzosdqyydeyutr  my-new-label-2    38.143.68.43    Ubuntu 20.04 x64        running
cqwdciprplqmkk  my-new-label-3    38.143.68.53    CentOS 6.8 x86          running
xiflugipmecvfo  my-new-label-4    38.143.68.29    Ubuntu 20.04 x64        running
```

## Installation

You can download pre-built binaries for Linux, macOS, and Windows on
the [releases page](https://github.com/letscloud-community/letscloud-cli/releases).

### Build Manually

If you have **Go and Makefile** installed, you can build and install the `letscloud` program with:

    make build
    
## Getting Started

1. Visit the LetsCloud Account at [my.letscloud.io](https://my.letscloud.io),

2. Go to [API tab](https://my.letscloud.io/profile/client-api), and press `Enable API` button.

3. You'll get your API Key. Copy it and save to a safe place.

4. Run the following command to set your API Key for the first time.

   `letscloud api-key set 'your_api_key'`
   Alternatively, you can export the API Key as an environment variable in linux OS.
   
   `export LETSCLOUD_API_KEY='your_api_key'`
    
5. That's it. Now you can use the program.

   For example, you can run the `profile` command to see your profile info
   
   `letscloud profile`
    
## Examples

### List all instances

```
$ letscloud instance list
IDENTIFIER      LABEL           IPv4            OS                      STATUS
rstpgafocoznzx  my-new-label-1    38.143.68.93    CentOS 6.8 x86          running
xzosdqyydeyutr  my-new-label-2    38.143.68.43    Ubuntu 20.04 x64        running
cqwdciprplqmkk  my-new-label-3    38.143.68.53    CentOS 6.8 x86          running
xiflugipmecvfo  my-new-label-4    38.143.68.29    Ubuntu 20.04 x64        running
```

### Create a New SSH Key

If you do not provide a public key, we will generate a new key and the private key will be shown. **Save this key as it is not stored**
```
$  letscloud ssh-key create --title=my-ssh-key
   SSH Key my-ssh-key successfully created!
   Here's your private key, please store it in a safe place
   -----BEGIN RSA PRIVATE KEY-----
   MIICWwIBAAKBgQDHPpb9xt+3X7FrvpzeZ7qyNXFz6Q0uGU7pKEahW4SfkQjV6dQ2
   ................................................................
   -----END RSA PRIVATE KEY-----
```

### Get Help

You can get all the available commands using

```shell script
letscloud help
NAME:
   Official LetsCloud CLI - This cli helps you to manage your LetsCloud infrastructure from your terminal

USAGE:
   letscloud [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
   api-key    Show or Set your API Key
   locations  Show All Locations
   plans      Show Plans by Location
   images     Show All the Images by Location
   ssh-key    Manage your SSH Keys
   instance   Manage your instances
   profile    Show your Profile Info
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

You can get help for any command using

```shell script
letscloud ssh-key help
NAME:
   Official LetsCloud CLI ssh-key - Manage your SSH Keys

USAGE:
   Official LetsCloud CLI ssh-key [global options] command [command options] [arguments...]

COMMANDS:
   list
   create
   delete
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## License

MIT license