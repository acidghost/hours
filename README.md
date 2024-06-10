# hours

✨ Overview
---

`hours` is a CLI app that allows you to track time on tasks you care about.

It is intended for users who want to do some sort of time tracking for
their projects, but don't want to use an overly complicated app or website to do
so. It has a simple and minimalistic UI; almost everything in `hours` can be
achieved with one or two keypresses.

💾 Install
---

**go**:

```sh
go install github.com/dhth/hours@latest
```

⚡️ Usage
---

```
Usage:
  hours [flags] [command]

Flags:
  -db-path string
        location where hours should create its DB file (default "/Users/dhruvthakur/hours.v1.db")

Commands:
  report
        outputs a report of time spent on tasks
  logreport
        outputs a report of recently added log entries
  active
        shows the task currently being tracked
```
