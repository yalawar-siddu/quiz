/*
 * Longest Compound Word.
 * Given a list of words like
 * https://github.com/NodePrime/quiz/blob/master/word.list find
 * the longest compound-word in the list, which is also a
 * concatenation of other sub-words that exist in the list.
 *
 * Usage "go run <program-name> <path-to-input-file>"
 *       Input file should contain list of words at each line
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/Workiva/go-datastructures/trie/ctrie"
)

const lcwDebug = false

var lcwCtrie *ctrie.Ctrie
var lcwOutputWord string
var lcwSubWords string
var lcwNotFound bool = true

func isValidCompundWord(word string) {
	wordLen := len(word)
	for i := 0; (lcwNotFound) && (i < (wordLen - 1)); i++ {
		wordPrefix := word[0 : i+1]
		wordSuffix := word[i+1 : wordLen]
		/* fmt.Println("P: ", wordPrefix, " S: ", wordSuffix) */
		_, ok := lcwCtrie.Lookup([]byte(wordPrefix))
		if ok {
			if lcwDebug {
				fmt.Println("Prefix is : ", wordPrefix)
			}
			lcwOutputWord += wordPrefix
			lcwSubWords +=  " " + wordPrefix
			_, ok = lcwCtrie.Lookup([]byte(wordSuffix))
			if ok {
				if lcwDebug {
					fmt.Println("Suffix  is : ", wordSuffix)
				}
				lcwOutputWord += wordSuffix
				lcwSubWords += " " + wordSuffix
				fmt.Println("")
				fmt.Println("LongestCompundWord: ", lcwOutputWord)
				fmt.Println("")
				fmt.Println("Length: ", len(lcwOutputWord))
				fmt.Println("")
				fmt.Println("lcwSubWords are: ", lcwSubWords)
				lcwOutputWord = ""
				lcwNotFound = false
				break
			}
			if lcwNotFound {
				isValidCompundWord(wordSuffix)
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\nUsage: go run <program> <path-to-input-file>")
		return;
	}
	start := time.Now()
	/* Initialize ctrie instance */
	lcwCtrie = ctrie.New(nil)

	/* Open file for reading */
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	/*
	 * Read all words from the file
	 */
	var wordMap map[int]string
	wordMap = make(map[int]string)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		/* Add word into a Trie data structure */
		lcwCtrie.Insert([]byte(scanner.Text()), nil)

		/* Add word into a HashMap data structure */
		wordMap[len(scanner.Text())] = scanner.Text()
	}

	isValidCompundWord("dichlorodiphenyltrichloroethanes")
	/*
	 * Create Integer Array with Keys of Length of words
	 * Then Sort that list
	 */
	var wordMapKeys []int
	for k := range wordMap {
		wordMapKeys = append(wordMapKeys, k)
	}
	/*
	 * Do sort the array in Reverse Order
	 */
	sort.Sort(sort.Reverse(sort.IntSlice(wordMapKeys)))
	for _, k := range wordMapKeys {
		lcwOutputWord = ""
		lcwSubWords = ""
		if lcwDebug {
			fmt.Println("Key:", k, "Value:", wordMap[k])
		}
		isValidCompundWord(wordMap[k])
		if !lcwNotFound {
			break
		}
	}

    elapsed := time.Since(start)
	if lcwNotFound {
		fmt.Println("\nNo Longest Compound Word found !!!")
	}
    fmt.Println("\nTotal Time taken", elapsed)
}
