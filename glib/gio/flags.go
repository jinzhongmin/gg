package gio

type GBusType int32
type GConverterResult int32
type GCredentialsType int32
type GDBusError int32
type GDBusMessageByteOrder int32
type GDBusMessageHeaderField int32
type GDBusMessageType int32
type GDataStreamByteOrder int32
type GDataStreamNewlineType int32
type GDriveStartStopType int32
type GEmblemOrigin int32
type GFileAttributeStatus int32
type GFileAttributeType int32
type GFileMonitorEvent int32
type GFileType int32
type GFilesystemPreviewType int32
type GIOErrorEnum int32
type GIOModuleScopeFlags int32
type GMemoryMonitorWarningLevel int32
type GMountOperationResult int32
type GNetworkConnectivity int32
type GNotificationPriority int32
type GPasswordSave int32
type GPollableReturn int32
type GResolverError int32
type GResolverRecordType int32
type GResourceError int32
type GSocketClientEvent int32
type GSocketFamily int32
type GSocketListenerEvent int32
type GSocketProtocol int32
type GSocketType int32
type GTlsAuthenticationMode int32
type GTlsCertificateRequestFlags int32
type GTlsChannelBindingError int32
type GTlsChannelBindingType int32
type GTlsDatabaseLookupFlags int32
type GTlsError int32
type GTlsInteractionResult int32
type GTlsProtocolVersion int32
type GTlsRehandshakeMode int32
type GUnixSocketAddressType int32
type GZlibCompressorFormat int32
type GAppInfoCreateFlags int32
type GApplicationFlags int32
type GAskPasswordFlags int32
type GBusNameOwnerFlags int32
type GBusNameWatcherFlags int32
type GConverterFlags int32
type GDBusCallFlags int32
type GDBusCapabilityFlags int32
type GDBusConnectionFlags int32
type GDBusInterfaceSkeletonFlags int32
type GDBusMessageFlags int32
type GDBusObjectManagerClientFlags int32
type GDBusPropertyInfoFlags int32
type GDBusProxyFlags int32
type GDBusSendMessageFlags int32
type GDBusServerFlags int32
type GDBusSignalFlags int32
type GDBusSubtreeFlags int32
type GDriveStartFlags int32
type GFileAttributeInfoFlags int32
type GFileCopyFlags int32
type GFileCreateFlags int32
type GFileMeasureFlags int32
type GFileMonitorFlags int32
type GFileQueryInfoFlags int32
type GIOStreamSpliceFlags int32
type GMountMountFlags int32
type GMountUnmountFlags int32
type GOutputStreamSpliceFlags int32
type GResolverNameLookupFlags int32
type GResourceFlags int32
type GResourceLookupFlags int32
type GSettingsBindFlags int32
type GSocketMsgFlags int32
type GSubprocessFlags int32
type GTestDBusFlags int32
type GTlsCertificateFlags int32
type GTlsDatabaseVerifyFlags int32
type GTlsPasswordFlags int32

const (
	GBusTypeStarter GBusType = -1
	GBusTypeNone    GBusType = 0
	GBusTypeSystem  GBusType = 1
	GBusTypeSession GBusType = 2

	GConverterResultError     GConverterResult = 0
	GConverterResultConverted GConverterResult = 1
	GConverterResultFinished  GConverterResult = 2
	GConverterResultFlushed   GConverterResult = 3

	GCredentialsTypeInvalid             GCredentialsType = 0
	GCredentialsTypeLinuxUcred          GCredentialsType = 1
	GCredentialsTypeFreebsdCmsgcred     GCredentialsType = 2
	GCredentialsTypeOpenbsdSockpeercred GCredentialsType = 3
	GCredentialsTypeSolarisUcred        GCredentialsType = 4
	GCredentialsTypeNetbsdUnpcbid       GCredentialsType = 5
	GCredentialsTypeAppleXucred         GCredentialsType = 6
	GCredentialsTypeWin32Pid            GCredentialsType = 7

	GDBusErrorFailed                        GDBusError = 0
	GDBusErrorNoMemory                      GDBusError = 1
	GDBusErrorServiceUnknown                GDBusError = 2
	GDBusErrorNameHasNoOwner                GDBusError = 3
	GDBusErrorNoReply                       GDBusError = 4
	GDBusErrorIoError                       GDBusError = 5
	GDBusErrorBadAddress                    GDBusError = 6
	GDBusErrorNotSupported                  GDBusError = 7
	GDBusErrorLimitsExceeded                GDBusError = 8
	GDBusErrorAccessDenied                  GDBusError = 9
	GDBusErrorAuthFailed                    GDBusError = 10
	GDBusErrorNoServer                      GDBusError = 11
	GDBusErrorTimeout                       GDBusError = 12
	GDBusErrorNoNetwork                     GDBusError = 13
	GDBusErrorAddressInUse                  GDBusError = 14
	GDBusErrorDisconnected                  GDBusError = 15
	GDBusErrorInvalidArgs                   GDBusError = 16
	GDBusErrorFileNotFound                  GDBusError = 17
	GDBusErrorFileExists                    GDBusError = 18
	GDBusErrorUnknownMethod                 GDBusError = 19
	GDBusErrorTimedOut                      GDBusError = 20
	GDBusErrorMatchRuleNotFound             GDBusError = 21
	GDBusErrorMatchRuleInvalid              GDBusError = 22
	GDBusErrorSpawnExecFailed               GDBusError = 23
	GDBusErrorSpawnForkFailed               GDBusError = 24
	GDBusErrorSpawnChildExited              GDBusError = 25
	GDBusErrorSpawnChildSignaled            GDBusError = 26
	GDBusErrorSpawnFailed                   GDBusError = 27
	GDBusErrorSpawnSetupFailed              GDBusError = 28
	GDBusErrorSpawnConfigInvalid            GDBusError = 29
	GDBusErrorSpawnServiceInvalid           GDBusError = 30
	GDBusErrorSpawnServiceNotFound          GDBusError = 31
	GDBusErrorSpawnPermissionsInvalid       GDBusError = 32
	GDBusErrorSpawnFileInvalid              GDBusError = 33
	GDBusErrorSpawnNoMemory                 GDBusError = 34
	GDBusErrorUnixProcessIdUnknown          GDBusError = 35
	GDBusErrorInvalidSignature              GDBusError = 36
	GDBusErrorInvalidFileContent            GDBusError = 37
	GDBusErrorSelinuxSecurityContextUnknown GDBusError = 38
	GDBusErrorAdtAuditDataUnknown           GDBusError = 39
	GDBusErrorObjectPathInUse               GDBusError = 40
	GDBusErrorUnknownObject                 GDBusError = 41
	GDBusErrorUnknownInterface              GDBusError = 42
	GDBusErrorUnknownProperty               GDBusError = 43
	GDBusErrorPropertyReadOnly              GDBusError = 44

	GDBusMessageByteOrderBigEndian    GDBusMessageByteOrder = 66
	GDBusMessageByteOrderLittleEndian GDBusMessageByteOrder = 108

	GDBusMessageHeaderFieldInvalid     GDBusMessageHeaderField = 0
	GDBusMessageHeaderFieldPath        GDBusMessageHeaderField = 1
	GDBusMessageHeaderFieldInterface   GDBusMessageHeaderField = 2
	GDBusMessageHeaderFieldMember      GDBusMessageHeaderField = 3
	GDBusMessageHeaderFieldErrorName   GDBusMessageHeaderField = 4
	GDBusMessageHeaderFieldReplySerial GDBusMessageHeaderField = 5
	GDBusMessageHeaderFieldDestination GDBusMessageHeaderField = 6
	GDBusMessageHeaderFieldSender      GDBusMessageHeaderField = 7
	GDBusMessageHeaderFieldSignature   GDBusMessageHeaderField = 8
	GDBusMessageHeaderFieldNumUnixFds  GDBusMessageHeaderField = 9

	GDBusMessageTypeInvalid      GDBusMessageType = 0
	GDBusMessageTypeMethodCall   GDBusMessageType = 1
	GDBusMessageTypeMethodReturn GDBusMessageType = 2
	GDBusMessageTypeError        GDBusMessageType = 3
	GDBusMessageTypeSignal       GDBusMessageType = 4

	GDataStreamByteOrderBigEndian    GDataStreamByteOrder = 0
	GDataStreamByteOrderLittleEndian GDataStreamByteOrder = 1
	GDataStreamByteOrderHostEndian   GDataStreamByteOrder = 2

	GDataStreamNewlineTypeLf   GDataStreamNewlineType = 0
	GDataStreamNewlineTypeCr   GDataStreamNewlineType = 1
	GDataStreamNewlineTypeCrLf GDataStreamNewlineType = 2
	GDataStreamNewlineTypeAny  GDataStreamNewlineType = 3

	GDriveStartStopTypeUnknown   GDriveStartStopType = 0
	GDriveStartStopTypeShutdown  GDriveStartStopType = 1
	GDriveStartStopTypeNetwork   GDriveStartStopType = 2
	GDriveStartStopTypeMultidisk GDriveStartStopType = 3
	GDriveStartStopTypePassword  GDriveStartStopType = 4

	GEmblemOriginUnknown      GEmblemOrigin = 0
	GEmblemOriginDevice       GEmblemOrigin = 1
	GEmblemOriginLivemetadata GEmblemOrigin = 2
	GEmblemOriginTag          GEmblemOrigin = 3

	GFileAttributeStatusUnset        GFileAttributeStatus = 0
	GFileAttributeStatusSet          GFileAttributeStatus = 1
	GFileAttributeStatusErrorSetting GFileAttributeStatus = 2

	GFileAttributeTypeInvalid    GFileAttributeType = 0
	GFileAttributeTypeString     GFileAttributeType = 1
	GFileAttributeTypeByteString GFileAttributeType = 2
	GFileAttributeTypeBoolean    GFileAttributeType = 3
	GFileAttributeTypeUint32     GFileAttributeType = 4
	GFileAttributeTypeInt32      GFileAttributeType = 5
	GFileAttributeTypeUint64     GFileAttributeType = 6
	GFileAttributeTypeInt64      GFileAttributeType = 7
	GFileAttributeTypeObject     GFileAttributeType = 8
	GFileAttributeTypeStringv    GFileAttributeType = 9

	GFileMonitorEventChanged          GFileMonitorEvent = 0
	GFileMonitorEventChangesDoneHint  GFileMonitorEvent = 1
	GFileMonitorEventDeleted          GFileMonitorEvent = 2
	GFileMonitorEventCreated          GFileMonitorEvent = 3
	GFileMonitorEventAttributeChanged GFileMonitorEvent = 4
	GFileMonitorEventPreUnmount       GFileMonitorEvent = 5
	GFileMonitorEventUnmounted        GFileMonitorEvent = 6
	GFileMonitorEventMoved            GFileMonitorEvent = 7
	GFileMonitorEventRenamed          GFileMonitorEvent = 8
	GFileMonitorEventMovedIn          GFileMonitorEvent = 9
	GFileMonitorEventMovedOut         GFileMonitorEvent = 10

	GFileTypeUnknown      GFileType = 0
	GFileTypeRegular      GFileType = 1
	GFileTypeDirectory    GFileType = 2
	GFileTypeSymbolicLink GFileType = 3
	GFileTypeSpecial      GFileType = 4
	GFileTypeShortcut     GFileType = 5
	GFileTypeMountable    GFileType = 6

	GFilesystemPreviewTypeIfAlways GFilesystemPreviewType = 0
	GFilesystemPreviewTypeIfLocal  GFilesystemPreviewType = 1
	GFilesystemPreviewTypeNever    GFilesystemPreviewType = 2

	GIOErrorEnumFailed             GIOErrorEnum = 0
	GIOErrorEnumNotFound           GIOErrorEnum = 1
	GIOErrorEnumExists             GIOErrorEnum = 2
	GIOErrorEnumIsDirectory        GIOErrorEnum = 3
	GIOErrorEnumNotDirectory       GIOErrorEnum = 4
	GIOErrorEnumNotEmpty           GIOErrorEnum = 5
	GIOErrorEnumNotRegularFile     GIOErrorEnum = 6
	GIOErrorEnumNotSymbolicLink    GIOErrorEnum = 7
	GIOErrorEnumNotMountableFile   GIOErrorEnum = 8
	GIOErrorEnumFilenameTooLong    GIOErrorEnum = 9
	GIOErrorEnumInvalidFilename    GIOErrorEnum = 10
	GIOErrorEnumTooManyLinks       GIOErrorEnum = 11
	GIOErrorEnumNoSpace            GIOErrorEnum = 12
	GIOErrorEnumInvalidArgument    GIOErrorEnum = 13
	GIOErrorEnumPermissionDenied   GIOErrorEnum = 14
	GIOErrorEnumNotSupported       GIOErrorEnum = 15
	GIOErrorEnumNotMounted         GIOErrorEnum = 16
	GIOErrorEnumAlreadyMounted     GIOErrorEnum = 17
	GIOErrorEnumClosed             GIOErrorEnum = 18
	GIOErrorEnumCancelled          GIOErrorEnum = 19
	GIOErrorEnumPending            GIOErrorEnum = 20
	GIOErrorEnumReadOnly           GIOErrorEnum = 21
	GIOErrorEnumCantCreateBackup   GIOErrorEnum = 22
	GIOErrorEnumWrongEtag          GIOErrorEnum = 23
	GIOErrorEnumTimedOut           GIOErrorEnum = 24
	GIOErrorEnumWouldRecurse       GIOErrorEnum = 25
	GIOErrorEnumBusy               GIOErrorEnum = 26
	GIOErrorEnumWouldBlock         GIOErrorEnum = 27
	GIOErrorEnumHostNotFound       GIOErrorEnum = 28
	GIOErrorEnumWouldMerge         GIOErrorEnum = 29
	GIOErrorEnumFailedHandled      GIOErrorEnum = 30
	GIOErrorEnumTooManyOpenFiles   GIOErrorEnum = 31
	GIOErrorEnumNotInitialized     GIOErrorEnum = 32
	GIOErrorEnumAddressInUse       GIOErrorEnum = 33
	GIOErrorEnumPartialInput       GIOErrorEnum = 34
	GIOErrorEnumInvalidData        GIOErrorEnum = 35
	GIOErrorEnumDbusError          GIOErrorEnum = 36
	GIOErrorEnumHostUnreachable    GIOErrorEnum = 37
	GIOErrorEnumNetworkUnreachable GIOErrorEnum = 38
	GIOErrorEnumConnectionRefused  GIOErrorEnum = 39
	GIOErrorEnumProxyFailed        GIOErrorEnum = 40
	GIOErrorEnumProxyAuthFailed    GIOErrorEnum = 41
	GIOErrorEnumProxyNeedAuth      GIOErrorEnum = 42
	GIOErrorEnumProxyNotAllowed    GIOErrorEnum = 43
	GIOErrorEnumBrokenPipe         GIOErrorEnum = 44
	GIOErrorEnumConnectionClosed   GIOErrorEnum = 44
	GIOErrorEnumNotConnected       GIOErrorEnum = 45
	GIOErrorEnumMessageTooLarge    GIOErrorEnum = 46
	GIOErrorEnumNoSuchDevice       GIOErrorEnum = 47
	GIOErrorEnumDestinationUnset   GIOErrorEnum = 48

	GIOModuleScopeFlagsNone            GIOModuleScopeFlags = 0
	GIOModuleScopeFlagsBlockDuplicates GIOModuleScopeFlags = 1

	GMemoryMonitorWarningLevelLow      GMemoryMonitorWarningLevel = 50
	GMemoryMonitorWarningLevelMedium   GMemoryMonitorWarningLevel = 100
	GMemoryMonitorWarningLevelCritical GMemoryMonitorWarningLevel = 255

	GMountOperationResultHandled   GMountOperationResult = 0
	GMountOperationResultAborted   GMountOperationResult = 1
	GMountOperationResultUnhandled GMountOperationResult = 2

	GNetworkConnectivityLocal   GNetworkConnectivity = 1
	GNetworkConnectivityLimited GNetworkConnectivity = 2
	GNetworkConnectivityPortal  GNetworkConnectivity = 3
	GNetworkConnectivityFull    GNetworkConnectivity = 4

	GNotificationPriorityNormal GNotificationPriority = 0
	GNotificationPriorityLow    GNotificationPriority = 1
	GNotificationPriorityHigh   GNotificationPriority = 2
	GNotificationPriorityUrgent GNotificationPriority = 3

	GPasswordSaveNever       GPasswordSave = 0
	GPasswordSaveForSession  GPasswordSave = 1
	GPasswordSavePermanently GPasswordSave = 2

	GPollableReturnFailed     GPollableReturn = 0
	GPollableReturnOk         GPollableReturn = 1
	GPollableReturnWouldBlock GPollableReturn = -27

	GResolverErrorNotFound         GResolverError = 0
	GResolverErrorTemporaryFailure GResolverError = 1
	GResolverErrorInternal         GResolverError = 2

	GResolverRecordTypeSrv GResolverRecordType = 1
	GResolverRecordTypeMx  GResolverRecordType = 2
	GResolverRecordTypeTxt GResolverRecordType = 3
	GResolverRecordTypeSoa GResolverRecordType = 4
	GResolverRecordTypeNs  GResolverRecordType = 5

	GResourceErrorNotFound GResourceError = 0
	GResourceErrorInternal GResourceError = 1

	GSocketClientEventResolving        GSocketClientEvent = 0
	GSocketClientEventResolved         GSocketClientEvent = 1
	GSocketClientEventConnecting       GSocketClientEvent = 2
	GSocketClientEventConnected        GSocketClientEvent = 3
	GSocketClientEventProxyNegotiating GSocketClientEvent = 4
	GSocketClientEventProxyNegotiated  GSocketClientEvent = 5
	GSocketClientEventTlsHandshaking   GSocketClientEvent = 6
	GSocketClientEventTlsHandshaked    GSocketClientEvent = 7
	GSocketClientEventComplete         GSocketClientEvent = 8

	GSocketFamilyInvalid GSocketFamily = 0
	GSocketFamilyUnix    GSocketFamily = 1
	GSocketFamilyIpv4    GSocketFamily = 2
	GSocketFamilyIpv6    GSocketFamily = 23

	GSocketListenerEventBinding   GSocketListenerEvent = 0
	GSocketListenerEventBound     GSocketListenerEvent = 1
	GSocketListenerEventListening GSocketListenerEvent = 2
	GSocketListenerEventListened  GSocketListenerEvent = 3

	GSocketProtocolUnknown GSocketProtocol = -1
	GSocketProtocolDefault GSocketProtocol = 0
	GSocketProtocolTcp     GSocketProtocol = 6
	GSocketProtocolUdp     GSocketProtocol = 17
	GSocketProtocolSctp    GSocketProtocol = 132

	GSocketTypeInvalid   GSocketType = 0
	GSocketTypeStream    GSocketType = 1
	GSocketTypeDatagram  GSocketType = 2
	GSocketTypeSeqpacket GSocketType = 3

	GTlsAuthenticationModeNone      GTlsAuthenticationMode = 0
	GTlsAuthenticationModeRequested GTlsAuthenticationMode = 1
	GTlsAuthenticationModeRequired  GTlsAuthenticationMode = 2

	GTlsCertificateRequestFlagsNone GTlsCertificateRequestFlags = 0

	GTlsChannelBindingErrorNotImplemented GTlsChannelBindingError = 0
	GTlsChannelBindingErrorInvalidState   GTlsChannelBindingError = 1
	GTlsChannelBindingErrorNotAvailable   GTlsChannelBindingError = 2
	GTlsChannelBindingErrorNotSupported   GTlsChannelBindingError = 3
	GTlsChannelBindingErrorGeneralError   GTlsChannelBindingError = 4

	GTlsChannelBindingTypeUnique         GTlsChannelBindingType = 0
	GTlsChannelBindingTypeServerEndPoint GTlsChannelBindingType = 1
	GTlsChannelBindingTypeExporter       GTlsChannelBindingType = 2

	GTlsDatabaseLookupFlagsNone    GTlsDatabaseLookupFlags = 0
	GTlsDatabaseLookupFlagsKeypair GTlsDatabaseLookupFlags = 1

	GTlsErrorUnavailable            GTlsError = 0
	GTlsErrorMisc                   GTlsError = 1
	GTlsErrorBadCertificate         GTlsError = 2
	GTlsErrorNotTls                 GTlsError = 3
	GTlsErrorHandshake              GTlsError = 4
	GTlsErrorCertificateRequired    GTlsError = 5
	GTlsErrorEof                    GTlsError = 6
	GTlsErrorInappropriateFallback  GTlsError = 7
	GTlsErrorBadCertificatePassword GTlsError = 8

	GTlsInteractionResultUnhandled GTlsInteractionResult = 0
	GTlsInteractionResultHandled   GTlsInteractionResult = 1
	GTlsInteractionResultFailed    GTlsInteractionResult = 2

	GTlsProtocolVersionUnknown GTlsProtocolVersion = 0
	GTlsProtocolVersionSsl30   GTlsProtocolVersion = 1
	GTlsProtocolVersionTls10   GTlsProtocolVersion = 2
	GTlsProtocolVersionTls11   GTlsProtocolVersion = 3
	GTlsProtocolVersionTls12   GTlsProtocolVersion = 4
	GTlsProtocolVersionTls13   GTlsProtocolVersion = 5
	GTlsProtocolVersionDtls10  GTlsProtocolVersion = 201
	GTlsProtocolVersionDtls12  GTlsProtocolVersion = 202

	GTlsRehandshakeModeNever    GTlsRehandshakeMode = 0
	GTlsRehandshakeModeSafely   GTlsRehandshakeMode = 1
	GTlsRehandshakeModeUnsafely GTlsRehandshakeMode = 2

	GUnixSocketAddressTypeInvalid        GUnixSocketAddressType = 0
	GUnixSocketAddressTypeAnonymous      GUnixSocketAddressType = 1
	GUnixSocketAddressTypePath           GUnixSocketAddressType = 2
	GUnixSocketAddressTypeAbstract       GUnixSocketAddressType = 3
	GUnixSocketAddressTypeAbstractPadded GUnixSocketAddressType = 4

	GZlibCompressorFormatZlib GZlibCompressorFormat = 0
	GZlibCompressorFormatGzip GZlibCompressorFormat = 1
	GZlibCompressorFormatRaw  GZlibCompressorFormat = 2

	GAppInfoCreateFlagsNone                        GAppInfoCreateFlags = 0
	GAppInfoCreateFlagsNeedsTerminal               GAppInfoCreateFlags = 1
	GAppInfoCreateFlagsSupportsUris                GAppInfoCreateFlags = 2
	GAppInfoCreateFlagsSupportsStartupNotification GAppInfoCreateFlags = 4

	GApplicationFlagsFlagsNone          GApplicationFlags = 0
	GApplicationFlagsDefaultFlags       GApplicationFlags = 0
	GApplicationFlagsIsService          GApplicationFlags = 1
	GApplicationFlagsIsLauncher         GApplicationFlags = 2
	GApplicationFlagsHandlesOpen        GApplicationFlags = 4
	GApplicationFlagsHandlesCommandLine GApplicationFlags = 8
	GApplicationFlagsSendEnvironment    GApplicationFlags = 16
	GApplicationFlagsNonUnique          GApplicationFlags = 32
	GApplicationFlagsCanOverrideAppId   GApplicationFlags = 64
	GApplicationFlagsAllowReplacement   GApplicationFlags = 128
	GApplicationFlagsReplace            GApplicationFlags = 256

	GAskPasswordFlagsNeedPassword       GAskPasswordFlags = 1
	GAskPasswordFlagsNeedUsername       GAskPasswordFlags = 2
	GAskPasswordFlagsNeedDomain         GAskPasswordFlags = 4
	GAskPasswordFlagsSavingSupported    GAskPasswordFlags = 8
	GAskPasswordFlagsAnonymousSupported GAskPasswordFlags = 16
	GAskPasswordFlagsTcrypt             GAskPasswordFlags = 32

	GBusNameOwnerFlagsNone             GBusNameOwnerFlags = 0
	GBusNameOwnerFlagsAllowReplacement GBusNameOwnerFlags = 1
	GBusNameOwnerFlagsReplace          GBusNameOwnerFlags = 2
	GBusNameOwnerFlagsDoNotQueue       GBusNameOwnerFlags = 4

	GBusNameWatcherFlagsNone      GBusNameWatcherFlags = 0
	GBusNameWatcherFlagsAutoStart GBusNameWatcherFlags = 1

	GConverterFlagsNone       GConverterFlags = 0
	GConverterFlagsInputAtEnd GConverterFlags = 1
	GConverterFlagsFlush      GConverterFlags = 2

	GDBusCallFlagsNone                          GDBusCallFlags = 0
	GDBusCallFlagsNoAutoStart                   GDBusCallFlags = 1
	GDBusCallFlagsAllowInteractiveAuthorization GDBusCallFlags = 2

	GDBusCapabilityFlagsNone          GDBusCapabilityFlags = 0
	GDBusCapabilityFlagsUnixFdPassing GDBusCapabilityFlags = 1

	GDBusConnectionFlagsNone                          GDBusConnectionFlags = 0
	GDBusConnectionFlagsAuthenticationClient          GDBusConnectionFlags = 1
	GDBusConnectionFlagsAuthenticationServer          GDBusConnectionFlags = 2
	GDBusConnectionFlagsAuthenticationAllowAnonymous  GDBusConnectionFlags = 4
	GDBusConnectionFlagsMessageBusConnection          GDBusConnectionFlags = 8
	GDBusConnectionFlagsDelayMessageProcessing        GDBusConnectionFlags = 16
	GDBusConnectionFlagsAuthenticationRequireSameUser GDBusConnectionFlags = 32
	GDBusConnectionFlagsCrossNamespace                GDBusConnectionFlags = 64

	GDBusInterfaceSkeletonFlagsNone                            GDBusInterfaceSkeletonFlags = 0
	GDBusInterfaceSkeletonFlagsHandleMethodInvocationsInThread GDBusInterfaceSkeletonFlags = 1

	GDBusMessageFlagsNone                          GDBusMessageFlags = 0
	GDBusMessageFlagsNoReplyExpected               GDBusMessageFlags = 1
	GDBusMessageFlagsNoAutoStart                   GDBusMessageFlags = 2
	GDBusMessageFlagsAllowInteractiveAuthorization GDBusMessageFlags = 4

	GDBusObjectManagerClientFlagsNone           GDBusObjectManagerClientFlags = 0
	GDBusObjectManagerClientFlagsDoNotAutoStart GDBusObjectManagerClientFlags = 1

	GDBusPropertyInfoFlagsNone     GDBusPropertyInfoFlags = 0
	GDBusPropertyInfoFlagsReadable GDBusPropertyInfoFlags = 1
	GDBusPropertyInfoFlagsWritable GDBusPropertyInfoFlags = 2

	GDBusProxyFlagsNone                         GDBusProxyFlags = 0
	GDBusProxyFlagsDoNotLoadProperties          GDBusProxyFlags = 1
	GDBusProxyFlagsDoNotConnectSignals          GDBusProxyFlags = 2
	GDBusProxyFlagsDoNotAutoStart               GDBusProxyFlags = 4
	GDBusProxyFlagsGetInvalidatedProperties     GDBusProxyFlags = 8
	GDBusProxyFlagsDoNotAutoStartAtConstruction GDBusProxyFlags = 16
	GDBusProxyFlagsNoMatchRule                  GDBusProxyFlags = 32

	GDBusSendMessageFlagsNone           GDBusSendMessageFlags = 0
	GDBusSendMessageFlagsPreserveSerial GDBusSendMessageFlags = 1

	GDBusServerFlagsNone                          GDBusServerFlags = 0
	GDBusServerFlagsRunInThread                   GDBusServerFlags = 1
	GDBusServerFlagsAuthenticationAllowAnonymous  GDBusServerFlags = 2
	GDBusServerFlagsAuthenticationRequireSameUser GDBusServerFlags = 4

	GDBusSignalFlagsNone               GDBusSignalFlags = 0
	GDBusSignalFlagsNoMatchRule        GDBusSignalFlags = 1
	GDBusSignalFlagsMatchArg0Namespace GDBusSignalFlags = 2
	GDBusSignalFlagsMatchArg0Path      GDBusSignalFlags = 4

	GDBusSubtreeFlagsNone                        GDBusSubtreeFlags = 0
	GDBusSubtreeFlagsDispatchToUnenumeratedNodes GDBusSubtreeFlags = 1

	GDriveStartFlagsNone GDriveStartFlags = 0

	GFileAttributeInfoFlagsNone          GFileAttributeInfoFlags = 0
	GFileAttributeInfoFlagsCopyWithFile  GFileAttributeInfoFlags = 1
	GFileAttributeInfoFlagsCopyWhenMoved GFileAttributeInfoFlags = 2

	GFileCopyFlagsNone                      GFileCopyFlags = 0
	GFileCopyFlagsOverwrite                 GFileCopyFlags = 1
	GFileCopyFlagsBackup                    GFileCopyFlags = 2
	GFileCopyFlagsNofollowSymlinks          GFileCopyFlags = 4
	GFileCopyFlagsAllMetadata               GFileCopyFlags = 8
	GFileCopyFlagsNoFallbackForMove         GFileCopyFlags = 16
	GFileCopyFlagsTargetDefaultPerms        GFileCopyFlags = 32
	GFileCopyFlagsTargetDefaultModifiedTime GFileCopyFlags = 64

	GFileCreateFlagsNone               GFileCreateFlags = 0
	GFileCreateFlagsPrivate            GFileCreateFlags = 1
	GFileCreateFlagsReplaceDestination GFileCreateFlags = 2

	GFileMeasureFlagsNone           GFileMeasureFlags = 0
	GFileMeasureFlagsReportAnyError GFileMeasureFlags = 2
	GFileMeasureFlagsApparentSize   GFileMeasureFlags = 4
	GFileMeasureFlagsNoXdev         GFileMeasureFlags = 8

	GFileMonitorFlagsNone           GFileMonitorFlags = 0
	GFileMonitorFlagsWatchMounts    GFileMonitorFlags = 1
	GFileMonitorFlagsSendMoved      GFileMonitorFlags = 2
	GFileMonitorFlagsWatchHardLinks GFileMonitorFlags = 4
	GFileMonitorFlagsWatchMoves     GFileMonitorFlags = 8

	GFileQueryInfoFlagsNone             GFileQueryInfoFlags = 0
	GFileQueryInfoFlagsNofollowSymlinks GFileQueryInfoFlags = 1

	GIOStreamSpliceFlagsNone         GIOStreamSpliceFlags = 0
	GIOStreamSpliceFlagsCloseStream1 GIOStreamSpliceFlags = 1
	GIOStreamSpliceFlagsCloseStream2 GIOStreamSpliceFlags = 2
	GIOStreamSpliceFlagsWaitForBoth  GIOStreamSpliceFlags = 4

	GMountMountFlagsNone GMountMountFlags = 0

	GMountUnmountFlagsNone  GMountUnmountFlags = 0
	GMountUnmountFlagsForce GMountUnmountFlags = 1

	GOutputStreamSpliceFlagsNone        GOutputStreamSpliceFlags = 0
	GOutputStreamSpliceFlagsCloseSource GOutputStreamSpliceFlags = 1
	GOutputStreamSpliceFlagsCloseTarget GOutputStreamSpliceFlags = 2

	GResolverNameLookupFlagsDefault  GResolverNameLookupFlags = 0
	GResolverNameLookupFlagsIpv4Only GResolverNameLookupFlags = 1
	GResolverNameLookupFlagsIpv6Only GResolverNameLookupFlags = 2

	GResourceFlagsNone       GResourceFlags = 0
	GResourceFlagsCompressed GResourceFlags = 1

	GResourceLookupFlagsNone GResourceLookupFlags = 0

	GSettingsBindFlagsDefault       GSettingsBindFlags = 0
	GSettingsBindFlagsGet           GSettingsBindFlags = 1
	GSettingsBindFlagsSet           GSettingsBindFlags = 2
	GSettingsBindFlagsNoSensitivity GSettingsBindFlags = 4
	GSettingsBindFlagsGetNoChanges  GSettingsBindFlags = 8
	GSettingsBindFlagsInvertBoolean GSettingsBindFlags = 16

	GSocketMsgFlagsNone      GSocketMsgFlags = 0
	GSocketMsgFlagsOob       GSocketMsgFlags = 1
	GSocketMsgFlagsPeek      GSocketMsgFlags = 2
	GSocketMsgFlagsDontroute GSocketMsgFlags = 4

	GSubprocessFlagsNone               GSubprocessFlags = 0
	GSubprocessFlagsStdinPipe          GSubprocessFlags = 1
	GSubprocessFlagsStdinInherit       GSubprocessFlags = 2
	GSubprocessFlagsStdoutPipe         GSubprocessFlags = 4
	GSubprocessFlagsStdoutSilence      GSubprocessFlags = 8
	GSubprocessFlagsStderrPipe         GSubprocessFlags = 16
	GSubprocessFlagsStderrSilence      GSubprocessFlags = 32
	GSubprocessFlagsStderrMerge        GSubprocessFlags = 64
	GSubprocessFlagsInheritFds         GSubprocessFlags = 128
	GSubprocessFlagsSearchPathFromEnvp GSubprocessFlags = 256

	GTestDBusFlagsNone GTestDBusFlags = 0

	GTlsCertificateFlagsNoFlags      GTlsCertificateFlags = 0
	GTlsCertificateFlagsUnknownCa    GTlsCertificateFlags = 1
	GTlsCertificateFlagsBadIdentity  GTlsCertificateFlags = 2
	GTlsCertificateFlagsNotActivated GTlsCertificateFlags = 4
	GTlsCertificateFlagsExpired      GTlsCertificateFlags = 8
	GTlsCertificateFlagsRevoked      GTlsCertificateFlags = 16
	GTlsCertificateFlagsInsecure     GTlsCertificateFlags = 32
	GTlsCertificateFlagsGenericError GTlsCertificateFlags = 64
	GTlsCertificateFlagsValidateAll  GTlsCertificateFlags = 127

	GTlsDatabaseVerifyFlagsNone GTlsDatabaseVerifyFlags = 0

	GTlsPasswordFlagsNone                  GTlsPasswordFlags = 0
	GTlsPasswordFlagsRetry                 GTlsPasswordFlags = 2
	GTlsPasswordFlagsManyTries             GTlsPasswordFlags = 4
	GTlsPasswordFlagsFinalTry              GTlsPasswordFlags = 8
	GTlsPasswordFlagsPkcs11User            GTlsPasswordFlags = 16
	GTlsPasswordFlagsPkcs11SecurityOfficer GTlsPasswordFlags = 32
	GTlsPasswordFlagsPkcs11ContextSpecific GTlsPasswordFlags = 64
)
