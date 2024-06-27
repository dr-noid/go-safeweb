# Report for Assignment 1

## Project chosen

Name: go-safeweb

URL: [<TODO>](https://github.com/google/go-safeweb)

Number of lines of code and the tool used to count it: 13k https://github.com/AlDanial/cloc

Programming language: go

## Coverage measurement

### Existing tool

go cover
https://pkg.go.dev/cmd/cover

<Show the coverage results provided by the existing tool with a screenshot>
run the tests and save the coverage data in 'coverage.out'
<code>go test -coverprofile coverage.out ./...</code>
use the cover tool to get a nice visual display:
<code>go tool cover -html=coverage.out</code>

![coverage1](existing-coverage/image1.png)
![coverage2](existing-coverage/image2.png)
![coverage3](existing-coverage/image3.png)
![coverage4](existing-coverage/image4.png)

### Your own coverage tool

<The following is supposed to be repeated for each group member>

Group member name: Nika Emadian

Function 1 name: SameSite

![](/existing-coverage/cookie.png)

Function 2 name: Write 

![](/existing-coverage/write.png)

[GitHub Commit](https://github.com/dr-noid/go-safeweb/commit/5fb77332308d05571fa3160aff046c01725fd559)

Group Member: Berat Kir

Function 1: echo

Function 2: uptime

![alt text](existing-coverage/dr-noid-coverage.png)

[GitHub Commit](https://github.com/dr-noid/go-safeweb/commit/b071a38bd809d8afcf3be28dace3da369f5fe4c2)

## Coverage improvement

### Individual tests

<The following is supposed to be repeated for each group member>

Group member name: Nika Emadian

Test 1: TestFlightValueNil

[GitHub Commit](https://github.com/dr-noid/go-safeweb/commit/7ad3e7c330db30aa69769a367bfd1039aa7af64a))

![Old coverage](/existing-coverage/FlightValuesold.png)

![New coverage](/existing-coverage/FlightValuesnew.png)

Fligth.go had a coverage of 89.5% initially, while by testing FlightValues it was increased to 92.1%.

Test 2: TestFlightAddCookie

[GitHub Commit](https://github.com/dr-noid/go-safeweb/commit/7ad3e7c330db30aa69769a367bfd1039aa7af64a))

![Old coverage](/existing-coverage/AddCookieOld.png)

![New coverage](/existing-coverage/AddCookieNew.png)

By tetsing the second function the coverage incearsed from 92.1% to 94.7%.

### Overall

<Provide a screenshot of the old coverage results by running an existing tool (the same as you already showed above)>

<Provide a screenshot of the new coverage results by running the existing tool using all test modifications made by the group>

## Statement of individual contributions

<Write what each group member did>
