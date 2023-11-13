section .text
global steps
steps:
    test edi, edi
    jle _exception
    mov rax, rdi
    mov rbx, 0
    jmp _loop
    
_exception:
    mov rax, -1
    ret

_loop:
    cmp rax, 1
    je _finalise
    test rax, 1
    jz _handle_even
    jmp _handle_odd

_finalise:
    mov rax, rbx
    ret

_handle_even:
    inc rbx
    shr rax, 1
    jmp _loop

_handle_odd:
    inc rbx
    imul rax, 3
    inc rax
    jmp _loop

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
