nasm bootloader.asm -f bin -o boot.bin
nasm kernel.asm -f bin -o kernel.bin
dd if=boot.bin of=floppy.img bs=512 count=1
dd if=kernel.bin of=floppy.img bs=512 count=1 seek=1
qemu-system-x86_64 -fda ./floppy.img
