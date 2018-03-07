# Introduction
Converts integers to Roman Numerals.  
Written in golang for speed.  
Part of a code challenge, but I considered using it in lemonbar.  

# Usage
Feed an integer as an argument  
```
# Integer as an argument
$ ./roman 10
X
```

Feed an integer from stdin  
```
# Integer from stdin
$ echo "33" | ./roman
XXXIII
```

# Compiling
There are no external dependencies. Simply run:  
```
$ go build -o ~/bin/roman
```
