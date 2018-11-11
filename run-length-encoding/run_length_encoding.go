package encode

import "strconv"

type Temp struct {
	letter   string
	count    int
	tmpCount string
	result   string
}

func RunLengthEncode(txt string) string {
	if len(txt) < 1 {
		return ""
	}

	var tmp Temp
	for _, l := range txt {
		if tmp.letter == string(l) {
			tmp.count++
			continue
		}
		if tmp.letter != "" {
			tmp.addEncodeLetter()
		}
		tmp.letter = string(l)
		tmp.count = 1
	}
	tmp.addEncodeLetter()

	return tmp.result
}

func RunLengthDecode(txt string) string {
	if len(txt) < 1 {
		return ""
	}

	var tmp Temp
	for _, s := range txt {
		if _, err := strconv.ParseInt(string(s), 10, 64); err == nil {
			tmp.tmpCount += string(s)
			continue
		}
		tmp.letter = string(s)
		tmp.addDecodeLetter()
	}

	return tmp.result
}

func (tmp *Temp) addEncodeLetter() {
	if tmp.count > 1 {
		tmp.result += strconv.Itoa(tmp.count) + tmp.letter
		return
	}
	tmp.result += tmp.letter
}

func (tmp *Temp) addDecodeLetter() {
	if tmp.tmpCount == "" {
		tmp.result += tmp.letter
	}

	n, _ := strconv.ParseInt(string(tmp.tmpCount), 10, 64)
	for i := 0; i < int(n); i++ {
		tmp.result += tmp.letter
	}
	tmp.tmpCount = ""
}
