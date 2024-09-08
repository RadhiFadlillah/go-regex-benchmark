package main

func findURIs(runes []rune) int {
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

		uri = [0-9A-Z_a-z]+:[/][/][^\t\n\f\r #/\?]+[^\t\n\f\r #\?]+(![?][^\t\n\f\r #]*)?(!#[^\t\n\f\r ]*)?;

		{uri} { count += 1; continue }
		*     { continue }
		$     { return count }
		*/
	}
}
