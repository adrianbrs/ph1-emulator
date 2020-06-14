package constants

// Errors contém as mensagens de erros
const (
	RuntimeChangeMemoryArea      = "Can't change instruction memory area during runtime"
	VirtualMemoryAddressOverflow = "virtualMemory address overflow: %d"
	VirtualMemoryWordOverflow    = "virtualMemory word length overflow: %d"
	FileNameEmpty                = "File name must not be empty"
	UnableToOpenFile             = "Unable to open file: %s"
	OperationNotFound            = "Operation %s not found"
	AcRangeOverflow              = "RegAC overflow: %d"
	InvalidInstruction           = "Invalid instruction: none found for opCode %s"
)
