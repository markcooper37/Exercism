default rel

section .rodata
    beginning: db "One for "
    beginning_len: equ $-beginning
    you: db "you"
    you_len: equ $-you
    ending: db ", one for me.", 0
    ending_len: equ $-ending

section .text
global two_fer
two_fer:
    mov rax, rdi
    mov rdi, rsi
    ; Beginning
    lea rsi, [beginning]
    mov rcx, beginning_len
    rep movsb
    ; Username
    cmp rax, 0
    jne user_name
    lea rsi, [you]
    mov rcx, you_len
    rep movsb

add_ending:
    ; Ending
    lea rsi, [ending]
    mov rcx, ending_len
    rep movsb
    ret

user_name:
    mov rsi, rax

move_char:
    cmp byte[rsi], 0
    je add_ending
    movsb
    loop move_char

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
