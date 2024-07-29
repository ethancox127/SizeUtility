This is a command line utility written in Golang that takes directories from the user and prints the size of those directories.
Users can specify if they want the sizes of the subdirectories under those specified to be printed as well using the --recursive flag.
Users can specify if the want the output to be easier to read using the --human flag.
To use this command line utility either run the main file with the directories and flags specified, or:
    Run "go build main.go".
    Run "go install" to install the cli.
    Use the "size" command the with necessary options to run the utility.

Ex: size .. --recursive --human will print out all the directories starting with and under the parent directory of current assuming the parent directory exists.