# https://leetcode.com/problems/word-frequency/

cat words.txt | tr ' ' '\n' | sort | uniq -c | sort -r | awk '{print $2, $1}'
