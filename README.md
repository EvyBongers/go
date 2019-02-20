# About

Tiny URL shortener for personal use

# Dependencies

This project is built using [bazel](https://bazel.build/). Check the [getting started page](https://docs.bazel.build/versions/master/getting-started.html) for installation instructions.

# Setup

The provided deployment script assumes that a file `redirects.json` with URL mappings is present in the deployments directory. The file contains a mapping of short URLs to destination URLs.

```
{
  "gh": "https://github.com/",
  "gl": "https://gitlab.com/",
  "hn": "https://news.ycombinator.com/"
}
```

Add a short entry (e.g.: go) to your system's `hosts` file to access the URL shortener.

```
::1		go
```

# Usage

A `docker-compose.yaml` is provided in the deployments directory, as well as a helper script `run.sh`. Simply invoking `run.sh` is the easiest way to get started.

Use the URL shortener by simply navigating to `go/<alias>` in your browser.

Note that the redirects file is only read on startup. If you change/add/remove redirects, a restart is required for them to come in effect.
