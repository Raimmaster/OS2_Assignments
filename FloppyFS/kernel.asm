;set print-registers
mov ah, 0x0E ;function nr
mov bh, 0x00 ;page
mov bl, 0x07 ;color

mov si, msg ;move msg to si-pointer
call PrintString ;call function to print si (msg)

jmp $ ;hang

PrintString:
.next_char:
mov al, [si] ;current character
or al, al
jz .print_done ;if current char is zero, go to end
int 0x10 ;print character
inc si ;increase pointer to msg (next character)
jmp .next_char
.print_done:
ret

msg db 'Hello world from the kernel!', 13, 10, 0

TIMES 512 - ($ - $$) db 0 ;fill the rest
