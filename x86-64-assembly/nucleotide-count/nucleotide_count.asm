section .text
global nucleotide_counts
nucleotide_counts:
    ; Set counts equal to zero
    xor rax, rax
    mov [rsi], rax
    mov [rsi+8], rax
    mov [rsi+16], rax
    mov [rsi+24], rax
    jmp _loop

_loop:
    mov al, [rdi]
    inc rdi
    ; End if all characters have been counted
    test al, al
    jz _valid_end
    ; Check each possible nucleotide for equality
    mov edx, -1
    mov ecx, 0
    cmp al, 'A'
    cmove edx, ecx
    inc ecx
    cmp al, 'C'
    cmove edx, ecx
    inc ecx
    cmp al, 'G'
    cmove edx, ecx
    inc ecx
    cmp al, 'T'
    cmove edx, ecx
    ; Check whether nucleotide is invalid
    test edx, edx
    jl _invalid_end
    inc qword [rsi+rdx*8]
    jmp _loop

_valid_end:
    ret

_invalid_end:
    ; Set counts equal to -1
    mov rax, -1
    mov [rsi], rax
    mov [rsi+8], rax
    mov [rsi+16], rax
    mov [rsi+24], rax
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
