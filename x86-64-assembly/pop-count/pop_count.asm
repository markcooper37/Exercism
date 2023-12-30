section .text
global egg_count
egg_count:
    mov rax, 0

_loop:
    cmp rdi, 0
    jz _end
    shr rdi, 1
    jc _inc_counter
    jmp _loop

_end:
    ret

_inc_counter:
    inc rax
    jmp _loop

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
