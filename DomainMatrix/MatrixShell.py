import os
import sys
import cmd
import signal
from Matrix import Matrix

class MatrixShell(cmd.Cmd):
    intro = 'Welcome to the access matrix shell.\tType help or ? to list commands.\n'
    Matrix = Matrix()
    prompt = '%s ' % Matrix.get_current_domain()

    # ----- basic turtle commands -----
    def do_su(self, name):
        'Set the current domain: su D1'
        print ("Domain name received: %s" % name)
        Matrix.set_current_domain(arg);
    def do_whoami(self, arg):
        'Show the current domain name'
    def do_create_domain(self, arg):
        'Create a domain: create_domain D1'
    def do_create_object(self, arg):
        'Create an object: create_object F1'
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
