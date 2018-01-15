# countdown-page
A single-file webserver that displays events on a schedule and a countdown until the next event

# Overview

`countdown-page` is a small Go program which, displays "events" (text or HTML) at a predetermined time. It also displays a countdown indicating the time remaining until the next event.

For example, if you wanted to show someone a notification on a website only after a certain time, this application could be configured to display a "Surprise will be here soon!" message and a countdown until that text will be replaced with the surprise in question.

When given an HTML page template and a JSON data file of a list of events (where each event is some text and a relative or absolute timestamp when it should be displayed), the countdown page application generates a single dependency-free binary for the specified target architecture. That binary will run a read-only webserver which serves a single page containing the current event and countdown until the next event, if any.

# Usage

```bash
# Make some changes to events and times
$EDITOR data.json
# Make some changes to how events are rendered
$EDITOR countdown.html 
# Build the server binary
./build.sh 
# Serve the countdown/event page at http://localhost:8081
./countdown-server -port 8081
```

`countdown-page` servers are configured and compiled/packaged in the same step. To make a server with customized events/times or HTML, you modify the required data/template files and compile the server from scratch. Cross-compilation is possible.

### Customizing a Server

There are two main inputs to the server which users can customize.

1. The `data.json` file contains an array of "events". Each event is a JSON object containing a `text` property, which represents the HTML content of the event, and a `time` property, representing when the event will occur. There are many valid formats for dates/times in `data.json`; the `data.json` included in this directory contains several examples of those formats. In general, the formats can be one of:
	- Any format parseable by the [`dateparse`](https://github.com/araddon/dateparse) library. This includes UNIX timestamps, local-timezone formats like `4/8/2014 22:05`, UTC formats, and formats with a specified timezone offset.
	- Durations in the future from build time. If a `time` value starts with a plus `+` character, the remainder of that value should be anything accepted by the [`ParseDuration`](https://golang.org/pkg/time/#ParseDuration) method, e.g. `+5h33m10s`. Durations are from *server build time*, not server start time; they are "frozen" into the generated server. Because of this, durations should primarily be used for testing; absolute dates should be used for servers that will run for a long time.
2. The `countdown.html` file is a template for the page it will display. You can modify the HTML content of the file as you see fit. You can add JavaScript as well, but the existing JavaScript at the end of the file should be left alone. The template string `{{ .Text }}` will be rendered with the text, verbatim (no escaping) of each event. The template string `{{ .Time }}` should not be used anywhere other than the "countdown" `div`, as its format may change.

When `data.json` is customized, events in the JSON array will be rendered by the server. When a built server is started, it will begin serving the text the event *preceeding* the *first* event in the array whose timestamp is still in the future. The countdown until the event after the one being displayed will also be rendered into the page by default. A time of 0 can be used to set a default event to show at startup.

For example, consider the following `data.json`:

```json
[
	{ "text": "event 1", "time": 0 },
	{ "text": "event 2", "time": "+1h" },
	{ "text": "event 3", "time": "12 Feb 2026, 01:00" }
]
```

If the server generated from that file is started immediately, it will display "event 1" and a countdown of one hour, after which it will display "event 2" and a countdown until 1AM (in the local time zone where the server is built) on February 12 2026, if that date is still in the future. If that server is not started until an hour has elapsed, it will just display "event 2" and the countdown until 1AM on Feb 12 '26.

### Building a Server

To build a server, ensure you have satisfied the requirements listed below, and run `./build.sh`.

If successful, that script will return 0 and place the generated server program at `countdown-server` in the directory where `build.sh` was invoked.

The build script will output the parsed messages and dates from `data.json` as the generated server will understand them. Users should verify that the dates output by `build.sh` are what they expect, since time zones and date format parsing can often yield unexpected results.

Cross-compilation can be performed by supplying 1 or 2 arguments to `build.sh`. The first argument will be used to specify operating system and the second will be used to specify architecture (e.g. `./build.sh linux` or `./build.sh linux amd64`). OS and architecture names must be among [those supported by Golang](https://github.com/golang/go/blob/master/src/go/build/syslist.go).

Errors regarding `GOPATH` can be ignored if the server builds correctly, though it is recommended to install and build this project inside your `GOPATH` workspace.

The build/configuration script for `countdown-page` is nonstandard by Golang conventions out of a desire to keep the build process as simple as possible, since build and configuration are the same in this project's case.

### Running a Server

The `countdown-server` binary has no dependencies and can be executed to start a webserver to show the countdown page. It takes one argument: `-port`, which accepts an integer value of what port the server should run on. If none is supplied, `8080` is used.

During runtime, `countdown-server` will print out the text and UNIX timestamp of any responses it renders.

# Requirements

### For building the server

- POSIX-compliant shell or WSL shell.
- Golang 1.5+

### For running the server

- Host operating system [supported as a Golang compilation target](https://github.com/golang/go/blob/master/src/go/build/syslist.go).

### For viewing the page

- A modern browser supporting the ES6 standard of JavaScript. No external sites/CDN connectivity is required.

# FAQ

### Why configure and compile at the same time/why the single-file approach?

The "frozen" single file approach is nonstandard; most web servers will keep HTML and perhaps JSON separately from the application binary. Why was this done?

Two reasons:

1. Deployment simplicity. Deploying a single file is easier than deploying more than one, and reasoning about "what will this app render" is easier when you know that HTML/JS/data assets are frozen into the app and not editable after the application starts.
2. Additionally, the frozen approach means that relative dates (e.g. "display this event 5 hours in the future") won't re-relativize themselves if a server is restarted. This is sometimes a desirable property.

# Issues and Improvements

Please report issues via the [GitHub issue tracker](https://github.com/zbentley/gdb-inject-perl/issues) for this project.

### TODO

- Sorting/verification of the order in `data.json` by timestamp.
- Support for timestamps "at current date".
- Better back-compat support for older browsers (JS too new).
- Customizable route (not just `/`).
- Better "JFDWIM" date parser.
- Better time zone support.
- Build-time warnings if detected dates are in the past.
- Make HTML more customizable by moving JS etc. out of the main file and into code/constant.

# License

See `LICENSE` in the root of the source code repository for this project for licensing details.