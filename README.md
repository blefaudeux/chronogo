# chronogo
[Complete WIP, come back in ten years] Like cron, for human beings

## Settings structure

```json
{
    "waitForInternetConnection" : true,
    "timedCommands" : {
        "daily": [
            "echo 'this is a daily test'"
        ],
        "weekly": [
            "echo 'this is a weekly test'"
        ],
        "monthly": [
            "echo 'this is a monthly test'"
        ]
    },
    "folderWatchCommands" : [
        {
            "folderToWatch": "home/user/blah",
            "commandToTrigger": "echo 'I just witnessed a change'"
        }
    ]
}
```