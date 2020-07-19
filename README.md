![mizukinana](https://user-images.githubusercontent.com/667169/45166942-b2076600-b22a-11e8-97ba-5af903d24b09.png)

<p align=center>
<a target="_blank" href="https://opensource.org/licenses/MIT" title="License: MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg"></a>
<a target="_blank" href="http://makeapullrequest.com" title="PRs Welcome"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg"></a>
<a href="https://codecov.io/gh/mkfsn/mizukinana">
  <img src="https://codecov.io/gh/mkfsn/mizukinana/branch/develop/graph/badge.svg" />
</a>
</p>  

# mizukinana

It's all about Mizuki Nana（水樹奈々）

# Documentation

See [godoc.org](https://godoc.org/github.com/mkfsn/mizukinana)

# Features

- [x] Provide a collection of resources from [Nana Party](https://www.mizukinana.jp/)
  - [x] List all News
  - [x] List all Schedules
  - [ ] List All Discographies (albums, singles, ... etc)
  - [x] List all Blog posts and details
  - [x] Collection of the Biographies (profile, voice, live, special event, others)
  - [x] Collection of the main page resources
- [x] List All concerts, in some format: table, json, and yaml.
- [ ] List All Set-List
- [ ] Run as a REST API service
- [ ] Show the latest tweets from Twitter

# Command Line Tool

## How To Install

```mizukinana
go install github.com/mkfsn/mizukinana/cmd/mizukinana
```

## Profile

> Warning: This is outdated

```bash
mizukinana profile [-p {table|yaml|json}]
```

![profile](https://user-images.githubusercontent.com/667169/45159966-ce9ba200-b21a-11e8-9e08-72902fd36ff8.gif)


## Concerts

```bash
mizukinana concerts [-o {table|yaml|json}] [-f FILTER]
```

![concerts](https://user-images.githubusercontent.com/667169/45159731-2dace700-b21a-11e8-8425-1ab37cf91b0d.gif)

![concerts-with-filter](https://user-images.githubusercontent.com/667169/45159647-ee7e9600-b219-11e8-9bba-3bdbce2d9e77.gif)
