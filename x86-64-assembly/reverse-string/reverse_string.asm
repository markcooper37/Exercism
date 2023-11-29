section .text
global reverse
reverse:
    ; Save start of string
    mov rsi, rdi

_find_end:
    ; Find end of string
	cmp byte [rsi], 0
    je _swap_loop
    inc rsi
    jmp _find_end

_swap_loop:
    ; End if all characters have been swapped
    dec rsi 
    cmp rsi, rdi
    jng _end
    ; Swap characters
	mov dl, [rsi]
    xchg dl, [rdi]
	mov [rsi], dl
	inc rdi
    jmp _swap_loop

_end:
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
