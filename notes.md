# Thoughts etc.

http://spec.commonmark.org/0.26/
Check the whitespace specification

Input is read line by line.

The order in which the block types are checked is important

When a "new" line is read an appropriate block is found via the specs. (if it starts with a number of # then it is a header block, etc.)
The block specifies weather it is expecting more lines, which the parser uses to figure out if the next line is "new" or not.
If it the next line is not new, it is fed into the current block until the block has had enough.
