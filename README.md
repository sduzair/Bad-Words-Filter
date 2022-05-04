# Flashcards-Memory-Building

Print the message Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats): each time the user is prompted for their next action. The action is read from the next line, processed, and the message is output again until the user decides to exit the program.

## The program's behavior depends on the user's input action:

-  add — create a new flashcard with a unique term and definition. After adding the card, output the message The pair ("term":"definition") has been added, where "term" is the term entered by the user and "definition" is the definition entered by the user. If a term or a definition already exists, output the line This <term/definition> already exists. Try again: and accept a new term or definition.
-  remove — ask the user for the term of the card they want to remove with the message Which card?, and read the user's input from the next line. If a matching flashcard exists, remove it from the set and output the message The card has been removed.. If there is no such flashcard in the set, output the message Can't remove "card": there is no such card., where "card" is the user's input.
import — print the line File name:, read the user's input from the next line, which is the file name, and import all the flashcards written to this file. If there is no file with such name, print the message File not found.. After importing the cards, print the message n cards have been loaded., where n is the number of cards in the file. The imported cards should be added to the ones that already exist in the memory of the program. However, the imported cards have priority: if you import a card with the name that already exists in the memory, the card from the file should overwrite the one in memory.
export — request the file name with the line File name: and write all currently available flashcards into this file. Print the message n cards have been saved., where n is the number of cards exported to the file.
ask — ask the user about the number of cards they want and then prompt them for definitions, like in the previous stage.
exit — print Bye bye! and finish the program.
log — ask the user where to save the log with the message File name:, save all the lines that have been input in/output to the console to the file, and print the message The log has been saved. Don't clear the log after saving it to the file.
hardest card — print a string that contains the term of the card with the highest number of wrong answers, for example, The hardest card is "term". You have N errors answering it. If there are several cards with the highest number of wrong answers, print all of the terms: The hardest cards are "term_1", "term_2". If there are no cards with errors in the user's answers, print the message There are no cards with errors.
reset stats — set the count of mistakes to 0 for all the cards and output the message Card statistics have been reset.

## When provided with command-line arguments, your program should do the following:

If --import_from=IMPORT is passed, where IMPORT is the file name, read the initial card set from the external file, and print the message n cards have been loaded. as the first line of the output, where n is the number of cards loaded from the external file. If such an argument is not provided, the set of cards should initially be empty and no message about card loading should be output.
If --export_to=EXPORT is passed, where EXPORT is the file name, write all cards that are in the program memory into this file after the user has entered exit, and the last line of the output should be n cards have been saved., where n is the number of flashcards in the set.

### Run arguments examples
```
./flashcards --import_from=derivatives.txt
./flashcards --export_to=animals.txt
./flashcards --import_from=words13june.txt --export_to=words14june.txt
./flashcards --export_to=vocab.txt --import_from=vocab.txt 
```
