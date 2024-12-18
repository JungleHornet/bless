# bless
A TUI library for go. https://github.com/fatih/color is used for colored/formatted text.

# Getting started
Once you import and `go get github.com/junglehornet/bless` use bless.New() to create a new blessing. Parameters are in order as follows:

Selector Open: Opening character when selecting an option.

Selector Close: Closing character when selecting an option.


Frame: Frame character that will be printed around the terminal.

Ex.
```
b := bless.New("{", "}", '#')
```

# Printing

Use methods to print to a blessing:
```
b.Println("Hello,")

line := b.Print("World!") // Println and Print return line number where printing started

b.Overwrite(line, "Gophers") // Overwrite takes a line number (int) and overwrites the line with new text

b.Print("!\n this is a new line")

b.RmLine(line + 1) // RmLine removes the line with the line number passed to it.
```
Output:
```
Hello,
Gophers!
```

# Color

Use https://github.com/fatih/color for inserting colored strings into the blessing.
```
red := color.New(color.FgRed).SprintFunc()
b.Print(red("Red text, "))
blue := color.New(color.FgBlue)
b.Print(blue.Sprint("Blue text"))
```


# UI Functions:

### Options:
You can use the LROptions function to provide the user with multiple options to choose between. The user can navigate the options with the arrow keys and submit their answer with the Enter/Return key. The first parameter is a prompt, and the rest are options. It returns the index of the chosen option:
```
b.Print(b.LROptions("Which option?", "option 1", "option 2"))
```
Output:
```
Which option?
 option 1 {OPTION 2} // user chooses option 2
1
```
or,
```
Which option?
{OPTION 1} option 2 // user chooses option 1
0
```
