// Code generated by re2c 3.1, DO NOT EDIT.
package main

func findEmails(bytes []byte) int {
	var count int
	var cur, mar int
	bytes = append(bytes, byte(0)) // add terminating null
	lim := len(bytes) - 1          // lim points at the terminating null

	for {
		{
			var yych byte
			yych = bytes[cur]
			switch yych {
			case '+':
				fallthrough
			case '-', '.':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy2
			default:
				if lim <= cur {
					goto yy11
				}
				goto yy1
			}
		yy1:
			cur += 1
			{
				continue
			}
		yy2:
			cur += 1
			mar = cur
			yych = bytes[cur]
			switch yych {
			case '+':
				fallthrough
			case '-', '.':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy2
			case '@':
				goto yy4
			default:
				goto yy3
			}
		yy3:
			{
				continue
			}
		yy4:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case 0x00:
				goto yy5
			case '.':
				goto yy6
			default:
				goto yy7
			}
		yy5:
			cur = mar
			goto yy3
		yy6:
			cur += 1
			yych = bytes[cur]
		yy7:
			switch yych {
			case '-':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy6
			case '.':
				goto yy8
			default:
				goto yy5
			}
		yy8:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case '-', '.':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy9
			default:
				goto yy5
			}
		yy9:
			cur += 1
			yych = bytes[cur]
			switch yych {
			case '-', '.':
				fallthrough
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				fallthrough
			case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
				fallthrough
			case '_':
				fallthrough
			case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
				goto yy9
			default:
				goto yy10
			}
		yy10:
			{
				count += 1
				continue
			}
		yy11:
			{
				return count
			}
		}

	}
}
