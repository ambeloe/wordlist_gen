# wordlist_gen
simple flexible wordlist generator

outputs to stdout

arguments:
-p the pattern
-e print estimates for keyspace and wordlist size
-o file to write wordlist to

[] to separate characters of the password; prepend a number to repeat
{} to specify specific characters to include in the character set for one password character (\ to escape)

@ number row symbols ~`!@#$%^&*()_-+=
: side symbols {[}]|\:;"'<,>.?/

a lower case letters
A upper case letters

# numbers 0123456789

% all letters
= all symbols

- no character

* everything except none

Examples:
	word that begins with P followed by a number and symbol-> [{P}]5[!a]2[#][@]
	cool, fool, mool, pool, tool, wool -> [{cfmptw}]2[{o}][{l}]
	custom letters or numbers -> [#{asdf}]
