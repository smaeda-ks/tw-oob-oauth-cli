# twitter-oobOAuth-cli
A simple CLI client for Twitter's OOB OAuth ([PIN-based OAuth](https://developer.twitter.com/en/docs/basics/authentication/overview/pin-based-oauth)).


# Installation

No need to install. This is just a small binary file so you can download and use it as needed. The download package only includes the following binary:

|Name|Type|
|---|---|
|darwin_amd64|for macOS 10.10 and above (64-bit x86)|
|windows_amd64|for Windows (64-bit x86)|
|windows_386|for Windows (32-bit x86)|
|linux_amd64|for Linux (64-bit x86)|
|linux_386|for Linux (32-bit x86)|

# Example

### Help

```text
$ ./tw-oob-oauth -h
NAME:
   tw-oob-oauth - Twitter PIN-based OAuth CLI client

USAGE:
   tw-oob-oauth [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --consumer-key value, --key value        App's consumer key
   --consumer-secret value, --secret value  App's consumer secret
   --help, -h                               show help
   --version, -v                            print the version
```

### Command

```text
$ ./tw-oob-oauth --key ${your_consumer_key} --secret ${your_consumer_secret}
1. Go to https://api.twitter.com/oauth/authorize?oauth_token=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
2. Authorize the application
3. Enter the supplied PIN here (and press Enter):
1234567
Access Token: 12345678-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
Token Secret: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

As you can see above, it will output your access token and secret there so you can save this value.

Enjoy!
