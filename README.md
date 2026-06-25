# sul

This is just a small collection of helpers/utility functions that may be used in various GOLANG projects. 

Are covered : 

- **strings helpers** 
- **files helpers**
- **slices/map helpers**
- **misc helpers**


## DEV Activities

### Push a new version

```bash
TAG="v0.0.7"
git add .
git commit -m"Preparing tag ${TAG}"
git push origin :refs/tags/${TAG}
git tag -f ${TAG}
git push origin --tags 
git push origin master
```

### TODO

- [ ] Write a more complete README.md (examples, ...)
- [ ] Split the content in buckets/sub-packages (to allow more fine-tuned imports)
- [ ] Recover associated lost tests (kept in original source projets for now)
- [ ] Put in place a Makefile
