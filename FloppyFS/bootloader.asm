[BITS 16]
[ORG 0x7C00]

mov dl, 0x0 ;drive 0 = floppy 1
mov dh, 0x0 ;head (0=base)
mov ch, 0x0 ;track/cylinder
mov cl, 0x02 ;sector (1=bootloader, apparently sectors starts counting at 1 instead of 0)
mov bx, 0x1000 ;place in RAM for kernel - I suppose randomly chosen on examples
mov es, bx ;place BX in pointer ES
mov bx, 0x0 ;back to zero - also has something to do with RAM position

ReadFloppy:
mov ah, 0x02
mov al, 0x01
int 0x13
jc ReadFloppy ;if it went wrong, try again

;pointers to RAM position (0x1000)
mov ax, 0x1000
mov ds, ax
mov es, ax
mov fs, ax
mov gs, ax
mov ss, ax

jmp 0x1000:0x0

;assuming we get never back here again, so no further coding needed (kernel handles everything now)

TIMES 510 - ($ - $$) db 0 ;fill resting bytes with zero
DW 0xAA55 ;end of bootloader (2 bytes)
