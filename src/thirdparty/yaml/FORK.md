This package is a fork of the [gopkg.in/yaml.v3](https://github.com/go-yaml/yaml/tree/v3) library, as of
commit [496545a6307b2a7d7a710fd516e5e16e8ab62dbc](https://github.com/go-yaml/yaml/commit/496545a6307b2a7d7a710fd516e5e16e8ab62dbc)

# Changes

Search for `FORK` in this directory to find forked code.

## Support for `json` struct tag

The only change in the fork is adding support for `json` struct tags:

```text
548a548,550
> 		if tag == "" {
> 			tag = field.Tag.Get("json")
> 		}
```

## Disable block encoding of scalars

As of the most recent commit, yaml.v3 has an encoding bug where scalars may not encode correctly in block form. This
fork disables block encoding of scalars as a result.

[This issue](https://github.com/go-yaml/yaml/issues/643) and
[unmerged yaml PR](https://github.com/go-yaml/yaml/pull/640) have details.
