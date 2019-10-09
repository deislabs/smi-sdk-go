# Release Process

This document describes the release process for this project. Releases are published via GitHub releases.

All releases will have a [semantic version](https://semver.org/) associated with them. This project will be versioned at 1.0.0 once the Service Mesh Interface(SMI) specficication is at 1.0.0 and all APIs are at `v1`. All releases until then will be `0.x.0` where `x` is incremented with each release.

To perform a release of the smi-sdk-go project:

1. Start by creating and pushing a git tag in the form: `v0.x.0`.
```console
$ git tag -a v0.1.0 -m "version 0.1.0"
$ git push origin v0.1.0
```
2. Then, generate all assets to upload with the git release:
```console
$ make dist
```

3. Last, visit the [releases page](https://github.com/deislabs/smi-sdk-go/releases) to `Draft a new release` using the tag you just created and pushed. Be sure to include release notes on what changes are included in the release and upload the assets created in the previous step.
