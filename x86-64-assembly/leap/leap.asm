section .text
global leap_year

leap_year:
    mov rax, rdi
    mov rdx, 0
    mov rcx, 400
    div rcx
    cmp rdx, 0
    jz _is_leap_year

    mov rax, rdi
    mov rdx, 0
    mov rcx, 100
    div rcx
    cmp rdx, 0
    jz _not_leap_year

    mov rax, rdi
    mov rdx, 0
    mov rcx, 4
    div rcx
    cmp rdx, 0
    jz _is_leap_year

_not_leap_year:
    mov rax, 0
    ret

_is_leap_year:
    mov rax, 1
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
