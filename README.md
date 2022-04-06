# MB/BALDIR CLI

Query stuff about me from an API / CLI.
Inspired by CLIs such as `docker` `git` `top`.

Purpose is to showcase my abilities and also be of use for geeks.

## Usage

`mb` or `baldir`

## Install

## Commands

`mb top`

Display ongoing stuff.
- Assigned tasks in various github/gitlab projects
- Current workplace (from json resume)
- Current volunteering
- ...
```
mb top

top - 20:29:35 up 2 min,  0 users,  load average: 0.00, 0.00, 0.00
Tasks:   6 total,   1 running,   5 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.0 us,  0.0 sy,  0.0 ni, 99.9 id,  0.1 wa,  0.0 hi,  0.0 si,  0.0 st
MiB Mem :  12735.6 total,  12524.8 free,    102.2 used,    108.6 buff/cache
MiB Swap:   4096.0 total,   4096.0 free,      0.0 used.  12431.0 avail Mem

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND
    1 root      20   0     896    528    464 S   0.0   0.0   0:00.01 init
   10 root      20   0     896     84     20 S   0.0   0.0   0:00.00 init
   11 root      20   0     896     84     20 S   0.0   0.0   0:00.00 init
   12 marco     20   0   12580   6208   4164 S   0.0   0.0   0:00.11 zsh
   62 marco     20   0   12412   4292   2260 S   0.0   0.0   0:00.00 zsh
  120 marco     20   0   10876   3732   3220 R   0.0   0.0   0:00.00 top

```

`mb search`

Search content in various locations I own
- Blog posts
- Side projects
- Resume
- ...


`mb about`

Quick Bio

`mb book`

Display and allow to book interactively a meeting from available spots in my calendar.
Requires to be authenticated.
Returns a link to an ical.

`mb book yyyy-mm-dd`

Book at the first available slot in the calendar for a given date.
Returns a link to an ical.

`mb login`

Login to be able to do some actions (like booking)

`mb meta`

Metadata about the tool itself

`mb meta stack`

Technologies used in this tool

`mb meta adr` (ls)

List architecture decision records about the tool

`mb meta adr #adr-number`

Show details of a peculiar ADR

`mb meta stats`

Various stats about the tool.
- github stars
- releases
- forks