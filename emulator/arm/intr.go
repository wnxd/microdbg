package arm

const (
	ARM_INTR_EXCP_UDEF = iota + 1
	ARM_INTR_EXCP_SWI
	ARM_INTR_EXCP_PREFETCH_ABORT
	ARM_INTR_EXCP_DATA_ABORT
	ARM_INTR_EXCP_IRQ
	ARM_INTR_EXCP_FIQ
	ARM_INTR_EXCP_BKPT
	ARM_INTR_EXCP_EXCEPTION_EXIT
	ARM_INTR_EXCP_KERNEL_TRAP
	ARM_INTR_EXCP_HVC
	ARM_INTR_EXCP_HYP_TRAP
	ARM_INTR_EXCP_SMC
	ARM_INTR_EXCP_VIRQ
	ARM_INTR_EXCP_VFIQ
	ARM_INTR_EXCP_SEMIHOST
	ARM_INTR_EXCP_NOCP
	ARM_INTR_EXCP_INVSTATE
	ARM_INTR_EXCP_STKOF
	ARM_INTR_EXCP_LAZYFP
	ARM_INTR_EXCP_LSERR
	ARM_INTR_EXCP_UNALIGNED
)
