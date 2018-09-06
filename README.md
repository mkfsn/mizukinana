<p align=center>
<a target="_blank" href="https://opensource.org/licenses/MIT" title="License: MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg"></a>
<a target="_blank" href="http://makeapullrequest.com" title="PRs Welcome"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg"></a>
</p>  

# mizukinana

It's all about Mizuki Nana（水樹奈々）

# Features

- [x] List All concerts, in some format: table, json, and yaml.
- [x] Display Profile
- [ ] List All Set-List
- [ ] List All Discographies (albums, singles, ... etc)
- [ ] Run as a REST API service
- [ ] Show the latest tweets from Twitter
- [ ] Show the latest schedule from the official website

# How To Install

```mizukinana
go install github.com/mkfsn/mizukinana
```

# Profile

```bash
mizukinana profile [-p {table|yaml|json}]
```

![profile](https://user-images.githubusercontent.com/667169/45159966-ce9ba200-b21a-11e8-9e08-72902fd36ff8.gif)


# Concerts

```bash
mizukinana concerts [-o {table|yaml|json}] [-f FILTER]
```

![concerts](https://user-images.githubusercontent.com/667169/45159731-2dace700-b21a-11e8-8425-1ab37cf91b0d.gif)

![concerts-with-filter](https://user-images.githubusercontent.com/667169/45159647-ee7e9600-b219-11e8-9bba-3bdbce2d9e77.gif)
