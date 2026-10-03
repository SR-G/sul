# sul

This is just a small collection of helpers/utility functions that may be used in various GOLANG projects. 

Are covered (one import path per sub-package, `github.com/SR-G/sul/<name>`): 

- **strings** : string, UUID, humanize (duration/file size) helpers
- **files** : filesystem helpers
- **collections** : slices/maps helpers
- **net** : URL and HTML helpers, HTTP fetch
- **streams** : counting `io.Reader`/`io.Writer` wrappers
- **threads** : generic thread pool executor
- **table** : generic CLI table rendering
- **tests** : generic `Assert` for tests
- **sul** (root) : misc helpers (env, JSON, hostname) and `Version`

## DEV Activities

### Push a new version

```bash
TAG="v0.1.0"
git add .
git commit -m"Preparing tag ${TAG}"
git push origin :refs/tags/${TAG}
git tag -f ${TAG}
git push origin --tags 
git push origin master
```

### TODO

- [x] Split the content in buckets/sub-packages (to allow more fine-tuned imports)
- [ ] Write a more complete README.md (examples, ...)
- [ ] Recover associated lost tests (kept in original source projets for now)
- [ ] Put in place a Makefile
