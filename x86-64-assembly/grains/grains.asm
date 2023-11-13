section .text
global square
square:
    cmp rdi, 1
    jl _exception
    cmp rdi, 64
    jg _exception
    mov rax, 1
    mov rcx, rdi
    sub rcx, 1
    shl rax, cl
    ret

_exception:
    mov rax, 0
    ret

global total
total:
    xor rdx, rdx
    mov rdi, 64

_total_loop:
    call square
    add rdx, rax
    dec rdi
    jg _total_loop
    mov rax, rdx
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
