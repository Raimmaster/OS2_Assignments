import os
import sys
import cmd
import signal
from Matrix import Matrix

class MatrixShell(cmd.Cmd):
    intro = 'Welcome to the access matrix shell.\tType help or ? to list commands.\n'
    matrix = Matrix()
    prompt = '(%s)>' % matrix.get_current_domain().name
    # ----- basic turtle commands -----
    def do_su(self, name):
        'Set the current domain: su D1'
        self.matrix.set_current_domain(name)
        self.prompt = '(%s)>' % self.matrix.get_current_domain().name
    
    def do_create_domain(self, arg):
        'Create a domain: create_domain D1'
        self.matrix.create_domain(arg)
    
    def do_create_object(self, arg):
        'Create an object: create_object F1'
        objeto = self.matrix.create_object(arg, self.matrix.get_current_domain())
        print('%s' % objeto.name)
        self.matrix.get_current_domain().print_current_objects()
    
    def do_delete_domain(self, arg):
        'Delete a domain: delete_domain D1'
    
    def do_delete_object(self, arg):
        'Delete an object: delete_object F1'
    
    def do_set_access_right(self, arg):
        'Set a domain/object access right: set_access_right read'
    
    def do_rm_access_right(self, arg):
        'Remove a domain/object access right: rm_access_right read'
    
    def do_verify_access_right(self, arg):
        'Verify a domain\'s access right, and get the objects where it can use it: verify_access_right read'
    
    def do_switch(self, arg):
        'Switch to another domain, if allowed: switch'
    
    def do_exit(self, arg):
        'Return to previous domain: exit'
    
    def do_quit(self, arg):
        'Exit from the program  '
        quit()

if __name__ == '__main__':
    MatrixShell().cmdloop()
