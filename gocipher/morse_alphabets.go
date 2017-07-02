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
	MorseProsignsMultiLine
	MorseAbbrNumbers
	MorseAbbrNumbers2
	MorseNonEnglish
	MorseGreek
	MorseRussian
	MorseBulgarian
	MorseHebrew
	MorseArabic
	MorsePersian
	Wabun
	//SKATS
	MorseThai
)

var morseAlphabets = [][][]string{
	MorseInternational:     morseInternational,
	MorseSymbols:           morseSymbols,
	MorseProsigns:          morseProsigns,
	MorseProsignsMultiLine: morseProsignsMultiLine,
	MorseAbbrNumbers:       morseAbbrNumbers,
	MorseAbbrNumbers2:      morseAbbrNumbers2,
	MorseNonEnglish:        morseNonEnglish,
	MorseGreek:             morseGreek,
	MorseRussian:           morseRussian,
	MorseBulgarian:         morseBulgarian,
	MorseHebrew:            morseHebrew,
	MorseArabic:            morseArabic,
	MorsePersian:           morsePersian,
	Wabun:                  wabun,
	MorseThai:              morseThai,
	//SKATS:                skats,
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
	{"'", "’", "‘", ".----."},          // Apostrophe
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
// https://infogalactic.com/info/Prosigns_for_Morse_code
// http://www.kent-engineers.com/prosigns.htm
// http://www.worldlibrary.org/articles/prosigns_for_morse_code
var morseProsigns = [][]string{
	{"<AA>", ".-.-"},                   // New Line (space down one line)
	{"<AR>", ".-.-."},                  // New Page (space down several lines). End of transmission
	{"<AS>", ".-..."},                  // Wait
	{"<BK>", "-...-.-"},                // Break. Invite receiving station to transmit
	{"<BT>", "-...-"},                  // New Paragraph (space down two lines). Pause
	{"<CL>", "-.-..-.."},               // Closing
	{"<CQ>", "-.-.--.-"},               // Calling any amateur radio station
	{"<CT>", "<KA>", "-.-.-"},          // Attention
	{"<HH>", "<EEEEEEEE>", "........"}, // Error
	{"<K>", "-.-"},                     // Invitation for any station to transmit
	{"<KN>", "-.--."},                  // Invitation for named station to transmit
	{"<NJ>", "<DO>", "-..---"},         // Shift to Wabun code
	{"<SK>", "<VA>", "...-.-"},         // End of contact
	{"<SN>", "<VE>", "...-."},          // Understood
	{"<SOS>", "...---..."},             // International distress signal
}

// https://en.wikipedia.org/wiki/Prosigns_for_Morse_code
var morseProsignsMultiLine = [][]string{
	{"CR-LF", "\r\n", "\n", ".-.-"},         // <AA> Typewritten as CR-LF
	{"CR-LF-LF", "\r\n\n", "\n\n", "-...-"}, // <BT> Typewritten as CR-LF-LF. Single-line display may use printed "=".
	{"CR-LF-LF-LF", "CR-LF-LF-LF-LF", "CR-LF-LF-LF-LF-LF", "\r\n\n\n", "\r\n\n\n\n", "\r\n\n\n\n\n",
		"\n\n\n", "\n\n\n\n", "\n\n\n\n\n", ".-.-."}, // <AR> Space down several lines. Single-line display may use printed "+".
}

// Abbreviated numbers
// Conflicts with International ABDENYUV
// http://cromwell-intl.com/radio/morse-code.html
var morseAbbrNumbers = [][]string{
	{"1", ".-"},    // A
	{"2", "..-"},   // U
	{"3", "...-"},  // V
	{"4", "....-"}, // 4 (no conflict)
	{"5", "....."}, // 5 (no conflict)
	{"5", "."},     // E
	{"6", "-...."}, // 6 (no conflict)
	{"7", "-..."},  // B
	{"8", "-.."},   // D
	{"9", "-."},    // N
	{"0", "-"},     // T
}

// Abbreviated numbers
// Conflicts with ABDGNSTUVW
// http://www.kent-engineers.com/thecode.htm
var morseAbbrNumbers2 = [][]string{
	{"1", ".-"},   // A
	{"2", "..-"},  // U
	{"3", ".--"},  // W
	{"4", "...-"}, // V
	{"5", "..."},  // S
	{"6", "-..."}, // B
	{"7", "--."},  // G
	{"8", "-.."},  // D
	{"9", "-."},   // N
	{"0", "-"},    // T
}

// US Navy Morse code
// https://web.archive.org/web/20101109183046/http://homepages.cwi.nl:80/~dik/english/codes/morse.html#usnavy
var morseUSNavy = [][]string{
	{"I", "."},              // E
	{"T", "-"},              // T (no conflict)
	{"N", ".."},             // I
	{"E", ".-"},             // A
	{"O", "-."},             // N
	{"A", "--"},             // M
	{"Y", "..."},            // S
	{"U", "..-"},            // U
	{"C", ".-."},            // R
	{"H", ".--"},            // W
	{"R", "-.."},            // D
	{"S", "-.-"},            // K
	{"L", "--."},            // G
	{"D", "---"},            // O
	{"1", "...."},           // H
	{"3", "...-"},           // V
	{"W", "..-."},           // F
	{"5", "..--"},           // Ü and Ŭ
	{"Q", ".-.."},           // L
	{"P", ".-.-"},           // <AA>
	{"M", "9", ".--."},      // P
	{"J", "V", "7", ".---"}, // J
	{"8", "-..."},           // B
	{"B", "X", "0", "-..-"}, // X (no conflict for X)
	{"K", "-.-."},           // C
	{"G", "6", "--.."},      // Z
	{"2", "--.-"},           // Q
	{"F", "4", "---."},      // Ó, Ö, and Ø
	{"Z", "----"},           // CH, Ĥ, and Š
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
	{"Ŝ", "...-."}, // <SN> and <VE>
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
// https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets#Cyrillic
// http://www.cvni.net/radio/nsnl/nsnl011/nsnl11morse.html
var morseCyrillicPartA = [][]string{
	{"A", ".-"},     // A
	{"Б", "-..."},   // B
	{"В", ".--"},    // W
	{"Г", "--."},    // G
	{"Д", "-.."},    // D
	{"Е", "Ё", "."}, // E
	{"Ж", "...-"},   // V
	{"З", "--.."},   // Z
	{"И", ".."},     // I
	{"Й", ".---"},   // J
	{"К", "-.-"},    // K
	{"Л", ".-.."},   // L
	{"М", "--"},     // M
	{"H", "-."},     // N
	{"О", "---"},    // O
	{"П", ".--."},   // P
	{"P", ".-."},    // R
	{"С", "..."},    // S
	{"Т", "-"},      // T
	{"У", "..-"},    // U
	{"Ф", "..-."},   // F
	{"Х", "...."},   // H
	{"Ц", "-.-."},   // C
	{"Ч", "---."},   // Ö
	{"Ш", "----"},   // CH
	{"Щ", "--.-"}}   // Q
var morseRussianPart = [][]string{
	{"Ъ", "--.--"}, // Ñ
	{"Ы", "-.--"},  // Y
	{"Ь", "-..-"}}  // X
var morseBulgarianPart = [][]string{
	{"Ъ", "-..-"}, // X
	{"Ь", "-.--"}} // Y
var morseCyrillicPartB = [][]string{
	{"Э", "..-.."}, // É
	{"Ю", "..--"},  // Ü
	{"Я", ".-.-"},  // Ä
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
	{"א", ".-"},   // A
	{"ב", "-..."}, // B
	{"ג", "--."},  // G
	{"ד", "-.."},  // D
	{"ה", "---"},  // O
	{"ו", "."},    // E
	{"ז", "--.."}, // Z
	{"ח", "...."}, // H
	{"ט", "..-"},  // U
	{"י", ".."},   // I
	{"כ", "-.-"},  // K
	{"ל", ".-.."}, // L
	{"מ", "--"},   // M
	{"נ", "-."},   // N
	{"ס", "-.-."}, // C
	{"ע", ".---"}, // J
	{"פ", ".--."}, // P
	{"צ", ".--"},  // W
	{"ק", "--.-"}, // Q
	{"ר", ".-."},  // R
	{"ש", "..."},  // S
	{"ת", "-"},    // T
}

// Arabic
// https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets#Arabic
var morseArabic = [][]string{
	{"ا", ".-"},    // A
	{"ب", "-..."},  // B
	{"ت", "-"},     // T
	{"ث", "-.-."},  // C
	{"ج", ".---"},  // J
	{"ح", "...."},  // H
	{"خ", "---"},   // O
	{"د", "-.."},   // D
	{"ذ", "--.."},  // Z
	{"ر", ".-."},   // R
	{"ز", "---."},  // Ö
	{"س", "..."},   // S
	{"ش", "----"},  // CH
	{"ص", "-..-"},  // X
	{"ض", "...-"},  // V
	{"ط", "..-"},   // U
	{"ظ", "-.--"},  // Y
	{"ع", ".-.-"},  // Ä
	{"غ", "--."},   // G
	{"ف", "..-."},  // F
	{"ق", "--.-"},  // Q
	{"ك", "-.-"},   // K
	{"ل", ".-.."},  // L
	{"م", "--"},    // M
	{"ن", "-."},    // N
	{"ه", "..-.."}, // É
	{"و", ".--"},   // W
	{"ي", ".."},    // I
	{"ﺀ", "."},     // E
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
	{"،", ".-.-.-"}, // ,
	{"؛", "-.-.-"},  // ‍;
	{":", "---..."},
	{"؟", "..--.."}, // ?
	{"!", "--..--"},
	{"\n", ".-.-."},
	{"-", "-....-"},
	{"/", "------"},
	{"ـ", "..--.-"}, // _
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

// Korean (SKATS - Standard Korean Alphabet Transliteration System)
// https://en.wikipedia.org/wiki/SKATS
// https://ko.wikipedia.org/wiki/%EB%AA%A8%EC%8A%A4_%EB%B6%80%ED%98%B8
var skats = [][]string{
	{"ㄱ", "L"},
	{"ㄴ", "F"},
	{"ㄷ", "B"},
	{"ㄹ", "V"},
	{"ㅁ", "M"},
	{"ㅂ", "W"},
	{"ㅅ", "G"},
	{"ㅇ", "K"},
	{"ㅈ", "P"},
	{"ㅊ", "C"},
	{"ㅋ", "X"},
	{"ㅌ", "Z"},
	{"ㅍ", "O"},
	{"ㅎ", "J"},
	{"ㅏ", "E"},
	{"ㅑ", "I"},
	{"ㅓ", "T"},
	{"ㅕ", "S"},
	{"ㅗ", "A"},
	{"ㅛ", "N"},
	{"ㅜ", "H"},
	{"ㅠ", "R"},
	{"ㅡ", "D"},
	{"ㅣ", "U"},
	{"ㅔ", "TU"},
	{"ㅐ", "EU"},
	{"ㅖ", "SU"},
	{"ㅒ", "IU"},
}

// Thai
// https://th.wikipedia.org/wiki/%E0%B8%A3%E0%B8%AB%E0%B8%B1%E0%B8%AA%E0%B8%A1%E0%B8%AD%E0%B8%A3%E0%B9%8C%E0%B8%AA
var morseThai = [][]string{
	// Consonants (พยัญชนะ)
	{"ก", "--."},     // G
	{"ข ฃ", "-.-."},  // C
	{"ค ฅ ฆ", "-.-"}, // K
	{"ง", "-.--."},   // (
	{"จ", "-..-."},   // /
	{"ฉ", "----"},    // CH, Ĥ, and Š
	{"ช", "-..-"},    // X
	{"ซ", "--.."},    // Z
	{"ญ", ".---"},    // J
	{"ด ฎ", "-.."},   // D
	{"ต ฏ", "-"},     // T
	{"ถ ฐ", "-.-.."}, // Ć, Ĉ, and Ç
	{"ท ธ ฑ ฒ", "-..--"},
	{"น ณ", "-."},    // N
	{"บ", "-..."},    // B
	{"ป", ".--."},    // P
	{"ผ", "--.-"},    // Q
	{"ฝ", "-.-.-"},   // <CT> and <KA>
	{"พ ภ", ".--.."}, // Þ
	{"ฟ", "..-."},    // F
	{"ม", "--"},      // M
	{"ย", "-.--"},    // Y
	{"ร", ".-."},     // R
	{"ล ฬ", ".-.."},  // L
	{"ว", ".--"},     // W
	{"ศ ษ ส", "..."}, // S
	{"ห", "...."},    // H
	{"อ", "-...-"},   // =
	{"ฮ", "--.--"},   // Ń and Ñ
	{"ฤ ฤๅ", ".-.--"},
	// สระ
	{"สระ ะ", ".-..."},  // <AS>
	{"สระ า", ".-"},     // A
	{"สระ อิ", "..-.."}, // Đ, É, and Ę
	{"สระ อี", ".."},    // I
	{"สระ อึ", "..--."}, // Ð
	{"สระ อื", "..--"},  // Ü and Ŭ
	{"สระ อุ", "..-.-"},
	{"สระ อู", "---."},  // Ó, Ö, and Ø
	{"สระ เ", "."},      // E
	{"สระ แ", ".-.-"},   // <AA>
	{"ไ ใ", ".-..-"},    // È, Ł
	{"โ", "---"},        // O
	{"สระ อำ", "...-."}, // <SN>, <VE>, and Ŝ
	// วรรณยุกต์
	{"ไม้เอก", "..-"},     // U
	{"ไม้โท", "...-"},     // V
	{"ไม้ตรี", "--..."},   // 7
	{"ไม้จัตวา", ".-.-."}, // <AR>
	// เครื่องหมาย
	{"ไม้หันอากาศ", ".--.-"}, // À and Å
	{"ไม้ไต่คู้", "---.."},   // 8
	{"การันต์", "--..-"},     // Ż
	{"ไม้ยมก", "-.---"},
	{"ฯ", "--.-."}, // Ĝ
	{"ฯลฯ", "---.-"},
	{" ", ".-..-."},   // "
	{"( )", "-.--.-"}, // )
}
