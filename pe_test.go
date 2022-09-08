package main

import (
	"log"
	"testing"
)

func TestEverything(t *testing.T) {
	output := GetAllMatches("?? alt/generateStore.php", "")
	if output[0] != "alt/generateStore.php" {
		t.Errorf("Doesnt match files" + output[0])
	}

	output = GetAllMatches("2.0G", "")
	if len(output) != 0 {
		t.Errorf("File size matches" + output[0])
	}

	output = GetAllMatches("4.0K", "")
	if len(output) != 0 {
		t.Errorf("File size matches" + output[0])
	}

	output = GetAllMatches("72K", "")
	if len(output) != 0 {
		t.Errorf("File size matches" + output[0])
	}

	output = GetAllMatches("remotes/origin/master", "")
	if len(output) != 0 {
		t.Errorf("Git branch matches" + output[0])
	}

	output = GetAllMatches("I have a cat.", "")
	if len(output) != 0 {
		t.Errorf("Matches sentence" + output[0])
	}

	output = GetAllMatches("0.0.0.0:3000", "")
	if len(output) != 0 {
		t.Errorf("Match ipadress" + output[0])
	}

	output = GetAllMatches("'/usr/bin", "")
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly" + output[0])
	}

	output = GetAllMatches("/usr_b/bin", "")
	if output[0] != "/usr_b/bin" {
		t.Errorf("Doesn't match statement correctly" + output[0])
	}

	output = GetAllMatches("\"/usr/bin", "")
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly" + output[0])
	}

	output = GetAllMatches("`/usr/bin", "")
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly" + output[0])
	}

	output = GetAllMatches("€/usr/bin", "")
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly" + output[0])
	}

	output = GetAllMatches("prefix=/usr/bin", "")
	if len(output) != 1 {
		t.Errorf("Should match =/usr/bin" + output[0])
	}
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly" + output[0])
	}

	output = GetAllMatches("/var//log", "")
	if len(output) != 0 {
		t.Errorf("Matches double //" + output[0])
	}

	output = GetAllMatches("s/+//", "")
	if len(output) != 0 {
		t.Errorf("Doesn't match substitute" + output[0])
	}

	output = GetAllMatches("s/^//", "")
	if len(output) != 0 {
		t.Errorf("Doesn't match substitute" + output[0])
	}

	output = GetAllMatches("/usr/bin/env\\", "")
	if len(output) != 1 {
		t.Errorf("Doesn't match escaped" + output[0])
	}
	if output[0] != "/usr/bin/env" {
		t.Errorf("Doesn't match escaped exactly" + output[0])
	}

	output = GetAllMatches("!#/usr/bin/env", "")
	if len(output) != 1 {
		t.Errorf("Doesn't match shebang" + output[0])
	}
	if output[0] != "/usr/bin/env" {
		t.Errorf("Doesn't match shebang exactly" + output[0])
	}

	output = GetAllMatches("hello .gitignore", "")
	if output[0] != ".gitignore" {
		t.Errorf("Doesnt match hidden files" + output[0])
	}

	output = GetAllMatches(" this.user ", "")
	if len(output) != 0 {
		t.Errorf("Matches this.user" + output[0])
	}

	output = GetAllMatches("To https://test@test.org/88/ls.git", "")
	if len(output) != 0 {
		t.Errorf("Matches email adresses" + output[0])
	}

	output = GetAllMatches(" mail@mail.com ", "")
	if len(output) != 0 {
		t.Errorf("Matches email adresses" + output[0])
	}

	output = GetAllMatches(" logo@2x.png ", "")
	if len(output) == 0 {
		t.Errorf("Doesn't match retina asset" + output[0])
	}

	output = GetAllMatches("and/or", "")
	if len(output) != 0 {
		t.Errorf("Matches and/or adresses" + output[0])
	}

	output = GetAllMatches("v1.2", "")
	if len(output) != 0 {
		t.Errorf("Matches version number" + output[0])
	}

	output = GetAllMatches("14.22.2", "")
	if len(output) != 0 {
		t.Errorf("Matches version number" + output[0])
	}

	output = GetAllMatches("~/v1.2/js", "")
	if len(output) != 1 {
		t.Errorf("Should match path with version inside" + output[0])
	}

	output = GetAllMatches("obj.slice()", "")
	if len(output) != 0 {
		t.Errorf("Matches function call" + output[0])
	}

	output = GetAllMatches("fs.read(arg)", "")
	if len(output) != 0 {
		t.Errorf("Matches function call" + output[0])
	}

	output = GetAllMatches("~/www", "")
	if len(output) == 0 || output[0] != "~/www" {
		t.Errorf("Doesnt match home" + output[0])
	}

	output = GetAllMatches("origin/master", "")
	if len(output) != 0 {
		t.Errorf("Matches remote name" + output[0])
	}

	output = GetAllMatches("john doe (dead on 28/04/2014)", "")
	if len(output) != 0 {
		t.Errorf("Matches date" + output[0])
	}

	output = GetAllMatches("john doe ,dead on 28/04/2014", "")
	if len(output) != 0 {
		t.Errorf("Matches date" + output[0])
	}

	output = GetAllMatches(".gitignore , ~/www", "")
	if len(output) != 2 {
		t.Errorf("Doesnt match multi" + output[0])
	}

	output = GetAllMatches("usçr.test.js", "")
	log.Println(output)
	if output[0] != "usçr.test.js" {
		t.Errorf("Doesnt match diacritics " + output[0])
	}

	output = GetAllMatches("user.test.js", "")
	if len(output) != 1 {
		t.Errorf("Doesnt match multiple extensions" + output[0])
	}

	output = GetAllMatches("[Error/foobar]", "")
	if len(output) == 1 {
		t.Errorf("Matches error" + output[0])
	}

	output = GetAllMatches("[Object.foo]", "")
	if len(output) == 1 {
		t.Errorf("Matches Object.foo" + output[0])
	}

	output = GetAllMatches("(user.js)", "")
	if len(output) != 1 {
		t.Errorf("Doesnt match surrounded by parens" + output[0])
	}
	if output[0] != "user.js" {
		t.Errorf("matches surrounded by parens badly" + output[0])
	}

	output = GetAllMatches("var/", "")
	if len(output) != 1 {
		t.Errorf("Doesnt match dir" + output[0])
	}

	output = GetAllMatches("//", "")
	if len(output) != 0 {
		t.Errorf("Comment matches" + output[0])
	}

	output = GetAllMatches("test.js:45", "ackmate")
	if len(output) != 1 {
		t.Errorf("Ackmate doesnt match" + output[0])
	}

	if output[0] == "test.js" {
		t.Errorf("Ackmate should not forget number" + output[0])
	}
	if output[0] != "test.js:45" {
		t.Errorf("Ackmate should output right line number" + output[0])
	}

	output = GetAllMatches("test.js:45:12", "ackmate")
	if len(output) != 1 {
		t.Errorf("Ackmate doesnt match" + output[0])
	}

	if output[0] == "test.js" {
		t.Errorf("Ackmate should not forget number" + output[0])
	}
	if output[0] != "test.js:45:12" {
		t.Errorf("Ackmate should output right line number" + output[0])
	}

	output = GetAllMatches("test.js:45:12 foo bar", "ackmate")
	if len(output) != 1 {
		t.Errorf("Ackmate doesnt match" + output[0])
	}

	if output[0] == "test.js" {
		t.Errorf("Ackmate should not forget number" + output[0])
	}
	if output[0] != "test.js:45:12" {
		t.Errorf("Ackmate should output right line number" + output[0])
	}
}
