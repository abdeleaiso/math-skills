# math-skills
## Description
Math Skills is a Go program that calculates four statistical measures from a given data file:
- **Average**
- **Median**
- **Variance**
- **Standard Deviation**

The program treats the provided data as a statistical population.
## Project Structure
```bash
.
|____README.md
|____main.go
|____go.mod
```

## Requirement
- **Go 1.24.4** or newer.

## Instalation
```bash
$ git clone <math-skills-repo>
$ cd math-skills
```

## Usage
Create a text file containing numerical data, with one value per line.

Example (data.txt):
```bash
189
113
121
114
145
110
```
Run the program:
```bash
$ go run . data.txt
```

## Feature
- Reads numerical data from a text file.
- Calculates the arithmetic mean.
- Calculates the median.
- Calculates population variance.
- Calculates population standard deviation.
- Supports positive and negative integer values.
- Accepts the input file path as a command-line argument.

## Limitations
- The program accepts only one input file per execution.
- Input data must contain valid integer values.
- The input file must not be empty.

## Authorof any given file

- **<a href="https://github.com/abdeleaiso"/>Me</a>**

## Licence
This project is licenced under the MIT. See the <a href="https://github.com/abdeleaiso/math-skills/blob/main/LICENSE"/>LICENCE</a> file for details