# observe

[![CircleCI](https://circleci.com/gh/dominikbraun/observe.svg?style=shield)](https://circleci.com/gh/dominikbraun/observe)
[![Go Report Card](https://goreportcard.com/badge/github.com/dominikbraun/observe)](https://goreportcard.com/report/github.com/dominikbraun/observe)
[![GitHub release](https://img.shields.io/github/v/release/dominikbraun/observe?include_prereleases&sort=semver)](https://github.com/dominikbraun/observe/releases)
[![License](https://img.shields.io/github/license/dominikbraun/observe)](https://github.com/dominikbraun/observe/blob/master/LICENSE)

:mag: Observe a website and get an e-mail if something changes.

## Example

Let's assume you want to get notified when a website for an event has been updated. After creating or re-using a `settings.yml`
file, you can start an observation like so:

```shell script
$ observe website https://example-event.com .
```

That's it! You'll get notified via e-mail as soon as something changes.

## Getting started

### Prerequisites

Since observe uses SendGrid to send e-mails, you just have to [create a free account](https://signup.sendgrid.com/) and
[create an API key](https://app.sendgrid.com/settings/api_keys). After that, create a `settings.yml` file in a directory
of your choice and fill in appropriate data.

```yaml
mail:
  from: mail@example.com
  to: mail@example.com
sendgrid:
  key: My-API-Key
```

If your mail provider does not allow `from` being the same address as `to` for security reasons, you may use a fake
address or your second e-mail address as sender.

### Installation

**Linux/macOS:** Download the [latest release](https://github.com/dominikbraun/observe/releases) and move the binary into
a directory like `/usr/local/bin`. Make sure the directory is in your `PATH`.

**Windows:** Download the [latest release](https://github.com/dominikbraun/observe/releases), create a directory like
`C:\Program Files\observe` and copy the executable into it. [Add the directory to `Path`](https://www.computerhope.com/issues/ch000549.htm).

**Docker:** Run `docker image pull dominikbraun/observe` to get the Docker image.

### Starting an observation

**Linux/macOS/Windows:** Let's start an observation that checks the website every 10 seconds, where `.` is the path to
your settings file:

```shell script
$ observe website --interval 10 https://example.com .
```

**Docker:** Mount the directory that contains your `settings.yml` file onto the container.

```shell script
$ docker container run -v /path/to/settings:/settings dominikbraun/observe website --interval 10 https://example.com .
```

You can run any observe sub-command by appending it to the image name.

## Disclaimer

I don't take any responsibility for misuse of this tool, even though the lookup interval is limited to 1 second.