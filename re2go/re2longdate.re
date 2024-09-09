package main

import (
	"unicode"
)

func findLongDatePattern(runes []rune) int {
	var cur, mar int
	lim := len(runes) - 1 // lim points at the terminating null
	count := 0

	// Insensitive search
	upperRunes := make([]rune, len(runes))
	for i, r := range runes {
		upperRunes[i] = unicode.ToUpper(r)
	}

	// Capturing groups
	/*!maxnmatch:re2c*/	yypmatch := make([]int, YYMAXNMATCH*2)
	var yynmatch int
	var yyt1, yyt2, yyt3, yyt4, yyt5, yyt6, yyt7, yyt8, yyt9, yyt10 int
	_ = yynmatch

	for { /*!re2c
		re2c:eof = 0;
		re2c:define:YYCTYPE     = rune;
		re2c:define:YYPEEK      = "upperRunes[cur]";
		re2c:define:YYSKIP      = "cur += 1";
		re2c:define:YYBACKUP    = "mar = cur";
		re2c:define:YYRESTORE   = "cur = mar";
		re2c:define:YYLESSTHAN  = "lim <= cur";
		re2c:define:YYSTAGP     = "@@{tag} = cur";
		re2c:define:YYSTAGN     = "@@{tag} = -1";
		re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";
		re2c:posix-captures = 1;
		re2c:yyfill:enable = 0;

		rxDay = [0-3]?[0-9];
		rxYear = 199[0-9]|20[0-3][0-9];
		rxMonth = JANUARY?|FEBRUARY?|MARCH|A[PV]RIL|MA[IY]|JU(!N[EI]|L[IY])|AUGUST|SEPTEMBER|O[CK]TOBER|NOVEMBER|DE[CSZſ]EMBER|JAN|FEB|M[AÄä]R|APR|JU[LN]|AUG|SEP|O[CK]T|NOV|DE[CZ]|JANUARI|FEBRUARI|M(!ARET|EI)|AGUSTUS|JÄNNER|FEBER|MÄRZ|JANVIER|FÉVRIER|MARS|JUI(!N|LLET)|AOUT|SEPTEMBRE|OCTOBRE|NOVEMBRE|DÉCEMBRE|OCAK|ŞUBAT|MART|NISAN|MAYıS|HAZIRAN|TEMMUZ|AĞUSTOS|E(!YLÜL|KIM)|KASıM|ARALıK|OCA|ŞUB|MAR|NIS|HAZ|TEM|AĞU|E(!YL|KI)|KAS|ARA;
		rxLongPattern = ({rxMonth})[\t\n\f\r ]({rxDay})(!ST|ND|RD|TH)?,?[\t\n\f\r ]({rxYear})|({rxDay})(!ST|ND|RD|TH|[\.])?[\t\n\f\r ](!OF[\t\n\f\r ])?({rxMonth})[,\.]?[\t\n\f\r ]({rxYear});

		{rxLongPattern} {
			count += 1
			continue
		}

		* { continue }
		$ { return count }
		*/
	}
}
