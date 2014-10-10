GroupMengine
============

GroupMengine is a golang based bot framework for GroupMe.  It is designed to be hosted on Google App Engine.

A GroupMengineConfig object needs to be created for this guy to run.  Right now main.go:init expects to find a function GetConfig which will return back the configuration.

Here is an example GroupMengineConfig object:
```
	config := GroupMengineConfig{
		Handlers: map[string]HandleFunc{
			"/groupmengine": randoText,
		},
		Bots: map[string]string{
			"GROUP_ID_HERE": "BOT_ID_HERE",
		},
	}
```	
	
The config serves two purposes right now: 1) Define a map of commands and their handlers 2) Define a map of groups and the bot to use for posting to tha group.

You will need to create a bot and register a callback url for the data to start flowing to your GroupMengine, see https://dev.groupme.com/bots for more.