---
title: "Development"
date: 2021-08-25T14:03:39+02:00
draft: false
weight: 4
---


Play by Post Dungeons and Dragons should not be a replacement for any platform for role playing games. This was designed to be a good tool to help some friends to have fun moments between sessions. 




## Postman Collection

We shared [here](https://github.com/betorvs/playbypost-dnd/blob/main/documentation/postman-collection/PlayByPost-DnD.postman_collection.json) a postman collection with all REST API. Just need to create a environment variable inside postman called server pointing to your local environment. For instance: `server = http://localhost:8080`.


## TODO List

- [ ] Add Hoard Calcutation for levels: 5-10, 11-16, >17.
- [ ] Finish creating all Core Features and Core Powers and migrate all features to use this centralised function.
  - [ ] Re write Use Potion usecase to use Power
- [ ] Add Slack Controller, usecase and gateway.
- [ ] End to End test with github actions.

