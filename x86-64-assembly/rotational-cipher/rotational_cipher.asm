section .text
global rotate
rotate:
    mov rcx, rdx
    ; Calculate key % 26 to get usable key value
    xor edx, edx
    mov al, 26
    movsx eax, al
    xchg eax, esi
    div esi

_loop:
    ; Rotate if character is an uppercase letter
    mov al, [rdi]
    sub al, 'A'
    cmp al, 25
    ja _not_uppercase
    call _rotate

_not_uppercase:
    ; Rotate if character is a lowercase letter
    add al, 'A'
    sub al, 'a'
    cmp al, 25
    ja _next
    call _rotate

_next:
    add al, 'a'
    mov [rcx], al
    inc rdi
    inc rcx
    cmp al, 0
    jnz _loop
    jmp _done

_rotate:
    add al, dl
    cmp al, 25
    jbe _done
    sub al, 26

_done:
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
