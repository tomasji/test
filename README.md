# Color occurrences

## Task
- there are XML files in the data directory, each having the same structure: `competitor_profile / jerseys / jersey`
- complete the `cmd/xmltool` program so that it:
  - reads and parses the XML files in the data directory
  - calculates the number of occurrences of given color in the XML files (attribute `colors` in the `jersey` tag may contain 0, 1 or more colors separated by `|`)
- the `TestCount` test shall pass (confirming there are 1000 occurrences of 'blue')
- think about how the program could be improved in terms of
  - performance
    - concurrent design patterns, allocations, batch processing ...
  - error handling
  - other


