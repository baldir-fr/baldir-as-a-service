# Tooling to organize ADRs

* Date: 2022-05-06

## Context and Problem Statement
ADRs can be hard to manage without tooling. Also, we need an overview of all the ADRs to have a global picture with their state.

## Decision Drivers
* Better if it can be collocated to git repository
* Cool if status is easily visible
* Must not be a pain to write ADRs

## Considered Options

* adr-tools
* log4brains
* obsidian with Kaban extension

## Decision Outcome

Chosen Obsidian with Kaban extension
- Visual
- Can drag and drop to change status
- Kaban board is markdown, so the list of ADRs is automatically readable in Github/Gitlab

### Positive Consequences

* Easy to manage status
* Can be readable on the go with Obsidian mobile client

### Negative Consequences

* Obsidian configuration may provoke conflicts in git or may require manual configuration on clients the first time
* Status is not visible in the ADR itself but on the index page
* Wikilinks between Adrs may not be exploitable outside of Obsidian
## Links

* https://obsidian.md/
* https://github.com/mgmeyers/obsidian-kanban