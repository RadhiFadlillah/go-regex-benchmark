package main

func findIPs(runes []rune) int {
	var cur, mar int
	lim := len(runes) - 1 // lim points at the terminating null
	count := 0

	for { /*!re2c
		re2c:eof = 0;
		re2c:define:YYCTYPE    = rune;
		re2c:define:YYPEEK     = "runes[cur]";
		re2c:define:YYSKIP     = "cur += 1";
		re2c:define:YYBACKUP   = "mar = cur";
		re2c:define:YYRESTORE  = "cur = mar";
		re2c:define:YYLESSTHAN = "lim <= cur";
		re2c:yyfill:enable = 0;

		ip = ((2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9])[\.])((2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9])[\.])((2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9])[\.])(!2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9]);

		{ip} { count += 1; continue }
		*    { continue }
		$    { return count }
		*/
	}
}
