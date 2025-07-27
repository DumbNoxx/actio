# Actio 📝

A ridiculously fast and elegant TUI for managing your tasks directly from the terminal. Stop messing with slow, bloated apps and get things done.

![Actio Demo](https://raw.githubusercontent.com/nxus-dev/actio/main/demo.gif)

## Features

-   **Create, edit, and delete tasks** with minimal keystrokes.
-   **Mark tasks as complete/pending** to track your progress.
-   **Vim-like keybindings** for efficient, mouse-free navigation.
-   **Persistent storage** so your tasks are saved between sessions.
-   **Clean and intuitive** interface that stays out of your way.

## Implementation Status

| Feature                                     | Status | Notes                                           |
| :------------------------------------------ | :----: | :---------------------------------------------- |
| **Core Functionality**                      |        |                                                 |
| Create, Edit, Delete Tasks                  |   ❌   | Basic functionality is stable.                  |
| Toggle Task Status (Complete/Pending)       |   ❌   | Uses `spacebar`.                                |
| Persistent Storage (e.g., local JSON file)  |   ❌   | Currently in progress. Saves state on exit.     |
| Vim Keybindings (j, k, etc.)                |   ❌   | Basic navigation implemented.                   |
|                                             |        |                                                 |
| **Enhancements**                            |        |                                                 |
| Add detailed task descriptions/notes        |   ❌   | Planned for v0.2.                               |
| Set task priorities (High, Medium, Low)     |   ❌   | Planned.                                        |
| Custom Theming / Colors                     |   ❌   | Future goal.                                    |
| Multiple Lists / Projects view              |   ❌   | Not started. A distant dream for you.           |

### Legend

-   `✅`: **Completed**
-   `⏳`: **In Progress**
-   `📝`: **Planned**
-   `❌`: **Not Started**

## Getting Started

### Prerequisites

-   [Go](https://golang.org/doc/install) (version 1.18 or higher)

### Installation

1.  Clone the repository:
    ```sh
    git clone https://github.com/your-username/actio.git
    ```
2.  Navigate to the project directory:
    ```sh
    cd actio
    ```
3.  Build the application:
    ```sh
    go build
    ```

## Usage

Run the application from your terminal:

```sh
./actio
```

### Keybindings

-   `j` / `↓` : Move down
-   `k` / `↑` : Move up
-   `n` : Create a new task
-   `e` : Edit the selected task
-   `d` : Delete the selected task
-   `space` : Toggle task status (complete/pending)
-   `q` / `ctrl+c` : Quit the application

## Technology Stack

-   **[Go](https://golang.org/)**: For the core application logic.
-   **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: For the terminal user interface.
-   **[Lipgloss](https://github.com/charmbracelet/lipgloss)**: For styling the TUI.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
