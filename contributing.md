# Welcome to universe contributing guide

In this guide you will get an overview of the contribution workflow from opening an issue, creating a PR, reviewing, and merging the PR.

## Issues

### Create a new issue

- [Check to make sure](https://docs.github.com/en/github/searching-for-information-on-github/searching-on-github/searching-issues-and-pull-requests#search-by-the-title-body-or-comments) someone hasn't already opened a similar [issue](https://github.com/1layar/universe/issues).
- If a similar issue doesn't exist, open a new issue using a relevant [issue form](https://github.com/1layar/universe/issues/new/choose).

### Pick up an issue to solve

- Scan through our [existing issues](https://github.com/1layar/universe/issues) to find one that interests you.
- The [👋 good first issue](https://github.com/1layar/universe/issues?q=is%3Aissue+is%3Aopen+label%3A%22%F0%9F%91%8B+good+first+issue%22) is a good place to start exploring issues that are well-groomed for newcomers.
- Do not hesitate to ask for more details or clarifying questions on the issue!
- Communicate on the issue you are intended to pick up _before_ starting working on it.
- Every issue that gets picked up will have an expected timeline for the implementation, the issue may be reassigned after the expected timeline. Please be responsible and proactive on the communication 🙇‍♂️

## Pull requests

When you're finished with the changes, create a pull request, or a series of pull requests if necessary.

Contributing to another codebase is not as simple as code changes, it is also about contributing influence to the design. Therefore, we kindly ask you that:

- Please acknowledge that no pull request is guaranteed to be merged.
- Please always do a self-review before requesting reviews from others.
- Please expect code review to be strict and may have multiple rounds.
- Please make self-contained incremental changes, pull requests with huge diff may be rejected for review.
- Please use English in code comments and docstring.
- Please do not force push unless absolutely necessary. Force pushes make review much harder in multiple rounds.

## Raising Pull Requests

- Please resolve linting and formatting issue produced by `golangci-lint run`.
- Please keep the PR's small and focused on one thing
- Please follow the format of creating branches
- feature/[feature name]: This branch should contain changes for a specific feature
    - Example: feature/email-service
- bugfix/[bug name]: This branch should contain only bug fixes for a specific bug
  - Example bugfix/db-connection