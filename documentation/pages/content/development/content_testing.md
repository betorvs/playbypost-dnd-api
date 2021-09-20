---
title: "How to help to test Dungeons and Dragons rules"
date: 2021-08-26T11:03:26+02:00
draft: false 
weight: 10
---


## Download PlaybyPost-DnD 

Download Play by Post DnD package from [releases page](https://github.com/betorvs/playbypost-dnd/releases) and uncompact it. 

## Install Postman

Install [Postman](https://www.postman.com/downloads/) and download our collection [here](https://github.com/betorvs/playbypost-dnd/blob/main/documentation/postman-collection/PlayByPost-DnD.postman_collection.json).

Configure Postman with an environment variable `server`. For instance: `server = http://localhost:8080`.

Change the values inside Postman and test any kind of resources from json files. Any issue, please report [here](https://github.com/betorvs/playbypost-dnd/issues) for any mistake in the mechanics of the game and report any wrong content [here](https://github.com/betorvs/playbypost-jsonparser/issues).