# Unexpected Zero Value from Go Map Access

This example demonstrates a common Go issue where accessing a non-existent key in a map returns the zero value for the map's value type instead of an error or nil.  This behavior can lead to subtle bugs that are difficult to track down if not handled carefully.

**Problem:**
The code attempts to access keys that don't exist in the map (`m[