# filter

[![Build Status](https://img.shields.io/travis/milgradesec/filter/master.svg?label=build)](https://travis-ci.org/milgradesec/filter)
[![Go Report Card](https://goreportcard.com/badge/milgradesec/filter)](https://goreportcard.com/badge/github.com/milgradesec/filter)

## Name

_filter_ - enables blocking requests based on lists and rules.

## Description

## Syntax

## Features

- Regex and simple string matching.
- Detects CNAME cloacking.
- Responses compatible with negative caching.

## Metrics

If monitoring is enabled (via the _prometheus_ plugin) then the following metric are exported:

- `coredns_filter_blocked_requests_total{server}` - count per server

## Examples

```corefile
.:53 {
    filter {
        allow ./lists/whitelist.txt
        block ./lists/blacklist.txt
        uncloak
    }
    forward . 1.1.1.1
}
```
