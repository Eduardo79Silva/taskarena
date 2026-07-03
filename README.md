# taskarena

A minimal push/pull task manager for Linux, built for keyboard-driven workflows on Hyprland/Omarchy.

The idea is simple: push tasks into an arena whenever they come to mind, pull one when you're ready to work. The arena picks your next task using a weighted random algorithm biased by priority and time estimate, so high-priority tasks surface more often, but nothing ever gets permanently buried.

## How it works

**Push** a task into the arena at any time:

```
taskarena push -n "Write unit tests" -p high -t 30
```

**Pull** the next task when you're ready, the arena selects one probabilistically based on priority and time estimate:

```
taskarena pull
```

**Done** the current task. Clear it from the current.json and put it in completed.json.

```
taskarena done
```

## Installation

### Prerequisites

- Go 1.21+

You need these if you want to integrate with hyprland:

- [`walker`](https://github.com/abenz1267/walker) (for the interactive push/pull shell scripts)
- `notify-send` (provided by `libnotify`)

### Build from source

```bash
git clone https://github.com/Eduardo79Silva/taskarena.git
cd taskarena
go build  .
```

### Deploy

Make sure the scripts have the executable bit set.

Copy the binary and helper scripts to `~/.local/bin`:

```bash
./deploy
```

To also wipe your current task list:

```bash
./deploy --clean
```

Make sure `~/.local/bin` is in your `$PATH`.

## Usage

### Commands

#### `push`

Add a task to the arena.

```
taskarena push -n <name> [-d <description>] [-p <priority>] [-t <time_estimate>]
```

| Flag         | Shorthand | Default  | Description                         |
| ------------ | --------- | -------- | ----------------------------------- |
| `--name`     | `-n`      | required | Task name                           |
| `--desc`     | `-d`      | `""`     | Optional description                |
| `--priority` | `-p`      | `medium` | `low`, `medium`, `high`, `veryhigh` |
| `--time`     | `-t`      | `25`     | Time estimate in minutes            |

#### `pull`

Pull the next task from the arena.

```
taskarena pull
```

The selection algorithm scores each pending task using a weighted sum of priority (70%) and inverse time estimate (30%), then applies a sharpness exponent to increase the spread between scores before doing weighted random sampling. Higher priority tasks surface more often — but lower priority tasks can still win.

## Hyprland integration

Two helper scripts are included for keybind integration: `task-push` and `task-pull`.

`task-push` opens a `walker` dmenu prompt sequence to fill in task name, description, priority, and time estimate interactively, then calls `taskarena push` and sends a confirmation notification.

`task-pull` calls `taskarena pull` and surfaces the selected task via `notify-send`.

Add to your Hyprland config (`~/.config/hypr/bindings.conf` or equivalent):

```
bind = SUPER CTRL, T, exec, task-push
bind = SUPER CTRL, P, exec, task-pull
```

## Data

Tasks are stored as a JSON array at:

```
~/.config/taskarena/tasks.json
```

## Contributing

Contributions are welcome. The project is intentionally small, the goal is to keep it that way.

If you want to contribute, a few areas that could use thought:

- An `info` subcommand to get statistics about current, completed and backlog tasks
- A `list` subcommand to inspect the current arena
- A `pause` subcommand that puts the current task in the backlog
- A `snooze` subcommand that puts the current task in the backlog and makes sure it is not picked up for the rest of the day
- A `status` subcommand outputting Waybar-compatible JSON for a current-task module
- A `--fits <minutes>` flag on `pull` to filter the pool to tasks that fit a given time slot
- A way to edit existing tasks without manually manipulating tasks.json
- Configurable scoring weights via `~/.config/taskarena/config.toml` among other configs
- Save the time worked on a task
- Task aging so that neglected low-priority tasks gradually increase their selection weight
- Task tags to enable filtering

Please open an issue before starting work on anything substantial so we can align on direction. Keep PRs focused and the binary dependency-light.

## License

GPL-3.0
