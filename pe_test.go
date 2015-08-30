package pathextractor

import "testing"

func TestEverything(t *testing.T) {
	output := GetAllMatches("?? alt/generateStore.php", MatchOptions{})
	if output[0] != "alt/generateStore.php" {
		t.Errorf("Doesnt match files", output)
	}

	output = GetAllMatches("I have a cat.", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches sentence", output)
	}

	output = GetAllMatches("!#/usr/bin/env", MatchOptions{})
	if len(output) != 1 {
		t.Errorf("Doesn't match shebang", output)
	}
	if output[0] != "/usr/bin/env" {
		t.Errorf("Doesn't match shebang", output)
	}

	output = GetAllMatches("hello .gitignore", MatchOptions{})
	if output[0] != ".gitignore" {
		t.Errorf("Doesnt match hidden files", output)
	}

	output = GetAllMatches(" this.user ", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches this.user", output)
	}

	output = GetAllMatches(" mail@mail.com ", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches email adresses", output)
	}

	output = GetAllMatches(" logo@2x.png ", MatchOptions{})
	if len(output) == 0 {
		t.Errorf("Doesn't match retina asset", output)
	}

	output = GetAllMatches("and/or", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches and/or adresses", output)
	}

	output = GetAllMatches("v1.2", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches version number", output)
	}

	output = GetAllMatches("obj.slice()", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches function call", output)
	}

	output = GetAllMatches("fs.read(arg)", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches function call", output)
	}

	output = GetAllMatches("~/www", MatchOptions{})
	if len(output) == 0 || output[0] != "~/www" {
		t.Errorf("Doesnt match home", output)
	}

	output = GetAllMatches("origin/master", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches remote name", output)
	}

	output = GetAllMatches("john doe (dead on 28/04/2014)", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches date", output)
	}

	output = GetAllMatches("john doe ,dead on 28/04/2014", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Matches date", output)
	}

	output = GetAllMatches(".gitignore , ~/www", MatchOptions{})
	if len(output) != 2 {
		t.Errorf("Doesnt match multi", output)
	}

	output = GetAllMatches("user.test.js", MatchOptions{})
	if len(output) != 1 {
		t.Errorf("Doesnt match multiple extensions", output)
	}

	output = GetAllMatches("(user.js)", MatchOptions{})
	if len(output) != 1 {
		t.Errorf("Doesnt match surrounded by parens", output)
	}
	if output[0] != "user.js" {
		t.Errorf("matches surrounded by parens badly", output)
	}

	output = GetAllMatches("var/", MatchOptions{})
	if len(output) != 1 {
		t.Errorf("Doesnt match dir", output)
	}

	output = GetAllMatches("//", MatchOptions{})
	if len(output) != 0 {
		t.Errorf("Comment matches", output)
	}

	output = GetAllMatches("test.js:45", MatchOptions{format: "ackmate"})
	if len(output) != 1 {
		t.Errorf("Ackmate doesnt match", output)
	}
	if output[0] == "test.js" {
		t.Errorf("Ackmate should not forget number", output)
	}
	if output[0] != "test.js:45" {
		t.Errorf("Ackmate should output right line number", output)
	}
}
