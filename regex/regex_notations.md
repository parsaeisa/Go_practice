# regex notations

**Regex** : cat

**Matches** : catalog , catherine , sophisticated 

you should specify beginning and ending of the line . 

### Counts

| char  | counts           |
|-------|------------------|
| *     | 0 or more times  |
| +     | 1 or more times  |
| {n}   | exactly n time   |
| {n,}  | at least n times |
| {n,m} | between n & m    |

example : a{2}

### Groupings
| groups | element |
|--------|---------|
 | [a-z]  | letters |
 | [0-9]  | numbers |

* for example in [4-9] you cannot enter 1,2,3

with () you can add a group of elements and seperate them by | . then 
only one of them is validated . 
example :
Good (morning|afternoon|night) ! 
"Good morning !" is validated .  

### Unclassified 
| notation | usage                           |
|----------|---------------------------------|
| $        | indicates end of the expression |
| ^        | NOT notation                    |