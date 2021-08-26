---
title: "Update Documentation"
date: 2021-08-26T11:03:26+02:00
draft: false 
weight: 9
---

## Setup Hugo

Install [hugo](https://gohugo.io/getting-started/installing/) and execute the following commands

```bash
cd documentation/pages
hugo new site playbypost-dnd-documentation
cd themes
git clone https://github.com/thingsym/hugo-theme-techdoc.git
```

## Execute it local

```bash
hugo server -D
```

## Create new pages

```bash
hugo new development/local-environment.md
```

Add `weight: X` and change `draft` to `false`.