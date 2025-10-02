package gobject

type GBindingFlags int32
type GTypeFundamentalFlags int32
type GTypeFlags int32
type GParamFlags int32
type GSignalFlags int32
type GSignalMatchType int32
type GConnectFlags int32
type GTypeDebugFlags int32

const (
	GBindingFlagDefault       GBindingFlags = 0
	GBindingFlagBidirectional GBindingFlags = 1 << 0
	GBindingFlagSyncCreate    GBindingFlags = 1 << 1
	GBindingFlagInvertBoolean GBindingFlags = 1 << 2

	GTypeFundamentalFlagClassed GTypeFundamentalFlags = 1 << iota
	GTypeFundamentalFlagInstantiatable
	GTypeFundamentalFlagDerivable
	GTypeFundamentalFlagDeepDerivable

	GTypeFlagNone          GTypeFlags = 0
	GTypeFlagAbstract      GTypeFlags = 1 << 4
	GTypeFlagValueAbstract GTypeFlags = 1 << 5
	GTypeFlagFinal         GTypeFlags = 1 << 6
	GTypeFlagDeprecated    GTypeFlags = 1 << 7

	GParamFlagReadable       GParamFlags = 1 << 0
	GParamFlagWritable       GParamFlags = 1 << 1
	GParamFlagReadWrite      GParamFlags = (GParamFlagReadable | GParamFlagWritable)
	GParamFlagConstruct      GParamFlags = 1 << 2
	GParamFlagConstructOnly  GParamFlags = 1 << 3
	GParamFlagLaxValidation  GParamFlags = 1 << 4
	GParamFlagStaticName     GParamFlags = 1 << 5
	GParamFlagPRIVATE        GParamFlags = GParamFlagStaticName
	GParamFlagStaticNick     GParamFlags = 1 << 6
	GParamFlagStaticBlurb    GParamFlags = 1 << 7
	GParamFlagExplicitNotify GParamFlags = 1 << 30
	// GParamFlagDeprecated     GParamFlags = 0x80000000

	GSignalFlagRunFirst            GSignalFlags = 1 << 0  // 1
	GSignalFlagRunLast             GSignalFlags = 1 << 1  // 2
	GSignalFlagRunCleanup          GSignalFlags = 1 << 2  // 4
	GSignalFlagNoRecurse           GSignalFlags = 1 << 3  // 8
	GSignalFlagDetailed            GSignalFlags = 1 << 4  // 16
	GSignalFlagAction              GSignalFlags = 1 << 5  // 32
	GSignalFlagNoHooks             GSignalFlags = 1 << 6  // 64
	GSignalFlagMustCollect         GSignalFlags = 1 << 7  // 128
	GSignalFlagDeprecated          GSignalFlags = 1 << 8  // 256
	GSignalFlagAccumulatorFirstRun GSignalFlags = 1 << 17 // 131072

	GSignalMatchTypeID        GSignalMatchType = 1 << 0
	GSignalMatchTypeDetail    GSignalMatchType = 1 << 1
	GSignalMatchTypeClosure   GSignalMatchType = 1 << 2
	GSignalMatchTypeFunc      GSignalMatchType = 1 << 3
	GSignalMatchTypeData      GSignalMatchType = 1 << 4
	GSignalMatchTypeUnblocked GSignalMatchType = 1 << 5

	GConnectFlagDefault GConnectFlags = 0
	GConnectFlagAfter   GConnectFlags = 1
	GConnectFlagSwapped GConnectFlags = 2

	GTypeDebugFlagsNone          GTypeDebugFlags = 0
	GTypeDebugFlagsObjects       GTypeDebugFlags = 1 << 0
	GTypeDebugFlagsSignals       GTypeDebugFlags = 1 << 1
	GTypeDebugFlagsInstanceCount GTypeDebugFlags = 1 << 2
	GTypeDebugFlagsMask          GTypeDebugFlags = 0x07
)
