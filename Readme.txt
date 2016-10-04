Follow these steps to build and run the program

Note:  I assume you have the complete set up for running Go programs.

1)  Get the Workiva package using below command
	go get github.com/Workiva/go-datastructures/...

2)  Update the Workiva package using below command
	go get -u github.com/Workiva/go-datastructures/...

3)  Add new directory 'longest_compound_word' in  path 'github.com\golang\example'
	You can create using right click button, create new Folder in Windows

4)  Copy below files into directory 'longest_compound_word'
	a) longest_compound_word.go
	b) Readme.txt
	c) nodeprime-input.txt
	d) nodeprime-input-negative.txt

5)  Copy any other your input files into the directory 'longest_compound_word'	

6) Goto Windows command prompt and run the program using below command
   go run <path-to-file-longest_compound_word.go> <path-to-input-file>

   For example
   C:\>go run C:\Projects\Go\src\github.com\golang\example\longest_compound_word\longest_compound_word.go  C:\Projects\Go\src\github.com\golang\example\longest_compound_word\nodeprime-input.txt

   ------------------------------------------------------------------------------
   You would see output something like below

   LongestCompundWord:  antidisestablishmentarianisms

   Length:  29

   lcwSubWords are:   an ti dis establishmentarianisms

   Total Time taken 1.9571119s
   ------------------------------------------------------------------------------

