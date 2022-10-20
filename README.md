# Discord Command Cleaner
This is a utility used to remove all active commands for a specific bot on
Discord. This is a very simple utility meant to be used in conjunction with
a more complex bot. The use of this utility is mainly to remove stale commands
without having to write additional migration code into a bot.

If you use Kubernetes, this bot can be used as an init container for your bot.

## Usage
This section describes how to use this utility. To get started, download
the correct [release](https://github.com/MrFlynn/discord-command-cleaner/releases)
for your system. Unpack the tarball and change directories into the resulting
folder.

Before you can run the bot, familiarize yourself with the following environment
variables and command line flags.

| Environment Variable     | Flag          | Description                                   |
| ------------------------ | ------------- | --------------------------------------------- |
| `DISCORD_CC_TOKEN`       | `-token`      | Discord bot token. Required value.            |
| `DISCORD_CC_GUILD_ID`    | `-guildID`    | ID of specific guild to remove commands from. |
| `DISCORD_CC_SHOW_STATUS` | `-showStatus` | Show a status message during cleanup.         |

You can either use environment variables or command line flags to configure
the utility. They are interchangable, but note that the command line flags take
precedence over the environment variables.

**Warning:** Make sure you have stopped the bot you wish to clean up commands 
for before running this utility.

### Running it Using the Command Line
To run the utility, just open a shell into the directory containing the binary
and run it. For example:

```bash
$ ./discord-command-cleaner -token $DISCORD_TOKEN
```

### Running it as an Init Container in Kubernetes
To include this utility as an init container, add the following yaml under the
`spec` key in the resource definition.

```yaml
initContainers:
  - name: clean-commands
    image: ghcr.io/mrflynn/discord-command-cleaner
    env:
      - name: DISCORD_CC_TOKEN
        valueFrom:
          secretKeyRef:
            name: discord-secrets
            key: discord-token
      - name: DISCORD_CC_GUILD_ID
        value: ""
      - name: DISCORD_CC_SHOW_STATUS
        value: false
```