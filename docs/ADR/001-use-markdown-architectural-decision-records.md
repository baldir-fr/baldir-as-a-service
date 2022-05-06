# Use Markdown Architectural Decision Records

- Date: 2022-05-06

## Context and Problem Statement

Architecture decisions can be hard to track and the context of when they were done is not always documented.

## Considered Options

* ADR
* Rely on git history
* Document in a structured document (Ms Word)

## Decision Outcome

Chosen ADR because
- Implicit assumptions should be made explicit. Design documentation is important to enable people understanding the decisions later on.
- Lives close to the code (in the same repository)
- Can be read without specific tooling (plain markdown)
- Is versionned
- 
### Positive Consequences->

* Improved onboarding
* Improved documentation
* Less context is lost
* Can be discussed in merge requests

### Negative Consequences->

* Requires more discipline in day to day work
* Lots of manual tasks


## Links

* https://adr.github.io/
* https://github.com/thomvaill/log4brains/blob/master/docs/adr/20200924-use-markdown-architectural-decision-records.md