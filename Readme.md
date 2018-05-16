The project is to create a new branch and exclude only the commits from a specific feature,
if you follow a commit pattern to add the features id, you will be able to remove them from
the `release` branch by just adding the id in the regex field.

To execute:

```
curl -X POST \
-H  'Content-Type: application/json' \
 -d '{"gitUrl":"gitToBeCherryPick", "regex": "regex","branchName": "{branchName}"}' \
localhost:1323/cherry-pick
```

