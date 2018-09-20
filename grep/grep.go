package grep

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	flagN            = "-n"
	flagL            = "-l"
	flagI            = "-i"
	flagV            = "-v"
	flagX            = "-x"
	flagM            = "-m"
	filesContentData = []string{
		"                                                     ",
		" iliad.txt                                           ",
		"   ---------------------------------------------     ",
		"   |Achilles sing, O Goddess! Peleus' son;     |     ",
		"   |His wrath pernicious, who ten thousand woes|     ",
		"   |Caused to Achaia's host, sent many a soul  |     ",
		"   |Illustrious into Ades premature,           |     ",
		"   |And Heroes gave (so stood the will of Jove)|     ",
		"   |To dogs and to all ravening fowls a prey,  |     ",
		"   |When fierce dispute had separated once     |     ",
		"   |The noble Chief Achilles from the son      |     ",
		"   |Of Atreus, Agamemnon, King of men.         |     ",
		"   ---------------------------------------------     ",
		"                                                     ",
		" midsummer-night.txt                                 ",
		"   -----------------------------------------------   ",
		"   |I do entreat your grace to pardon me.        |   ",
		"   |I know not by what power I am made bold,     |   ",
		"   |Nor how it may concern my modesty,           |   ",
		"   |In such a presence here to plead my thoughts;|   ",
		"   |But I beseech your grace that I may know     |   ",
		"   |The worst that may befall me in this case,   |   ",
		"   |If I refuse to wed Demetrius.                |   ",
		"   -----------------------------------------------   ",
		"                                                     ",
		" paradise-lost.txt                                   ",
		"   ------------------------------------------------- ",
		"   |Of Mans First Disobedience, and the Fruit      | ",
		"   |Of that Forbidden Tree, whose mortal tast      | ",
		"   |Brought Death into the World, and all our woe, | ",
		"   |With loss of Eden, till one greater Man        | ",
		"   |Restore us, and regain the blissful Seat,      | ",
		"   |Sing Heav'nly Muse, that on the secret top     | ",
		"   |Of Oreb, or of Sinai, didst inspire            | ",
		"   |That Shepherd, who first taught the chosen Seed| ",
		"   ------------------------------------------------- ",
	}
	filesData = make(map[string]File)
)

type File struct {
	name    string
	content []string
}

type Result struct {
	flags map[string]bool
	found []string
}

// Search searches a file for lines matching a regular expression pattern
func Search(pattern string, flags, files []string) []string {
	var r Result
	r.setData(flags, len(files))
	parseFiles()
	for _, f := range files {
		r.findInOne(pattern, f)
	}

	return r.found
}

func (f *File) addLine(l string) {
	f.content = append(f.content, strings.TrimSpace(strings.Trim(l, " | ")))
}

func (r *Result) addFound(file, line string, n int) {
	switch {
	case r.isFlagActive(flagN):
		line = strconv.Itoa(n) + ":" + line
	case r.isFlagActive(flagL):
		line = file
		if r.isFoundExists(line) {
			return
		}
	}

	if r.isFlagActive(flagM) && !r.isFlagActive(flagL) {
		line = file + ":" + line
	}
	r.found = append(r.found, line)
}

func (r *Result) findInOne(pattern, file string) {
	content := filesData[file].content
	for i := 0; i < len(content); i++ {
		line := content[i]
		lCompare := line
		pCompare := pattern
		if r.isFlagActive(flagI) {
			lCompare = strings.ToLower(line)
			pCompare = strings.ToLower(pattern)
		}

		switch {
		case r.isFlagActive(flagX):
			if lCompare == pCompare {
				r.addFound(file, line, i+1)
			}
		case r.isFlagActive(flagV):
			if !strings.Contains(lCompare, pCompare) {
				r.addFound(file, line, i+1)
			}
		default:
			if strings.Contains(lCompare, pCompare) {
				r.addFound(file, line, i+1)
			}
		}
	}
}

func (r *Result) isFlagActive(flag string) bool {
	if r.flags[flag] {
		return true
	}
	return false
}

func (r *Result) isFoundExists(check string) bool {
	for _, f := range r.found {
		if f == check {
			return true
		}
	}
	return false
}

func (r *Result) setData(flags []string, nFiles int) {
	r.flags = make(map[string]bool)
	for _, f := range flags {
		r.flags[f] = true
	}

	if nFiles > 1 {
		r.flags[flagM] = true
	}
	r.found = []string{}
}

func parseFiles() {
	start := false
	var f File
	for i := 0; i < len(filesContentData); i++ {
		line := filesContentData[i]
		matched, _ := regexp.MatchString("^.*txt", line)
		if matched {
			f.name = strings.TrimSpace(line)
			start = true
			i++
			continue
		}

		if strings.Contains(line, "-") {
			filesData[f.name] = f
			f = File{}
			start = false
			continue
		}

		if start {
			f.addLine(line)
		}
	}
}
