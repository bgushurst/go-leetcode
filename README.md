# GO Leetcode Development Environment

### 1. Setup question files

Example: `123-question-name/123-question-name.go`

Use package `leetcode`

### 2. Copy Question from Leetcode

Copy the current question in from leetcode and provide a default return.

### 3. Create Tests

Select the question in VSCode and choose `GO: Generate Unit Tests for Selected Function`

Update the test with supplied test cases from Leetcode.

### 4. Use GOW to run tests

Run the following command to use a filewatcher to rerun tests as changes are made to the file.

`gow test 123-question-name/123-question-name.go 123-question-name/123-question-name_test.go`
