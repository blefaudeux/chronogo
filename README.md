# chronogo
[Complete WIP, come back in ten years] Like cron, for human beings

Execute asynchronously a given set of commands, on a (timed) regular basis, log the outputs and handle things like computer restarts. Not there yet, but should in the end handle a folderwatch, as well as conditional execution (execute this if this command returns that). Should already be functional

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
    "DBPath": "chronoDB"
}
```