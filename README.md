# code-snippet-manager

The Code Snippet Manager is a command-line tool that allows you to manage and organize your code snippets efficiently. It enables you to store, retrieve, list, and remove code snippets with ease.

# Installation


# Prerequisites


Before installing the Code Snippet Manager, ensure you have the following prerequisites:

Go Programming Language (Golang) installed on your system.



# Install from Source
To install the Code Snippet Manager from source, follow these steps:

# Clone the repository:

``` bash

git clone https://github.com/mobster1425/code-snippet-manager.git
```
Change directory to the project folder:

``` bash

cd code-snippet-manager

```
Build the executable:

``` bash

go build

```


Optionally, add the executable to your system's PATH for easier access.

Download Executable (Windows and macOS)
Alternatively, you can download pre-built executables for Windows and macOS , and place the executable in a directory accessible from the command line.

# Usage
The Code Snippet Manager provides the following commands:

# Add a Snippet
# To add a new code snippet to the manager, use the add command:

``` bash

go-code-snippet-manager add "Snippet Name" "Category" "Your code snippet here"

```

# Get a Snippet
# To retrieve a code snippet from the manager, use the get command:

``` bash

go-code-snippet-manager get "Snippet Name" "Category"

```
# List Snippets
# To list all the stored code snippets, use the list command:

``` bash

go-code-snippet-manager list

```

# Remove a Snippet
# To remove a code snippet from the manager, use the remove command:

``` bash

go-code-snippet-manager remove "Snippet Name" "Category"

```

# Examples
# Adding a Snippet
``` bash

go-code-snippet-manager add "Hello World" "Printf" "fmt.Println(\"Hello, World!\")"

```
# Getting a Snippet
``` bash

go-code-snippet-manager get "Hello World" "Printf"

```
# Listing Snippets


```bash

go-code-snippet-manager list

```
# Removing a Snippet


```bash

go-code-snippet-manager remove "Hello World" "Printf"

```
