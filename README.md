# aws-vault-switch 

A simple Go CLI app that extends [aws-vault](https://github.com/99designs/aws-vault) functionality to enable seamless switching between AWS profiles without the need to exit the current shell session. 

Written by [Kobe Subramaniam](https://kobebigs.com).

## Key Features

- **Profile Listing**: List all available AWS profiles managed by `aws-vault`.
- **Profile Switching**: Switch between AWS profiles without exiting the current shell session, maintaining your working context and improving efficiency.
- **Easy Integration**: Built using Go and [Cobra](https://github.com/spf13/cobra-cli) library, this tool integrates smoothly with existing `aws-vault` setup and enhances their capabilities.

## Installation

For developers, clone the repository and build the application

``` bash
git clone https://github.com/kobebigs/aws-vault-switch
cd aws-vault-switch
go build -o aws-vault-switch
```

## Usage

1. List profiles

    ``` bash
    ./aws-vault-switch ls
    ```

2. Switch profile

    ``` bash
    ./aws-vault-switch switch -p <profile name>
    ```
