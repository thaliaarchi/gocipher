package gocipher

/*
 * https://en.wikipedia.org/wiki/Morse_code
 * https://en.wikipedia.org/wiki/Morse_code_for_non-Latin_alphabets
 * https://en.wikipedia.org/wiki/Russian_Morse_code
 * https://fa.wikipedia.org/wiki/%DA%A9%D8%AF_%D9%85%D9%88%D8%B1%D8%B3
 * https://en.wikipedia.org/wiki/Wabun_code
 */

type MorseAlphabet uint

const (
	// MorseInternational is the International Morse code Remommendation ITU-R M.1677-1
	// http://www.itu.int/rec/R-REC-M.1677-1-200910-I/
	MorseInternational MorseAlphabet = iota
	MorseSymbols
	MorseProsigns
	MorseNonEnglish
	MorseGreek
	MorseRussian
	MorseBulgarian
	MorseHebrew
	MorseArabic
	MorsePersian
)

var morseAlphabets = [][][]string{
	MorseInternational: morseInternational,
	MorseSymbols:       morseSymbols,
	MorseProsigns:      morseProsigns,
	MorseNonEnglish:    morseNonEnglish,
	MorseGreek:         morseGreek,
	MorseRussian:       morseRussian,
	MorseBulgarian:     morseBulgarian,
	MorseHebrew:        morseHebrew,
	MorseArabic:        morseArabic,
	MorsePersian:       morsePersian,
}

var morseInternational = [][]string{
	// Letters
	{"A", ".-"},
	{"B", "-..."},
	{"C", "-.-."},
	{"D", "-.."},
	{"E", "."},
	{"É", "..-.."}, // accented
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
	{".", ".-.-.-"},                // Full stop (period)
	{",", "--..--"},                // Comma
	{":", "---..."},                // Colon or division sign
	{"?", "..--.."},                // Question mark (note of interrogation or request for repetition of a transmission not understood)
	{"'", ".----."},                // Apostrophe ’
	{"-", "-....-"},                // Hyphen or dash or subtraction sign –
	{"/", "-..-."},                 // Fraction bar or division sign
	{"(", "-.--."},                 // Left-hand bracket (parenthesis)
	{")", "-.--.-"},                // Right-hand bracket (parenthesis)
	{"\"", ".-..-."},               // Inverted commas (quotation marks) (before and after the words) “”
	{"=", "-...-"},                 // Double hyphen
	{"<Understood>", "...-."},      // Understood
	{"<Error>", "........"},        // Error (eight dots)
	{"+", ".-.-."},                 // Cross or addition sign
	{"K", "-.-"},                   // Invitation to transmit
	{"<Wait>", ".-..."},            // Wait
	{"<End of work>", "...-.-"},    // End of work
	{"<Starting signal>", "-.-.-"}, // Starting signal (to precede every transmission)
	{"×", "-..-"},                  // Multiplication sign
	{"@", ".--.-."},                // Commercial at
}

var morseSymbols = [][]string{
	{"!", "-.-.--"}, // <KW>
	{"&", ".-..."},  // <AS>, Not in ITU-R recommendation
	{";", "-.-.-."},
	{"+", ".-.-."},   // <AR>
	{"_", "..--.-"},  // Not in ITU-R recommendation
	{"$", "...-..-"}, // <SX>, Not in ITU-R recommendation
	{"@", ".--.-."},  // <AC>
}

// https://en.wikipedia.org/wiki/Prosigns_for_Morse_code
// http://www.kent-engineers.com/prosigns.htm
var morseProsigns = [][]string{
	{"<AA>", ".-.-"},                     // New Line (space down one line)
	{"<AR>", ".-.-."},                    // New Page (space down several lines); End of transmission
	{"<AS>", ".-..."},                    // Wait
	{"<BK>", "-...-.-"},                  // Break; Invite receiving station to transmit
	{"<BT>", "-...-"},                    // New Paragraph (space down two lines); Pause
	{"<CL>", "-.-..-.."},                 // Closing
	{"<CQ>", "-.-.--.-"},                 // Calling any amateur radio station
	{"<CT>", "-.-.-"}, {"<KA>", "-.-.-"}, // Attention
	{"<HH>", "........"}, {"<EEEEEEEE>", "........"}, // Error
	{"<K>", "-.-"},                         // Invitation for any station to transmit
	{"<KN>", "-.--."},                      // Invitation for named station to transmit
	{"<NJ>", "-..---"}, {"<DO>", "-..---"}, // Shift to Wabun code
	{"<R>", ".-."},                         // All received OK
	{"<SK>", "...-.-"}, {"<VA>", "...-.-"}, // End of contact
	{"<SN>", "...-."}, {"<VE>", "...-."}, // Understood
	{"<SOS>", "...---..."}, // International distress signal
}

var morseNonEnglish = [][]string{
	{"À", ".--.-"}, {"Å", ".--.-"},
	{"Ä", ".-.-"}, {"Æ", ".-.-"}, {"Ą", ".-.-"},
	{"Ć", "-.-.."}, {"Ĉ", "-.-.."}, {"Ç", "-.-.."},
	{"CH", "----"}, {"Ĥ", "----"}, {"Š", "----"},
	{"Đ", "..-.."}, {"É", "..-.."}, {"Ę", "..-.."},
	{"Ð", "..--."},
	{"È", ".-..-"}, {"Ł", ".-..-"},
	{"Ĝ", "--.-."},
	{"Ĵ", ".---."},
	{"Ń", "--.--"}, {"Ñ", "--.--"},
	{"Ó", "---."}, {"Ö", "---."}, {"Ø", "---."},
	{"Ś", "...-..."},
	{"Ŝ", "...-."}, // <SN> <VE>
	{"Þ", ".--.."},
	{"Ü", "..--"}, {"Ŭ", "..--"},
	{"Ź", "--..-."},
	{"Ż", "--..-"},
}

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
	// The Greek diphthongs are specified in old Greek Morse-code tables
	// but they are never used in actual communication,
	// the two vowels being sent separately.
	{"HY", "...-"},
	{"OI", "---.."},
	{"AY", "..--"},
	{"YI", ".---"},
	{"EI", "..."},
	{"EY", "---."},
	{"OY", "..-"},
	{"AI", ".-.-"},
}

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
	{"Hyphen", "-...-"}, // Hyphen is above, so...
	{"Error/redo", "........"},
	{"@", ".--.-."},
}

var morseRussian = append(append(morseCyrillicPartA, morseRussianPart...), morseCyrillicPartB...)
var morseBulgarian = append(append(morseCyrillicPartA, morseBulgarianPart...), morseCyrillicPartB...)

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

var morsePersian = [][]string{
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
	{"ق", "---..."},
	{"ک", "-.-"},  // K
	{"گ", "--.-"}, // Q
	{"ل", ".-.."}, // L
	{"م", "--"},   // M
	{"ن", "-."},   // N
	{"و", ".--"},  // W
	{"ه", "."},    // E
	{"ی", ".."},   // I

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

var wabun = [][]string{}
