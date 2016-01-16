package pathextractor

import "testing"

func TestEverything(t *testing.T) {
	output := GetAllMatches("?? alt/generateStore.php", "")
	if output[0] != "alt/generateStore.php" {
		t.Errorf("Doesnt match files", output)
	}

	output = GetAllMatches("2.0G", "")
	if len(output) != 0 {
		t.Errorf("File size matches", output)
	}

	output = GetAllMatches("4.0K", "")
	if len(output) != 0 {
		t.Errorf("File size matches", output)
	}

	output = GetAllMatches("72K", "")
	if len(output) != 0 {
		t.Errorf("File size matches", output)
	}

	output = GetAllMatches("remotes/origin/master", "")
	if len(output) != 0 {
		t.Errorf("Git branch matches", output)
	}

	output = GetAllMatches("I have a cat.", "")
	if len(output) != 0 {
		t.Errorf("Matches sentence", output)
	}

	output = GetAllMatches("'/usr/bin", "")
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly", output)
	}

	output = GetAllMatches("/usr_b/bin", "")
	if output[0] != "/usr_b/bin" {
		t.Errorf("Doesn't match statement correctly", output)
	}

	output = GetAllMatches("\"/usr/bin", "")
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly", output)
	}

	output = GetAllMatches("`/usr/bin", "")
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly", output)
	}

	output = GetAllMatches("€/usr/bin", "")
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly", output)
	}

	output = GetAllMatches("prefix=/usr/bin", "")
	if len(output) != 1 {
		t.Errorf("Should match =/usr/bin", output)
	}
	if output[0] != "/usr/bin" {
		t.Errorf("Doesn't match statement correctly", output)
	}

	output = GetAllMatches("/var//log", "")
	if len(output) != 0 {
		t.Errorf("Matches double //", output)
	}

	output = GetAllMatches("s/+//", "")
	if len(output) != 0 {
		t.Errorf("Doesn't match substitute", output)
	}

	output = GetAllMatches("s/^//", "")
	if len(output) != 0 {
		t.Errorf("Doesn't match substitute", output)
	}

	output = GetAllMatches("/usr/bin/env\\", "")
	if len(output) != 1 {
		t.Errorf("Doesn't match escaped", output)
	}
	if output[0] != "/usr/bin/env" {
		t.Errorf("Doesn't match escaped exactly", output)
	}

	output = GetAllMatches("!#/usr/bin/env", "")
	if len(output) != 1 {
		t.Errorf("Doesn't match shebang", output)
	}
	if output[0] != "/usr/bin/env" {
		t.Errorf("Doesn't match shebang exactly", output)
	}

	output = GetAllMatches("hello .gitignore", "")
	if output[0] != ".gitignore" {
		t.Errorf("Doesnt match hidden files", output)
	}

	output = GetAllMatches(" this.user ", "")
	if len(output) != 0 {
		t.Errorf("Matches this.user", output)
	}

	output = GetAllMatches("To https://test@test.org/88/ls.git", "")
	if len(output) != 0 {
		t.Errorf("Matches email adresses", output)
	}

	output = GetAllMatches(" mail@mail.com ", "")
	if len(output) != 0 {
		t.Errorf("Matches email adresses", output)
	}

	output = GetAllMatches(" logo@2x.png ", "")
	if len(output) == 0 {
		t.Errorf("Doesn't match retina asset", output)
	}

	output = GetAllMatches("and/or", "")
	if len(output) != 0 {
		t.Errorf("Matches and/or adresses", output)
	}

	output = GetAllMatches("v1.2", "")
	if len(output) != 0 {
		t.Errorf("Matches version number", output)
	}

	output = GetAllMatches("14.22.2", "")
	if len(output) != 0 {
		t.Errorf("Matches version number", output)
	}

	output = GetAllMatches("~/v1.2/js", "")
	if len(output) != 1 {
		t.Errorf("Should match path with version inside", output)
	}

	output = GetAllMatches("obj.slice()", "")
	if len(output) != 0 {
		t.Errorf("Matches function call", output)
	}

	output = GetAllMatches("fs.read(arg)", "")
	if len(output) != 0 {
		t.Errorf("Matches function call", output)
	}

	output = GetAllMatches("~/www", "")
	if len(output) == 0 || output[0] != "~/www" {
		t.Errorf("Doesnt match home", output)
	}

	output = GetAllMatches("origin/master", "")
	if len(output) != 0 {
		t.Errorf("Matches remote name", output)
	}

	output = GetAllMatches("john doe (dead on 28/04/2014)", "")
	if len(output) != 0 {
		t.Errorf("Matches date", output)
	}

	output = GetAllMatches("john doe ,dead on 28/04/2014", "")
	if len(output) != 0 {
		t.Errorf("Matches date", output)
	}

	output = GetAllMatches(".gitignore , ~/www", "")
	if len(output) != 2 {
		t.Errorf("Doesnt match multi", output)
	}

	output = GetAllMatches("user.test.js", "")
	if len(output) != 1 {
		t.Errorf("Doesnt match multiple extensions", output)
	}

	output = GetAllMatches("[Error/foobar]", "")
	if len(output) == 1 {
		t.Errorf("Matches error", output)
	}

	output = GetAllMatches("[Object.foo]", "")
	if len(output) == 1 {
		t.Errorf("Matches Object.foo", output)
	}

	output = GetAllMatches("(user.js)", "")
	if len(output) != 1 {
		t.Errorf("Doesnt match surrounded by parens", output)
	}
	if output[0] != "user.js" {
		t.Errorf("matches surrounded by parens badly", output)
	}

	output = GetAllMatches("var/", "")
	if len(output) != 1 {
		t.Errorf("Doesnt match dir", output)
	}

	output = GetAllMatches("//", "")
	if len(output) != 0 {
		t.Errorf("Comment matches", output)
	}

	output = GetAllMatches("test.js:45", "ackmate")
	if len(output) != 1 {
		t.Errorf("Ackmate doesnt match", output)
	}

	if output[0] == "test.js" {
		t.Errorf("Ackmate should not forget number", output)
	}
	if output[0] != "test.js:45" {
		t.Errorf("Ackmate should output right line number", output)
	}

	output = GetAllMatches("test.js:45:12", "ackmate")
	if len(output) != 1 {
		t.Errorf("Ackmate doesnt match", output)
	}

	if output[0] == "test.js" {
		t.Errorf("Ackmate should not forget number", output)
	}
	if output[0] != "test.js:45:12" {
		t.Errorf("Ackmate should output right line number", output)
	}

	output = GetAllMatches("test.js:45:12 foo bar", "ackmate")
	if len(output) != 1 {
		t.Errorf("Ackmate doesnt match", output)
	}

	if output[0] == "test.js" {
		t.Errorf("Ackmate should not forget number", output)
	}
	if output[0] != "test.js:45:12" {
		t.Errorf("Ackmate should output right line number", output)
	}
}
