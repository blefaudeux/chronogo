# chronogo [![Build Status](https://travis-ci.org/blefaudeux/chronogo.svg?branch=master)](https://travis-ci.org/blefaudeux/chronogo)

Like cron, but hopefully easier to set up and in userspace.

Execute asynchronously a given set of commands, on a (timed) regular basis, log the outputs and handle things like computer restarts.

This also handles a folderwatch, in that you can add as many folders to watch as you like, and associate a command triggered by any change in them.

Future plans include:
- unit testing, adding error messages if the config is broken
- adding pre-built binaries for major platforms
- better log handling

## Install
For now, you'll need golang installed. Once done (using brew, apt-get, chocolatey or other) just run :

`go get -u github.com/blefaudeux/chronogo`

## Settings structure

This structure is an example, it can be generated from scratch with the '--initSettings' flag.

```json
{
    "WaitForInternetConnection" : true,
    "TimedCommands" : {
        "Daily": [
            {
                "Command" : "echo",
                "Args" : [
                    "daily test 1"
                ]
            },
            {
                "Command" : "cat",
                "Args" : [
                    "/home"
                ]
            }
        ],
        "Weekly": [
            {
                "Command" : "echo",
                "Args" : [
                    "This is the weekly",
                    " test number 1"
                ]
            }
        ],
        "Monthly": [
            {
                "Command" : "echo",
                "Args" : [
                    "Monthly test"
                ]
            },
            {
                "Command" : "sleep",
                "Args" : [
                    "200"
                ]
            }
        ]
    },
    "MaxCommandsInFlight": 2,
    "FolderWatch": [
        {
            "FolderToWatch": "~/test",
            "CommandToTrigger": {
                "Command": "echo",
                "Args": [
                    "Just detected a change in ~/test"
                ]
            }
        }
    ],
    "DBPath": "chronoDB",
    "LogPath": "logFromTheLastRun.log"
}
```
