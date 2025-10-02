package glib

type GBookmarkFileError int32
type GChecksumType int32
type GConvertError int32
type GDateDMY int32
type GDateMonth int32
type GDateWeekday int32
type GErrorType int32
type GFileError int32
type GIOChannelError int32
type GIOError int32
type GIOStatus int32
type GKeyFileError int32
type GLogWriterOutput int32
type GMarkupError int32
type GNormalizeMode int32
type GNumberParserError int32
type GOnceStatus int32
type GOptionArg int32
type GOptionError int32
type GRegexError int32
type GSeekType int32
type GShellError int32
type GSliceConfig int32
type GSpawnError int32
type GTestFileType int32
type GTestLogType int32
type GTestResult int32
type GThreadError int32
type GTimeType int32
type GTokenType int32
type GTraverseType int32
type GUnicodeBreakType int32
type GUnicodeScript int32
type GUnicodeType int32
type GUriError int32
type GUserDirectory int32
type GVariantClass int32
type GVariantParseError int32
type GWin32OSType int32
type GAsciiType int32
type GFileSetContentsFlags int32
type GFileTest int32
type GFormatSizeFlags int32
type GHookFlagMask int32
type GIOCondition int32
type GIOFlags int32
type GKeyFileFlags int32
type GLogLevelFlags int32
type GMainContextFlags int32
type GMarkupCollectType int32
type GMarkupParseFlags int32
type GOptionFlags int32
type GRegexCompileFlags int32
type GRegexMatchFlags int32
type GSpawnFlags int32
type GTestSubprocessFlags int32
type GTestTrapFlags int32
type GTraverseFlags int32
type GUriFlags int32
type GUriHideFlags int32
type GUriParamsFlags int32

const (
	GBookmarkFileErrorInvalidUri       GBookmarkFileError = 0
	GBookmarkFileErrorInvalidValue     GBookmarkFileError = 1
	GBookmarkFileErrorAppNotRegistered GBookmarkFileError = 2
	GBookmarkFileErrorUriNotFound      GBookmarkFileError = 3
	GBookmarkFileErrorRead             GBookmarkFileError = 4
	GBookmarkFileErrorUnknownEncoding  GBookmarkFileError = 5
	GBookmarkFileErrorWrite            GBookmarkFileError = 6
	GBookmarkFileErrorFileNotFound     GBookmarkFileError = 7

	GChecksumTypeMd5    GChecksumType = 0
	GChecksumTypeSha1   GChecksumType = 1
	GChecksumTypeSha256 GChecksumType = 2
	GChecksumTypeSha512 GChecksumType = 3
	GChecksumTypeSha384 GChecksumType = 4

	GConvertErrorNoConversion    GConvertError = 0
	GConvertErrorIllegalSequence GConvertError = 1
	GConvertErrorFailed          GConvertError = 2
	GConvertErrorPartialInput    GConvertError = 3
	GConvertErrorBadUri          GConvertError = 4
	GConvertErrorNotAbsolutePath GConvertError = 5
	GConvertErrorNoMemory        GConvertError = 6
	GConvertErrorEmbeddedNul     GConvertError = 7

	GDateDMYDay   GDateDMY = 0
	GDateDMYMonth GDateDMY = 1
	GDateDMYYear  GDateDMY = 2

	GDateMonthBadMonth  GDateMonth = 0
	GDateMonthJanuary   GDateMonth = 1
	GDateMonthFebruary  GDateMonth = 2
	GDateMonthMarch     GDateMonth = 3
	GDateMonthApril     GDateMonth = 4
	GDateMonthMay       GDateMonth = 5
	GDateMonthJune      GDateMonth = 6
	GDateMonthJuly      GDateMonth = 7
	GDateMonthAugust    GDateMonth = 8
	GDateMonthSeptember GDateMonth = 9
	GDateMonthOctober   GDateMonth = 10
	GDateMonthNovember  GDateMonth = 11
	GDateMonthDecember  GDateMonth = 12

	GDateWeekdayBadWeekday GDateWeekday = 0
	GDateWeekdayMonday     GDateWeekday = 1
	GDateWeekdayTuesday    GDateWeekday = 2
	GDateWeekdayWednesday  GDateWeekday = 3
	GDateWeekdayThursday   GDateWeekday = 4
	GDateWeekdayFriday     GDateWeekday = 5
	GDateWeekdaySaturday   GDateWeekday = 6
	GDateWeekdaySunday     GDateWeekday = 7

	GErrorTypeUnknown           GErrorType = 0
	GErrorTypeUnexpEof          GErrorType = 1
	GErrorTypeUnexpEofInString  GErrorType = 2
	GErrorTypeUnexpEofInComment GErrorType = 3
	GErrorTypeNonDigitInConst   GErrorType = 4
	GErrorTypeDigitRadix        GErrorType = 5
	GErrorTypeFloatRadix        GErrorType = 6
	GErrorTypeFloatMalformed    GErrorType = 7

	GFileErrorExist       GFileError = 0
	GFileErrorIsdir       GFileError = 1
	GFileErrorAcces       GFileError = 2
	GFileErrorNametoolong GFileError = 3
	GFileErrorNoent       GFileError = 4
	GFileErrorNotdir      GFileError = 5
	GFileErrorNxio        GFileError = 6
	GFileErrorNodev       GFileError = 7
	GFileErrorRofs        GFileError = 8
	GFileErrorTxtbsy      GFileError = 9
	GFileErrorFault       GFileError = 10
	GFileErrorLoop        GFileError = 11
	GFileErrorNospc       GFileError = 12
	GFileErrorNomem       GFileError = 13
	GFileErrorMfile       GFileError = 14
	GFileErrorNfile       GFileError = 15
	GFileErrorBadf        GFileError = 16
	GFileErrorInval       GFileError = 17
	GFileErrorPipe        GFileError = 18
	GFileErrorAgain       GFileError = 19
	GFileErrorIntr        GFileError = 20
	GFileErrorIo          GFileError = 21
	GFileErrorPerm        GFileError = 22
	GFileErrorNosys       GFileError = 23
	GFileErrorFailed      GFileError = 24

	GIOChannelErrorFbig     GIOChannelError = 0
	GIOChannelErrorInval    GIOChannelError = 1
	GIOChannelErrorIo       GIOChannelError = 2
	GIOChannelErrorIsdir    GIOChannelError = 3
	GIOChannelErrorNospc    GIOChannelError = 4
	GIOChannelErrorNxio     GIOChannelError = 5
	GIOChannelErrorOverflow GIOChannelError = 6
	GIOChannelErrorPipe     GIOChannelError = 7
	GIOChannelErrorFailed   GIOChannelError = 8

	GIOErrorNone    GIOError = 0
	GIOErrorAgain   GIOError = 1
	GIOErrorInval   GIOError = 2
	GIOErrorUnknown GIOError = 3

	GIOStatusError  GIOStatus = 0
	GIOStatusNormal GIOStatus = 1
	GIOStatusEof    GIOStatus = 2
	GIOStatusAgain  GIOStatus = 3

	GKeyFileErrorUnknownEncoding GKeyFileError = 0
	GKeyFileErrorParse           GKeyFileError = 1
	GKeyFileErrorNotFound        GKeyFileError = 2
	GKeyFileErrorKeyNotFound     GKeyFileError = 3
	GKeyFileErrorGroupNotFound   GKeyFileError = 4
	GKeyFileErrorInvalidValue    GKeyFileError = 5

	GLogWriterOutputHandled   GLogWriterOutput = 1
	GLogWriterOutputUnhandled GLogWriterOutput = 0

	GMarkupErrorBadUtf8          GMarkupError = 0
	GMarkupErrorEmpty            GMarkupError = 1
	GMarkupErrorParse            GMarkupError = 2
	GMarkupErrorUnknownElement   GMarkupError = 3
	GMarkupErrorUnknownAttribute GMarkupError = 4
	GMarkupErrorInvalidContent   GMarkupError = 5
	GMarkupErrorMissingAttribute GMarkupError = 6

	GNormalizeModeDefault        GNormalizeMode = 0
	GNormalizeModeNfd            GNormalizeMode = 0
	GNormalizeModeDefaultCompose GNormalizeMode = 1
	GNormalizeModeNfc            GNormalizeMode = 1
	GNormalizeModeAll            GNormalizeMode = 2
	GNormalizeModeNfkd           GNormalizeMode = 2
	GNormalizeModeAllCompose     GNormalizeMode = 3
	GNormalizeModeNfkc           GNormalizeMode = 3

	GNumberParserErrorInvalid     GNumberParserError = 0
	GNumberParserErrorOutOfBounds GNumberParserError = 1

	GOnceStatusNotcalled GOnceStatus = 0
	GOnceStatusProgress  GOnceStatus = 1
	GOnceStatusReady     GOnceStatus = 2

	GOptionArgNone          GOptionArg = 0
	GOptionArgString        GOptionArg = 1
	GOptionArgInt           GOptionArg = 2
	GOptionArgCallback      GOptionArg = 3
	GOptionArgFilename      GOptionArg = 4
	GOptionArgStringArray   GOptionArg = 5
	GOptionArgFilenameArray GOptionArg = 6
	GOptionArgDouble        GOptionArg = 7
	GOptionArgInt64         GOptionArg = 8

	GOptionErrorUnknownOption GOptionError = 0
	GOptionErrorBadValue      GOptionError = 1
	GOptionErrorFailed        GOptionError = 2

	GRegexErrorCompile                                  GRegexError = 0
	GRegexErrorOptimize                                 GRegexError = 1
	GRegexErrorReplace                                  GRegexError = 2
	GRegexErrorMatch                                    GRegexError = 3
	GRegexErrorInternal                                 GRegexError = 4
	GRegexErrorStrayBackslash                           GRegexError = 101
	GRegexErrorMissingControlChar                       GRegexError = 102
	GRegexErrorUnrecognizedEscape                       GRegexError = 103
	GRegexErrorQuantifiersOutOfOrder                    GRegexError = 104
	GRegexErrorQuantifierTooBig                         GRegexError = 105
	GRegexErrorUnterminatedCharacterClass               GRegexError = 106
	GRegexErrorInvalidEscapeInCharacterClass            GRegexError = 107
	GRegexErrorRangeOutOfOrder                          GRegexError = 108
	GRegexErrorNothingToRepeat                          GRegexError = 109
	GRegexErrorUnrecognizedCharacter                    GRegexError = 112
	GRegexErrorPosixNamedClassOutsideClass              GRegexError = 113
	GRegexErrorUnmatchedParenthesis                     GRegexError = 114
	GRegexErrorInexistentSubpatternReference            GRegexError = 115
	GRegexErrorUnterminatedComment                      GRegexError = 118
	GRegexErrorExpressionTooLarge                       GRegexError = 120
	GRegexErrorMemoryError                              GRegexError = 121
	GRegexErrorVariableLengthLookbehind                 GRegexError = 125
	GRegexErrorMalformedCondition                       GRegexError = 126
	GRegexErrorTooManyConditionalBranches               GRegexError = 127
	GRegexErrorAssertionExpected                        GRegexError = 128
	GRegexErrorUnknownPosixClassName                    GRegexError = 130
	GRegexErrorPosixCollatingElementsNotSupported       GRegexError = 131
	GRegexErrorHexCodeTooLarge                          GRegexError = 134
	GRegexErrorInvalidCondition                         GRegexError = 135
	GRegexErrorSingleByteMatchInLookbehind              GRegexError = 136
	GRegexErrorInfiniteLoop                             GRegexError = 140
	GRegexErrorMissingSubpatternNameTerminator          GRegexError = 142
	GRegexErrorDuplicateSubpatternName                  GRegexError = 143
	GRegexErrorMalformedProperty                        GRegexError = 146
	GRegexErrorUnknownProperty                          GRegexError = 147
	GRegexErrorSubpatternNameTooLong                    GRegexError = 148
	GRegexErrorTooManySubpatterns                       GRegexError = 149
	GRegexErrorInvalidOctalValue                        GRegexError = 151
	GRegexErrorTooManyBranchesInDefine                  GRegexError = 154
	GRegexErrorDefineRepetion                           GRegexError = 155
	GRegexErrorInconsistentNewlineOptions               GRegexError = 156
	GRegexErrorMissingBackReference                     GRegexError = 157
	GRegexErrorInvalidRelativeReference                 GRegexError = 158
	GRegexErrorBacktrackingControlVerbArgumentForbidden GRegexError = 159
	GRegexErrorUnknownBacktrackingControlVerb           GRegexError = 160
	GRegexErrorNumberTooBig                             GRegexError = 161
	GRegexErrorMissingSubpatternName                    GRegexError = 162
	GRegexErrorMissingDigit                             GRegexError = 163
	GRegexErrorInvalidDataCharacter                     GRegexError = 164
	GRegexErrorExtraSubpatternName                      GRegexError = 165
	GRegexErrorBacktrackingControlVerbArgumentRequired  GRegexError = 166
	GRegexErrorInvalidControlChar                       GRegexError = 168
	GRegexErrorMissingName                              GRegexError = 169
	GRegexErrorNotSupportedInClass                      GRegexError = 171
	GRegexErrorTooManyForwardReferences                 GRegexError = 172
	GRegexErrorNameTooLong                              GRegexError = 175
	GRegexErrorCharacterValueTooLarge                   GRegexError = 176

	GSeekTypeCur GSeekType = 0
	GSeekTypeSet GSeekType = 1
	GSeekTypeEnd GSeekType = 2

	GShellErrorBadQuoting  GShellError = 0
	GShellErrorEmptyString GShellError = 1
	GShellErrorFailed      GShellError = 2

	GSliceConfigAlwaysMalloc      GSliceConfig = 1
	GSliceConfigBypassMagazines   GSliceConfig = 2
	GSliceConfigWorkingSetMsecs   GSliceConfig = 3
	GSliceConfigColorIncrement    GSliceConfig = 4
	GSliceConfigChunkSizes        GSliceConfig = 5
	GSliceConfigContentionCounter GSliceConfig = 6

	GSpawnErrorFork        GSpawnError = 0
	GSpawnErrorRead        GSpawnError = 1
	GSpawnErrorChdir       GSpawnError = 2
	GSpawnErrorAcces       GSpawnError = 3
	GSpawnErrorPerm        GSpawnError = 4
	GSpawnErrorTooBig      GSpawnError = 5
	GSpawnError2big        GSpawnError = 5
	GSpawnErrorNoexec      GSpawnError = 6
	GSpawnErrorNametoolong GSpawnError = 7
	GSpawnErrorNoent       GSpawnError = 8
	GSpawnErrorNomem       GSpawnError = 9
	GSpawnErrorNotdir      GSpawnError = 10
	GSpawnErrorLoop        GSpawnError = 11
	GSpawnErrorTxtbusy     GSpawnError = 12
	GSpawnErrorIo          GSpawnError = 13
	GSpawnErrorNfile       GSpawnError = 14
	GSpawnErrorMfile       GSpawnError = 15
	GSpawnErrorInval       GSpawnError = 16
	GSpawnErrorIsdir       GSpawnError = 17
	GSpawnErrorLibbad      GSpawnError = 18
	GSpawnErrorFailed      GSpawnError = 19

	GTestFileTypeDist  GTestFileType = 0
	GTestFileTypeBuilt GTestFileType = 1

	GTestLogTypeNone        GTestLogType = 0
	GTestLogTypeError       GTestLogType = 1
	GTestLogTypeStartBinary GTestLogType = 2
	GTestLogTypeListCase    GTestLogType = 3
	GTestLogTypeSkipCase    GTestLogType = 4
	GTestLogTypeStartCase   GTestLogType = 5
	GTestLogTypeStopCase    GTestLogType = 6
	GTestLogTypeMinResult   GTestLogType = 7
	GTestLogTypeMaxResult   GTestLogType = 8
	GTestLogTypeMessage     GTestLogType = 9
	GTestLogTypeStartSuite  GTestLogType = 10
	GTestLogTypeStopSuite   GTestLogType = 11

	GTestResultSuccess    GTestResult = 0
	GTestResultSkipped    GTestResult = 1
	GTestResultFailure    GTestResult = 2
	GTestResultIncomplete GTestResult = 3

	GThreadErrorThreadErrorAgain GThreadError = 0

	GTimeTypeStandard  GTimeType = 0
	GTimeTypeDaylight  GTimeType = 1
	GTimeTypeUniversal GTimeType = 2

	GTokenTypeEof            GTokenType = 0
	GTokenTypeLeftParen      GTokenType = 40
	GTokenTypeRightParen     GTokenType = 41
	GTokenTypeLeftCurly      GTokenType = 123
	GTokenTypeRightCurly     GTokenType = 125
	GTokenTypeLeftBrace      GTokenType = 91
	GTokenTypeRightBrace     GTokenType = 93
	GTokenTypeEqualSign      GTokenType = 61
	GTokenTypeComma          GTokenType = 44
	GTokenTypeNone           GTokenType = 256
	GTokenTypeError          GTokenType = 257
	GTokenTypeChar           GTokenType = 258
	GTokenTypeBinary         GTokenType = 259
	GTokenTypeOctal          GTokenType = 260
	GTokenTypeInt            GTokenType = 261
	GTokenTypeHex            GTokenType = 262
	GTokenTypeFloat          GTokenType = 263
	GTokenTypeString         GTokenType = 264
	GTokenTypeSymbol         GTokenType = 265
	GTokenTypeIdentifier     GTokenType = 266
	GTokenTypeIdentifierNull GTokenType = 267
	GTokenTypeCommentSingle  GTokenType = 268
	GTokenTypeCommentMulti   GTokenType = 269

	GTraverseTypeInOrder    GTraverseType = 0
	GTraverseTypePreOrder   GTraverseType = 1
	GTraverseTypePostOrder  GTraverseType = 2
	GTraverseTypeLevelOrder GTraverseType = 3

	GUnicodeBreakTypeMandatory                  GUnicodeBreakType = 0
	GUnicodeBreakTypeCarriageReturn             GUnicodeBreakType = 1
	GUnicodeBreakTypeLineFeed                   GUnicodeBreakType = 2
	GUnicodeBreakTypeCombiningMark              GUnicodeBreakType = 3
	GUnicodeBreakTypeSurrogate                  GUnicodeBreakType = 4
	GUnicodeBreakTypeZeroWidthSpace             GUnicodeBreakType = 5
	GUnicodeBreakTypeInseparable                GUnicodeBreakType = 6
	GUnicodeBreakTypeNonBreakingGlue            GUnicodeBreakType = 7
	GUnicodeBreakTypeContingent                 GUnicodeBreakType = 8
	GUnicodeBreakTypeSpace                      GUnicodeBreakType = 9
	GUnicodeBreakTypeAfter                      GUnicodeBreakType = 10
	GUnicodeBreakTypeBefore                     GUnicodeBreakType = 11
	GUnicodeBreakTypeBeforeAndAfter             GUnicodeBreakType = 12
	GUnicodeBreakTypeHyphen                     GUnicodeBreakType = 13
	GUnicodeBreakTypeNonStarter                 GUnicodeBreakType = 14
	GUnicodeBreakTypeOpenPunctuation            GUnicodeBreakType = 15
	GUnicodeBreakTypeClosePunctuation           GUnicodeBreakType = 16
	GUnicodeBreakTypeQuotation                  GUnicodeBreakType = 17
	GUnicodeBreakTypeExclamation                GUnicodeBreakType = 18
	GUnicodeBreakTypeIdeographic                GUnicodeBreakType = 19
	GUnicodeBreakTypeNumeric                    GUnicodeBreakType = 20
	GUnicodeBreakTypeInfixSeparator             GUnicodeBreakType = 21
	GUnicodeBreakTypeSymbol                     GUnicodeBreakType = 22
	GUnicodeBreakTypeAlphabetic                 GUnicodeBreakType = 23
	GUnicodeBreakTypePrefix                     GUnicodeBreakType = 24
	GUnicodeBreakTypePostfix                    GUnicodeBreakType = 25
	GUnicodeBreakTypeComplexContext             GUnicodeBreakType = 26
	GUnicodeBreakTypeAmbiguous                  GUnicodeBreakType = 27
	GUnicodeBreakTypeUnknown                    GUnicodeBreakType = 28
	GUnicodeBreakTypeNextLine                   GUnicodeBreakType = 29
	GUnicodeBreakTypeWordJoiner                 GUnicodeBreakType = 30
	GUnicodeBreakTypeHangulLJamo                GUnicodeBreakType = 31
	GUnicodeBreakTypeHangulVJamo                GUnicodeBreakType = 32
	GUnicodeBreakTypeHangulTJamo                GUnicodeBreakType = 33
	GUnicodeBreakTypeHangulLvSyllable           GUnicodeBreakType = 34
	GUnicodeBreakTypeHangulLvtSyllable          GUnicodeBreakType = 35
	GUnicodeBreakTypeCloseParanthesis           GUnicodeBreakType = 36
	GUnicodeBreakTypeCloseParenthesis           GUnicodeBreakType = 36
	GUnicodeBreakTypeConditionalJapaneseStarter GUnicodeBreakType = 37
	GUnicodeBreakTypeHebrewLetter               GUnicodeBreakType = 38
	GUnicodeBreakTypeRegionalIndicator          GUnicodeBreakType = 39
	GUnicodeBreakTypeEmojiBase                  GUnicodeBreakType = 40
	GUnicodeBreakTypeEmojiModifier              GUnicodeBreakType = 41
	GUnicodeBreakTypeZeroWidthJoiner            GUnicodeBreakType = 42
	GUnicodeBreakTypeAksara                     GUnicodeBreakType = 43
	GUnicodeBreakTypeAksaraPreBase              GUnicodeBreakType = 44
	GUnicodeBreakTypeAksaraStart                GUnicodeBreakType = 45
	GUnicodeBreakTypeViramaFinal                GUnicodeBreakType = 46
	GUnicodeBreakTypeVirama                     GUnicodeBreakType = 47

	GUnicodeScriptInvalidCode           GUnicodeScript = -1
	GUnicodeScriptCommon                GUnicodeScript = 0
	GUnicodeScriptInherited             GUnicodeScript = 1
	GUnicodeScriptArabic                GUnicodeScript = 2
	GUnicodeScriptArmenian              GUnicodeScript = 3
	GUnicodeScriptBengali               GUnicodeScript = 4
	GUnicodeScriptBopomofo              GUnicodeScript = 5
	GUnicodeScriptCherokee              GUnicodeScript = 6
	GUnicodeScriptCoptic                GUnicodeScript = 7
	GUnicodeScriptCyrillic              GUnicodeScript = 8
	GUnicodeScriptDeseret               GUnicodeScript = 9
	GUnicodeScriptDevanagari            GUnicodeScript = 10
	GUnicodeScriptEthiopic              GUnicodeScript = 11
	GUnicodeScriptGeorgian              GUnicodeScript = 12
	GUnicodeScriptGothic                GUnicodeScript = 13
	GUnicodeScriptGreek                 GUnicodeScript = 14
	GUnicodeScriptGujarati              GUnicodeScript = 15
	GUnicodeScriptGurmukhi              GUnicodeScript = 16
	GUnicodeScriptHan                   GUnicodeScript = 17
	GUnicodeScriptHangul                GUnicodeScript = 18
	GUnicodeScriptHebrew                GUnicodeScript = 19
	GUnicodeScriptHiragana              GUnicodeScript = 20
	GUnicodeScriptKannada               GUnicodeScript = 21
	GUnicodeScriptKatakana              GUnicodeScript = 22
	GUnicodeScriptKhmer                 GUnicodeScript = 23
	GUnicodeScriptLao                   GUnicodeScript = 24
	GUnicodeScriptLatin                 GUnicodeScript = 25
	GUnicodeScriptMalayalam             GUnicodeScript = 26
	GUnicodeScriptMongolian             GUnicodeScript = 27
	GUnicodeScriptMyanmar               GUnicodeScript = 28
	GUnicodeScriptOgham                 GUnicodeScript = 29
	GUnicodeScriptOldItalic             GUnicodeScript = 30
	GUnicodeScriptOriya                 GUnicodeScript = 31
	GUnicodeScriptRunic                 GUnicodeScript = 32
	GUnicodeScriptSinhala               GUnicodeScript = 33
	GUnicodeScriptSyriac                GUnicodeScript = 34
	GUnicodeScriptTamil                 GUnicodeScript = 35
	GUnicodeScriptTelugu                GUnicodeScript = 36
	GUnicodeScriptThaana                GUnicodeScript = 37
	GUnicodeScriptThai                  GUnicodeScript = 38
	GUnicodeScriptTibetan               GUnicodeScript = 39
	GUnicodeScriptCanadianAboriginal    GUnicodeScript = 40
	GUnicodeScriptYi                    GUnicodeScript = 41
	GUnicodeScriptTagalog               GUnicodeScript = 42
	GUnicodeScriptHanunoo               GUnicodeScript = 43
	GUnicodeScriptBuhid                 GUnicodeScript = 44
	GUnicodeScriptTagbanwa              GUnicodeScript = 45
	GUnicodeScriptBraille               GUnicodeScript = 46
	GUnicodeScriptCypriot               GUnicodeScript = 47
	GUnicodeScriptLimbu                 GUnicodeScript = 48
	GUnicodeScriptOsmanya               GUnicodeScript = 49
	GUnicodeScriptShavian               GUnicodeScript = 50
	GUnicodeScriptLinearB               GUnicodeScript = 51
	GUnicodeScriptTaiLe                 GUnicodeScript = 52
	GUnicodeScriptUgaritic              GUnicodeScript = 53
	GUnicodeScriptNewTaiLue             GUnicodeScript = 54
	GUnicodeScriptBuginese              GUnicodeScript = 55
	GUnicodeScriptGlagolitic            GUnicodeScript = 56
	GUnicodeScriptTifinagh              GUnicodeScript = 57
	GUnicodeScriptSylotiNagri           GUnicodeScript = 58
	GUnicodeScriptOldPersian            GUnicodeScript = 59
	GUnicodeScriptKharoshthi            GUnicodeScript = 60
	GUnicodeScriptUnknown               GUnicodeScript = 61
	GUnicodeScriptBalinese              GUnicodeScript = 62
	GUnicodeScriptCuneiform             GUnicodeScript = 63
	GUnicodeScriptPhoenician            GUnicodeScript = 64
	GUnicodeScriptPhagsPa               GUnicodeScript = 65
	GUnicodeScriptNko                   GUnicodeScript = 66
	GUnicodeScriptKayahLi               GUnicodeScript = 67
	GUnicodeScriptLepcha                GUnicodeScript = 68
	GUnicodeScriptRejang                GUnicodeScript = 69
	GUnicodeScriptSundanese             GUnicodeScript = 70
	GUnicodeScriptSaurashtra            GUnicodeScript = 71
	GUnicodeScriptCham                  GUnicodeScript = 72
	GUnicodeScriptOlChiki               GUnicodeScript = 73
	GUnicodeScriptVai                   GUnicodeScript = 74
	GUnicodeScriptCarian                GUnicodeScript = 75
	GUnicodeScriptLycian                GUnicodeScript = 76
	GUnicodeScriptLydian                GUnicodeScript = 77
	GUnicodeScriptAvestan               GUnicodeScript = 78
	GUnicodeScriptBamum                 GUnicodeScript = 79
	GUnicodeScriptEgyptianHieroglyphs   GUnicodeScript = 80
	GUnicodeScriptImperialAramaic       GUnicodeScript = 81
	GUnicodeScriptInscriptionalPahlavi  GUnicodeScript = 82
	GUnicodeScriptInscriptionalParthian GUnicodeScript = 83
	GUnicodeScriptJavanese              GUnicodeScript = 84
	GUnicodeScriptKaithi                GUnicodeScript = 85
	GUnicodeScriptLisu                  GUnicodeScript = 86
	GUnicodeScriptMeeteiMayek           GUnicodeScript = 87
	GUnicodeScriptOldSouthArabian       GUnicodeScript = 88
	GUnicodeScriptOldTurkic             GUnicodeScript = 89
	GUnicodeScriptSamaritan             GUnicodeScript = 90
	GUnicodeScriptTaiTham               GUnicodeScript = 91
	GUnicodeScriptTaiViet               GUnicodeScript = 92
	GUnicodeScriptBatak                 GUnicodeScript = 93
	GUnicodeScriptBrahmi                GUnicodeScript = 94
	GUnicodeScriptMandaic               GUnicodeScript = 95
	GUnicodeScriptChakma                GUnicodeScript = 96
	GUnicodeScriptMeroiticCursive       GUnicodeScript = 97
	GUnicodeScriptMeroiticHieroglyphs   GUnicodeScript = 98
	GUnicodeScriptMiao                  GUnicodeScript = 99
	GUnicodeScriptSharada               GUnicodeScript = 100
	GUnicodeScriptSoraSompeng           GUnicodeScript = 101
	GUnicodeScriptTakri                 GUnicodeScript = 102
	GUnicodeScriptBassaVah              GUnicodeScript = 103
	GUnicodeScriptCaucasianAlbanian     GUnicodeScript = 104
	GUnicodeScriptDuployan              GUnicodeScript = 105
	GUnicodeScriptElbasan               GUnicodeScript = 106
	GUnicodeScriptGrantha               GUnicodeScript = 107
	GUnicodeScriptKhojki                GUnicodeScript = 108
	GUnicodeScriptKhudawadi             GUnicodeScript = 109
	GUnicodeScriptLinearA               GUnicodeScript = 110
	GUnicodeScriptMahajani              GUnicodeScript = 111
	GUnicodeScriptManichaean            GUnicodeScript = 112
	GUnicodeScriptMendeKikakui          GUnicodeScript = 113
	GUnicodeScriptModi                  GUnicodeScript = 114
	GUnicodeScriptMro                   GUnicodeScript = 115
	GUnicodeScriptNabataean             GUnicodeScript = 116
	GUnicodeScriptOldNorthArabian       GUnicodeScript = 117
	GUnicodeScriptOldPermic             GUnicodeScript = 118
	GUnicodeScriptPahawhHmong           GUnicodeScript = 119
	GUnicodeScriptPalmyrene             GUnicodeScript = 120
	GUnicodeScriptPauCinHau             GUnicodeScript = 121
	GUnicodeScriptPsalterPahlavi        GUnicodeScript = 122
	GUnicodeScriptSiddham               GUnicodeScript = 123
	GUnicodeScriptTirhuta               GUnicodeScript = 124
	GUnicodeScriptWarangCiti            GUnicodeScript = 125
	GUnicodeScriptAhom                  GUnicodeScript = 126
	GUnicodeScriptAnatolianHieroglyphs  GUnicodeScript = 127
	GUnicodeScriptHatran                GUnicodeScript = 128
	GUnicodeScriptMultani               GUnicodeScript = 129
	GUnicodeScriptOldHungarian          GUnicodeScript = 130
	GUnicodeScriptSignwriting           GUnicodeScript = 131
	GUnicodeScriptAdlam                 GUnicodeScript = 132
	GUnicodeScriptBhaiksuki             GUnicodeScript = 133
	GUnicodeScriptMarchen               GUnicodeScript = 134
	GUnicodeScriptNewa                  GUnicodeScript = 135
	GUnicodeScriptOsage                 GUnicodeScript = 136
	GUnicodeScriptTangut                GUnicodeScript = 137
	GUnicodeScriptMasaramGondi          GUnicodeScript = 138
	GUnicodeScriptNushu                 GUnicodeScript = 139
	GUnicodeScriptSoyombo               GUnicodeScript = 140
	GUnicodeScriptZanabazarSquare       GUnicodeScript = 141
	GUnicodeScriptDogra                 GUnicodeScript = 142
	GUnicodeScriptGunjalaGondi          GUnicodeScript = 143
	GUnicodeScriptHanifiRohingya        GUnicodeScript = 144
	GUnicodeScriptMakasar               GUnicodeScript = 145
	GUnicodeScriptMedefaidrin           GUnicodeScript = 146
	GUnicodeScriptOldSogdian            GUnicodeScript = 147
	GUnicodeScriptSogdian               GUnicodeScript = 148
	GUnicodeScriptElymaic               GUnicodeScript = 149
	GUnicodeScriptNandinagari           GUnicodeScript = 150
	GUnicodeScriptNyiakengPuachueHmong  GUnicodeScript = 151
	GUnicodeScriptWancho                GUnicodeScript = 152
	GUnicodeScriptChorasmian            GUnicodeScript = 153
	GUnicodeScriptDivesAkuru            GUnicodeScript = 154
	GUnicodeScriptKhitanSmallScript     GUnicodeScript = 155
	GUnicodeScriptYezidi                GUnicodeScript = 156
	GUnicodeScriptCyproMinoan           GUnicodeScript = 157
	GUnicodeScriptOldUyghur             GUnicodeScript = 158
	GUnicodeScriptTangsa                GUnicodeScript = 159
	GUnicodeScriptToto                  GUnicodeScript = 160
	GUnicodeScriptVithkuqi              GUnicodeScript = 161
	GUnicodeScriptMath                  GUnicodeScript = 162
	GUnicodeScriptKawi                  GUnicodeScript = 163
	GUnicodeScriptNagMundari            GUnicodeScript = 164
	GUnicodeScriptTodhri                GUnicodeScript = 165
	GUnicodeScriptGaray                 GUnicodeScript = 166
	GUnicodeScriptTuluTigalari          GUnicodeScript = 167
	GUnicodeScriptSunuwar               GUnicodeScript = 168
	GUnicodeScriptGurungKhema           GUnicodeScript = 169
	GUnicodeScriptKiratRai              GUnicodeScript = 170
	GUnicodeScriptOlOnal                GUnicodeScript = 171

	GUnicodeTypeControl            GUnicodeType = 0
	GUnicodeTypeFormat             GUnicodeType = 1
	GUnicodeTypeUnassigned         GUnicodeType = 2
	GUnicodeTypePrivateUse         GUnicodeType = 3
	GUnicodeTypeSurrogate          GUnicodeType = 4
	GUnicodeTypeLowercaseLetter    GUnicodeType = 5
	GUnicodeTypeModifierLetter     GUnicodeType = 6
	GUnicodeTypeOtherLetter        GUnicodeType = 7
	GUnicodeTypeTitlecaseLetter    GUnicodeType = 8
	GUnicodeTypeUppercaseLetter    GUnicodeType = 9
	GUnicodeTypeSpacingMark        GUnicodeType = 10
	GUnicodeTypeEnclosingMark      GUnicodeType = 11
	GUnicodeTypeNonSpacingMark     GUnicodeType = 12
	GUnicodeTypeDecimalNumber      GUnicodeType = 13
	GUnicodeTypeLetterNumber       GUnicodeType = 14
	GUnicodeTypeOtherNumber        GUnicodeType = 15
	GUnicodeTypeConnectPunctuation GUnicodeType = 16
	GUnicodeTypeDashPunctuation    GUnicodeType = 17
	GUnicodeTypeClosePunctuation   GUnicodeType = 18
	GUnicodeTypeFinalPunctuation   GUnicodeType = 19
	GUnicodeTypeInitialPunctuation GUnicodeType = 20
	GUnicodeTypeOtherPunctuation   GUnicodeType = 21
	GUnicodeTypeOpenPunctuation    GUnicodeType = 22
	GUnicodeTypeCurrencySymbol     GUnicodeType = 23
	GUnicodeTypeModifierSymbol     GUnicodeType = 24
	GUnicodeTypeMathSymbol         GUnicodeType = 25
	GUnicodeTypeOtherSymbol        GUnicodeType = 26
	GUnicodeTypeLineSeparator      GUnicodeType = 27
	GUnicodeTypeParagraphSeparator GUnicodeType = 28
	GUnicodeTypeSpaceSeparator     GUnicodeType = 29

	GUriErrorFailed        GUriError = 0
	GUriErrorBadScheme     GUriError = 1
	GUriErrorBadUser       GUriError = 2
	GUriErrorBadPassword   GUriError = 3
	GUriErrorBadAuthParams GUriError = 4
	GUriErrorBadHost       GUriError = 5
	GUriErrorBadPort       GUriError = 6
	GUriErrorBadPath       GUriError = 7
	GUriErrorBadQuery      GUriError = 8
	GUriErrorBadFragment   GUriError = 9

	GUserDirectoryDirectoryDesktop     GUserDirectory = 0
	GUserDirectoryDirectoryDocuments   GUserDirectory = 1
	GUserDirectoryDirectoryDownload    GUserDirectory = 2
	GUserDirectoryDirectoryMusic       GUserDirectory = 3
	GUserDirectoryDirectoryPictures    GUserDirectory = 4
	GUserDirectoryDirectoryPublicShare GUserDirectory = 5
	GUserDirectoryDirectoryTemplates   GUserDirectory = 6
	GUserDirectoryDirectoryVideos      GUserDirectory = 7
	GUserDirectoryNDirectories         GUserDirectory = 8

	GVariantClassBoolean    GVariantClass = 98
	GVariantClassByte       GVariantClass = 121
	GVariantClassInt16      GVariantClass = 110
	GVariantClassUint16     GVariantClass = 113
	GVariantClassInt32      GVariantClass = 105
	GVariantClassUint32     GVariantClass = 117
	GVariantClassInt64      GVariantClass = 120
	GVariantClassUint64     GVariantClass = 116
	GVariantClassHandle     GVariantClass = 104
	GVariantClassDouble     GVariantClass = 100
	GVariantClassString     GVariantClass = 115
	GVariantClassObjectPath GVariantClass = 111
	GVariantClassSignature  GVariantClass = 103
	GVariantClassVariant    GVariantClass = 118
	GVariantClassMaybe      GVariantClass = 109
	GVariantClassArray      GVariantClass = 97
	GVariantClassTuple      GVariantClass = 40
	GVariantClassDictEntry  GVariantClass = 123

	GVariantParseErrorFailed                     GVariantParseError = 0
	GVariantParseErrorBasicTypeExpected          GVariantParseError = 1
	GVariantParseErrorCannotInferType            GVariantParseError = 2
	GVariantParseErrorDefiniteTypeExpected       GVariantParseError = 3
	GVariantParseErrorInputNotAtEnd              GVariantParseError = 4
	GVariantParseErrorInvalidCharacter           GVariantParseError = 5
	GVariantParseErrorInvalidFormatString        GVariantParseError = 6
	GVariantParseErrorInvalidObjectPath          GVariantParseError = 7
	GVariantParseErrorInvalidSignature           GVariantParseError = 8
	GVariantParseErrorInvalidTypeString          GVariantParseError = 9
	GVariantParseErrorNoCommonType               GVariantParseError = 10
	GVariantParseErrorNumberOutOfRange           GVariantParseError = 11
	GVariantParseErrorNumberTooBig               GVariantParseError = 12
	GVariantParseErrorTypeError                  GVariantParseError = 13
	GVariantParseErrorUnexpectedToken            GVariantParseError = 14
	GVariantParseErrorUnknownKeyword             GVariantParseError = 15
	GVariantParseErrorUnterminatedStringConstant GVariantParseError = 16
	GVariantParseErrorValueExpected              GVariantParseError = 17
	GVariantParseErrorRecursion                  GVariantParseError = 18

	GWin32OSTypeAny         GWin32OSType = 0
	GWin32OSTypeWorkstation GWin32OSType = 1
	GWin32OSTypeServer      GWin32OSType = 2

	GAsciiTypeAlnum  GAsciiType = 1
	GAsciiTypeAlpha  GAsciiType = 2
	GAsciiTypeCntrl  GAsciiType = 4
	GAsciiTypeDigit  GAsciiType = 8
	GAsciiTypeGraph  GAsciiType = 16
	GAsciiTypeLower  GAsciiType = 32
	GAsciiTypePrint  GAsciiType = 64
	GAsciiTypePunct  GAsciiType = 128
	GAsciiTypeSpace  GAsciiType = 256
	GAsciiTypeUpper  GAsciiType = 512
	GAsciiTypeXdigit GAsciiType = 1024

	GFileSetContentsFlagsNone         GFileSetContentsFlags = 0
	GFileSetContentsFlagsConsistent   GFileSetContentsFlags = 1
	GFileSetContentsFlagsDurable      GFileSetContentsFlags = 2
	GFileSetContentsFlagsOnlyExisting GFileSetContentsFlags = 4

	GFileTestIsRegular    GFileTest = 1
	GFileTestIsSymlink    GFileTest = 2
	GFileTestIsDir        GFileTest = 4
	GFileTestIsExecutable GFileTest = 8
	GFileTestExists       GFileTest = 16

	GFormatSizeFlagsDefault    GFormatSizeFlags = 0
	GFormatSizeFlagsLongFormat GFormatSizeFlags = 1
	GFormatSizeFlagsIecUnits   GFormatSizeFlags = 2
	GFormatSizeFlagsBits       GFormatSizeFlags = 4
	GFormatSizeFlagsOnlyValue  GFormatSizeFlags = 8
	GFormatSizeFlagsOnlyUnit   GFormatSizeFlags = 16

	GHookFlagMaskActive GHookFlagMask = 1
	GHookFlagMaskInCall GHookFlagMask = 2
	GHookFlagMaskMask   GHookFlagMask = 15

	GIOConditionIn   GIOCondition = 1
	GIOConditionOut  GIOCondition = 4
	GIOConditionPri  GIOCondition = 2
	GIOConditionErr  GIOCondition = 8
	GIOConditionHup  GIOCondition = 16
	GIOConditionNval GIOCondition = 32

	GIOFlagsNone        GIOFlags = 0
	GIOFlagsAppend      GIOFlags = 1
	GIOFlagsNonblock    GIOFlags = 2
	GIOFlagsIsReadable  GIOFlags = 4
	GIOFlagsIsWritable  GIOFlags = 8
	GIOFlagsIsWriteable GIOFlags = 8
	GIOFlagsIsSeekable  GIOFlags = 16
	GIOFlagsMask        GIOFlags = 31
	GIOFlagsGetMask     GIOFlags = 31
	GIOFlagsSetMask     GIOFlags = 3

	GKeyFileFlagsNone             GKeyFileFlags = 0
	GKeyFileFlagsKeepComments     GKeyFileFlags = 1
	GKeyFileFlagsKeepTranslations GKeyFileFlags = 2

	GLogLevelFlagsFlagRecursion GLogLevelFlags = 1
	GLogLevelFlagsFlagFatal     GLogLevelFlags = 2
	GLogLevelFlagsLevelError    GLogLevelFlags = 4
	GLogLevelFlagsLevelCritical GLogLevelFlags = 8
	GLogLevelFlagsLevelWarning  GLogLevelFlags = 16
	GLogLevelFlagsLevelMessage  GLogLevelFlags = 32
	GLogLevelFlagsLevelInfo     GLogLevelFlags = 64
	GLogLevelFlagsLevelDebug    GLogLevelFlags = 128
	GLogLevelFlagsLevelMask     GLogLevelFlags = -4

	GMainContextFlagsNone             GMainContextFlags = 0
	GMainContextFlagsOwnerlessPolling GMainContextFlags = 1

	GMarkupCollectTypeInvalid  GMarkupCollectType = 0
	GMarkupCollectTypeString   GMarkupCollectType = 1
	GMarkupCollectTypeStrdup   GMarkupCollectType = 2
	GMarkupCollectTypeBoolean  GMarkupCollectType = 3
	GMarkupCollectTypeTristate GMarkupCollectType = 4
	GMarkupCollectTypeOptional GMarkupCollectType = 65536

	GMarkupParseFlagsDefaultFlags                GMarkupParseFlags = 0
	GMarkupParseFlagsDoNotUseThisUnsupportedFlag GMarkupParseFlags = 1
	GMarkupParseFlagsTreatCdataAsText            GMarkupParseFlags = 2
	GMarkupParseFlagsPrefixErrorPosition         GMarkupParseFlags = 4
	GMarkupParseFlagsIgnoreQualified             GMarkupParseFlags = 8

	GOptionFlagsNone        GOptionFlags = 0
	GOptionFlagsHidden      GOptionFlags = 1
	GOptionFlagsInMain      GOptionFlags = 2
	GOptionFlagsReverse     GOptionFlags = 4
	GOptionFlagsNoArg       GOptionFlags = 8
	GOptionFlagsFilename    GOptionFlags = 16
	GOptionFlagsOptionalArg GOptionFlags = 32
	GOptionFlagsNoalias     GOptionFlags = 64
	GOptionFlagsDeprecated  GOptionFlags = 128

	GRegexCompileFlagsDefault          GRegexCompileFlags = 0
	GRegexCompileFlagsCaseless         GRegexCompileFlags = 1
	GRegexCompileFlagsMultiline        GRegexCompileFlags = 2
	GRegexCompileFlagsDotall           GRegexCompileFlags = 4
	GRegexCompileFlagsExtended         GRegexCompileFlags = 8
	GRegexCompileFlagsAnchored         GRegexCompileFlags = 16
	GRegexCompileFlagsDollarEndonly    GRegexCompileFlags = 32
	GRegexCompileFlagsUngreedy         GRegexCompileFlags = 512
	GRegexCompileFlagsRaw              GRegexCompileFlags = 2048
	GRegexCompileFlagsNoAutoCapture    GRegexCompileFlags = 4096
	GRegexCompileFlagsOptimize         GRegexCompileFlags = 8192
	GRegexCompileFlagsFirstline        GRegexCompileFlags = 262144
	GRegexCompileFlagsDupnames         GRegexCompileFlags = 524288
	GRegexCompileFlagsNewlineCr        GRegexCompileFlags = 1048576
	GRegexCompileFlagsNewlineLf        GRegexCompileFlags = 2097152
	GRegexCompileFlagsNewlineCrlf      GRegexCompileFlags = 3145728
	GRegexCompileFlagsNewlineAnycrlf   GRegexCompileFlags = 5242880
	GRegexCompileFlagsBsrAnycrlf       GRegexCompileFlags = 8388608
	GRegexCompileFlagsJavascriptCompat GRegexCompileFlags = 33554432

	GRegexMatchFlagsDefault         GRegexMatchFlags = 0
	GRegexMatchFlagsAnchored        GRegexMatchFlags = 16
	GRegexMatchFlagsNotbol          GRegexMatchFlags = 128
	GRegexMatchFlagsNoteol          GRegexMatchFlags = 256
	GRegexMatchFlagsNotempty        GRegexMatchFlags = 1024
	GRegexMatchFlagsPartial         GRegexMatchFlags = 32768
	GRegexMatchFlagsNewlineCr       GRegexMatchFlags = 1048576
	GRegexMatchFlagsNewlineLf       GRegexMatchFlags = 2097152
	GRegexMatchFlagsNewlineCrlf     GRegexMatchFlags = 3145728
	GRegexMatchFlagsNewlineAny      GRegexMatchFlags = 4194304
	GRegexMatchFlagsNewlineAnycrlf  GRegexMatchFlags = 5242880
	GRegexMatchFlagsBsrAnycrlf      GRegexMatchFlags = 8388608
	GRegexMatchFlagsBsrAny          GRegexMatchFlags = 16777216
	GRegexMatchFlagsPartialSoft     GRegexMatchFlags = 32768
	GRegexMatchFlagsPartialHard     GRegexMatchFlags = 134217728
	GRegexMatchFlagsNotemptyAtstart GRegexMatchFlags = 268435456

	GSpawnFlagsDefault              GSpawnFlags = 0
	GSpawnFlagsLeaveDescriptorsOpen GSpawnFlags = 1
	GSpawnFlagsDoNotReapChild       GSpawnFlags = 2
	GSpawnFlagsSearchPath           GSpawnFlags = 4
	GSpawnFlagsStdoutToDevNull      GSpawnFlags = 8
	GSpawnFlagsStderrToDevNull      GSpawnFlags = 16
	GSpawnFlagsChildInheritsStdin   GSpawnFlags = 32
	GSpawnFlagsFileAndArgvZero      GSpawnFlags = 64
	GSpawnFlagsSearchPathFromEnvp   GSpawnFlags = 128
	GSpawnFlagsCloexecPipes         GSpawnFlags = 256
	GSpawnFlagsChildInheritsStdout  GSpawnFlags = 512
	GSpawnFlagsChildInheritsStderr  GSpawnFlags = 1024
	GSpawnFlagsStdinFromDevNull     GSpawnFlags = 2048

	GTestSubprocessFlagsDefault       GTestSubprocessFlags = 0
	GTestSubprocessFlagsInheritStdin  GTestSubprocessFlags = 1
	GTestSubprocessFlagsInheritStdout GTestSubprocessFlags = 2
	GTestSubprocessFlagsInheritStderr GTestSubprocessFlags = 4

	GTestTrapFlagsDefault       GTestTrapFlags = 0
	GTestTrapFlagsSilenceStdout GTestTrapFlags = 128
	GTestTrapFlagsSilenceStderr GTestTrapFlags = 256
	GTestTrapFlagsInheritStdin  GTestTrapFlags = 512

	GTraverseFlagsLeaves    GTraverseFlags = 1
	GTraverseFlagsNonLeaves GTraverseFlags = 2
	GTraverseFlagsAll       GTraverseFlags = 3
	GTraverseFlagsMask      GTraverseFlags = 3
	GTraverseFlagsLeafs     GTraverseFlags = 1
	GTraverseFlagsNonLeafs  GTraverseFlags = 2

	GUriFlagsNone            GUriFlags = 0
	GUriFlagsParseRelaxed    GUriFlags = 1
	GUriFlagsHasPassword     GUriFlags = 2
	GUriFlagsHasAuthParams   GUriFlags = 4
	GUriFlagsEncoded         GUriFlags = 8
	GUriFlagsNonDns          GUriFlags = 16
	GUriFlagsEncodedQuery    GUriFlags = 32
	GUriFlagsEncodedPath     GUriFlags = 64
	GUriFlagsEncodedFragment GUriFlags = 128
	GUriFlagsSchemeNormalize GUriFlags = 256

	GUriHideFlagsNone       GUriHideFlags = 0
	GUriHideFlagsUserinfo   GUriHideFlags = 1
	GUriHideFlagsPassword   GUriHideFlags = 2
	GUriHideFlagsAuthParams GUriHideFlags = 4
	GUriHideFlagsQuery      GUriHideFlags = 8
	GUriHideFlagsFragment   GUriHideFlags = 16

	GUriParamsFlagsNone            GUriParamsFlags = 0
	GUriParamsFlagsCaseInsensitive GUriParamsFlags = 1
	GUriParamsFlagsWwwForm         GUriParamsFlags = 2
	GUriParamsFlagsParseRelaxed    GUriParamsFlags = 4
)
