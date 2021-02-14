/*
 * Copyright © 2021 Peter M. Stahl pemistahl@gmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either expressed or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lingua

import "regexp"

const maxNgramLength = 5

var japaneseCharacterSet = regexp.MustCompile("^[\\p{Hiragana}\\p{Katakana}\\p{Han}]+$")
var letter = regexp.MustCompile("^\\p{L}+$")
var multipleWhitespace = regexp.MustCompile("\\s+")
var noLetter = regexp.MustCompile("^[^\\p{L}]+$")
var numbers = regexp.MustCompile("\\p{N}")
var punctuation = regexp.MustCompile("\\p{P}")

var charsToLanguagesMapping = map[string][]Language{
	"Ãã":     {Portuguese, Vietnamese},
	"ĄąĘę":   {Lithuanian, Polish},
	"Żż":     {Polish, Romanian},
	"Îî":     {French, Romanian},
	"Ññ":     {Basque, Spanish},
	"ŇňŤť":   {Czech, Slovak},
	"Ăă":     {Romanian, Vietnamese},
	"İıĞğ":   {Azerbaijani, Turkish},
	"ЈјЉљЊњ": {Macedonian, Serbian},
	"ẸẹỌọ":   {Vietnamese, Yoruba},
	"Ōō":     {Maori, Yoruba},

	"ĀāĒēĪī": {Latvian, Maori, Yoruba},
	"Şş":     {Azerbaijani, Romanian, Turkish},
	"Ďď":     {Czech, Romanian, Slovak},
	"ÐðÞþ":   {Icelandic, Latvian, Turkish},
	"Ûû":     {French, Hungarian, Latvian},
	"Ćć":     {Bosnian, Croatian, Polish},
	"Đđ":     {Bosnian, Croatian, Vietnamese},
	"Іі":     {Belarusian, Kazakh, Ukrainian},
	"Ìì":     {Italian, Vietnamese, Yoruba},
	"Øø":     {Bokmal, Danish, Nynorsk},

	"Ūū":     {Latvian, Lithuanian, Maori, Yoruba},
	"Ëë":     {Afrikaans, Albanian, Dutch, French},
	"ÈèÙù":   {French, Italian, Vietnamese, Yoruba},
	"Êê":     {Afrikaans, French, Portuguese, Vietnamese},
	"Õõ":     {Estonian, Hungarian, Portuguese, Vietnamese},
	"Ôô":     {French, Portuguese, Slovak, Vietnamese},
	"ЁёЫыЭэ": {Belarusian, Kazakh, Mongolian, Russian},
	"ЩщЪъ":   {Bulgarian, Kazakh, Mongolian, Russian},
	"Ææ":     {Bokmal, Danish, Icelandic, Nynorsk},
	"Åå":     {Bokmal, Danish, Nynorsk, Swedish},

	"Òò": {Catalan, Italian, Latvian, Vietnamese, Yoruba},
	"Ýý": {Czech, Icelandic, Slovak, Turkish, Vietnamese},
	"Ää": {Estonian, Finnish, German, Slovak, Swedish},
	"Ââ": {Latvian, Portuguese, Romanian, Turkish, Vietnamese},
	"Àà": {Catalan, French, Italian, Portuguese, Vietnamese},

	"Üü":     {Azerbaijani, Catalan, Estonian, German, Hungarian, Spanish, Turkish},
	"ČčŠšŽž": {Bosnian, Czech, Croatian, Latvian, Lithuanian, Slovak, Slovene},

	"Çç": {Albanian, Azerbaijani, Basque, Catalan, French, Latvian, Portuguese, Turkish},
	"Öö": {Azerbaijani, Estonian, Finnish, German, Hungarian, Icelandic, Swedish, Turkish},

	"Óó":     {Catalan, Hungarian, Icelandic, Irish, Polish, Portuguese, Slovak, Spanish, Vietnamese, Yoruba},
	"ÁáÍíÚú": {Catalan, Czech, Icelandic, Irish, Hungarian, Portuguese, Slovak, Spanish, Vietnamese, Yoruba},

	"Éé": {Catalan, Czech, French, Hungarian, Icelandic, Irish, Italian, Portuguese, Slovak, Spanish, Vietnamese, Yoruba},
}
