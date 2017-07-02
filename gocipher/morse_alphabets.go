package gocipher

/*
 * https://en.wikipedia.org/wiki/Morse_code
 * https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
 */

type MorseAlphabet uint

const (
	MorseInternational MorseAlphabet = iota
	MorseSymbols
	MorseProsigns
	MorseAbbrNumbers
	MorseNonEnglish
	MorseGreek
	MorseRussian
	MorseBulgarian
	MorseHebrew
	MorseArabic
	MorsePersian
	Wabun
)

var morseAlphabets = [][][]string{
	MorseInternational: morseInternational,
	MorseSymbols:       morseSymbols,
	MorseProsigns:      morseProsigns,
	MorseAbbrNumbers:   morseAbbrNumbers,
	MorseNonEnglish:    morseNonEnglish,
	MorseGreek:         morseGreek,
	MorseRussian:       morseRussian,
	MorseBulgarian:     morseBulgarian,
	MorseHebrew:        morseHebrew,
	MorseArabic:        morseArabic,
	MorsePersian:       morsePersian,
	Wabun:              wabun,
}

// International Morse code Remommendation ITU-R M.1677-1
// http://www.itu.int/rec/R-REC-M.1677-1-200910-I/
// https://en.wikipedia.org/wiki/Morse_code
var morseInternational = [][]string{
	// Letters
	{"A", ".-"},
	{"B", "-..."},
	{"C", "-.-."},
	{"D", "-.."},
	{"E", "."},
	{"É", "..-.."}, // Accented
	{"F", "..-."},
	{"G", "--."},
	{"H", "...."},
	{"I", ".."},
	{"J", ".---"},
	{"K", "-.-"},
	{"L", ".-.."},
	{"M", "--"},
	{"N", "-."},
	{"O", "---"},
	{"P", ".--."},
	{"Q", "--.-"},
	{"R", ".-."},
	{"S", "..."},
	{"T", "-"},
	{"U", "..-"},
	{"V", "...-"},
	{"W", ".--"},
	{"X", "-..-"},
	{"Y", "-.--"},
	{"Z", "--.."},
	// Figures
	{"1", ".----"},
	{"2", "..---"},
	{"3", "...--"},
	{"4", "....-"},
	{"5", "....."},
	{"6", "-...."},
	{"7", "--..."},
	{"8", "---.."},
	{"9", "----."},
	{"0", "-----"},
	// Punctuation marks and miscellaneous signs
	{".", ".-.-.-"},                    // Full stop (period)
	{",", "--..--"},                    // Comma
	{":", "---..."},                    // Colon or division sign
	{"?", "..--.."},                    // Question mark (note of interrogation or request for repetition of a transmission not understood)
	{"'", "’", ".----."},               // Apostrophe
	{"-", "–", "-....-"},               // Hyphen or dash or subtraction sign
	{"/", "-..-."},                     // Fraction bar or division sign
	{"(", "-.--."},                     // Left-hand bracket (parenthesis)
	{")", "-.--.-"},                    // Right-hand bracket (parenthesis)
	{"\"", "“", "”", ".-..-."},         // Inverted commas (quotation marks) (before and after the words)
	{"=", "-...-"},                     // Double hyphen
	{"<SN>", "<VE>", "...-."},          // Understood
	{"<HH>", "<EEEEEEEE>", "........"}, // Error (eight dots)
	{"+", ".-.-."},                     // Cross or addition sign (<AR> digraph)
	{"<K>", "-.-"},                     // Invitation to transmit
	{"<AS>", ".-..."},                  // Wait
	{"<SK>", "<VA>", "...-.-"},         // End of work
	{"<CT>", "<KA>", "-.-.-"},          // Starting signal (to precede every transmission)
	{"×", "-..-"},                      // Multiplication sign
	{"@", ".--.-."},                    // Commercial at (<AC> digraph)
}

// Additional symbols not in ITU-R recommendation
// https://en.wikipedia.org/wiki/Morse_code
var morseSymbols = [][]string{
	{"!", "-.-.--"}, // <KW> digraph, Not in ITU-R recommendation
	{"&", ".-..."},  // <AS> digraph, Not in ITU-R recommendation
	{";", "-.-.-."},
	{"_", "..--.-"},  // Not in ITU-R recommendation
	{"$", "...-..-"}, // <SX>, Not in ITU-R recommendation
}

// Prosigns
// https://en.wikipedia.org/wiki/Prosigns_for_Morse_code
// http://www.kent-engineers.com/prosigns.htm
var morseProsigns = [][]string{
	{"<AA>", ".-.-"},                   // New Line (space down one line)
	{"<AR>", ".-.-."},                  // New Page (space down several lines); End of transmission
	{"<AS>", ".-..."},                  // Wait
	{"<BK>", "-...-.-"},                // Break; Invite receiving station to transmit
	{"<BT>", "-...-"},                  // New Paragraph (space down two lines); Pause
	{"<CL>", "-.-..-.."},               // Closing
	{"<CQ>", "-.-.--.-"},               // Calling any amateur radio station
	{"<CT>", "<KA>", "-.-.-"},          // Attention
	{"<HH>", "<EEEEEEEE>", "........"}, // Error
	{"<K>", "-.-"},                     // Invitation for any station to transmit
	{"<KN>", "-.--."},                  // Invitation for named station to transmit
	{"<NJ>", "<DO>", "-..---"},         // Shift to Wabun code
	{"<R>", ".-."},                     // All received OK
	{"<SK>", "<VA>", "...-.-"},         // End of contact
	{"<SN>", "<VE>", "...-."},          // Understood
	{"<SOS>", "...---..."},             // International distress signal
}

// Abbreviated numbers
// Conflicts with ABDGNSTUVW
// http://www.kent-engineers.com/thecode.htm
var morseAbbrNumbers = [][]string{
	{"1", ".-"},   // Conflicts with A
	{"2", "..-"},  // Conflicts with U
	{"3", ".--"},  // Conflicts with W
	{"4", "...-"}, // Conflicts with V
	{"5", "..."},  // Conflicts with S
	{"6", "-..."}, // Conflicts with B
	{"7", "--."},  // Conflicts with G
	{"8", "-.."},  // Conflicts with D
	{"9", "-."},   // Conflicts with N
	{"0", "-"},    // Conflicts with T
}

// US Navy Morse code
// https://web.archive.org/web/20101109183046/http://homepages.cwi.nl:80/~dik/english/codes/morse.html#usnavy
var morseUSNavy = [][]string{
	{"I", "."},              // Conflicts with E
	{"T", "-"},              // Also T in International
	{"N", ".."},             // Conflicts with I
	{"E", ".-"},             // Conflicts with A
	{"O", "-."},             // Conflicts with N
	{"A", "--"},             // Conflicts with M
	{"Y", "..."},            // Conflicts with S
	{"U", "..-"},            // Conflicts with U
	{"C", ".-."},            // Conflicts with R
	{"H", ".--"},            // Conflicts with W
	{"R", "-.."},            // Conflicts with D
	{"S", "-.-"},            // Conflicts with K
	{"L", "--."},            // Conflicts with G
	{"D", "---"},            // Conflicts with O
	{"1", "...."},           // Conflicts with H
	{"3", "...-"},           // Conflicts with V
	{"W", "..-."},           // Conflicts with F
	{"5", "..--"},           // Conflicts with Ü and Ŭ
	{"Q", ".-.."},           // Conflicts with L
	{"P", ".-.-"},           // Conflicts with <AA>
	{"M", "9", ".--."},      // Conflicts with P
	{"J", "V", "7", ".---"}, // Conflicts with J
	{"8", "-..."},           // Conflicts with B
	{"B", "X", "0", "-..-"}, // Conflicts with X (no conflict for X)
	{"K", "-.-."},           // Conflicts with C
	{"G", "6", "--.."},      // Conflicts with Z
	{"2", "--.-"},           // Conflicts with Q
	{"F", "4", "---."},      // Conflicts with Ó, Ö, and Ø
	{"Z", "----"},           // Conflicts with CH, Ĥ, and Š
}

var morseNonEnglish = [][]string{
	{"À", "Å", ".--.-"},
	{"Ä", "Æ", "Ą", ".-.-"},
	{"Ć", "Ĉ", "Ç", "-.-.."},
	{"CH", "Ĥ", "Š", "----"},
	{"Đ", "É", "Ę", "..-.."},
	{"Ð", "..--."},
	{"È", "Ł", ".-..-"},
	{"Ĝ", "--.-."},
	{"Ĵ", ".---."},
	{"Ń", "Ñ", "--.--"},
	{"Ó", "Ö", "Ø", "---."},
	{"Ś", "...-..."},
	{"Ŝ", "...-."}, // Conflicts with <SN> and <VE>
	{"Þ", ".--.."},
	{"Ü", "Ŭ", "..--"},
	{"Ź", "--..-."},
	{"Ż", "--..-"},
}

// Greek
// The Greek diphthongs are specified in old Greek Morse-code tables but they are
// never used in actual communication, the two vowels being sent separately.
var morseGreek = [][]string{
	// Letters
	{"Α", ".-"},
	{"Β", "-..."},
	{"Γ", "--."},
	{"Δ", "-.."},
	{"Ε", "."},
	{"Ζ", "--.."},
	{"Η", "...."},
	{"Θ", "-.-."},
	{"Ι", ".."},
	{"Κ", "-.-"},
	{"Λ", ".-.."},
	{"Μ", "--"},
	{"Ν", "-."},
	{"Ξ", "-..-"},
	{"Ο", "---"},
	{"Π", ".--."},
	{"Ρ", ".-."},
	{"Σ", "..."},
	{"Τ", "-"},
	{"Υ", "-.--"},
	{"Φ", "..-."},
	{"Χ", "----"},
	{"Ψ", "--.-"},
	{"Ω", ".--"},
	// Diphthongs
	{"HY", "...-"},
	{"OI", "---.."},
	{"AY", "..--"},
	{"YI", ".---"},
	{"EI", "..."},
	{"EY", "---."},
	{"OY", "..-"},
	{"AI", ".-.-"},
}

// Russian and Bulgarian Morse have only a few differences
var morseCyrillicPartA = [][]string{
	{"A", ".-"},
	{"Б", "-..."},
	{"В", ".--"},
	{"Г", "--."},
	{"Д", "-.."},
	{"Е", "."},
	{"Ж", "...-"},
	{"З", "--.."},
	{"И", ".."},
	{"Й", ".---"},
	{"К", "-.-"},
	{"Л", ".-.."},
	{"М", "--"},
	{"H", "-."},
	{"О", "---"},
	{"П", ".--."},
	{"P", ".-."},
	{"С", "..."},
	{"Т", "-"},
	{"У", "..-"},
	{"Ф", "..-."},
	{"Х", "...."},
	{"Ц", "-.-."},
	{"Ч", "---."},
	{"Ш", "----"},
	{"Щ", "--.-"}}
var morseRussianPart = [][]string{
	{"Ъ", "--.--"},
	{"Ы", "-.--"},
	{"Ь", "-..-"}}
var morseBulgarianPart = [][]string{
	{"Ъ", "-..-"},
	{"Ь", "-.--"}}
var morseCyrillicPartB = [][]string{
	{"Э", "..-.."},
	{"Ю", "..--"},
	{"Я", ".-.-"},
	{"1", ".----"},
	{"2", "..---"},
	{"3", "...--"},
	{"4", "....-"},
	{"5", "....."},
	{"6", "-...."},
	{"7", "--..."},
	{"8", "---.."},
	{"9", "----."},
	{"0", "-----"},
	{".", "......"},
	{",", ".-.-.-"},
	{":", "---..."},
	{";", "-.-.-"},
	{"(", "-.--.-"},
	{")", "-.--.-"},
	{"'", ".----."},
	{"\"", ".-..-."},
	{"-", "-....-"},
	{"/", "-..-."},
	{"?", "..--.."},
	{"!", "--..--"},
	{"=", "-...-"},                     // Hyphen
	{"<HH>", "<EEEEEEEE>", "........"}, // Error/redo
	{"@", ".--.-."},
}

// Russian Morse code
// https://en.wikipedia.org/wiki/Russian_Morse_code
// https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets#Cyrillic
var morseRussian = append(append(morseCyrillicPartA, morseRussianPart...), morseCyrillicPartB...)

// Bulgarian
// https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets#Cyrillic
var morseBulgarian = append(append(morseCyrillicPartA, morseBulgarianPart...), morseCyrillicPartB...)

// Hebrew
// https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets#Hebrew
var morseHebrew = [][]string{
	{"א", ".-"},
	{"ב", "-..."},
	{"ג", "--."},
	{"ד", "-.."},
	{"ה", "---"},
	{"ו", "."},
	{"ז", "--.."},
	{"ח", "...."},
	{"ט", "..-"},
	{"י", ".."},
	{"כ", "-.-"},
	{"ל", ".-.."},
	{"מ", "--"},
	{"נ", "-."},
	{"ס", "-.-."},
	{"ע", ".---"},
	{"פ", ".--."},
	{"צ", ".--"},
	{"ק", "--.-"},
	{"ר", ".-."},
	{"ש", "..."},
	{"ת", "-"},
}

// Arabic
// https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets#Arabic
var morseArabic = [][]string{
	{"ا", ".-"},
	{"ب", "-..."},
	{"ت", "-"},
	{"ث", "-.-."},
	{"ج", ".---"},
	{"ح", "...."},
	{"خ", "---"},
	{"د", "-.."},
	{"ذ", "--.."},
	{"ر", ".-."},
	{"ز", "---."},
	{"س", "..."},
	{"ش", "----"},
	{"ص", "-..-"},
	{"ض", "...-"},
	{"ط", "..-"},
	{"ظ", "-.--"},
	{"ع", ".-.-"},
	{"غ", "--."},
	{"ف", "..-."},
	{"ق", "--.-"},
	{"ك", "-.-"},
	{"ل", ".-.."},
	{"م", "--"},
	{"ن", "-."},
	{"ه", "..-.."},
	{"و", ".--"},
	{"ي", ".."},
	{"ﺀ", "."},
}

// Persian
// https://fa.wikipedia.org/wiki/%DA%A9%D8%AF_%D9%85%D9%88%D8%B1%D8%B3
// https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets#Persian
var morsePersian = [][]string{
	// Alphabet
	{"ا", ".-"},    // A
	{"ب", "-..."},  // B
	{"پ", ".--."},  // P
	{"ت", "-"},     // T
	{"ث", "-.-."},  // C
	{"ج", ".---"},  // J
	{"چ", "---."},  // Ö
	{"ح", "...."},  // H
	{"خ", "-..-"},  // X
	{"د", "-.."},   // D
	{"ذ", "...-"},  // V
	{"ر", ".-."},   // R
	{"ز", "--.."},  // Z
	{"ژ", "--."},   // G
	{"س", "..."},   // S
	{"ش", "----"},  // Š
	{"ص", ".-.-"},  // Ä
	{"ض", "..-.."}, // É
	{"ط", "..-"},   // U
	{"ظ", "-.--"},  // Y
	{"ع", "---"},   // O
	{"غ", "..--"},  // Ü
	{"ف", "..-."},  // F
	{"ق", "...---"},
	{"ک", "-.-"},  // K
	{"گ", "--.-"}, // Q
	{"ل", ".-.."}, // L
	{"م", "--"},   // M
	{"ن", "-."},   // N
	{"و", ".--"},  // W
	{"ه", "."},    // E
	{"ی", ".."},   // I
	// Numbers
	{"۰", "-----"}, // 0
	{"۱", ".----"}, // 1
	{"۲", "..---"}, // 2
	{"۳", "...--"}, // 3
	{"۴", "....-"}, // 4
	{"۵", "....."}, // 5
	{"۶", "-...."}, // 6
	{"۷", "--..."}, // 7
	{"۸", "---.."}, // 8
	{"۹", "----."}, // 9
	// Punctuation
	{".", "......"},
	{"،", ".-.-.-"}, // translates to ,
	{"؛", "-.-.-"},  // ‍translates to ;
	{":", "---..."},
	{"؟", "..--.."}, // translates to ?
	{"!", "--..--"},
	{"\n", ".-.-."},
	{"-", "-....-"},
	{"/", "------"},
	{"ـ", "..--.-"}, // translates to _
	{"(", "-.--.-"},
	{")", "-.--.-"},
}

// Japanese (Wabun code)
// https://en.wikipedia.org/wiki/Wabun_code#Expanded_katakana_Wabun_chart
var wabun = [][]string{
	// Monographs (gojūon)
	{"ア", "A", "--.--"},
	{"イ", "I", ".-"},
	{"ウ", "U", "..-"},
	{"エ", "E", "-.---"},
	{"オ", "O", ".-..."},
	{"カ", "KA", ".-.."},
	{"キ", "KI", "-.-.."},
	{"ク", "KU", "...-"},
	{"ケ", "KE", "-.--"},
	{"コ", "KO", "----"},
	{"サ", "SA", "-.-.-"},
	{"シ", "SHI", "--.-."},
	{"ス", "SU", "---.-"},
	{"セ", "SE", ".---."},
	{"ソ", "SO", "---."},
	{"タ", "TA", "-."},
	{"チ", "CHI", "..-."},
	{"ツ", "TSU", ".--."},
	{"テ", "TE", ".-.--"},
	{"ト", "TO", "..-.."},
	{"ナ", "NA", ".-."},
	{"ニ", "NI", "-.-."},
	{"ヌ", "NU", "...."},
	{"ネ", "NE", "--.-"},
	{"ノ", "NO", "..--"},
	{"ハ", "HA", "-..."},
	{"ヒ", "HI", "--..-"},
	{"フ", "FU", "--.."},
	{"ヘ", "HE", "."},
	{"ホ", "HO", "-.."},
	{"マ", "MA", "-..-"},
	{"ミ", "MI", "..-.-"},
	{"ム", "MU", "-"},
	{"メ", "ME", "-...-"},
	{"モ", "MO", "-..-."},
	{"ヤ", "YA", ".--"},
	{"ユ", "YU", "-..--"},
	{"ヨ", "YO", "--"},
	{"ラ", "RA", "..."},
	{"リ", "RI", "--."},
	{"ル", "RU", "-.--."},
	{"レ", "RE", "---"},
	{"ロ", "RO", ".-.-"},
	{"ワ", "WA", "-.-"},
	{"ヰ", "WI", ".-..-"}, // (wi)
	{"ヱ", "WE", ".--.."}, // (we)
	{"ヲ", "WO", ".---"},
	{"ン", "N", ".-.-."},
	{"゛", "", ".."},      // Dakuten ◌゛ (Diacritic)     Combining version:  ゙
	{"゜", "", "..--."},   // Handakuten ◌゜ (Diacritic)  Combining version:  ゚
	{"ー", "", ".--.-"},   // Chōonpuー (Long Vowel)      Combining version: ̄
	{"、", ",", ".-.-.-"}, // Comma
	{"。", ".", ".-.-.."}, // Full stop

	// Diacritics (gojūon with dakuten)
	{"ガ", "GA", ".-.. .."},
	{"ギ", "GI", "-.-.. .."},
	{"グ", "GU", "...- .."},
	{"ゲ", "GE", "-.-- .."},
	{"ゴ", "GO", "---- .."},
	{"ザ", "ZA", "-.-.- .."},
	{"ジ", "JI", "--.-. .."},
	{"ズ", "ZU", "---.- .."},
	{"ゼ", "ZE", ".---. .."},
	{"ゾ", "ZO", "---. .."},
	{"ダ", "DA", "-. .."},
	{"ヂ", "JI", "..-. .."},
	{"ヅ", "ZU", ".--. .."},
	{"デ", "DE", ".-.-- .."},
	{"ド", "DO", "..-.. .."},
	{"バ", "BA", "-... .."},
	{"ビ", "BI", "--..- .."},
	{"ブ", "BU", "--.. .."},
	{"ベ", "BE", ". .."},
	{"ボ", "BO", "-.. .."},

	// Diacritics (gojūon with handakuten)
	{"パ", "PA", "-... ..--."},
	{"ピ", "PI", "--.. ..--."},
	{"プ", "PU", "--.. ..--."},
	{"ペ", "PE", ". ..--."},
	{"ポ", "PO", "-.. ..--."},

	// Digraphs (yōon)
	{"キャ", "KYA", "-.-.. .--"},
	{"キュ", "KYU", "-.-.. -..--"},
	{"キョ", "KYO", "-.-.. --"},
	{"シャ", "SHA", "--.-. .--"},
	{"シュ", "SHU", "--.-. -..--"},
	{"ショ", "SHO", "--.-. --"},
	{"チャ", "CHA", "..-. .--"},
	{"チュ", "CHU", "..-. -..--"},
	{"チョ", "CHO", "..-. --"},
	{"ニャ", "NYA", "-.-. .--"},
	{"ニュ", "NYU", "-.-. -..--"},
	{"ニョ", "NYO", "-.-. --"},
	{"ヒャ", "HYA", "--..- .--"},
	{"ヒュ", "HYU", "--..- -..--"},
	{"ヒョ", "HYO", "--..- --"},
	{"ミャ", "MYA", "..-.- .--"},
	{"ミュ", "MYU", "..-.- -..--"},
	{"ミョ", "MYO", "..-.- --"},
	{"リャ", "RYA", "--. .--"},
	{"リュ", "RYU", "--. -..--"},
	{"リョ", "RYO", "--. --"},

	// Digraphs with diacritics (yōon with dakuten)
	{"ギャ", "GYA", "-.-.. .. .--"},
	{"ギュ", "GYU", "-.-.. .. -..--"},
	{"ギョ", "GYO", "-.-.. .. --"},
	{"ジャ", "JA", "--.-. .. .--"},
	{"ジュ", "JU", "--.-. .. -..--"},
	{"ジョ", "JO", "--.-. .. --"},
	{"ヂャ", "JA", "..-. .. .--"},
	{"ヂュ", "JU", "..-. .. -..--"},
	{"ヂョ", "JO", "..-. .. --"},
	{"ビャ", "BYA", "--..- .. .--"},
	{"ビュ", "BYU", "--..- .. -..--"},
	{"ビョ", "BYO", "--..- .. --"},

	// Digraphs with diacritics (yōon with handakuten)
	{"ピャ", "PYA", "--..- ..--. .--"},
	{"ピュ", "PYU", "--..- ..--. -..--"},
	{"ピョ", "PYO", "--..- ..--. --"},
}

// Chinese (Chinese telegraph code)
// https://en.wikipedia.org/wiki/Chinese_telegraph_code
var morseChinese = [][]string{}

// Korean (SKATS - Standard Korean Alphabet Transliteration System)
// https://en.wikipedia.org/wiki/SKATS
var morseKorean = [][]string{}

// Thai
// https://th.wikipedia.org/wiki/%E0%B8%A3%E0%B8%AB%E0%B8%B1%E0%B8%AA%E0%B8%A1%E0%B8%AD%E0%B8%A3%E0%B9%8C%E0%B8%AA
var morseThai = [][]string{}
