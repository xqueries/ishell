# ishell
ishell - interactive shell, is a customised shell that allows us to power the xdb commands.

## Motivation
When we were trying to use third party solutions for having a shell that parses commands and executes them, it wasn't a straightforward task. Thus we decided to use existing solutions and create a solution of our own that behaves how we want it to.

## Final goal
Existing solutions like [cobra](https://github.com/spf13/cobra) can only parse commands/arguments/flags but once we start a terminal, control can't be handed over to it. A CUI can be created using [promptui](https://github.com/manifoldco/promptui) but it doesn't have a command/argument/flag parser built within.

Our final goal should be something where we have can plug in both solution like cobra and promptui and the main engine uses solutions like these to have a final CUI that is powerful enough to parse commands/arguments/flags and also run them appropriately! Auto-completion for commands is also expected.

### Contributing

We're happy if you contribute!
Currently this is the very beginning of the project, we'd love to hear things that'd make this better :D
