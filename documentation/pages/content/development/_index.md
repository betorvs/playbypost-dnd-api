---
title: "Development"
date: 2021-08-25T14:03:39+02:00
draft: false
weight: 4
---


Play by Post Dungeons and Dragons should not be a replacement for any platform for role playing games. This was designed to be a good tool to help some friends to have fun moments between sessions. 

Play by Post Dungeons and Dragons have 3 main components:  
- CRUD API to create campaigns, adventures, encounters and register players (requires MongoDB).   
- Bot integration (Requires Ngrok and a chat account).   
- Rules mechanisms (can run without the others).   


## Postman Collection

We shared [here](https://github.com/betorvs/playbypost-dnd/blob/main/documentation/postman-collection/PlayByPost-DnD.postman_collection.json) a postman collection with all REST API. Just need to create a environment variable inside postman called server pointing to your local environment. For instance: `server = http://localhost:8080`.


## TODO List

- [ ] Add Hoard Calcutation for levels: 5-10, 11-16, >17.
- [ ] Finish creating all Core Features and Core Powers and migrate race and class features to use this centralised function.
  - [ ] Re write Use Potion usecase to use Power
  - [ ] Re write character related code to use Core Features
- [ ] Migrate Background, Race, Skills, Class to use struct
  - [ ] Create a extended database to allow adding customised Background, Race, Skills and Class.
- [ ] Add Slack Controller, usecase and gateway.
- [ ] Add rest client code in gateway to be used for e2e tests 
  - [ ] End to End test with github actions.
- [ ] Create a turn control for Encounters
- [ ] Add goreleaser to create the packages with binaries and json files

