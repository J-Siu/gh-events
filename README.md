# gh-events

Github api `users/{user}/received_events` listing.

## Usage

### gh cli

```sh
gh events
```

### Standalone

```sh
List Github api 'users/<USER>/received_events' output.

Usage:
  gh-events [flags]
  gh-events [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dump        Dump raw json of Github api 'users/<USER>/received_events'.
  help        Help about any command

Flags:
  -a, --all           show skipped event
  -c, --create-time   show create time
  -h, --help          help for gh-events
  -t, --type          show event type
  -u, --url           show full url
  -v, --version       version for gh-events

Use "gh-events [command] --help" for more information about a command.
```
