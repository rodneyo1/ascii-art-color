# ASCII Art Color
```
                         _    _          _   _          
                        | |  | |        | | | |         
                        | |__| |   ___  | | | |   ___   
                        |  __  |  / _ \ | | | |  / _ \  
                        | |  | | |  __/ | | | | | (_) | 
                        |_|  |_|  \___| |_| |_|  \___/  
                                                                                      
```

This Go program generates ASCII art from a given string using specified fonts and colors. It includes various features such as color customization and font selection. This README provides details on how to use the program, including different ways to run it, and lists its limitations.

## Table of Contents
1. [Installation](#installation)
2. [Usage](#usage)
3. [Flags](#flags)
4. [Examples](#examples)
5. [Invalid Color Handling](#invalid-color-handling)
6. [Limitations](#limitations)

## Installation

### Prerequisites

- **Go**: This program requires Go to be installed on your system. If you don't have Go installed, follow the steps below.

### Installing Go

1. **Download Go**: Visit the [official Go downloads page](https://golang.org/dl/) and download the installer for your operating system.

2. **Install Go**:
    - **Windows**: Run the `.msi` installer and follow the prompts.
    - **macOS**: Open the `.pkg` file and follow the prompts.
    - **Linux**:
      ```sh
      wget https://dl.google.com/go/go1.16.5.linux-amd64.tar.gz
      sudo tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
      export PATH=$PATH:/usr/local/go/bin
      ```

3. **Verify Installation**: Open a terminal and run the following command to verify that Go is installed:
    ```sh
    go version
    ```
    You should see the Go version information.

### Setting Up the ASCII Art Generator

1. **Clone the repository** (if applicable) or download the source code files.
2. **Navigate to the project directory**:
    ```sh
    cd path/to/project
    ```

## Usage

The program can be run using different command-line arguments and options. Here are the general usage instructions:

```sh
go run . [OPTION] [STRING]
```

### Flags

- `--color`: Specifies the color of the text.
- `--help`: Displays usage instructions.

### Color Options

- Predefined colors: `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`, `orange`.
- Hexadecimal format: `#RRGGBB` (e.g., `#FF5733`).
- RGB format: `rgb(R, G, B)` (e.g., `rgb(255,87,51)`).

## Examples

1. **Basic usage**:

```sh
go run . "Hello, World!"
```

2. **Using color**:

```sh
go run . --color=red "Hello, World!"
```

3. **Using hex color**:

```sh
go run . --color=#FF5733 "Hello, World!"
```

4. **Using RGB color**:

```sh
go run . --color="rgb(255,87,51)" "Hello, World!"
```

5. **Coloring specific characters**:

```sh
go run . --color=blue H "Hello, World!"
```

6. **Using different fonts**:

```sh
go run . --color=green "Hello, World!" standard
go run . --color=green "Hello, World!" shadow
go run . --color=green "Hello, World!" thinkertoy
```

7. **Coloring specific characters with different fonts**:

```sh
go run . --color=yellow l "Hello, World!" standard
go run . --color=yellow l "Hello, World!" shadow
go run . --color=yellow l "Hello, World!" thinkertoy
```

## Invalid Color Handling

The program has built-in error handling for invalid color inputs. Here's how it handles different invalid color scenarios:

1. **RGB values larger than 255**:
    - The program ignores the invalid RGB values and defaults to no color.
    - Example:
      ```sh
      go run . --color="rgb(300,400,500)" "Hello, World!"
      ```
    - Output: The text will be displayed without any color.

2. **Invalid hex color**:
    - The program ignores the invalid hex value and defaults to no color.
    - Example:
      ```sh
      go run . --color="#ZZZZZZ" "Hello, World!"
      ```
    - Output: The text will be displayed without any color.

3. **Predefined color not present**:
    - The program ignores the unrecognized predefined color and defaults to no color.
    - Example:
      ```sh
      go run . --color="unknowncolor" "Hello, World!"
      ```
    - Output: The text will be displayed without any color.

## Limitations

1. **File Availability**: The program relies on font files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`). These files must be present in the same directory as the executable or source code.
2. **Color Support**: The terminal must support ANSI escape codes for color formatting to work correctly.
3. **Character Set**: The program only supports characters present in the font files, typically limited to standard ASCII characters (32-126).

## Conclusion


Feel free to modify and enhance the program as needed. If you encounter any issues or have suggestions for improvement, please submit an issue or pull request on the project's repository. Happy ASCII art creation!