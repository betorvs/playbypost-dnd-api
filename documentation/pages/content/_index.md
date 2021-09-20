---
title: "Home"
date: 2021-08-25T15:30:08+02:00
draft: false
weight: 1
---

"A play-by-post role-playing game (or sim) is an online text-based role-playing game in which players interact with each other and a predefined environment via text. It is a subset of the online role-playing community which caters to both gamers and creative writers. Play-by-post games may be based on other role-playing games, non-game fiction including books, television and movies, or original settings. This activity is closely related to both interactive fiction and collaborative writing. Compared to other roleplaying game formats, this type tends to have the loosest rulesets." - [wikipedia](https://en.wikipedia.org/wiki/Play-by-post_role-playing_game)

Play by Post Dungeons and Dragons is design to be a bot with REST API created to help friends to have fun together without thinking about rulesets. Anyone should be able to start this service and connect it to a chat and start having fun with yours friends. 

Goals:   
[X] Easy to start and connect to chat.  
[X] Asynchronous format. It means all actions can be taken by each player in it own time without needed to have all players online.  
[X] Possibility to add custom resources as magic items, backgrounds, skills, races and more.  


## How it works

Play by Post DnD requires a MongoDB account to keep all campaign data like encounters and players and a chat account. We right now have plans to add only [Slack](https://slack.com/), but we will add more chat options in the future. And because of it we need to have a slack account with admins rights.

After downloading our package and configure it. You need to start it and run ngrok to expose it throught internet (enabling access for Slack) and configure Slack to send messages to our playbypost bot.