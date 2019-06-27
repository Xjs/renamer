# renamer
Read source and target filenames from files and rename files correspondingly

## Usage

```
$ renamer listA listB
```

`renamer` expects listA to contain a list of existing file names (relative to renamer's working directory) and listB to contain a list of target names.
`renamer` will attempt to rename each entry from listA to the corresponding entry (based on line numbers) in listB.
