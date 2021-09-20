---
title: "Masters"
date: 2021-08-25T15:30:13+02:00
draft: false
weight: 3
---

It was designed to match with Dungeons and Dragons 5th edition terminology that means to start your campaign you need to:
- Create a campaign
- Add an adventure on it
- Add many encounters as you wanted on it
- Add players to your new campaign


## Dependencies

MongoDB - Recommended use an external service like [Mongo Atlas DB](https://www.mongodb.com/cloud)

NGROK - Public endpoint for your bots. It's easy and have a free plan. Check it [here](https://ngrok.com/).

## How to run

Download Play by Post DnD package from [releases page](https://github.com/betorvs/playbypost-dnd/releases) and uncompact it. 

```bash
./playbypost-dnd
{"level":"info","timestamp":"2021-09-17T10:36:56.720+0200","caller":"customlog/logger.go:25","msg":"JSONs Database ready"}
{"level":"info","timestamp":"2021-09-17T10:36:56.721+0200","caller":"customlog/logger.go:25","msg":"D20 are ready"}
{"level":"info","timestamp":"2021-09-17T10:36:56.722+0200","caller":"customlog/logger.go:25","msg":"MongoDB ready"}

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.5.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```

Then redirect your ngrok to use that port:
```bash
./ngrok http 8080
```

Use https endpoint from ngrok command to connect with Slack.