# Please ensure you adhere to every item in this list

+ The PR title is formatted as follows: `<type>[optional scope]: <description>` For example, `fix(os/gtime): fix time zone issue`
  + `<type>` is mandatory and can be one of `fix`, `feat`, `build`, `ci`, `docs`, `style`, `refactor`, `perf`, `test`, `chore`
    + fix: Used when a bug has been fixed.
    + feat: Used when a new feature has been added.
    + build: Used for modifications to the project build system, such as changes to dependencies, external interfaces, or upgrading Node version.
    + ci: Used for modifications to continuous integration processes, such as changes to Travis, Jenkins workflow configurations.
    + docs: Used for modifications to documentation, such as changes to README files, API documentation, etc.
    + style: Used for changes to code style, such as adjustments to indentation, spaces, blank lines, etc.
    + refactor: Used for code refactoring, such as changes to code structure, variable names, function names, without altering functionality.
    + perf: Used for performance optimization, such as improving code performance, reducing memory usage, etc.
    + test: Used for modifications to test cases, such as adding, deleting, or modifying test cases for code.
    + chore: Used for modifications to non-business-related code, such as changes to build processes or tool configurations.
  + After `<type>`, specify the affected package name or scope in parentheses, for example, `(os/gtime)`.
  + The part after the colon uses the verb tense + phrase that completes the blank in
  + Lowercase verb after the colon
  + No trailing period
  + Keep the title as short as possible. ideally under 76 characters or shorter
  + [Reference Documentation](https://www.conventionalcommits.org/en/v1.0.0/)
+ If there is a corresponding issue, add either `Fixes #1234` or `Updates #1234`
  (the latter if this is not a complete fix) to this comment
+ Delete these instructions once you have read and applied them
