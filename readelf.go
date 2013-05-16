package main

import (
        "encoding/binary"
        "errors"
        "fmt"
        "os"
        "unsafe"
)

const (
        EI_MAG0 = 0     /* File identification byte 0 index */
        ELFMAG0 = 0x7f  /* Magic number byte 0 */

        EI_MAG1 = 1     /* File identification byte 1 index */
        ELFMAG1 = 'E'   /* Magic number byte 1 */

        EI_MAG2 = 2     /* File identification byte 2 index */
        ELFMAG2 = 'L'   /* Magic number byte 2 */

        EI_MAG3 = 3     /* File identification byte 3 index */
        ELFMAG3 = 'F'   /* Magic number byte 3 */

        /* Conglomeration of the identification bytes, for easy testing as a word.  */
        ELFMAG  = "\177ELF"
        SELFMAG = 4

        EI_CLASS     = 4     /* File class byte index */
        ELFCLASSNONE = 0     /* Invalid class */
        ELFCLASS32   = 1     /* 32-bit objects */
        ELFCLASS64   = 2     /* 64-bit objects */
        ELFCLASSNUM  = 3

        EI_DATA     = 5     /* Data encoding byte index */
        ELFDATANONE = 0     /* Invalid data encoding */
        ELFDATA2LSB = 1     /* 2's complement, little endian */
        ELFDATA2MSB = 2     /* 2's complement, big endian */
        ELFDATANUM  = 3

        EI_VERSION = 6     /* File version byte index */
        /* Value must be EV_CURRENT */

        EI_OSABI            = 7     /* OS ABI identification */
        ELFOSABI_NONE       = 0     /* UNIX System V ABI */
        ELFOSABI_SYSV       = 0     /* Alias.  */
        ELFOSABI_HPUX       = 1     /* HP-UX */
        ELFOSABI_NETBSD     = 2     /* NetBSD.  */
        ELFOSABI_LINUX      = 3     /* Linux.  */
        ELFOSABI_SOLARIS    = 6     /* Sun Solaris.  */
        ELFOSABI_AIX        = 7     /* IBM AIX.  */
        ELFOSABI_IRIX       = 8     /* SGI Irix.  */
        ELFOSABI_FREEBSD    = 9     /* FreeBSD.  */
        ELFOSABI_TRU64      = 10    /* Compaq TRU64 UNIX.  */
        ELFOSABI_MODESTO    = 11    /* Novell Modesto.  */
        ELFOSABI_OPENBSD    = 12    /* OpenBSD.  */
        ELFOSABI_ARM        = 97    /* ARM */
        ELFOSABI_STANDALONE = 255   /* Standalone (embedded) application */
        EI_NIDENT           = 16
        EI_ABIVERSION       = 8     /* ABI version */

        EI_PAD  = 9     /* Byte index of padding bytes */

        /* Legal values for e_type (object file type).  */

        ET_NONE   = 0      /* No file type */
        ET_REL    = 1      /* Relocatable file */
        ET_EXEC   = 2      /* Executable file */
        ET_DYN    = 3      /* Shared object file */
        ET_CORE   = 4      /* Core file */
        ET_NUM    = 5      /* Number of defined types */
        ET_LOOS   = 0xfe00 /* OS-specific range start */
        ET_HIOS   = 0xfeff /* OS-specific range end */
        ET_LOPROC = 0xff00 /* Processor-specific range start */
        ET_HIPROC = 0xffff /* Processor-specific range end */

        /* Legal values for e_machine (architecture).  */

        EM_NONE        = 0     /* No machine */
        EM_M32         = 1     /* AT&T WE 32100 */
        EM_SPARC       = 2     /* SUN SPARC */
        EM_386         = 3     /* Intel 80386 */
        EM_68K         = 4     /* Motorola m68k family */
        EM_88K         = 5     /* Motorola m88k family */
        EM_860         = 7     /* Intel 80860 */
        EM_MIPS        = 8     /* MIPS R3000 big-endian */
        EM_S370        = 9     /* IBM System/370 */
        EM_MIPS_RS3_LE = 10    /* MIPS R3000 little-endian */

        EM_PARISC      = 15    /* HPPA */
        EM_VPP500      = 17    /* Fujitsu VPP500 */
        EM_SPARC32PLUS = 18    /* Sun's "v8plus" */
        EM_960         = 19    /* Intel 80960 */
        EM_PPC         = 20    /* PowerPC */
        EM_PPC64       = 21    /* PowerPC 64-bit */
        EM_S390        = 22    /* IBM S390 */
        EM_V800        = 36    /* NEC V800 series */
        EM_FR20        = 37    /* Fujitsu FR20 */
        EM_RH32        = 38    /* TRW RH-32 */
        EM_RCE         = 39    /* Motorola RCE */
        EM_ARM         = 40    /* ARM */
        EM_FAKE_ALPHA  = 41    /* Digital Alpha */
        EM_SH          = 42    /* Hitachi SH */
        EM_SPARCV9     = 43    /* SPARC v9 64-bit */
        EM_TRICORE     = 44    /* Siemens Tricore */
        EM_ARC         = 45    /* Argonaut RISC Core */
        EM_H8_300      = 46    /* Hitachi H8/300 */
        EM_H8_300H     = 47    /* Hitachi H8/300H */
        EM_H8S         = 48    /* Hitachi H8S */
        EM_H8_500      = 49    /* Hitachi H8/500 */
        EM_IA_64       = 50    /* Intel Merced */
        EM_MIPS_X      = 51    /* Stanford MIPS-X */
        EM_COLDFIRE    = 52    /* Motorola Coldfire */
        EM_68HC12      = 53    /* Motorola M68HC12 */
        EM_MMA         = 54    /* Fujitsu MMA Multimedia Accelerator*/
        EM_PCP         = 55    /* Siemens PCP */
        EM_NCPU        = 56    /* Sony nCPU embeeded RISC */
        EM_NDR1        = 57    /* Denso NDR1 microprocessor */
        EM_STARCORE    = 58    /* Motorola Start*Core processor */
        EM_ME16        = 59    /* Toyota ME16 processor */
        EM_ST100       = 60    /* STMicroelectronic ST100 processor */
        EM_TINYJ       = 61    /* Advanced Logic Corp. Tinyj emb.fam*/
        EM_X86_64      = 62    /* AMD x86-64 architecture */
        EM_PDSP        = 63    /* Sony DSP Processor */
        EM_FX66        = 66    /* Siemens FX66 microcontroller */
        EM_ST9PLUS     = 67    /* STMicroelectronics ST9+ 8/16 mc */
        EM_ST7         = 68    /* STmicroelectronics ST7 8 bit mc */
        EM_68HC16      = 69    /* Motorola MC68HC16 microcontroller */
        EM_68HC11      = 70    /* Motorola MC68HC11 microcontroller */
        EM_68HC08      = 71    /* Motorola MC68HC08 microcontroller */
        EM_68HC05      = 72    /* Motorola MC68HC05 microcontroller */
        EM_SVX         = 73    /* Silicon Graphics SVx */
        EM_ST19        = 74    /* STMicroelectronics ST19 8 bit mc */
        EM_VAX         = 75    /* Digital VAX */
        EM_CRIS        = 76    /* Axis Communications 32-bit embedded processor */
        EM_JAVELIN     = 77    /* Infineon Technologies 32-bit embedded processor */
        EM_FIREPATH    = 78    /* Element 14 64-bit DSP Processor */
        EM_ZSP         = 79    /* LSI Logic 16-bit DSP Processor */
        EM_MMIX        = 80    /* Donald Knuth's educational 64-bit processor */
        EM_HUANY       = 81    /* Harvard University machine-independent object files */
        EM_PRISM       = 82    /* SiTera Prism */
        EM_AVR         = 83    /* Atmel AVR 8-bit microcontroller */
        EM_FR30        = 84    /* Fujitsu FR30 */
        EM_D10V        = 85    /* Mitsubishi D10V */
        EM_D30V        = 86    /* Mitsubishi D30V */
        EM_V850        = 87    /* NEC v850 */
        EM_M32R        = 88    /* Mitsubishi M32R */
        EM_MN10300     = 89    /* Matsushita MN10300 */
        EM_MN10200     = 90    /* Matsushita MN10200 */
        EM_PJ          = 91    /* picoJava */
        EM_OPENRISC    = 92    /* OpenRISC 32-bit embedded processor */
        EM_ARC_A5      = 93    /* ARC Cores Tangent-A5 */
        EM_XTENSA      = 94    /* Tensilica Xtensa Architecture */
        EM_NUM         = 95

        /* If it is necessary to assign new unofficial EM_* values, please
           pick large random numbers (0x8523, 0xa7f2, etc.) to minimize the
           chances of collision with official or non-GNU unofficial values.  */

        EM_ALPHA = 0x9026

        /* Legal values for e_version (version).  */

        EV_NONE    = 0     /* Invalid ELF version */
        EV_CURRENT = 1     /* Current version */
        EV_NUM     = 2

        /* Legal values for p_type (segment type).  */

        PT_NULL         = 0          /* Program header table entry unused */
        PT_LOAD         = 1          /* Loadable program segment */
        PT_DYNAMIC      = 2          /* Dynamic linking information */
        PT_INTERP       = 3          /* Program interpreter */
        PT_NOTE         = 4          /* Auxiliary information */
        PT_SHLIB        = 5          /* Reserved */
        PT_PHDR         = 6          /* Entry for header table itself */
        PT_TLS          = 7          /* Thread-local storage segment */
        PT_NUM          = 8          /* Number of defined types */
        PT_LOOS         = 0x60000000 /* Start of OS-specific */
        PT_GNU_EH_FRAME = 0x6474e550 /* GCC .eh_frame_hdr segment */
        PT_GNU_STACK    = 0x6474e551 /* Indicates stack executability */
        PT_GNU_RELRO    = 0x6474e552 /* Read-only after relocation */
        PT_LOSUNW       = 0x6ffffffa
        PT_SUNWBSS      = 0x6ffffffa /* Sun Specific segment */
        PT_SUNWSTACK    = 0x6ffffffb /* Stack segment */
        PT_HISUNW       = 0x6fffffff
        PT_HIOS         = 0x6fffffff /* End of OS-specific */
        PT_LOPROC       = 0x70000000 /* Start of processor-specific */
        PT_HIPROC       = 0x7fffffff /* End of processor-specific */

        /* Legal values for p_flags (segment flags).  */

        PF_X        = (1 << 0)   /* Segment is executable */
        PF_W        = (1 << 1)   /* Segment is writable */
        PF_R        = (1 << 2)   /* Segment is readable */
        PF_MASKOS   = 0x0ff00000 /* OS-specific */
        PF_MASKPROC = 0xf0000000 /* Processor-specific */

        /* Special section indices.  */

        SHN_UNDEF     = 0      /* Undefined section */
        SHN_LORESERVE = 0xff00 /* Start of reserved indices */
        SHN_LOPROC    = 0xff00 /* Start of processor-specific */
        SHN_BEFORE    = 0xff00 /* Order section before all others
           =                           (Solaris).  */
        SHN_AFTER = 0xff01 /* Order section after all others
           =                           (Solaris).  */
        SHN_HIPROC    = 0xff1f /* End of processor-specific */
        SHN_LOOS      = 0xff20 /* Start of OS-specific */
        SHN_HIOS      = 0xff3f /* End of OS-specific */
        SHN_ABS       = 0xfff1 /* Associated symbol is absolute */
        SHN_COMMON    = 0xfff2 /* Associated symbol is common */
        SHN_XINDEX    = 0xffff /* Index is in extra table.  */
        SHN_HIRESERVE = 0xffff /* End of reserved indices */
        /* Legal values for sh_type (section type).  */

        SHT_NULL          = 0          /* Section header table entry unused */
        SHT_PROGBITS      = 1          /* Program data */
        SHT_SYMTAB        = 2          /* Symbol table */
        SHT_STRTAB        = 3          /* String table */
        SHT_RELA          = 4          /* Relocation entries with addends */
        SHT_HASH          = 5          /* Symbol hash table */
        SHT_DYNAMIC       = 6          /* Dynamic linking information */
        SHT_NOTE          = 7          /* Notes */
        SHT_NOBITS        = 8          /* Program space with no data (bss) */
        SHT_REL           = 9          /* Relocation entries, no addends */
        SHT_SHLIB         = 10         /* Reserved */
        SHT_DYNSYM        = 11         /* Dynamic linker symbol table */
        SHT_INIT_ARRAY    = 14         /* Array of constructors */
        SHT_FINI_ARRAY    = 15         /* Array of destructors */
        SHT_PREINIT_ARRAY = 16         /* Array of pre-constructors */
        SHT_GROUP         = 17         /* Section group */
        SHT_SYMTAB_SHNDX  = 18         /* Extended section indeces */
        SHT_NUM           = 19         /* Number of defined types.  */
        SHT_LOOS          = 0x60000000 /* Start OS-specific.  */
        SHT_GNU_HASH      = 0x6ffffff6 /* GNU-style hash table.  */
        SHT_GNU_LIBLIST   = 0x6ffffff7 /* Prelink library list */
        SHT_CHECKSUM      = 0x6ffffff8 /* Checksum for DSO content.  */
        SHT_LOSUNW        = 0x6ffffffa /* Sun-specific low bound.  */
        SHT_SUNW_move     = 0x6ffffffa
        SHT_SUNW_COMDAT   = 0x6ffffffb
        SHT_SUNW_syminfo  = 0x6ffffffc
        SHT_GNU_verdef    = 0x6ffffffd /* Version definition section.  */
        SHT_GNU_verneed   = 0x6ffffffe /* Version needs section.  */
        SHT_GNU_versym    = 0x6fffffff /* Version symbol table.  */
        SHT_HISUNW        = 0x6fffffff /* Sun-specific high bound.  */
        SHT_HIOS          = 0x6fffffff /* End OS-specific type */
        SHT_LOPROC        = 0x70000000 /* Start of processor-specific */
        SHT_HIPROC        = 0x7fffffff /* End of processor-specific */
        SHT_LOUSER        = 0x80000000 /* Start of application-specific */
        SHT_HIUSER        = 0x8fffffff /* End of application-specific */

        /* Legal values for sh_flags (section flags).  */

        SHF_WRITE            = (1 << 0) /* Writable */
        SHF_ALLOC            = (1 << 1) /* Occupies memory during execution */
        SHF_EXECINSTR        = (1 << 2) /* Executable */
        SHF_MERGE            = (1 << 4) /* Might be merged */
        SHF_STRINGS          = (1 << 5) /* Contains nul-terminated strings */
        SHF_INFO_LINK        = (1 << 6) /* `sh_info' contains SHT index */
        SHF_LINK_ORDER       = (1 << 7) /* Preserve order after combining */
        SHF_OS_NONCONFORMING = (1 << 8) /* Non-standard OS specific handling
           =                      required */
        SHF_GROUP    = (1 << 9)   /* Section is member of a group.  */
        SHF_TLS      = (1 << 10)  /* Section hold thread-local data.  */
        SHF_MASKOS   = 0x0ff00000 /* OS-specific.  */
        SHF_MASKPROC = 0xf0000000 /* Processor-specific */
        SHF_ORDERED  = (1 << 30)  /* Special ordering requirement
           =                      (Solaris).  */
        SHF_EXCLUDE = (1 << 31) /* Section is excluded unless
           referenced or allocated (Solaris).*/

        /* Section group handling.  */
        GRP_COMDAT = 0x1   /* Mark group as COMDAT.  */

)

type Elf32_Ehdr struct {
        e_ident     [EI_NIDENT]byte
        e_type      uint16  /* file type */
        e_machine   uint16  /* architecture */
        e_version   uint32
        e_entry     uint32  /* entry point */
        e_phoff     uint32  /* PH table offset */
        e_shoff     uint32  /* SH table offset */
        e_flags     uint32
        e_ehsize    uint16  /* ELF header size in bytes */
        e_phentsize uint16  /* PH size */
        e_phnum     uint16  /* PH number */
        e_shentsize uint16  /* SH size */
        e_shnum     uint16  /* SH number */
        e_shstrndx  uint16  /* SH name string table index */
}
type Elf64_Ehdr struct {
        e_ident     [EI_NIDENT]byte
        e_type      uint16  /* file type */
        e_machine   uint16  /* architecture */
        e_version   uint32
        e_entry     uint64  /* entry point */
        e_phoff     uint64  /* PH table offset */
        e_shoff     uint64  /* SH table offset */
        e_flags     uint32
        e_ehsize    uint16  /* ELF header size in bytes */
        e_phentsize uint16  /* PH size */
        e_phnum     uint16  /* PH number */
        e_shentsize uint16  /* SH size */
        e_shnum     uint16  /* SH number */
        e_shstrndx  uint16  /* SH name string table index */
}

type Elf32_Phdr struct {
        p_type   uint32  /* Segment type */
        p_offset uint32  /* Segment file offset */
        p_vaddr  uint32  /* Segment virtual address */
        p_paddr  uint32  /* Segment physical address */
        p_filesz uint32  /* Segment size in file */
        p_memsz  uint32  /* Segment size in memory */
        p_flags  uint32  /* Segment flags */
        p_align  uint32  /* Segment alignment */
}

type Elf64_Phdr struct {
        p_type   uint32  /* Segment type */
        p_flags  uint32  /* Segment flags */
        p_offset uint64  /* Segment file offset */
        p_vaddr  uint64  /* Segment virtual address */
        p_paddr  uint64  /* Segment physical address */
        p_filesz uint64  /* Segment size in file */
        p_memsz  uint64  /* Segment size in memory */
        p_align  uint64  /* Segment alignment */
}

type Elf32_Shdr struct {
        sh_name      uint32  /* Section name (string tbl index) */
        sh_type      uint32  /* Section type */
        sh_flags     uint32  /* Section flags */
        sh_addr      uint32  /* Section virtual addr at execution */
        sh_offset    uint32  /* Section file offset */
        sh_size      uint32  /* Section size in bytes */
        sh_link      uint32  /* Link to another section */
        sh_info      uint32  /* Additional section information */
        sh_addralign uint32  /* Section alignment */
        sh_entsize   uint32  /* Entry size if section holds table */
}

type Elf64_Shdr struct {
        sh_name      uint32  /* Section name (string tbl index) */
        sh_type      uint32  /* Section type */
        sh_flags     uint64  /* Section flags */
        sh_addr      uint64  /* Section virtual addr at execution */
        sh_offset    uint64  /* Section file offset */
        sh_size      uint64  /* Section size in bytes */
        sh_link      uint32  /* Link to another section */
        sh_info      uint32  /* Additional section information */
        sh_addralign uint64  /* Section alignment */
        sh_entsize   uint64  /* Entry size if section holds table */
}

/* Symbol table entry.  */

type Elf32_Sym struct {
        st_name  uint32  /* Symbol name (string tbl index) */
        st_value uint32  /* Symbol value */
        st_size  uint32  /* Symbol size */
        st_info  byte    /* Symbol type and binding */
        st_other byte    /* Symbol visibility */
        st_shndx uint16  /* Section index */
}

type Elf64_Sym struct {
        st_name  uint32  /* Symbol name (string tbl index) */
        st_info  byte    /* Symbol type and binding */
        st_other byte    /* Symbol visibility */
        st_shndx uint16  /* Section index */
        st_value uint64  /* Symbol value */
        st_size  uint64  /* Symbol size */
}

type Elf32_EhdrBytes struct {
        e_ident     [EI_NIDENT]byte
        e_type      [2]byte /* file type */
        e_machine   [2]byte /* architecture */
        e_version   [4]byte
        e_entry     [4]byte /* entry point */
        e_phoff     [4]byte /* PH table offset */
        e_shoff     [4]byte /* SH table offset */
        e_flags     [4]byte
        e_ehsize    [2]byte /* ELF header size in bytes */
        e_phentsize [2]byte /* PH size */
        e_phnum     [2]byte /* PH number */
        e_shentsize [2]byte /* SH size */
        e_shnum     [2]byte /* SH number */
        e_shstrndx  [2]byte /* SH name string table index */
}
type Elf64_EhdrBytes struct {
        e_ident     [EI_NIDENT]byte
        e_type      [2]byte /* file type */
        e_machine   [2]byte /* architecture */
        e_version   [4]byte
        e_entry     [8]byte /* entry point */
        e_phoff     [8]byte /* PH table offset */
        e_shoff     [8]byte /* SH table offset */
        e_flags     [4]byte
        e_ehsize    [2]byte /* ELF header size in bytes */
        e_phentsize [2]byte /* PH size */
        e_phnum     [2]byte /* PH number */
        e_shentsize [2]byte /* SH size */
        e_shnum     [2]byte /* SH number */
        e_shstrndx  [2]byte /* SH name string table index */
}

type Elf32_PhdrBytes struct {
        p_type   [4]byte /* Segment type */
        p_offset [4]byte /* Segment file offset */
        p_vaddr  [4]byte /* Segment virtual address */
        p_paddr  [4]byte /* Segment physical address */
        p_filesz [4]byte /* Segment size in file */
        p_memsz  [4]byte /* Segment size in memory */
        p_flags  [4]byte /* Segment flags */
        p_align  [4]byte /* Segment alignment */
}

type Elf64_PhdrBytes struct {
        p_type   [4]byte /* Segment type */
        p_flags  [4]byte /* Segment flags */
        p_offset [8]byte /* Segment file offset */
        p_vaddr  [8]byte /* Segment virtual address */
        p_paddr  [8]byte /* Segment physical address */
        p_filesz [8]byte /* Segment size in file */
        p_memsz  [8]byte /* Segment size in memory */
        p_align  [8]byte /* Segment alignment */
}

type Elf32_ShdrBytes struct {
        sh_name      [4]byte /* Section name (string tbl index) */
        sh_type      [4]byte /* Section type */
        sh_flags     [4]byte /* Section flags */
        sh_addr      [4]byte /* Section virtual addr at execution */
        sh_offset    [4]byte /* Section file offset */
        sh_size      [4]byte /* Section size in bytes */
        sh_link      [4]byte /* Link to another section */
        sh_info      [4]byte /* Additional section information */
        sh_addralign [4]byte /* Section alignment */
        sh_entsize   [4]byte /* Entry size if section holds table */
}

type Elf64_ShdrBytes struct {
        sh_name      [4]byte /* Section name (string tbl index) */
        sh_type      [4]byte /* Section type */
        sh_flags     [8]byte /* Section flags */
        sh_addr      [8]byte /* Section virtual addr at execution */
        sh_offset    [8]byte /* Section file offset */
        sh_size      [8]byte /* Section size in bytes */
        sh_link      [4]byte /* Link to another section */
        sh_info      [4]byte /* Additional section information */
        sh_addralign [8]byte /* Section alignment */
        sh_entsize   [8]byte /* Entry size if section holds table */
}

/* Symbol table entry.  */

type Elf32_SymBytes struct {
        st_name  [4]byte /* Symbol name (string tbl index) */
        st_value [4]byte /* Symbol value */
        st_size  [4]byte /* Symbol size */
        st_info  byte    /* Symbol type and binding */
        st_other byte    /* Symbol visibility */
        st_shndx [2]byte /* Section index */
}

type Elf64_SymBytes struct {
        st_name  [4]byte /* Symbol name (string tbl index) */
        st_info  byte    /* Symbol type and binding */
        st_other byte    /* Symbol visibility */
        st_shndx [2]byte /* Section index */
        st_value [8]byte /* Symbol value */
        st_size  [8]byte /* Symbol size */
}

type ELFReader struct {
        i32         bool
        isBigEndian bool    /* 真为大端，假为小端 */
        ehdr64      *Elf64_Ehdr
        ehdr32      *Elf32_Ehdr
        f           *os.File
}

func NewELFReader(s string) (*ELFReader, error) {
        v := new(ELFReader)
        if f, err := os.Open(s); err != nil {
                return nil, err
        } else {
                v.f = f
        }
        return v, nil
}

func (this *ELFReader) _GetELFHeaderStr32() (string, error) {
        var ehdr32 *Elf32_EhdrBytes
        cc := make([]byte, unsafe.Sizeof(*ehdr32))
        if _, err := this.f.ReadAt(cc, 0); err != nil {
                return "", err
        }
        cc1 := &cc[0]
        pntStr := ""
        ehdr32 = (*Elf32_EhdrBytes)(unsafe.Pointer(cc1))
        this.ehdr32 = new(Elf32_Ehdr)
        if this.isBigEndian {
                this.ehdr32.e_type = binary.BigEndian.Uint16(ehdr32.e_type[:])
                this.ehdr32.e_machine = binary.BigEndian.Uint16(ehdr32.e_machine[:])
                this.ehdr32.e_version = binary.BigEndian.Uint32(ehdr32.e_version[:])
                this.ehdr32.e_entry = binary.BigEndian.Uint32(ehdr32.e_entry[:])
                this.ehdr32.e_phoff = binary.BigEndian.Uint32(ehdr32.e_phoff[:])
                this.ehdr32.e_shoff = binary.BigEndian.Uint32(ehdr32.e_shoff[:])
                this.ehdr32.e_flags = binary.BigEndian.Uint32(ehdr32.e_flags[:])
                this.ehdr32.e_ehsize = binary.BigEndian.Uint16(ehdr32.e_ehsize[:])
                this.ehdr32.e_phentsize = binary.BigEndian.Uint16(ehdr32.e_phentsize[:])
                this.ehdr32.e_phnum = binary.BigEndian.Uint16(ehdr32.e_phnum[:])
                this.ehdr32.e_shentsize = binary.BigEndian.Uint16(ehdr32.e_shentsize[:])
                this.ehdr32.e_shnum = binary.BigEndian.Uint16(ehdr32.e_shnum[:])
                this.ehdr32.e_shstrndx = binary.BigEndian.Uint16(ehdr32.e_shstrndx[:])
        } else {
                this.ehdr32.e_type = binary.LittleEndian.Uint16(ehdr32.e_type[:])
                this.ehdr32.e_machine = binary.LittleEndian.Uint16(ehdr32.e_machine[:])
                this.ehdr32.e_version = binary.LittleEndian.Uint32(ehdr32.e_version[:])
                this.ehdr32.e_entry = binary.LittleEndian.Uint32(ehdr32.e_entry[:])
                this.ehdr32.e_phoff = binary.LittleEndian.Uint32(ehdr32.e_phoff[:])
                this.ehdr32.e_shoff = binary.LittleEndian.Uint32(ehdr32.e_shoff[:])
                this.ehdr32.e_flags = binary.LittleEndian.Uint32(ehdr32.e_flags[:])
                this.ehdr32.e_ehsize = binary.LittleEndian.Uint16(ehdr32.e_ehsize[:])
                this.ehdr32.e_phentsize = binary.LittleEndian.Uint16(ehdr32.e_phentsize[:])
                this.ehdr32.e_phnum = binary.LittleEndian.Uint16(ehdr32.e_phnum[:])
                this.ehdr32.e_shentsize = binary.LittleEndian.Uint16(ehdr32.e_shentsize[:])
                this.ehdr32.e_shnum = binary.LittleEndian.Uint16(ehdr32.e_shnum[:])
                this.ehdr32.e_shstrndx = binary.LittleEndian.Uint16(ehdr32.e_shstrndx[:])
        }
        pntStr = pntStr + fmt.Sprintf("  %-35s", "Type:")
        switch this.ehdr32.e_type {
        case ET_REL:
                pntStr = pntStr + "REL (Relocatable file)\n"
        case ET_EXEC:
                pntStr = pntStr + "EXEC (Executable file)\n"
        case ET_DYN:
                pntStr = pntStr + "DYN (Shared object file)\n"
        case ET_CORE:
                pntStr = pntStr + "CORE (Core file)\n"
        default:
                pntStr = pntStr + "No file type\n"
        }
        pntStr = pntStr + fmt.Sprintf("  %-35s", "Machine:")
        switch this.ehdr32.e_machine {
        case EM_M32:
                pntStr = pntStr + "AT&T WE 32100\n"
        case EM_SPARC:
                pntStr = pntStr + "SUN SPARC\n"
        case EM_386:
                pntStr = pntStr + "Intel 80386\n"
        case EM_68K:
                pntStr = pntStr + "Motorola m68k family\n"
        case EM_88K:
                pntStr = pntStr + "Motorola m88k family\n"
        case EM_860:
                pntStr = pntStr + "Intel 80860\n"
        case EM_MIPS:
                pntStr = pntStr + "MIPS R3000 big-endian\n"
        case EM_S370:
                pntStr = pntStr + "IBM System/370\n"
        case EM_MIPS_RS3_LE:
                pntStr = pntStr + "MIPS R3000 little-endian\n"
        case EM_PARISC:
                pntStr = pntStr + "HPPA\n"
        case EM_VPP500:
                pntStr = pntStr + "Fujitsu VPP500\n"
        case EM_SPARC32PLUS:
                pntStr = pntStr + "Sun's \"v8plus\"\n"
        case EM_960:
                pntStr = pntStr + "Intel 80960\n"
        case EM_PPC:
                pntStr = pntStr + "PowerPC\n"
        case EM_PPC64:
                pntStr = pntStr + "PowerPC 64-bit\n"
        case EM_S390:
                pntStr = pntStr + "IBM S390\n"
        case EM_V800:
                pntStr = pntStr + "NEC V800 series\n"
        case EM_FR20:
                pntStr = pntStr + "Fujitsu FR20\n"
        case EM_RH32:
                pntStr = pntStr + "TRW RH-32\n"
        case EM_RCE:
                pntStr = pntStr + "Motorola RCE\n"
        case EM_ARM:
                pntStr = pntStr + "ARM\n"
        case EM_FAKE_ALPHA:
                pntStr = pntStr + "Digital Alpha\n"
        case EM_SH:
                pntStr = pntStr + "Hitachi SH\n"
        case EM_SPARCV9:
                pntStr = pntStr + "SPARC v9 64-bit\n"
        case EM_TRICORE:
                pntStr = pntStr + "Siemens Tricore\n"
        case EM_ARC:
                pntStr = pntStr + "Argonaut RISC Core\n"
        case EM_H8_300:
                pntStr = pntStr + "Hitachi H8/300\n"
        case EM_H8_300H:
                pntStr = pntStr + "Hitachi H8/300H\n"
        case EM_H8S:
                pntStr = pntStr + "Hitachi H8S\n"
        case EM_H8_500:
                pntStr = pntStr + "Hitachi H8/500\n"
        case EM_IA_64:
                pntStr = pntStr + "Intel Merced\n"
        case EM_MIPS_X:
                pntStr = pntStr + "Stanford MIPS-X\n"
        case EM_COLDFIRE:
                pntStr = pntStr + "Motorola Coldfire\n"
        case EM_68HC12:
                pntStr = pntStr + "Motorola M68HC12\n"
        case EM_MMA:
                pntStr = pntStr + "Fujitsu MMA Multimedia Accelerator\n"
        case EM_PCP:
                pntStr = pntStr + "Siemens PCP\n"
        case EM_NCPU:
                pntStr = pntStr + "Sony nCPU embeeded RISC\n"
        case EM_NDR1:
                pntStr = pntStr + "Denso NDR1 microprocessor\n"
        case EM_STARCORE:
                pntStr = pntStr + "Motorola Start*Core processor\n"
        case EM_ME16:
                pntStr = pntStr + "Toyota ME16 processor\n"
        case EM_ST100:
                pntStr = pntStr + "STMicroelectronic ST100 processor\n"
        case EM_TINYJ:
                pntStr = pntStr + "Advanced Logic Corp. Tinyj emb.fam\n"
        case EM_X86_64:
                pntStr = pntStr + "Advanced Micro Devices X86-64\n"
        case EM_PDSP:
                pntStr = pntStr + "Sony DSP Processor\n"
        case EM_FX66:
                pntStr = pntStr + "Siemens FX66 microcontroller\n"
        case EM_ST9PLUS:
                pntStr = pntStr + "STMicroelectronics ST9+ 8/16 mc\n"
        case EM_ST7:
                pntStr = pntStr + "STmicroelectronics ST7 8 bit mc\n"
        case EM_68HC16:
                pntStr = pntStr + "Motorola MC68HC16 microcontroller\n"
        case EM_68HC11:
                pntStr = pntStr + "Motorola MC68HC11 microcontroller\n"
        case EM_68HC08:
                pntStr = pntStr + "Motorola MC68HC08 microcontroller\n"
        case EM_68HC05:
                pntStr = pntStr + "Motorola MC68HC05 microcontroller\n"
        case EM_SVX:
                pntStr = pntStr + "Silicon Graphics SVx\n"
        case EM_ST19:
                pntStr = pntStr + "STMicroelectronics ST19 8 bit mc\n"
        case EM_VAX:
                pntStr = pntStr + "Digital VAX\n"
        case EM_CRIS:
                pntStr = pntStr + "Axis Communications 32-bit embedded processor\n"
        case EM_JAVELIN:
                pntStr = pntStr + "Infineon Technologies 32-bit embedded processor\n"
        case EM_FIREPATH:
                pntStr = pntStr + "Element 14 64-bit DSP Processor\n"
        case EM_ZSP:
                pntStr = pntStr + "LSI Logic 16-bit DSP Processor\n"
        case EM_MMIX:
                pntStr = pntStr + "Donald Knuth's educational 64-bit processor\n"
        case EM_HUANY:
                pntStr = pntStr + "Harvard University machine-independent object files\n"
        case EM_PRISM:
                pntStr = pntStr + "SiTera Prism\n"
        case EM_AVR:
                pntStr = pntStr + "Atmel AVR 8-bit microcontroller\n"
        case EM_FR30:
                pntStr = pntStr + "Fujitsu FR30\n"
        case EM_D10V:
                pntStr = pntStr + "Mitsubishi D10V\n"
        case EM_D30V:
                pntStr = pntStr + "Mitsubishi D30V\n"
        case EM_V850:
                pntStr = pntStr + "NEC v850\n"
        case EM_M32R:
                pntStr = pntStr + "Mitsubishi M32R\n"
        case EM_MN10300:
                pntStr = pntStr + "Matsushita MN10300\n"
        case EM_MN10200:
                pntStr = pntStr + "Matsushita MN10200\n"
        case EM_PJ:
                pntStr = pntStr + "picoJava\n"
        case EM_OPENRISC:
                pntStr = pntStr + "OpenRISC 32-bit embedded processor\n"
        case EM_ARC_A5:
                pntStr = pntStr + "ARC Cores Tangent-A5\n"
        case EM_XTENSA:
                pntStr = pntStr + "Tensilica Xtensa Architecture\n"
        default:
                pntStr = pntStr + "No machine\n"
        }
        pntStr = pntStr + fmt.Sprintf("  %-35s0x%x\n", "Version:", this.ehdr32.e_version)
        pntStr = pntStr + fmt.Sprintf("  %-35s0x%x\n", "Entry point address:", this.ehdr32.e_entry)
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Start of program headers:", this.ehdr32.e_phoff, "(bytes into file)")
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Start of section headers:", this.ehdr32.e_shoff, "(bytes into file)")
        pntStr = pntStr + fmt.Sprintf("  %-35s0x%x\n", "Flags:", this.ehdr32.e_flags)
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Size of this header:", this.ehdr32.e_ehsize, "(bytes)")
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Size of program headers:", this.ehdr32.e_phentsize, "(bytes)")
        pntStr = pntStr + fmt.Sprintf("  %-35s%d\n", "Number of program headers:", this.ehdr32.e_phnum)
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Size of section headers:", this.ehdr32.e_shentsize, "(bytes)")
        pntStr = pntStr + fmt.Sprintf("  %-35s%d\n", "Number of section headers:", this.ehdr32.e_shnum)
        pntStr = pntStr + fmt.Sprintf("  %-35s%d\n", "Section header string table index:", this.ehdr32.e_shstrndx)

        return pntStr, nil
}

func (this *ELFReader) _GetELFHeaderStr64() (string, error) {
        var ehdr64 *Elf64_EhdrBytes

        cc := make([]byte, unsafe.Sizeof(*ehdr64))
        if _, err := this.f.ReadAt(cc, 0); err != nil {
                return "", err
        }

        cc1 := &cc[0]
        pntStr := ""
        ehdr64 = (*Elf64_EhdrBytes)(unsafe.Pointer(cc1))
        this.ehdr64 = new(Elf64_Ehdr)
        if this.isBigEndian { /* 大端 */
                this.ehdr64.e_type = binary.BigEndian.Uint16(ehdr64.e_type[:])
                this.ehdr64.e_machine = binary.BigEndian.Uint16(ehdr64.e_machine[:])
                this.ehdr64.e_version = binary.BigEndian.Uint32(ehdr64.e_version[:])
                this.ehdr64.e_entry = binary.BigEndian.Uint64(ehdr64.e_entry[:])
                this.ehdr64.e_phoff = binary.BigEndian.Uint64(ehdr64.e_phoff[:])
                this.ehdr64.e_shoff = binary.BigEndian.Uint64(ehdr64.e_shoff[:])
                this.ehdr64.e_flags = binary.BigEndian.Uint32(ehdr64.e_flags[:])
                this.ehdr64.e_ehsize = binary.BigEndian.Uint16(ehdr64.e_ehsize[:])
                this.ehdr64.e_phentsize = binary.BigEndian.Uint16(ehdr64.e_phentsize[:])
                this.ehdr64.e_phnum = binary.BigEndian.Uint16(ehdr64.e_phnum[:])
                this.ehdr64.e_shentsize = binary.BigEndian.Uint16(ehdr64.e_shentsize[:])
                this.ehdr64.e_shnum = binary.BigEndian.Uint16(ehdr64.e_shnum[:])
                this.ehdr64.e_shstrndx = binary.BigEndian.Uint16(ehdr64.e_shstrndx[:])
        } else {
                this.ehdr64.e_type = binary.LittleEndian.Uint16(ehdr64.e_type[:])
                this.ehdr64.e_machine = binary.LittleEndian.Uint16(ehdr64.e_machine[:])
                this.ehdr64.e_version = binary.LittleEndian.Uint32(ehdr64.e_version[:])
                this.ehdr64.e_entry = binary.LittleEndian.Uint64(ehdr64.e_entry[:])
                this.ehdr64.e_phoff = binary.LittleEndian.Uint64(ehdr64.e_phoff[:])
                this.ehdr64.e_shoff = binary.LittleEndian.Uint64(ehdr64.e_shoff[:])
                this.ehdr64.e_flags = binary.LittleEndian.Uint32(ehdr64.e_flags[:])
                this.ehdr64.e_ehsize = binary.LittleEndian.Uint16(ehdr64.e_ehsize[:])
                this.ehdr64.e_phentsize = binary.LittleEndian.Uint16(ehdr64.e_phentsize[:])
                this.ehdr64.e_phnum = binary.LittleEndian.Uint16(ehdr64.e_phnum[:])
                this.ehdr64.e_shentsize = binary.LittleEndian.Uint16(ehdr64.e_shentsize[:])
                this.ehdr64.e_shnum = binary.LittleEndian.Uint16(ehdr64.e_shnum[:])
                this.ehdr64.e_shstrndx = binary.LittleEndian.Uint16(ehdr64.e_shstrndx[:])
        }

        pntStr = pntStr + fmt.Sprintf("  %-35s", "Type:")
        switch this.ehdr64.e_type {
        case ET_REL:
                pntStr = pntStr + "REL (Relocatable file)\n"
        case ET_EXEC:
                pntStr = pntStr + "EXEC (Executable file)\n"
        case ET_DYN:
                pntStr = pntStr + "DYN (Shared object file)\n"
        case ET_CORE:
                pntStr = pntStr + "CORE (Core file)\n"
        default:
                pntStr = pntStr + "No file type\n"
        }
        pntStr = pntStr + fmt.Sprintf("  %-35s", "Machine:")
        switch this.ehdr64.e_machine {
        case EM_M32:
                pntStr = pntStr + "AT&T WE 32100\n"
        case EM_SPARC:
                pntStr = pntStr + "SUN SPARC\n"
        case EM_386:
                pntStr = pntStr + "Intel 80386\n"
        case EM_68K:
                pntStr = pntStr + "Motorola m68k family\n"
        case EM_88K:
                pntStr = pntStr + "Motorola m88k family\n"
        case EM_860:
                pntStr = pntStr + "Intel 80860\n"
        case EM_MIPS:
                pntStr = pntStr + "MIPS R3000 big-endian\n"
        case EM_S370:
                pntStr = pntStr + "IBM System/370\n"
        case EM_MIPS_RS3_LE:
                pntStr = pntStr + "MIPS R3000 little-endian\n"
        case EM_PARISC:
                pntStr = pntStr + "HPPA\n"
        case EM_VPP500:
                pntStr = pntStr + "Fujitsu VPP500\n"
        case EM_SPARC32PLUS:
                pntStr = pntStr + "Sun's \"v8plus\"\n"
        case EM_960:
                pntStr = pntStr + "Intel 80960\n"
        case EM_PPC:
                pntStr = pntStr + "PowerPC\n"
        case EM_PPC64:
                pntStr = pntStr + "PowerPC 64-bit\n"
        case EM_S390:
                pntStr = pntStr + "IBM S390\n"
        case EM_V800:
                pntStr = pntStr + "NEC V800 series\n"
        case EM_FR20:
                pntStr = pntStr + "Fujitsu FR20\n"
        case EM_RH32:
                pntStr = pntStr + "TRW RH-32\n"
        case EM_RCE:
                pntStr = pntStr + "Motorola RCE\n"
        case EM_ARM:
                pntStr = pntStr + "ARM\n"
        case EM_FAKE_ALPHA:
                pntStr = pntStr + "Digital Alpha\n"
        case EM_SH:
                pntStr = pntStr + "Hitachi SH\n"
        case EM_SPARCV9:
                pntStr = pntStr + "SPARC v9 64-bit\n"
        case EM_TRICORE:
                pntStr = pntStr + "Siemens Tricore\n"
        case EM_ARC:
                pntStr = pntStr + "Argonaut RISC Core\n"
        case EM_H8_300:
                pntStr = pntStr + "Hitachi H8/300\n"
        case EM_H8_300H:
                pntStr = pntStr + "Hitachi H8/300H\n"
        case EM_H8S:
                pntStr = pntStr + "Hitachi H8S\n"
        case EM_H8_500:
                pntStr = pntStr + "Hitachi H8/500\n"
        case EM_IA_64:
                pntStr = pntStr + "Intel Merced\n"
        case EM_MIPS_X:
                pntStr = pntStr + "Stanford MIPS-X\n"
        case EM_COLDFIRE:
                pntStr = pntStr + "Motorola Coldfire\n"
        case EM_68HC12:
                pntStr = pntStr + "Motorola M68HC12\n"
        case EM_MMA:
                pntStr = pntStr + "Fujitsu MMA Multimedia Accelerator\n"
        case EM_PCP:
                pntStr = pntStr + "Siemens PCP\n"
        case EM_NCPU:
                pntStr = pntStr + "Sony nCPU embeeded RISC\n"
        case EM_NDR1:
                pntStr = pntStr + "Denso NDR1 microprocessor\n"
        case EM_STARCORE:
                pntStr = pntStr + "Motorola Start*Core processor\n"
        case EM_ME16:
                pntStr = pntStr + "Toyota ME16 processor\n"
        case EM_ST100:
                pntStr = pntStr + "STMicroelectronic ST100 processor\n"
        case EM_TINYJ:
                pntStr = pntStr + "Advanced Logic Corp. Tinyj emb.fam\n"
        case EM_X86_64:
                pntStr = pntStr + "Advanced Micro Devices X86-64\n"
        case EM_PDSP:
                pntStr = pntStr + "Sony DSP Processor\n"
        case EM_FX66:
                pntStr = pntStr + "Siemens FX66 microcontroller\n"
        case EM_ST9PLUS:
                pntStr = pntStr + "STMicroelectronics ST9+ 8/16 mc\n"
        case EM_ST7:
                pntStr = pntStr + "STmicroelectronics ST7 8 bit mc\n"
        case EM_68HC16:
                pntStr = pntStr + "Motorola MC68HC16 microcontroller\n"
        case EM_68HC11:
                pntStr = pntStr + "Motorola MC68HC11 microcontroller\n"
        case EM_68HC08:
                pntStr = pntStr + "Motorola MC68HC08 microcontroller\n"
        case EM_68HC05:
                pntStr = pntStr + "Motorola MC68HC05 microcontroller\n"
        case EM_SVX:
                pntStr = pntStr + "Silicon Graphics SVx\n"
        case EM_ST19:
                pntStr = pntStr + "STMicroelectronics ST19 8 bit mc\n"
        case EM_VAX:
                pntStr = pntStr + "Digital VAX\n"
        case EM_CRIS:
                pntStr = pntStr + "Axis Communications 32-bit embedded processor\n"
        case EM_JAVELIN:
                pntStr = pntStr + "Infineon Technologies 32-bit embedded processor\n"
        case EM_FIREPATH:
                pntStr = pntStr + "Element 14 64-bit DSP Processor\n"
        case EM_ZSP:
                pntStr = pntStr + "LSI Logic 16-bit DSP Processor\n"
        case EM_MMIX:
                pntStr = pntStr + "Donald Knuth's educational 64-bit processor\n"
        case EM_HUANY:
                pntStr = pntStr + "Harvard University machine-independent object files\n"
        case EM_PRISM:
                pntStr = pntStr + "SiTera Prism\n"
        case EM_AVR:
                pntStr = pntStr + "Atmel AVR 8-bit microcontroller\n"
        case EM_FR30:
                pntStr = pntStr + "Fujitsu FR30\n"
        case EM_D10V:
                pntStr = pntStr + "Mitsubishi D10V\n"
        case EM_D30V:
                pntStr = pntStr + "Mitsubishi D30V\n"
        case EM_V850:
                pntStr = pntStr + "NEC v850\n"
        case EM_M32R:
                pntStr = pntStr + "Mitsubishi M32R\n"
        case EM_MN10300:
                pntStr = pntStr + "Matsushita MN10300\n"
        case EM_MN10200:
                pntStr = pntStr + "Matsushita MN10200\n"
        case EM_PJ:
                pntStr = pntStr + "picoJava\n"
        case EM_OPENRISC:
                pntStr = pntStr + "OpenRISC 32-bit embedded processor\n"
        case EM_ARC_A5:
                pntStr = pntStr + "ARC Cores Tangent-A5\n"
        case EM_XTENSA:
                pntStr = pntStr + "Tensilica Xtensa Architecture\n"
        default:
                pntStr = pntStr + "No machine\n"
        }
        pntStr = pntStr + fmt.Sprintf("  %-35s0x%x\n", "Version:", this.ehdr64.e_version)
        pntStr = pntStr + fmt.Sprintf("  %-35s0x%x\n", "Entry point address:", this.ehdr64.e_entry)
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Start of program headers:", this.ehdr64.e_phoff, "(bytes into file)")
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Start of section headers:", this.ehdr64.e_shoff, "(bytes into file)")
        pntStr = pntStr + fmt.Sprintf("  %-35s0x%x\n", "Flags:", this.ehdr64.e_flags)
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Size of this header:", this.ehdr64.e_ehsize, "(bytes)")
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Size of program headers:", this.ehdr64.e_phentsize, "(bytes)")
        pntStr = pntStr + fmt.Sprintf("  %-35s%d\n", "Number of program headers:", this.ehdr64.e_phnum)
        pntStr = pntStr + fmt.Sprintf("  %-35s%d %s\n", "Size of section headers:", this.ehdr64.e_shentsize, "(bytes)")
        pntStr = pntStr + fmt.Sprintf("  %-35s%d\n", "Number of section headers:", this.ehdr64.e_shnum)
        pntStr = pntStr + fmt.Sprintf("  %-35s%d\n", "Section header string table index:", this.ehdr64.e_shstrndx)

        return pntStr, nil
}

func (this *ELFReader) PrintELFHeader() error {
        pntStr := "ELF Header:\n"
        ch := make([]byte, 16)
        if _, err := this.f.Read(ch); err != nil {
                return err
        } else {
                pntStr = pntStr + "  Magic:  "
                pntStr = pntStr + fmt.Sprintf("%02x %02x %02x %02x %02x %02x %02x "+
                        "%02x %02x %02x %02x %02x %02x %02x %02x %02x\n",
                        ch[0], ch[1], ch[2], ch[3], ch[4], ch[5], ch[6], ch[7],
                        ch[8], ch[9], ch[10], ch[11], ch[12], ch[13], ch[14], ch[15])
                pntStr = pntStr + fmt.Sprintf("  %-35s", "Class:")
                switch ch[EI_CLASS] {
                case ELFCLASS32:
                        this.i32 = true
                        pntStr = pntStr + "ELF32\n"
                case ELFCLASS64:
                        pntStr = pntStr + "ELF64\n"
                default:
                        return errors.New("Invalid class")
                }
                pntStr = pntStr + fmt.Sprintf("  %-35s", "Data:")
                switch ch[EI_DATA] {
                case ELFDATA2LSB:
                        pntStr = pntStr + "2's complement, little endian\n"
                case ELFDATA2MSB:
                        this.isBigEndian = true
                        pntStr = pntStr + "2's complement, big endian\n"
                default:
                        pntStr = pntStr + "Invalid data encoding\n"
                }
                pntStr = pntStr + fmt.Sprintf("  %-35s", "Version:")
                pntStr = pntStr + fmt.Sprintf("%d (current)\n", EV_CURRENT)
                pntStr = pntStr + fmt.Sprintf("  %-35s", "OS/ABI:")
                switch ch[EI_OSABI] {
                case ELFOSABI_SYSV:
                        pntStr = pntStr + "UNIX - System V\n"
                case ELFOSABI_HPUX:
                        pntStr = pntStr + "HP - UNIX\n"
                case ELFOSABI_NETBSD:
                        pntStr = pntStr + "NetBSD\n"
                case ELFOSABI_LINUX:
                        pntStr = pntStr + "Linux\n"
                case ELFOSABI_SOLARIS:
                        pntStr = pntStr + "Sun Solaris\n"
                case ELFOSABI_AIX:
                        pntStr = pntStr + "IBM AIX\n"
                case ELFOSABI_IRIX:
                        pntStr = pntStr + "SGI Irix\n"
                case ELFOSABI_FREEBSD:
                        pntStr = pntStr + "FreeBSD\n"
                case ELFOSABI_TRU64:
                        pntStr = pntStr + "Compaq TRU64 UNIX\n"
                case ELFOSABI_MODESTO:
                        pntStr = pntStr + "Novell Modesto\n"
                case ELFOSABI_OPENBSD:
                        pntStr = pntStr + "OpenBSD\n"
                case ELFOSABI_ARM:
                        pntStr = pntStr + "ARM\n"
                case ELFOSABI_STANDALONE:
                        pntStr = pntStr + "Standalone (embedded) application\n"
                default:
                        pntStr = pntStr + "Invalid OS ABI identification\n"
                }
                pntStr = pntStr + fmt.Sprintf("  %-35s%d\n", "ABI Version:", ch[EI_ABIVERSION])
        }

        if this.i32 {
                myStr, err := this._GetELFHeaderStr32()
                if err != nil {
                        return err
                }
                fmt.Printf("%s\n", pntStr+myStr)
        } else {
                myStr, err := this._GetELFHeaderStr64()
                if err != nil {
                        return err
                }
                fmt.Printf("%s\n", pntStr+myStr)
        }

        return nil
}

func (this *ELFReader) _GetProgramHeaderStr32() (string, error) {
        pntStr := "Program Headers:\n"
        st := int(this.ehdr32.e_phentsize * this.ehdr32.e_phnum)
        b := make([]byte, st)
        if _, err := this.f.ReadAt(b, int64(this.ehdr32.e_phoff)); err != nil {
                return "", err
        }
        pntStr = pntStr + fmt.Sprintf("  %-15s%-10s %-10s %-10s\n  %15s%-10s %-10s  %-6s %s\n",
                "Type", "Offset", "VirtAddr", "PhysAddr", "", "FileSiz", "MemSiz", "Flags", "Align")
        j := int(this.ehdr32.e_phnum)
        pos := 0
        phdr := new(Elf32_Phdr)
        if this.isBigEndian {
                for i := 0; i < j; i++ {
                        pntStr = pntStr + "  "
                        phdrbyte := (*Elf32_PhdrBytes)(unsafe.Pointer(&b[pos]))

                        phdr.p_type = binary.BigEndian.Uint32(phdrbyte.p_type[:])
                        phdr.p_offset = binary.BigEndian.Uint32(phdrbyte.p_offset[:])
                        phdr.p_vaddr = binary.BigEndian.Uint32(phdrbyte.p_vaddr[:])
                        phdr.p_paddr = binary.BigEndian.Uint32(phdrbyte.p_paddr[:])
                        phdr.p_filesz = binary.BigEndian.Uint32(phdrbyte.p_filesz[:])
                        phdr.p_memsz = binary.BigEndian.Uint32(phdrbyte.p_memsz[:])
                        phdr.p_flags = binary.BigEndian.Uint32(phdrbyte.p_flags[:])
                        phdr.p_align = binary.BigEndian.Uint32(phdrbyte.p_align[:])
                        if phdr.p_type < PT_LOOS {
                                switch int(phdr.p_type) {
                                case PT_LOAD:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "LOAD")
                                case PT_DYNAMIC:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "DYNAMIC")
                                case PT_INTERP:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "INTERP")
                                case PT_NOTE:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "NOTE")
                                case PT_SHLIB:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "SHLIB")
                                case PT_PHDR:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "PHDR")
                                case PT_TLS:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "TLS")
                                case PT_NUM:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "NUM")
                                case PT_NULL:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "UNUSED")
                                default:
                                        return "", fmt.Errorf("未知Segment类型[0x%x]", phdr.p_type)
                                }
                        } else {
                                if phdr.p_type < PT_LOPROC {
                                        switch int(phdr.p_type) {
                                        case PT_LOOS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "LOOS")
                                        case PT_GNU_EH_FRAME:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_EH_FRAME")
                                        case PT_GNU_STACK:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_STACK")
                                        case PT_GNU_RELRO:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_RELRO")
                                        case PT_SUNWBSS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "SUNWBSS")
                                        case PT_SUNWSTACK:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "SUNWSTACK")
                                        case PT_HIOS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "HIOS")
                                        default:
                                                t := fmt.Sprintf("%s+%x", "LOOS", phdr.p_type-PT_LOOS)
                                                pntStr = pntStr + fmt.Sprintf("%-15s", t)
                                        }
                                } else {
                                        if phdr.p_type < PT_HIPROC {
                                                t := fmt.Sprintf("%s+%x", "PT_LOPROC", phdr.p_type-PT_LOPROC)
                                                pntStr = pntStr + fmt.Sprintf("%-15s", t)
                                        } else {
                                                return "", fmt.Errorf("未知Segment类型[0x%x]", phdr.p_type)
                                        }
                                }
                        }
                        pntStr = pntStr + fmt.Sprintf("0x%08x 0x%08x 0x%08x\n",
                                phdr.p_offset, phdr.p_vaddr, phdr.p_paddr)
                        pntStr = pntStr + fmt.Sprintf("  %15s0x%08x 0x%08x ", "", phdr.p_filesz, phdr.p_memsz)
                        tmpStr := ""
                        if phdr.p_flags&PF_R == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "R"
                        }
                        if phdr.p_flags&PF_W == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "W"
                        }
                        if phdr.p_flags&PF_X == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "E"
                        }
                        pntStr = pntStr + fmt.Sprintf(" %-6s %x\n", tmpStr, phdr.p_align)
                        pos += int(this.ehdr64.e_phentsize)
                }
        } else {
                for i := 0; i < j; i++ {
                        pntStr = pntStr + "  "
                        phdrbyte := (*Elf32_PhdrBytes)(unsafe.Pointer(&b[pos]))

                        phdr.p_type = binary.LittleEndian.Uint32(phdrbyte.p_type[:])
                        phdr.p_offset = binary.LittleEndian.Uint32(phdrbyte.p_offset[:])
                        phdr.p_vaddr = binary.LittleEndian.Uint32(phdrbyte.p_vaddr[:])
                        phdr.p_paddr = binary.LittleEndian.Uint32(phdrbyte.p_paddr[:])
                        phdr.p_filesz = binary.LittleEndian.Uint32(phdrbyte.p_filesz[:])
                        phdr.p_memsz = binary.LittleEndian.Uint32(phdrbyte.p_memsz[:])
                        phdr.p_flags = binary.LittleEndian.Uint32(phdrbyte.p_flags[:])
                        phdr.p_align = binary.LittleEndian.Uint32(phdrbyte.p_align[:])

                        if phdr.p_type < PT_LOOS {
                                switch int(phdr.p_type) {
                                case PT_LOAD:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "LOAD")
                                case PT_DYNAMIC:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "DYNAMIC")
                                case PT_INTERP:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "INTERP")
                                case PT_NOTE:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "NOTE")
                                case PT_SHLIB:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "SHLIB")
                                case PT_PHDR:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "PHDR")
                                case PT_TLS:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "TLS")
                                case PT_NUM:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "NUM")
                                case PT_NULL:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "UNUSED")
                                default:
                                        return "", fmt.Errorf("未知Segment类型[0x%x]", phdr.p_type)
                                }
                        } else {
                                if phdr.p_type < PT_LOPROC {
                                        switch int(phdr.p_type) {
                                        case PT_LOOS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "LOOS")
                                        case PT_GNU_EH_FRAME:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_EH_FRAME")
                                        case PT_GNU_STACK:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_STACK")
                                        case PT_GNU_RELRO:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_RELRO")
                                        case PT_SUNWBSS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "SUNWBSS")
                                        case PT_SUNWSTACK:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "SUNWSTACK")
                                        case PT_HIOS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "HIOS")
                                        default:
                                                t := fmt.Sprintf("%s+%x", "LOOS", phdr.p_type-PT_LOOS)
                                                pntStr = pntStr + fmt.Sprintf("%-15s", t)
                                        }
                                } else {
                                        if phdr.p_type < PT_HIPROC {
                                                t := fmt.Sprintf("%s+%x", "PT_LOPROC", phdr.p_type-PT_LOPROC)
                                                pntStr = pntStr + fmt.Sprintf("%-15s", t)
                                        } else {
                                                return "", fmt.Errorf("未知Segment类型[0x%x]", phdr.p_type)
                                        }
                                }
                        }
                        pntStr = pntStr + fmt.Sprintf("0x%08x 0x%08x 0x%08x\n",
                                phdr.p_offset, phdr.p_vaddr, phdr.p_paddr)
                        pntStr = pntStr + fmt.Sprintf("  %15s0x%08x 0x%08x ", "", phdr.p_filesz, phdr.p_memsz)
                        tmpStr := ""
                        if phdr.p_flags&PF_R == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "R"
                        }
                        if phdr.p_flags&PF_W == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "W"
                        }
                        if phdr.p_flags&PF_X == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "E"
                        }
                        pntStr = pntStr + fmt.Sprintf(" %-6s %x\n", tmpStr, phdr.p_align)
                        pos += int(this.ehdr32.e_phentsize)
                }
        }
        return pntStr, nil
}

func (this *ELFReader) _GetProgramHeaderStr64() (string, error) {
        pntStr := "Program Headers:\n"
        st := int(this.ehdr64.e_phentsize * this.ehdr64.e_phnum)
        b := make([]byte, st)
        if _, err := this.f.ReadAt(b, int64(this.ehdr64.e_phoff)); err != nil {
                return "", err
        }
        pntStr = pntStr + fmt.Sprintf("  %-15s%-18s %-18s %-18s\n  %15s%-18s %-18s  %-6s %s\n",
                "Type", "Offset", "VirtAddr", "PhysAddr", "", "FileSiz", "MemSiz", "Flags", "Align")
        j := int(this.ehdr64.e_phnum)
        pos := 0
        phdr := new(Elf64_Phdr)
        if this.isBigEndian {
                for i := 0; i < j; i++ {
                        pntStr = pntStr + "  "
                        phdrbyte := (*Elf64_PhdrBytes)(unsafe.Pointer(&b[pos]))
                        phdr.p_type = binary.BigEndian.Uint32(phdrbyte.p_type[:])
                        phdr.p_flags = binary.BigEndian.Uint32(phdrbyte.p_flags[:])
                        phdr.p_offset = binary.BigEndian.Uint64(phdrbyte.p_offset[:])
                        phdr.p_vaddr = binary.BigEndian.Uint64(phdrbyte.p_vaddr[:])
                        phdr.p_paddr = binary.BigEndian.Uint64(phdrbyte.p_paddr[:])
                        phdr.p_filesz = binary.BigEndian.Uint64(phdrbyte.p_filesz[:])
                        phdr.p_memsz = binary.BigEndian.Uint64(phdrbyte.p_memsz[:])
                        phdr.p_align = binary.BigEndian.Uint64(phdrbyte.p_align[:])
                        if phdr.p_type < PT_LOOS {
                                switch int(phdr.p_type) {
                                case PT_LOAD:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "LOAD")
                                case PT_DYNAMIC:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "DYNAMIC")
                                case PT_INTERP:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "INTERP")
                                case PT_NOTE:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "NOTE")
                                case PT_SHLIB:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "SHLIB")
                                case PT_PHDR:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "PHDR")
                                case PT_TLS:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "TLS")
                                case PT_NUM:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "NUM")
                                case PT_NULL:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "UNUSED")
                                default:
                                        return "", fmt.Errorf("未知Segment类型[0x%x]", phdr.p_type)
                                }
                        } else {
                                if phdr.p_type < PT_LOPROC {
                                        switch int(phdr.p_type) {
                                        case PT_LOOS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "LOOS")
                                        case PT_GNU_EH_FRAME:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_EH_FRAME")
                                        case PT_GNU_STACK:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_STACK")
                                        case PT_GNU_RELRO:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_RELRO")
                                        case PT_SUNWBSS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "SUNWBSS")
                                        case PT_SUNWSTACK:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "SUNWSTACK")
                                        case PT_HIOS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "HIOS")
                                        default:
                                                t := fmt.Sprintf("%s+%x", "LOOS", phdr.p_type-PT_LOOS)
                                                pntStr = pntStr + fmt.Sprintf("%-15s", t)
                                        }
                                } else {
                                        if phdr.p_type < PT_HIPROC {
                                                t := fmt.Sprintf("%s+%x", "PT_LOPROC", phdr.p_type-PT_LOPROC)
                                                pntStr = pntStr + fmt.Sprintf("%-15s", t)
                                        } else {
                                                return "", fmt.Errorf("未知Segment类型[0x%x]", phdr.p_type)
                                        }
                                }
                        }
                        pntStr = pntStr + fmt.Sprintf("0x%016x 0x%016x 0x%016x\n",
                                phdr.p_offset, phdr.p_vaddr, phdr.p_paddr)
                        pntStr = pntStr + fmt.Sprintf("  %15s0x%016x 0x%016x ", "", phdr.p_filesz, phdr.p_memsz)
                        tmpStr := ""
                        if phdr.p_flags&PF_R == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "R"
                        }
                        if phdr.p_flags&PF_W == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "W"
                        }
                        if phdr.p_flags&PF_X == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "E"
                        }
                        pntStr = pntStr + fmt.Sprintf(" %-6s %x\n", tmpStr, phdr.p_align)
                        pos += int(this.ehdr64.e_phentsize)
                }
        } else {
                for i := 0; i < j; i++ {
                        pntStr = pntStr + "  "
                        phdrbyte := (*Elf64_PhdrBytes)(unsafe.Pointer(&b[pos]))
                        phdr.p_type = binary.LittleEndian.Uint32(phdrbyte.p_type[:])
                        phdr.p_flags = binary.LittleEndian.Uint32(phdrbyte.p_flags[:])
                        phdr.p_offset = binary.LittleEndian.Uint64(phdrbyte.p_offset[:])
                        phdr.p_vaddr = binary.LittleEndian.Uint64(phdrbyte.p_vaddr[:])
                        phdr.p_paddr = binary.LittleEndian.Uint64(phdrbyte.p_paddr[:])
                        phdr.p_filesz = binary.LittleEndian.Uint64(phdrbyte.p_filesz[:])
                        phdr.p_memsz = binary.LittleEndian.Uint64(phdrbyte.p_memsz[:])
                        phdr.p_align = binary.LittleEndian.Uint64(phdrbyte.p_align[:])
                        if phdr.p_type < PT_LOOS {
                                switch int(phdr.p_type) {
                                case PT_LOAD:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "LOAD")
                                case PT_DYNAMIC:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "DYNAMIC")
                                case PT_INTERP:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "INTERP")
                                case PT_NOTE:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "NOTE")
                                case PT_SHLIB:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "SHLIB")
                                case PT_PHDR:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "PHDR")
                                case PT_TLS:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "TLS")
                                case PT_NUM:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "NUM")
                                case PT_NULL:
                                        pntStr = pntStr + fmt.Sprintf("%-15s", "UNUSED")
                                default:
                                        return "", fmt.Errorf("未知Segment类型[0x%x]", phdr.p_type)
                                }
                        } else {
                                if phdr.p_type < PT_LOPROC {
                                        switch int(phdr.p_type) {
                                        case PT_LOOS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "LOOS")
                                        case PT_GNU_EH_FRAME:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_EH_FRAME")
                                        case PT_GNU_STACK:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_STACK")
                                        case PT_GNU_RELRO:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "GNU_RELRO")
                                        case PT_SUNWBSS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "SUNWBSS")
                                        case PT_SUNWSTACK:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "SUNWSTACK")
                                        case PT_HIOS:
                                                pntStr = pntStr + fmt.Sprintf("%-15s", "HIOS")
                                        default:
                                                t := fmt.Sprintf("%s+%x", "LOOS", phdr.p_type-PT_LOOS)
                                                pntStr = pntStr + fmt.Sprintf("%-15s", t)
                                        }
                                } else {
                                        if phdr.p_type < PT_HIPROC {
                                                t := fmt.Sprintf("%s+%x", "PT_LOPROC", phdr.p_type-PT_LOPROC)
                                                pntStr = pntStr + fmt.Sprintf("%-15s", t)
                                        } else {
                                                return "", fmt.Errorf("未知Segment类型[0x%x]", phdr.p_type)
                                        }
                                }
                        }
                        pntStr = pntStr + fmt.Sprintf("0x%016x 0x%016x 0x%016x\n",
                                phdr.p_offset, phdr.p_vaddr, phdr.p_paddr)
                        pntStr = pntStr + fmt.Sprintf("  %15s0x%016x 0x%016x ", "", phdr.p_filesz, phdr.p_memsz)
                        tmpStr := ""
                        if phdr.p_flags&PF_R == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "R"
                        }
                        if phdr.p_flags&PF_W == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "W"
                        }
                        if phdr.p_flags&PF_X == 0x0 {
                                tmpStr = tmpStr + " "
                        } else {
                                tmpStr = tmpStr + "E"
                        }
                        pntStr = pntStr + fmt.Sprintf(" %-6s %x\n", tmpStr, phdr.p_align)
                        pos += int(this.ehdr64.e_phentsize)
                }
        }
        return pntStr, nil
}

func (this *ELFReader) PrintProgramHeader() error {
        if this.i32 {
                if myStr, err := this._GetProgramHeaderStr32(); err != nil {
                        return err
                } else {
                        fmt.Printf("%s\n", myStr)
                }

        } else {
                if myStr, err := this._GetProgramHeaderStr64(); err != nil {
                        return err
                } else {
                        fmt.Printf("%s\n", myStr)
                }
        }
        return nil
}

func (this *ELFReader) _GetSectionHeaderStr32() (string, error) {
        var sh_size int
        var sh_offset int64
        pntStr := "Section Headers:\n"
        st := int(this.ehdr32.e_shentsize * this.ehdr32.e_shnum)
        b := make([]byte, st)
        if _, err := this.f.ReadAt(b, int64(this.ehdr32.e_shoff)); err != nil {
                return "", err
        }
        pntStr = pntStr + fmt.Sprintf("  %4s %-17s %-8s %-8s  %s\n  %4s %-8s  % -8s %s  %s  %s  %s\n",
                "[Nr]", "Name", "Type", "Address", "Offset", "", "Size", "EntSize", "Flags", "Link", "Info", "Align")
        j := int(this.ehdr32.e_shnum)
        pos := 0
        pos = int(this.ehdr32.e_shstrndx * this.ehdr32.e_shentsize)
        shdrbyte := (*Elf32_ShdrBytes)(unsafe.Pointer(&b[pos]))

        if this.isBigEndian {
                sh_size = int(binary.BigEndian.Uint32(shdrbyte.sh_size[:]))
                sh_offset = int64(binary.BigEndian.Uint32(shdrbyte.sh_offset[:]))
        } else {
                sh_size = int(binary.LittleEndian.Uint32(shdrbyte.sh_size[:]))
                sh_offset = int64(binary.LittleEndian.Uint32(shdrbyte.sh_offset[:]))
        }
        strTab := make([]byte, sh_size)
        if _, err := this.f.ReadAt(strTab, sh_offset); err != nil {
                return "", err
        }
        shdr := new(Elf32_Shdr)
        pos = 0
        if this.isBigEndian {
                for i := 0; i < j; i++ {
                        pntStr = pntStr + "  "
                        shdrbyte := (*Elf32_ShdrBytes)(unsafe.Pointer(&b[pos]))

                        shdr.sh_name = binary.BigEndian.Uint32(shdrbyte.sh_name[:])
                        shdr.sh_type = binary.BigEndian.Uint32(shdrbyte.sh_type[:])
                        shdr.sh_flags = binary.BigEndian.Uint32(shdrbyte.sh_flags[:])
                        shdr.sh_addr = binary.BigEndian.Uint32(shdrbyte.sh_addr[:])
                        shdr.sh_offset = binary.BigEndian.Uint32(shdrbyte.sh_offset[:])
                        shdr.sh_size = binary.BigEndian.Uint32(shdrbyte.sh_size[:])
                        shdr.sh_link = binary.BigEndian.Uint32(shdrbyte.sh_link[:])
                        shdr.sh_info = binary.BigEndian.Uint32(shdrbyte.sh_info[:])
                        shdr.sh_addralign = binary.BigEndian.Uint32(shdrbyte.sh_addralign[:])
                        shdr.sh_entsize = binary.BigEndian.Uint32(shdrbyte.sh_entsize[:])

                        pntStr = pntStr + fmt.Sprintf("[%02d] ", i)
                        pntStr = pntStr + fmt.Sprintf("%-17.17s ", getStr(int(shdr.sh_name), strTab))
                        t := ""
                        switch shdr.sh_type {
                        case SHT_NULL:
                                t = t + "NULL"
                        case SHT_PROGBITS:
                                t = t + "PROGBITS"
                        case SHT_SYMTAB:
                                t = t + "SYMTAB"
                        case SHT_STRTAB:
                                t = t + "STRTAB"
                        case SHT_RELA:
                                t = t + "RELA"
                        case SHT_HASH:
                                t = t + "HASH"
                        case SHT_DYNAMIC:
                                t = t + "DYNAMIC"
                        case SHT_NOTE:
                                t = t + "NOTE"
                        case SHT_NOBITS:
                                t = t + "NOBITS"
                        case SHT_REL:
                                t = t + "REL"
                        case SHT_SHLIB:
                                t = t + "SHLIB"
                        case SHT_DYNSYM:
                                t = t + "DYNSYM"
                        case SHT_INIT_ARRAY:
                                t = t + "INIT_ARRAY"
                        case SHT_FINI_ARRAY:
                                t = t + "FINI_ARRAY"
                        case SHT_PREINIT_ARRAY:
                                t = t + "PREINIT_ARRAY"
                        case SHT_GROUP:
                                t = t + "GROUP"
                        case SHT_SYMTAB_SHNDX:
                                t = t + "SYMTAB_SHNDX"
                        case SHT_NUM:
                                t = t + "NUM"
                        case SHT_LOOS:
                                t = t + "LOOS"
                        case SHT_GNU_HASH:
                                t = t + "GNU_HASH"
                        case SHT_GNU_LIBLIST:
                                t = t + "GNU_LIBLIST"
                        case SHT_CHECKSUM:
                                t = t + "CHECKSUM"
                        case SHT_LOSUNW:
                                t = t + "LOSUNW"
                        case SHT_SUNW_COMDAT:
                                t = t + "SUNW_COMDAT"
                        case SHT_SUNW_syminfo:
                                t = t + "SUNW_syminfo"
                        case SHT_GNU_verdef:
                                t = t + "GNU_verdef"
                        case SHT_GNU_verneed:
                                t = t + "GNU_verneed"
                        case SHT_HISUNW:
                                t = t + "HISUNW"
                        case SHT_LOPROC:
                                t = t + "LOPROC"
                        case SHT_HIPROC:
                                t = t + "HIPROC"
                        case SHT_LOUSER:
                                t = t + "LOUSER"
                        case SHT_HIUSER:
                                t = t + "HIUSER"
                        default:
                                return "", fmt.Errorf("未知Section类型[0x%x]", shdr.sh_type)
                        }
                        pntStr = pntStr + fmt.Sprintf(" %-16s %08x %08x\n", t, shdr.sh_addr, shdr.sh_offset)
                        pntStr = pntStr + fmt.Sprintf("  %4s %08x  %08x %-5s  %4d  %4d  %5d\n",
                                "", shdr.sh_size, shdr.sh_entsize, getFlags(uint(shdr.sh_flags)),
                                shdr.sh_link, shdr.sh_info, shdr.sh_addralign)
                        pos += int(this.ehdr64.e_shentsize)
                }
        } else {
                for i := 0; i < j; i++ {
                        pntStr = pntStr + "  "
                        shdrbyte := (*Elf32_ShdrBytes)(unsafe.Pointer(&b[pos]))

                        shdr.sh_name = binary.LittleEndian.Uint32(shdrbyte.sh_name[:])
                        shdr.sh_type = binary.LittleEndian.Uint32(shdrbyte.sh_type[:])
                        shdr.sh_flags = binary.LittleEndian.Uint32(shdrbyte.sh_flags[:])
                        shdr.sh_addr = binary.LittleEndian.Uint32(shdrbyte.sh_addr[:])
                        shdr.sh_offset = binary.LittleEndian.Uint32(shdrbyte.sh_offset[:])
                        shdr.sh_size = binary.LittleEndian.Uint32(shdrbyte.sh_size[:])
                        shdr.sh_link = binary.LittleEndian.Uint32(shdrbyte.sh_link[:])
                        shdr.sh_info = binary.LittleEndian.Uint32(shdrbyte.sh_info[:])
                        shdr.sh_addralign = binary.LittleEndian.Uint32(shdrbyte.sh_addralign[:])
                        shdr.sh_entsize = binary.LittleEndian.Uint32(shdrbyte.sh_entsize[:])

                        pntStr = pntStr + fmt.Sprintf("[%02d] ", i)
                        pntStr = pntStr + fmt.Sprintf("%-17.17s", getStr(int(shdr.sh_name), strTab))
                        t := ""
                        switch shdr.sh_type {
                        case SHT_NULL:
                                t = t + "NULL"
                        case SHT_PROGBITS:
                                t = t + "PROGBITS"
                        case SHT_SYMTAB:
                                t = t + "SYMTAB"
                        case SHT_STRTAB:
                                t = t + "STRTAB"
                        case SHT_RELA:
                                t = t + "RELA"
                        case SHT_HASH:
                                t = t + "HASH"
                        case SHT_DYNAMIC:
                                t = t + "DYNAMIC"
                        case SHT_NOTE:
                                t = t + "NOTE"
                        case SHT_NOBITS:
                                t = t + "NOBITS"
                        case SHT_REL:
                                t = t + "REL"
                        case SHT_SHLIB:
                                t = t + "SHLIB"
                        case SHT_DYNSYM:
                                t = t + "DYNSYM"
                        case SHT_INIT_ARRAY:
                                t = t + "INIT_ARRAY"
                        case SHT_FINI_ARRAY:
                                t = t + "FINI_ARRAY"
                        case SHT_PREINIT_ARRAY:
                                t = t + "PREINIT_ARRAY"
                        case SHT_GROUP:
                                t = t + "GROUP"
                        case SHT_SYMTAB_SHNDX:
                                t = t + "SYMTAB_SHNDX"
                        case SHT_NUM:
                                t = t + "NUM"
                        case SHT_LOOS:
                                t = t + "LOOS"
                        case SHT_GNU_HASH:
                                t = t + "GNU_HASH"
                        case SHT_GNU_LIBLIST:
                                t = t + "GNU_LIBLIST"
                        case SHT_CHECKSUM:
                                t = t + "CHECKSUM"
                        case SHT_LOSUNW:
                                t = t + "LOSUNW"
                        case SHT_SUNW_COMDAT:
                                t = t + "SUNW_COMDAT"
                        case SHT_SUNW_syminfo:
                                t = t + "SUNW_syminfo"
                        case SHT_GNU_verdef:
                                t = t + "GNU_verdef"
                        case SHT_GNU_verneed:
                                t = t + "GNU_verneed"
                        case SHT_HISUNW:
                                t = t + "HISUNW"
                        case SHT_LOPROC:
                                t = t + "LOPROC"
                        case SHT_HIPROC:
                                t = t + "HIPROC"
                        case SHT_LOUSER:
                                t = t + "LOUSER"
                        case SHT_HIUSER:
                                t = t + "HIUSER"
                        default:
                                return "", fmt.Errorf("未知Section类型[0x%x]", shdr.sh_type)
                        }
                        pntStr = pntStr + fmt.Sprintf(" %-16s %08x %08x\n", t, shdr.sh_addr, shdr.sh_offset)
                        pntStr = pntStr + fmt.Sprintf("  %4s %08x  %08x %-5s  %4d  %4d  %5d\n",
                                "", shdr.sh_size, shdr.sh_entsize, getFlags(uint(shdr.sh_flags)),
                                shdr.sh_link, shdr.sh_info, shdr.sh_addralign)
                        pos += int(this.ehdr32.e_shentsize)
                }
        }
        pntStr = pntStr + "Key to Flags:\n" +
                "  W (write), A (alloc), X (execute), M (merge), S (strings)\n" +
                "  I (info), L (link order), G (group), x (unknown)\n" +
                "  O (extra OS processing required) o (OS specific), p (processor specific)\n"
        return pntStr, nil
}

func (this *ELFReader) _GetSectionHeaderStr64() (string, error) {
        var sh_size int
        var sh_offset int64
        pntStr := "Section Headers:\n"
        st := int(this.ehdr64.e_shentsize * this.ehdr64.e_shnum)
        b := make([]byte, st)
        if _, err := this.f.ReadAt(b, int64(this.ehdr64.e_shoff)); err != nil {
                return "", err
        }
        pntStr = pntStr + fmt.Sprintf("  %4s %-17s %-16s %-16s  %s\n  %4s %-16s  % -16s %s  %s  %s  %s\n",
                "[Nr]", "Name", "Type", "Address", "Offset", "", "Size", "EntSize", "Flags", "Link", "Info", "Align")
        j := int(this.ehdr64.e_shnum)
        pos := 0
        pos = int(this.ehdr64.e_shstrndx * this.ehdr64.e_shentsize)
        shdrbyte := (*Elf64_ShdrBytes)(unsafe.Pointer(&b[pos]))

        if this.isBigEndian {
                sh_size = int(binary.BigEndian.Uint64(shdrbyte.sh_size[:]))
                sh_offset = int64(binary.BigEndian.Uint64(shdrbyte.sh_offset[:]))
        } else {
                sh_size = int(binary.LittleEndian.Uint64(shdrbyte.sh_size[:]))
                sh_offset = int64(binary.LittleEndian.Uint64(shdrbyte.sh_offset[:]))
        }
        strTab := make([]byte, sh_size)
        if _, err := this.f.ReadAt(strTab, sh_offset); err != nil {
                return "", err
        }
        shdr := new(Elf64_Shdr)
        pos = 0
        if this.isBigEndian {
                for i := 0; i < j; i++ {
                        pntStr = pntStr + "  "
                        shdrbyte := (*Elf64_ShdrBytes)(unsafe.Pointer(&b[pos]))

                        shdr.sh_name = binary.BigEndian.Uint32(shdrbyte.sh_name[:])
                        shdr.sh_type = binary.BigEndian.Uint32(shdrbyte.sh_type[:])
                        shdr.sh_flags = binary.BigEndian.Uint64(shdrbyte.sh_flags[:])
                        shdr.sh_addr = binary.BigEndian.Uint64(shdrbyte.sh_addr[:])
                        shdr.sh_offset = binary.BigEndian.Uint64(shdrbyte.sh_offset[:])
                        shdr.sh_size = binary.BigEndian.Uint64(shdrbyte.sh_size[:])
                        shdr.sh_link = binary.BigEndian.Uint32(shdrbyte.sh_link[:])
                        shdr.sh_info = binary.BigEndian.Uint32(shdrbyte.sh_info[:])
                        shdr.sh_addralign = binary.BigEndian.Uint64(shdrbyte.sh_addralign[:])
                        shdr.sh_entsize = binary.BigEndian.Uint64(shdrbyte.sh_entsize[:])

                        pntStr = pntStr + fmt.Sprintf("[%02d] ", i)
                        pntStr = pntStr + fmt.Sprintf("%-17.17s ", getStr(int(shdr.sh_name), strTab))
                        t := ""
                        switch shdr.sh_type {
                        case SHT_NULL:
                                t = t + "NULL"
                        case SHT_PROGBITS:
                                t = t + "PROGBITS"
                        case SHT_SYMTAB:
                                t = t + "SYMTAB"
                        case SHT_STRTAB:
                                t = t + "STRTAB"
                        case SHT_RELA:
                                t = t + "RELA"
                        case SHT_HASH:
                                t = t + "HASH"
                        case SHT_DYNAMIC:
                                t = t + "DYNAMIC"
                        case SHT_NOTE:
                                t = t + "NOTE"
                        case SHT_NOBITS:
                                t = t + "NOBITS"
                        case SHT_REL:
                                t = t + "REL"
                        case SHT_SHLIB:
                                t = t + "SHLIB"
                        case SHT_DYNSYM:
                                t = t + "DYNSYM"
                        case SHT_INIT_ARRAY:
                                t = t + "INIT_ARRAY"
                        case SHT_FINI_ARRAY:
                                t = t + "FINI_ARRAY"
                        case SHT_PREINIT_ARRAY:
                                t = t + "PREINIT_ARRAY"
                        case SHT_GROUP:
                                t = t + "GROUP"
                        case SHT_SYMTAB_SHNDX:
                                t = t + "SYMTAB_SHNDX"
                        case SHT_NUM:
                                t = t + "NUM"
                        case SHT_LOOS:
                                t = t + "LOOS"
                        case SHT_GNU_HASH:
                                t = t + "GNU_HASH"
                        case SHT_GNU_LIBLIST:
                                t = t + "GNU_LIBLIST"
                        case SHT_CHECKSUM:
                                t = t + "CHECKSUM"
                        case SHT_LOSUNW:
                                t = t + "LOSUNW"
                        case SHT_SUNW_COMDAT:
                                t = t + "SUNW_COMDAT"
                        case SHT_SUNW_syminfo:
                                t = t + "SUNW_syminfo"
                        case SHT_GNU_verdef:
                                t = t + "GNU_verdef"
                        case SHT_GNU_verneed:
                                t = t + "GNU_verneed"
                        case SHT_HISUNW:
                                t = t + "HISUNW"
                        case SHT_LOPROC:
                                t = t + "LOPROC"
                        case SHT_HIPROC:
                                t = t + "HIPROC"
                        case SHT_LOUSER:
                                t = t + "LOUSER"
                        case SHT_HIUSER:
                                t = t + "HIUSER"
                        default:
                                return "", fmt.Errorf("未知Section类型[0x%x]", shdr.sh_type)
                        }
                        pntStr = pntStr + fmt.Sprintf(" %-16s %016x %08x\n", t, shdr.sh_addr, shdr.sh_offset)
                        pntStr = pntStr + fmt.Sprintf("  %4s %016x  %016x %-5s  %4d  %4d  %5d\n",
                                "", shdr.sh_size, shdr.sh_entsize, getFlags(uint(shdr.sh_flags)),
                                shdr.sh_link, shdr.sh_info, shdr.sh_addralign)
                        pos += int(this.ehdr64.e_shentsize)
                }
        } else {
                for i := 0; i < j; i++ {
                        pntStr = pntStr + "  "
                        shdrbyte := (*Elf64_ShdrBytes)(unsafe.Pointer(&b[pos]))

                        shdr.sh_name = binary.LittleEndian.Uint32(shdrbyte.sh_name[:])
                        shdr.sh_type = binary.LittleEndian.Uint32(shdrbyte.sh_type[:])
                        shdr.sh_flags = binary.LittleEndian.Uint64(shdrbyte.sh_flags[:])
                        shdr.sh_addr = binary.LittleEndian.Uint64(shdrbyte.sh_addr[:])
                        shdr.sh_offset = binary.LittleEndian.Uint64(shdrbyte.sh_offset[:])
                        shdr.sh_size = binary.LittleEndian.Uint64(shdrbyte.sh_size[:])
                        shdr.sh_link = binary.LittleEndian.Uint32(shdrbyte.sh_link[:])
                        shdr.sh_info = binary.LittleEndian.Uint32(shdrbyte.sh_info[:])
                        shdr.sh_addralign = binary.LittleEndian.Uint64(shdrbyte.sh_addralign[:])
                        shdr.sh_entsize = binary.LittleEndian.Uint64(shdrbyte.sh_entsize[:])

                        pntStr = pntStr + fmt.Sprintf("[%02d] ", i)
                        pntStr = pntStr + fmt.Sprintf("%-17.17s", getStr(int(shdr.sh_name), strTab))
                        t := ""
                        switch shdr.sh_type {
                        case SHT_NULL:
                                t = t + "NULL"
                        case SHT_PROGBITS:
                                t = t + "PROGBITS"
                        case SHT_SYMTAB:
                                t = t + "SYMTAB"
                        case SHT_STRTAB:
                                t = t + "STRTAB"
                        case SHT_RELA:
                                t = t + "RELA"
                        case SHT_HASH:
                                t = t + "HASH"
                        case SHT_DYNAMIC:
                                t = t + "DYNAMIC"
                        case SHT_NOTE:
                                t = t + "NOTE"
                        case SHT_NOBITS:
                                t = t + "NOBITS"
                        case SHT_REL:
                                t = t + "REL"
                        case SHT_SHLIB:
                                t = t + "SHLIB"
                        case SHT_DYNSYM:
                                t = t + "DYNSYM"
                        case SHT_INIT_ARRAY:
                                t = t + "INIT_ARRAY"
                        case SHT_FINI_ARRAY:
                                t = t + "FINI_ARRAY"
                        case SHT_PREINIT_ARRAY:
                                t = t + "PREINIT_ARRAY"
                        case SHT_GROUP:
                                t = t + "GROUP"
                        case SHT_SYMTAB_SHNDX:
                                t = t + "SYMTAB_SHNDX"
                        case SHT_NUM:
                                t = t + "NUM"
                        case SHT_LOOS:
                                t = t + "LOOS"
                        case SHT_GNU_HASH:
                                t = t + "GNU_HASH"
                        case SHT_GNU_LIBLIST:
                                t = t + "GNU_LIBLIST"
                        case SHT_CHECKSUM:
                                t = t + "CHECKSUM"
                        case SHT_LOSUNW:
                                t = t + "LOSUNW"
                        case SHT_SUNW_COMDAT:
                                t = t + "SUNW_COMDAT"
                        case SHT_SUNW_syminfo:
                                t = t + "SUNW_syminfo"
                        case SHT_GNU_verdef:
                                t = t + "GNU_verdef"
                        case SHT_GNU_verneed:
                                t = t + "GNU_verneed"
                        case SHT_HISUNW:
                                t = t + "HISUNW"
                        case SHT_LOPROC:
                                t = t + "LOPROC"
                        case SHT_HIPROC:
                                t = t + "HIPROC"
                        case SHT_LOUSER:
                                t = t + "LOUSER"
                        case SHT_HIUSER:
                                t = t + "HIUSER"
                        default:
                                return "", fmt.Errorf("未知Section类型[0x%x]", shdr.sh_type)
                        }
                        pntStr = pntStr + fmt.Sprintf(" %-16s %016x %08x\n", t, shdr.sh_addr, shdr.sh_offset)
                        pntStr = pntStr + fmt.Sprintf("  %4s %016x  %016x %-5s  %4d  %4d  %5d\n",
                                "", shdr.sh_size, shdr.sh_entsize, getFlags(uint(shdr.sh_flags)),
                                shdr.sh_link, shdr.sh_info, shdr.sh_addralign)
                        pos += int(this.ehdr64.e_shentsize)
                }
        }
        pntStr = pntStr + "Key to Flags:\n" +
                "  W (write), A (alloc), X (execute), M (merge), S (strings)\n" +
                "  I (info), L (link order), G (group), x (unknown)\n" +
                "  O (extra OS processing required) o (OS specific), p (processor specific)\n"
        return pntStr, nil
}

func getStr(index int, bStr []byte) string {
        sLen := len(bStr)
        for i := index; i < sLen; i++ {
                if bStr[i] == 0x0 {
                        return string(bStr[index:i])
                }
        }
        return ""
}

func getFlags(f uint) string {
        str := ""
        if f&SHF_WRITE != 0 {
                str = str + "W"
        }
        if f&SHF_ALLOC != 0 {
                str = str + "A"
        }
        if f&SHF_EXECINSTR != 0 {
                str = str + "X"
        }
        if f&SHF_MERGE != 0 {
                str = str + "M"
        }
        if f&SHF_STRINGS != 0 {
                str = str + "S"
        }
        if f&SHF_INFO_LINK != 0 {
                str = str + "I"
        }
        if f&SHF_LINK_ORDER != 0 {
                str = str + "L"
        }
        if f&SHF_OS_NONCONFORMING != 0 {
                str = str + "O"
        }
        if f&SHF_GROUP != 0 {
                str = str + "G"
        }
        if f&SHF_TLS != 0 {
                str = str + "W"
        }
        if f&SHF_ORDERED != 0 {
                str = str + "x"
        }
        if f&SHF_MASKOS != 0 {
                str = str + "o"
        }
        if f&SHF_MASKPROC != 0 {
                str = str + "p"
        }

        return str
}

func (this *ELFReader) PrintSectionHeader() error {
        if this.i32 {
                if myStr, err := this._GetSectionHeaderStr32(); err != nil {
                        return err
                } else {
                        fmt.Printf("%s\n", myStr)
                }

        } else {
                if myStr, err := this._GetSectionHeaderStr64(); err != nil {
                        return err
                } else {
                        fmt.Printf("%s\n", myStr)
                }
        }
        return nil
}
func (this *ELFReader) Destroy() error {
        return this.f.Close()
}

func main() {
        if len(os.Args) == 1 {
                fmt.Printf("%s: elf-file(s)\n", os.Args[0])
                return
        }

        er, err := NewELFReader(os.Args[1])
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        defer er.Destroy()

        if err := er.PrintELFHeader(); err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        if err := er.PrintProgramHeader(); err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
        if err := er.PrintSectionHeader(); err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }
}
