// Code generated by re2c 3.1, DO NOT EDIT.
package main

func findURIs(runes []rune) int {
	var cur, mar int
	lim := len(runes) - 1 // lim points at the terminating null
	count := 0

	for {
		{
			var yych rune
			yych = runes[cur]
			switch yych {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy3
			default:
				if lim <= cur {
					goto yy13
				}
				goto yy1
			}
		yy1:
			cur += 1
		yy2:
			{
				continue
			}
		yy3:
			cur += 1
			mar = cur
			yych = runes[cur]
			switch yych {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ':':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy5
			default:
				goto yy2
			}
		yy4:
			cur += 1
			yych = runes[cur]
		yy5:
			switch yych {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy4
			case ':':
				goto yy7
			default:
				goto yy6
			}
		yy6:
			cur = mar
			goto yy2
		yy7:
			cur += 1
			yych = runes[cur]
			switch yych {
			case '/':
				goto yy8
			default:
				goto yy6
			}
		yy8:
			cur += 1
			yych = runes[cur]
			switch yych {
			case '/':
				goto yy9
			default:
				goto yy6
			}
		yy9:
			cur += 1
			yych = runes[cur]
			switch yych {
			case '\t', '\n':
				fallthrough
			case '\f', '\r':
				fallthrough
			case ' ':
				fallthrough
			case '#':
				fallthrough
			case '/':
				fallthrough
			case '?':
				goto yy6
			default:
				if lim <= cur {
					goto yy6
				}
				goto yy10
			}
		yy10:
			cur += 1
			yych = runes[cur]
			switch yych {
			case '\t', '\n':
				fallthrough
			case '\f', '\r':
				fallthrough
			case ' ':
				fallthrough
			case '#':
				fallthrough
			case '?':
				goto yy6
			default:
				if lim <= cur {
					goto yy6
				}
				goto yy11
			}
		yy11:
			cur += 1
			yych = runes[cur]
			switch yych {
			case '\t', '\n':
				fallthrough
			case '\f', '\r':
				fallthrough
			case ' ':
				goto yy12
			default:
				if lim <= cur {
					goto yy12
				}
				goto yy11
			}
		yy12:
			{
				count += 1
				continue
			}
		yy13:
			{
				return count
			}
		}

	}
}
