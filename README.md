# chronogo
[Complete WIP, come back in ten years] Like cron, for human beings

## Settings structure

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
                "Command" : "echo",
                "Args" : [
                    "daily test 1"
                ]
            }
        ],
        "Weekly": [
            "echo 'this is a weekly test'"
        ],
        "Monthly": [
            "echo 'this is a monthly test'"
        ]
    },
    "FolderWatchCommands" : [
        {
            "FolderToWatch": "home/user/blah",
            "CommandToTrigger": "echo 'I just witnessed a change'"
        }
    ]
}
```