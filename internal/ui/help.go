package ui

import "fmt"

var (
	helpText = fmt.Sprintf(`
  %s
%s
  %s

  %s
%s
  %s
%s
  %s
%s
  %s
%s
  %s
%s
  %s
%s
`,
		helpHeaderStyle.Render("\"hours\" Reference Manual"),
		helpSectionStyle.Render(`
  (scroll line by line with j/k/arrow keys or by half a page with <c-d>/<c-u>)

  "hours" has 5 panes:
    - Tasks List View                      Shows your tasks
    - Task Management View                 Allows you to create/update tasks
    - Task Log List View                   Shows your task log entries
    - Inactive Tasks List View             Shows inactive tasks
    - Help View (this one)
`),
		helpHeaderStyle.Render("Keyboard Shortcuts"),
		helpHeaderStyle.Render("General"),
		helpSectionStyle.Render(`
    1                                       Switch to Tasks List View
    2                                       Switch to Task Log List View
    3                                       Switch to Inactive Task Log List View
    <tab>                                   Go to next view/form entry
    <shift+tab>                             Go to previous view/form entry
      ?                                     Show help view
`),
		helpHeaderStyle.Render("General List Controls"),
		helpSectionStyle.Render(`
    h/<Up>                                  Move cursor up
    k/<Down>                                Move cursor down
    h<Left>                                 Go to previous page
    l<Right>                                Go to next page
    <ctrl+r>                                Refresh list
`),
		helpHeaderStyle.Render("Task List View"),
		helpSectionStyle.Render(`
    a                                       Add a task
    u                                       Update task details
    s                                       Toggle recording time on the currently selected task,
                                                will open up a form to record a comment on the
                                                second "s" keypress
    <ctrl+s>                                Add a manual task log entry
    <ctrl+t>                                Go to currently tracked item
    <ctrl+d>                                Deactivate task
`),
		helpHeaderStyle.Render("Task Log List View"),
		helpSectionStyle.Render(`
    <ctrl+d>                                Delete task log entry
`),
		helpHeaderStyle.Render("Inactive Task List View"),
		helpSectionStyle.Render(`
    <ctrl+d>                                Activate task
`),
		helpHeaderStyle.Render("Task Log Entry View"),
		helpSectionStyle.Render(`
    enter                                   Save task log entry
`),
	)
)
