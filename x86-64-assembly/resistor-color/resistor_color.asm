default rel
section .rodata
black: db "black", 0
brown: db "brown", 0
red: db "red", 0
orange: db "orange", 0
yellow: db "yellow", 0
green: db "green", 0
blue: db "blue", 0
violet: db "violet", 0
grey: db "grey", 0
white: db "white", 0

section .data
all_colors: dq black, brown, red, orange, yellow, green, blue, violet, grey, white, 0

section .text
global color_code
color_code:
    lea rdx, [all_colors]
    xor rcx, rcx

_while:
    ; Test each word in turn to see if it matches
    mov rsi, [rdx + rcx * 8]
    call _compare_colours
    test rax, rax
    jnz _end
    inc rcx
    jmp _while

_end:
    mov rax, rcx
    ret

_compare_colours:
    mov al, byte[rsi]
    cmp	al, byte[rdi]
    jnz	_colours_no_match

_compare_colours_while:
    ; Compare each byte of the colours
    cmp	[rdi], byte 0
    je	_colours_match
    inc	rdi
    inc	rsi
    mov al, byte[rsi]
    cmp	al, byte[rdi]
    je	_compare_colours_while

_colours_no_match:
    xor   rax, rax
    ret

_colours_match:
    mov   rax, 1
    ret

global colors
colors:
    lea rax, [all_colors]
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
