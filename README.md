# kconfig
kconfig

# resources used
https://stackoverflow.com/questions/53256373/sending-file-over-ssh-in-go

# Requirements
research best input file format - pass as command in running

no input file save dir

for rke2 and k3s compatible and other orchestration platform compatible. Use dicts - e.g cluster type

make linux and .exe compatible

efficient

# workflow
input reqs
loop through
ssh in with username and password or ssh key
retrieve config
save configs in structs
add to struct array
then store in compiled file

# use yaml format
save to - 

- name: dev
    type: rke2 - dictionary to map for config dir
    ip: 
    username:
    password:
- name


# have two commads
# first command - config will runn config script
# second command - edit-config will open txt edito to edit the config file
# config file - have the dir to save the config
#  and the node ip, uisername, password and cluster name
# delete config before creating