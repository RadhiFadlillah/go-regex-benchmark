package main

/*!rules:re2c:base_template
re2c:eof              = 0;
re2c:yyfill:enable    = 0;
re2c:posix-captures   = 0;
re2c:case-insensitive = 0;

re2c:define:YYCTYPE     = byte;
re2c:define:YYPEEK      = "input[cursor]";
re2c:define:YYSKIP      = "cursor++";
re2c:define:YYBACKUP    = "marker = cursor";
re2c:define:YYRESTORE   = "cursor = marker";
re2c:define:YYLESSTHAN  = "limit <= cursor";
re2c:define:YYSTAGP     = "@@{tag} = cursor";
re2c:define:YYSTAGN     = "@@{tag} = -1";
re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";
*/

func findEmails(input []byte) int {
	var count int
	var cursor, marker int
	input = append(input, byte(0)) // add terminating null
	limit := len(input) - 1        // lim points at the terminating null

	for { /*!use:re2c:base_template
		quant1 = [+\-.0-9A-Z_a-z]+;
		quant2 = [+\-.0-9A-Z_a-z]+@[\-.0-9A-Z_a-z]+;
		email  = [+\-.0-9A-Z_a-z]+@[\-.0-9A-Z_a-z]+[.][\-.0-9A-Z_a-z]+;

		{email}  { count++; continue }
		{quant1} { continue }
		{quant2} { continue }
		*        { continue }
		$        { return count }
		*/
	}
}

func findURIs(input []byte) int {
	var count int
	var cursor, marker int
	input = append(input, byte(0)) // add terminating null
	limit := len(input) - 1        // lim points at the terminating null

	for { /*!use:re2c:base_template
		quant1 = [0-9A-Z_a-z]+;
		quant2 = [0-9A-Z_a-z]+:[/][/][^\t\n\f\r #/?]+;
		quant3 = [0-9A-Z_a-z]+:[/][/][^\t\n\f\r #/?]+[^\t\n\f\r #?]+;
		quant4 = [0-9A-Z_a-z]+:[/][/][^\t\n\f\r #/?]+[^\t\n\f\r #?]+(![?][^\t\n\f\r #]*)?;
		uri    = [0-9A-Z_a-z]+:[/][/][^\t\n\f\r #/?]+[^\t\n\f\r #?]+(![?][^\t\n\f\r #]*)?(!#[^\t\n\f\r ]*)?;

		{uri}    { count++; continue }
		{quant1} { continue }
		{quant2} { continue }
		{quant3} { continue }
		{quant4} { continue }
		*        { continue }
		$        { return count }
		*/
	}
}

func findIPs(input []byte) int {
	var count int
	var cursor, marker int
	input = append(input, byte(0)) // add terminating null
	limit := len(input) - 1        // lim points at the terminating null

	for { /*!use:re2c:base_template
		ip = ((2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9])[.])((2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9])[.])((2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9])[.])(!2(!5[0-5]|[0-4][0-9])|[01]?[0-9][0-9]);

		{ip} { count++; continue }
		*    { continue }
		$    { return count }
		*/
	}
}

func findLongDatePattern(input []byte) int {
	var count int
	var cursor, marker int
	input = append(input, byte(0)) // add terminating null
	limit := len(input) - 1        // lim points at the terminating null

	// Variable for capturing parentheses (twice the number of groups).
	/*!maxnmatch:re2c*/
	yypmatch := make([]int, YYMAXNMATCH*2)
	var yynmatch int
	_ = yynmatch

	// Autogenerated tag variables used by the lexer to track tag values.
	/*!stags:re2c format = 'var @@ int; _ = @@\n'; */

	for { /*!use:re2c:base_template
		re2c:posix-captures   = 1;
		re2c:case-insensitive = 1;

		rxDay = [0-3]?[0-9];
		rxYear = 199[0-9]|20[0-3][0-9];
		rxMonth = January?|February?|March|A[pv]ril|Ma[iy]|Ju(!n[ei]|l[iy])|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Ju[ln]|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|M(!aret|ei)|Agustus|Jänner|Feber|März|janvier|février|mars|jui(!n|llet)|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|E(!ylül|kim)|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|E(!yl|ki)|Kas|Ara;
		rxMDY = ({rxMonth})[\t\n\f\r ]({rxDay})(!st|nd|rd|th)?,?[\t\n\f\r ]({rxYear});
		rxDMY = ({rxDay})(!st|nd|rd|th|[.])?[\t\n\f\r ](!of[\t\n\f\r ])?({rxMonth})[,.]?[\t\n\f\r ]({rxYear});

		{rxMDY} { count++; continue }
		{rxDMY} { count++; continue }

		* { continue }
		$ { return count }
		*/
	}
}
